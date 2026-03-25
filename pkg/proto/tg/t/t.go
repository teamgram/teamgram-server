// Copyright © 2025 The Teamgooo Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package main

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func Match(m *tg.InputUser, f ...interface{}) {
	switch c := m.InputUserClazz.(type) {
	case *tg.TLInputUserEmpty:
		for _, v := range f {
			if f1, ok := v.(func(c *tg.TLInputUserEmpty) interface{}); ok {
				f1(c)
			}
		}
	case *tg.TLInputUserSelf:
		for _, v := range f {
			if f1, ok := v.(func(c *tg.TLInputUserSelf) interface{}); ok {
				f1(c)
			}
		}
	case *tg.TLInputUser:
		for _, v := range f {
			if f1, ok := v.(func(c *tg.TLInputUser) interface{}); ok {
				f1(c)
			}
		}
	case *tg.TLInputUserFromMessage:
		for _, v := range f {
			if f1, ok := v.(func(c *tg.TLInputUserFromMessage) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

func Match2(m *tg.InputUser,
	f1 func(c *tg.TLInputUserEmpty) interface{},
	f2 func(c *tg.TLInputUserSelf) interface{},
	f3 func(c *tg.TLInputUser) interface{},
	f4 func(c *tg.TLInputUserFromMessage) interface{}) {
	if m == nil || m.InputUserClazz == nil {
		return
	}

	switch c := m.InputUserClazz.(type) {
	case *tg.TLInputUserEmpty:
		f1(c)
	case *tg.TLInputUserSelf:
		f2(c)
	case *tg.TLInputUser:
		f3(c)
	case *tg.TLInputUserFromMessage:
		f4(c)
	default:
		fmt.Printf("Match2: unknown InputUserClazz type: %T\n", m.InputUserClazz)
	}
}

func MatchTN[T1, T2, T3, T4 tg.InputUserClazz](m *tg.InputUser, f1 func(c T1) interface{}, f2 func(c T2) interface{}, f3 func(c T3) interface{}, f4 func(c T4) interface{}) {
	if m == nil || m.InputUserClazz == nil {
		return
	}

	switch c := m.InputUserClazz.(type) {
	case T1:
		f1(c)
	case T2:
		f2(c)
	case T3:
		f3(c)
	case T4:
		f4(c)
	default:
		fmt.Printf("MatchT: unknown InputUserClazz type: %T\n", m.InputUserClazz)
		// TODO(@benqi): handle error
		// return fmt.Errorf("unknown InputUserClazz type: %T", m.InputUserClazz)
		return
	}
}

func main() {
	//r := &tg.TLAuthCancelCode{
	//	PhoneNumber:   "12345",
	//	PhoneCodeHash: "12345",
	//}
	//
	//fmt.Printf("%s\n", iface.WithNameWrapper{ClazzName: "", TLObject: r})

	//var f = []func[T tg.InputUserClazz](c T) interface{}{
	//
	//}
	//	func(c *tg.TLInputUserEmpty) interface{} {
	//		fmt.Printf("f1: %v\n", c)
	//		return nil
	//	},
	//	func(c *tg.TLInputUserEmpty) interface{} {
	//		fmt.Printf("f2: %v\n", c)
	//		return nil
	//	},
	//}

	m := tg.MakeInputUser(&tg.TLInputUserEmpty{
		//
	})

	Match(m,
		func(c *tg.TLInputUserEmpty) interface{} {
			fmt.Printf("f1: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUserSelf) interface{} {
			fmt.Printf("f2: %v\n", c)
			return nil
		})

	Match2(m,
		func(c *tg.TLInputUserEmpty) interface{} {
			fmt.Printf("f1: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUserSelf) interface{} {
			fmt.Printf("f2: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUser) interface{} {
			fmt.Printf("f3: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUserFromMessage) interface{} {
			fmt.Printf("f4: %v\n", c)
			return nil
		})

	MatchTN(m,
		func(c *tg.TLInputUserEmpty) interface{} {
			fmt.Printf("f1: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUserSelf) interface{} {
			fmt.Printf("f2: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUser) interface{} {
			fmt.Printf("f3: %v\n", c)
			return nil
		},
		func(c *tg.TLInputUserFromMessage) interface{} {
			fmt.Printf("f4: %v\n", c)
			return nil
		})
}
