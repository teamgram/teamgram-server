// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package tg

import (
	"fmt"
	"testing"
)

func TestNewRpcError(t *testing.T) {
	fmt.Println(NewRpcError(nil))
	fmt.Println(NewRpcError(fmt.Errorf("test error")))
	e := NewRpcError(ErrInternalServerError)
	fmt.Println(e)
	fmt.Println(NewRpcError(e))
}
