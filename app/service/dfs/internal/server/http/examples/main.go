// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func main() {
	srv, err := rest.NewServer(rest.RestConf{
		Port: 9091, // 侦听端口
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{Path: "./logs"}, // 日志路径
		},
	})
	if err != nil {
		panic(err)
	}
	defer srv.Stop()
	// 注册路由
	srv.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/dfs/file/:file",
			Handler: getDfsFile,
		},
	})

	srv.Start() // 启动服务
}

func getDfsFile(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		creator int64
		fileId  int64
	)

	var (
		fileName struct {
			FileName string `path:"file"`
		}
	)

	err = httpx.ParsePath(r, &fileName)
	if err != nil {
		logx.Errorf("getDfsFile - error: %v", err)
		httpx.Error(w, err)
		return
	}

	v := strings.Split(fileName.FileName, "_")
	if len(v) != 2 {
		logx.Errorf("getDfsFile - error: invalid fileName: %s", fileName)
		http.NotFound(w, r)
		return
	}

	if creator, err = strconv.ParseInt(v[0], 10, 64); err != nil {
		logx.Errorf("getDfsFile - error: invalid fileName: %s", fileName)
		http.NotFound(w, r)
		return
	}
	_ = creator

	if fileId, err = strconv.ParseInt(strings.Split(v[1], ".")[0], 10, 64); err != nil {
		logx.Errorf("getDfsFile - error: invalid fileName: %s", fileName)
		http.NotFound(w, r)
		return
	}
	_ = fileId

	httpx.WriteJson(w, http.StatusOK, fileName) // 返回结果
}
