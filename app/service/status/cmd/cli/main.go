// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"context"

	status_client "github.com/teamgram/teamgram-server/app/service/status/client"
	"github.com/teamgram/teamgram-server/app/service/status/status"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	StatusClient zrpc.RpcClientConf
}

func main() {
	var (
		c Config
	)
	conf.MustLoad("./cli.yaml", &c)
	logx.Info(c)

	cli := status_client.NewStatusClient(zrpc.MustNewClient(c.StatusClient))
	rList, _ := cli.StatusGetUsersOnlineSessionsList(context.Background(), &status.TLStatusGetUsersOnlineSessionsList{
		Users: []int64{
			136917740,
			136917738,
			136817694,
			136917734,
			136917739,
			136847696,
		},
	})
	for _, rV := range rList.GetDatas() {
		logx.Infof("rV : %s", rV)
	}
}
