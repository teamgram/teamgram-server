// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"context"
	"flag"

	media_client "github.com/teamgram/teamgram-server/app/service/media/client"
	"github.com/teamgram/teamgram-server/app/service/media/media"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	documentId = flag.Int64("id", 0, "the document id")
	configFile = flag.String("f", "./cli.yaml", "the config file")
)

type Config struct {
	MediaClient zrpc.RpcClientConf
}

func main() {
	flag.Parse()

	var (
		c Config
	)
	conf.MustLoad(*configFile, &c)
	logx.Info(c)

	cli := media_client.NewMediaClient(zrpc.MustNewClient(c.MediaClient))
	document, err := cli.MediaGetDocument(context.Background(), &media.TLMediaGetDocument{
		Id: *documentId,
	})
	if err != nil {
		logx.Errorf("getDocument - error: %v", err)
	} else {
		logx.Infof("getDocument - reply: %s", document)
	}
}
