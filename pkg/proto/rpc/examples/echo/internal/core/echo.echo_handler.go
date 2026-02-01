// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"github.com/teamgooo/teamgooo-server/pkg/proto/rpc/examples/echo/echo"
)

func (c *EchoCore) EchoEcho(in *echo.TLEchoEcho) (resp *echo.Echo, err error) {
	// do something here
	/*
		temp, ok1 := metainfo.GetValue(c.ctx, "temp")
		logid, ok2 := metainfo.GetPersistentValue(c.ctx, "logid")

		if !(ok1 && ok2) {
			// panic("It looks like the protocol does not support transmitting meta information")
		}
		klog.Debug(temp)  // "temp-value"
		klog.Debug(logid) // "12345"
		klog.Debug("echo called")
	*/

	return echo.MakeEcho(&echo.TLEcho{
		Message: in.Message,
	}), nil
}
