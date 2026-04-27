// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iface

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type clazzNamer interface {
	ClazzName() string
}

// DebugString returns a compact, flat JSON representation intended for logs.
// It keeps TL identity in @type/@id and avoids the MarshalJSON _name/_object
// wrapper so nested TL objects are easier to inspect.
func DebugString(obj any) string {
	return DebugStringWithName("", obj)
}

// DebugStringWithName is DebugString with a constructor-name hint for generated
// String methods where ClazzName2 may not be populated on hand-built values.
func DebugStringWithName(clazzName string, obj any) string {
	data, err := MarshalDebugJSONWithName(clazzName, obj)
	if err != nil {
		return fmt.Sprintf("<debug json error: %v>", err)
	}
	return string(data)
}

// MarshalDebugJSON returns DebugString's JSON bytes.
func MarshalDebugJSON(obj any) ([]byte, error) {
	return MarshalDebugJSONWithName("", obj)
}

// MarshalDebugJSONWithName returns DebugStringWithName's JSON bytes.
func MarshalDebugJSONWithName(clazzName string, obj any) ([]byte, error) {
	return json.Marshal(buildDebugValue(reflect.ValueOf(obj), clazzName))
}

type DebugLazyValue struct {
	Obj any
}

func (v DebugLazyValue) String() string {
	return DebugString(v.Obj)
}

func DebugLazy(obj any) fmt.Stringer {
	return DebugLazyValue{Obj: obj}
}

func buildDebugValue(v reflect.Value, clazzNameHint string) any {
	v = unwrapNilable(v)
	if !v.IsValid() {
		return nil
	}

	if unwrapped, ok := unwrapClazzWrapper(v); ok {
		return buildDebugValue(unwrapped, clazzNameHint)
	}

	switch v.Kind() {
	case reflect.Struct:
		return buildDebugStruct(v, clazzNameHint)
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return v.Bytes()
		}
		out := make([]any, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			out = append(out, buildDebugValue(v.Index(i), ""))
		}
		return out
	case reflect.Array:
		out := make([]any, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			out = append(out, buildDebugValue(v.Index(i), ""))
		}
		return out
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			if v.CanInterface() {
				return v.Interface()
			}
			return nil
		}
		out := make(map[string]any, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			out[iter.Key().String()] = buildDebugValue(iter.Value(), "")
		}
		return out
	default:
		if v.CanInterface() {
			return v.Interface()
		}
		return nil
	}
}

func buildDebugStruct(v reflect.Value, clazzNameHint string) any {
	out := make(map[string]any, v.NumField()+2)
	if clazzName := debugClazzName(v, clazzNameHint); clazzName != "" {
		out["@type"] = clazzName
	}
	if clazzID, ok := debugClazzID(v); ok {
		out["@id"] = fmt.Sprintf("0x%08x", clazzID)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		name, ok := debugJSONFieldName(field)
		if !ok || name == "_id" || name == "_name" || field.Name == "ClazzID" || field.Name == "ClazzName2" {
			continue
		}

		fv := v.Field(i)
		if isDebugZero(fv) {
			continue
		}
		out[name] = buildDebugValue(fv, "")
	}
	return out
}

func unwrapNilable(v reflect.Value) reflect.Value {
	for v.IsValid() && (v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface) {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
}

func unwrapClazzWrapper(v reflect.Value) (reflect.Value, bool) {
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	t := v.Type()
	exported := 0
	clazzIndex := -1
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		exported++
		if field.Name == "Clazz" {
			clazzIndex = i
		}
	}
	if exported == 1 && clazzIndex >= 0 {
		fv := v.Field(clazzIndex)
		if !isDebugZero(fv) {
			return fv, true
		}
	}
	return reflect.Value{}, false
}

func debugClazzName(v reflect.Value, fallback string) string {
	if v.CanInterface() {
		if n, ok := v.Interface().(clazzNamer); ok {
			if name := n.ClazzName(); name != "" {
				return name
			}
		}
	}
	if v.CanAddr() {
		if n, ok := v.Addr().Interface().(clazzNamer); ok {
			if name := n.ClazzName(); name != "" {
				return name
			}
		}
	}
	return fallback
}

func debugClazzID(v reflect.Value) (uint32, bool) {
	if v.Kind() != reflect.Struct {
		return 0, false
	}
	f := v.FieldByName("ClazzID")
	if !f.IsValid() || f.Kind() != reflect.Uint32 {
		return 0, false
	}
	id := uint32(f.Uint())
	return id, id != 0
}

func debugJSONFieldName(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("json")
	if tag == "-" {
		return "", false
	}
	if tag == "" {
		return field.Name, true
	}
	if idx := strings.IndexByte(tag, ','); idx >= 0 {
		tag = tag[:idx]
	}
	if tag == "" {
		return field.Name, true
	}
	return tag, true
}

func isDebugZero(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Func:
		return true
	default:
		return false
	}
}
