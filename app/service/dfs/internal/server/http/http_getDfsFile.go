/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDfsFile(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		f, err := DfsOpenFile(ctx, creator, fileId)
		if err != nil {
			logx.Errorf("getDfsFile(%d, %d) - error: invalid fileName: %s", creator, fileId, fileName)
			http.NotFound(w, r)
			return
		}

		http.ServeContent(w, r, fileName.FileName, f.Modtime, f.ReadSeeker)
	}
}

func DfsOpenFile(ctx *svc.ServiceContext, creatorId, fileId int64) (f *model.DfsHttpFileInfo, err error) {
	fileInfo, err := ctx.Dao.GetFileInfo(context.Background(), creatorId, fileId)
	if err != nil {
		return nil, err
	}

	return &model.DfsHttpFileInfo{
		ReadSeeker: ctx.Dao.NewSSDBReader(fileInfo),
		Name:       fileInfo.FileName,
		Modtime:    time.Unix(fileInfo.Mtime, 0),
	}, nil
}
