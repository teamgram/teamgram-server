// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package tg

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
)

type (
	FutureSalt       = mt.FutureSalt
	FutureSaltClazz  = mt.FutureSaltClazz
	TLFutureSalt     = mt.TLFutureSalt
	FutureSalts      = mt.FutureSalts
	FutureSaltsClazz = mt.FutureSaltsClazz
	TLFutureSalts    = mt.TLFutureSalts
)

var (
	MakeTLFutureSalts     = mt.MakeTLFutureSalts
	MakeTLFutureSalt      = mt.MakeTLFutureSalt
	DecodeFutureSaltClazz = mt.DecodeFutureSaltClazz
)
