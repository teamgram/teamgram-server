// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package service

import (
	"context"
	"encoding/json"

	"github.com/teamgooo/teamgooo-server/pkg/proto/rpc/examples/echo/echo"
	"github.com/teamgooo/teamgooo-server/pkg/proto/rpc/examples/echo/internal/core"

	"github.com/cloudwego/kitex/pkg/klog"
)

// EchoEcho
// echo.echo message:string = Echo;
func (s *Service) EchoEcho(ctx context.Context, request *echo.TLEchoEcho) (*echo.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	klog.Infof("echos.echo - metadata: {}, request: %v", request)

	r, err := c.EchoEcho(request)
	if err != nil {
		return nil, err
	}

	txt, _ := json.Marshal(r)
	klog.Infof("echos.echo - reply: %s", string(txt))
	return r, err
}
