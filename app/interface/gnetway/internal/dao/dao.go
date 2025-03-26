// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/config"
)

type Dao struct {
}

func New(c config.Config) *Dao {
	return &Dao{}
}
