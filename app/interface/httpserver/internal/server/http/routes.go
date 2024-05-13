// Copyright (c) 2024-present, Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package http

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/core"
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api",
				Handler: apiw1(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apiw1",
				Handler: apiw1(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apiw1_test",
				Handler: apiw1(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apiw1_premium",
				Handler: apiw1(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/apiw2",
				Handler: apiw1Debug(serverCtx),
			},
		},
	)
}

func apiw1(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength <= 0 || r.Body == nil {
			http.NotFound(w, r)
		} else {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.NotFound(w, r)
			} else {
				authKeyId := int64(binary.LittleEndian.Uint64(body))
				message := mtproto.NewMTPRawMessage(authKeyId, 0, 2)
				message.Decode(body)

				c := core.New2(
					r.Context(),
					ctx,
					&metadata.RpcMetadata{
						ServerId:      ctx.Dao.GetGatewayId(),
						ClientAddr:    httpx.GetRemoteAddr(r),
						AuthId:        authKeyId,
						SessionId:     0,
						ReceiveTime:   time.Now().Unix(),
						UserId:        0,
						ClientMsgId:   0,
						IsBot:         false,
						Layer:         0,
						Client:        "",
						IsAdmin:       false,
						Takeout:       nil,
						Langpack:      "",
						PermAuthKeyId: 0,
					})
				rData, err2 := c.HttpserverApiw1(message)
				if err2 != nil {
					http.Error(w, err2.Error(), http.StatusInternalServerError)
				} else {
					w.Header().Set("Access-Control-Allow-Headers", "origin, content-type")
					w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Max-Age", "1728000")
					w.Header().Set("Cache-control", "no-store")
					w.Header().Set("Connection", "keep-alive")
					w.Header().Set("Content-type", "application/octet-stream")
					w.Header().Set("Pragma", "no-cache")
					w.Header().Set("Strict-Transport-Security", "max-age=15768000")
					w.WriteHeader(200)
					w.Write(rData.Payload)
				}
			}
		}
	}
}

func apiw1Debug(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		http.NotFound(w, r)
	}
}
