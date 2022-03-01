// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package server

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type Logger struct {
	logx.Logger
}

func NewLogger() Logger {
	return Logger{}
}
