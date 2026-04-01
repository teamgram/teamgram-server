package iface

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type TLObjectSizer interface {
	CalcSize(layer int32) int
}

type TLObjectValidator interface {
	Validate(layer int32) error
}

func MarshalWithName(clazzName string, obj TLObject) ([]byte, error) {
	if obj == nil {
		return []byte("null"), nil
	}

	rv := reflect.ValueOf(obj)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return []byte("null"), nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return json.Marshal(WithNameWrapper{
			ClazzName: clazzName,
			TLObject:  obj,
		})
	}

	out := make(map[string]any, rv.NumField())
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		if field.PkgPath != "" {
			continue
		}
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := field.Name
		if tag != "" {
			if idx := strings.IndexByte(tag, ','); idx >= 0 {
				tag = tag[:idx]
			}
			if tag != "" {
				name = tag
			}
		}
		out[name] = rv.Field(i).Interface()
	}

	return json.Marshal(map[string]any{
		"_name":   clazzName,
		"_object": out,
	})
}

func CalcStringSize(v string) int {
	return calcTLBytesSize(len(v))
}

func CalcBytesSize(v []byte) int {
	return calcTLBytesSize(len(v))
}

func CalcObjectSize(obj TLObject, layer int32) int {
	if obj == nil {
		return 0
	}
	if s, ok := obj.(TLObjectSizer); ok {
		return s.CalcSize(layer)
	}

	data, err := EncodeObject(obj, layer)
	if err != nil {
		return 0
	}
	return len(data)
}

func CalcObjectListSize[T TLObject](vList []T, layer int32) int {
	size := 8
	for _, obj := range vList {
		size += CalcObjectSize(obj, layer)
	}
	return size
}

func CalcBareObjectVectorSize[T TLObject](vList []T, layer int32) int {
	size := 4
	for _, obj := range vList {
		size += CalcObjectSize(obj, layer)
	}
	return size
}

func CalcInt32ListSize(vList []int32) int {
	return 8 + len(vList)*4
}

func CalcInt64ListSize(vList []int64) int {
	return 8 + len(vList)*8
}

func CalcStringListSize(vList []string) int {
	size := 8
	for _, v := range vList {
		size += CalcStringSize(v)
	}
	return size
}

func CalcBytesListSize(vList [][]byte) int {
	size := 8
	for _, v := range vList {
		size += CalcBytesSize(v)
	}
	return size
}

func ValidateRequiredString(field string, v string) error {
	if v == "" {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidateRequiredBytes(field string, v []byte) error {
	if v == nil {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidateRequiredObject(field string, v TLObject) error {
	if v == nil {
		return fmt.Errorf("%s is required", field)
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map:
		if rv.IsNil() {
			return fmt.Errorf("%s is required", field)
		}
	}
	if !rv.IsValid() {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidateRequiredObjectWithLayer(field string, v TLObject, layer int32) error {
	if err := ValidateRequiredObject(field, v); err != nil {
		return err
	}
	if validator, ok := v.(TLObjectValidator); ok {
		return validator.Validate(layer)
	}
	return nil
}

func ValidateRequiredSlice(field string, v any) error {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return fmt.Errorf("%s is required", field)
	}
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("%s must be a slice", field)
	}
	if rv.IsNil() {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidateRequiredSliceWithLayer(field string, v any, layer int32) error {
	if err := ValidateRequiredSlice(field, v); err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i)
		if !item.IsValid() {
			return fmt.Errorf("%s[%d] is required", field, i)
		}

		if item.Kind() == reflect.Interface || item.Kind() == reflect.Pointer {
			if item.IsNil() {
				return fmt.Errorf("%s[%d] is required", field, i)
			}
		}

		if item.CanInterface() {
			if validator, ok := item.Interface().(TLObjectValidator); ok {
				if err := validator.Validate(layer); err != nil {
					return err
				}
				continue
			}
		}

		for item.IsValid() && (item.Kind() == reflect.Interface || item.Kind() == reflect.Pointer) {
			if item.IsNil() {
				return fmt.Errorf("%s[%d] is required", field, i)
			}
			item = item.Elem()
		}
		if !item.IsValid() || !item.CanInterface() {
			continue
		}
		if validator, ok := item.Interface().(TLObjectValidator); ok {
			if err := validator.Validate(layer); err != nil {
				return err
			}
		}
	}
	return nil
}

func calcTLBytesSize(n int) int {
	headerLen := 1
	if n > 253 {
		headerLen = 4
	}
	size := headerLen + n
	if rem := size % 4; rem != 0 {
		size += 4 - rem
	}
	return size
}
