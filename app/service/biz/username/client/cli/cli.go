// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: Teamgram (teamgram.io@gmail.com)
//

package main

import (
	"context"

	username_client "github.com/teamgram/teamgram-server/app/service/biz/username/client"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Username zrpc.RpcClientConf
}

func main() {
	var (
		c Config
	)

	conf.MustLoad("./cli.yaml", &c)
	cli := username_client.NewUsernameClient(zrpc.MustNewClient(c.Username))
	for i := 0; i < 1; i++ {
		rValue, err := cli.UsernameGetAccountUsername(context.Background(), &username.TLUsernameGetAccountUsername{
			UserId: 6,
		})
		if err != nil {
			logx.Errorf("username.getAccountUsername - error: %v", err)
		} else {
			logx.Infof("username.getAccountUsername - reply: %v", rValue)
		}
	}
}
