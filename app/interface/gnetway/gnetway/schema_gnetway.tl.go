/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package gnetway

import (
	// "encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields
