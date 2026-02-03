// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"encoding/hex"
	"fmt"

	"github.com/teamgooo/teamgooo-server/pkg/proto/bin"
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"
	"github.com/teamgooo/teamgooo-server/pkg/proto/mt"
)

const (
	ClazzID_echo       = 0x05162463
	ClazzID_echo2      = 0x05162464
	ClazzID_echos_echo = 0x05162465
)

func init() {
	iface.RegisterClazzID(ClazzID_echo, func() iface.TLObject { return &TLEcho{} })
	iface.RegisterClazzID(ClazzID_echo2, func() iface.TLObject { return &TLEcho2{} })
	iface.RegisterClazzID(ClazzID_echos_echo, func() iface.TLObject { return &TLEchosEcho{} })
}

type TLEcho struct {
	Message string `json:"message"`
}

func (m *TLEcho) ClazzName() string {
	return "echo"
}

func (m *TLEcho) EchoName() string {
	return m.ClazzName()
}

func (m *TLEcho) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echo)
	x.PutString(m.Message)
	return nil
}

func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type TLEcho2 struct {
	Message string `json:"message"`
}

func (m *TLEcho2) ClazzName() string {
	return "echo2"
}

func (m *TLEcho2) EchoName() string {
	return m.ClazzName()
}

func (m *TLEcho2) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echo2)
	x.PutString(m.Message)
	return nil
}

func (m *TLEcho2) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type EchoClazz interface {
	EchoName() string
}

type Echo struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	// EchoName() string
	EchoClazz
}

func MakeEcho(e EchoClazz) *Echo {
	switch c := e.(type) {
	case *TLEcho:
		return &Echo{
			ClazzID:   ClazzID_echo,
			ClazzName: "echo",
			EchoClazz: c,
		}
	case *TLEcho2:
		return &Echo{
			ClazzID:   ClazzID_echo2,
			ClazzName: "echo2",
			EchoClazz: c,
		}
	default:
		//
	}
	return nil
}

func (m *Echo) Encode(x *bin.Encoder, layer int32) error {
	clazzName := m.ClazzName
	if clazzName == "" {
		clazzName = iface.GetClazzNameByID(m.ClazzID)
	}

	switch clazzName {
	case "echo":
		t := &TLEcho{}
		return t.Encode(x, layer)
	case "echo2":
		t := &TLEcho2{}
		return t.Encode(x, layer)
	default:
		return nil
	}
}

func (m *Echo) Decode(d *bin.Decoder) (err error) {
	m.ClazzID, err = d.ClazzID()
	if err != nil {
		return
	}

	switch m.ClazzID {
	case ClazzID_echo:
		m.ClazzName = "echo"
	case ClazzID_echo2:
		m.ClazzName = "echo2"
	default:
		err = fmt.Errorf("invalid constructorId: 0x%x", m.ClazzID)
	}

	return
}

func (m *Echo) Match(f1 func(c *TLEcho) interface{}, f2 func(c *TLEcho2) interface{}) interface{} {
	switch c := m.EchoClazz.(type) {
	case *TLEcho:
		return f1(c)
	case *TLEcho2:
		return f2(c)
	default:
		//
	}
	return nil
}

func (m *Echo) Match2(f ...interface{}) {
	switch c := m.EchoClazz.(type) {
	case *TLEcho:
		for _, v := range f {
			if f1, ok := v.(func(c *TLEcho) interface{}); ok {
				f1(c)
			}
		}
	case *TLEcho2:
		for _, v := range f {
			if f1, ok := v.(func(c *TLEcho2) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

type TLEchosEcho struct {
	Message string `json:"message"`
}

func (m *TLEchosEcho) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echos_echo)
	x.PutString(m.Message)
	return nil
}

func (m *TLEchosEcho) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type RPCEchos interface {
	EchosEcho(in *TLEchosEcho) (*Echo, error)
}

func main() {
	//var (
	//	echo Echo
	//)
	//
	//echo = &TLEcho{Message: "hello"}
	//
	//switch t := echo.(type) {
	//case nil:
	//	fmt.Println("nil")
	//case *TLEcho:
	//	fmt.Println(t.EchoName())
	//case *TLEcho2:
	//	fmt.Println(t.EchoName())
	//default:
	//	panic("unknown type")
	//}

	echo := MakeEcho(&TLEcho{Message: "echo"})
	echo.Match(
		func(c *TLEcho) interface{} {
			fmt.Println(c.Message)
			return nil
		},
		func(c *TLEcho2) interface{} {
			fmt.Println(c.Message)
			return nil
		})

	echo = MakeEcho(&TLEcho2{Message: "echo2"})
	echo.Match2(
		func(c *TLEcho) interface{} {
			fmt.Println(c.Message)
			return nil
		},
		func(c *TLEcho2) interface{} {
			fmt.Println(c.Message)
			return nil
		})
	reqPQ := &mt.TLReqPq{
		ClazzID: mt.ClazzID_resPQ,
		Nonce:   bin.Int128{},
	}

	x := bin.NewEncoder()
	_ = reqPQ.Encode(x, 0)
	fmt.Println(hex.EncodeToString(x.Bytes()))
}
