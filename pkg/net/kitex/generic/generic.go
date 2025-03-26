// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package generic

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type Generic = generic.Generic

type binaryTgGeneric struct {
}

// BinaryTgGeneric tg Generic
func BinaryTgGeneric() Generic {
	return &binaryTgGeneric{}
}

func (g *binaryTgGeneric) Close() error {
	return nil
}

func (g *binaryTgGeneric) PayloadCodec() remote.PayloadCodec {
	return nil
}

func (g *binaryTgGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	return 0
}

func (g *binaryTgGeneric) Framed() bool {
	return true
}

func (g *binaryTgGeneric) GetMethod(req interface{}, method string) (*generic.Method, error) {
	return nil, nil
}

func (g *binaryTgGeneric) IDLServiceName() string {
	return ""
}

func (g *binaryTgGeneric) MessageReaderWriter() interface{} {
	return nil
}
