// Copyright (c) 2021-present,  Teamgooo Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

func removeAllNil(vList []tg.AuthorizationClazz) []tg.AuthorizationClazz {
	for i := 0; i < len(vList); {
		if vList[i] != nil {
			i++
			continue
		}

		if i < len(vList)-1 {
			copy(vList[i:], vList[i+1:])
		}

		vList[len(vList)-1] = nil
		vList = vList[:len(vList)-1]
	}

	return vList
}
