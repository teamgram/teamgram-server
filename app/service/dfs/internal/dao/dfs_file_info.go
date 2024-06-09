// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"strconv"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

func (d *Dao) SetFileInfo(ctx context.Context, fileInfo *model.DfsFileInfo) (err error) {
	var (
		key = getFileInfoKey(fileInfo.Creator, fileInfo.FileId)
		// args = append([]interface{}{key}, fileInfo.ToArgs()...)
	)

	// TODO(@benqi): args error??
	if err = d.ssdb.HmsetCtx(ctx, key, fileInfo.ToArgs()); err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(HMSET %s,%v) error(%v)", key, fileInfo, err)
		return
	}

	if _, err = d.ssdb.ExpireCtx(ctx, key, ssdbExpire); err != nil {
		logx.WithContext(ctx).Error("conn.Send(EXPIRE %d,%d) error(%v)", key, ssdbExpire, err)
	}

	return
}

func (d *Dao) GetFileInfo(ctx context.Context, ownerId, fileId int64) (fileInfo *model.DfsFileInfo, err error) {
	var (
		values map[string]string
		k      = getFileInfoKey(ownerId, fileId)
	)

	if values, err = d.ssdb.HgetallCtx(ctx, k); err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(HGETALL %s) error(%v)", k, err)
		return
	} else if len(values) == 0 {
		err = model.ErrorDfsFileNotFound
		logx.WithContext(ctx).Infof("conn.Do(HGETALL %s) is error(%v)", k, err)
		return
	}

	fileInfo = &model.DfsFileInfo{
		Creator:           ownerId,
		FileId:            fileId,
		Big:               false,
		FileName:          "",
		FileTotalParts:    0,
		FirstFilePartSize: 0,
		FilePartSize:      0,
		LastFilePartSize:  0,
	}

	for k2, v := range values {
		switch k2 {
		case "big":
			if v == "true" {
				fileInfo.Big = true
			}
		case "file_name":
			fileInfo.FileName = v
		case "file_total_parts":
			fileTotalParts, _ := strconv.Atoi(v)
			fileInfo.FileTotalParts = fileTotalParts
		case "first_file_part_size":
			firstFilePartSize, _ := strconv.Atoi(v)
			fileInfo.FirstFilePartSize = firstFilePartSize
		case "file_part_size":
			filePartSize, _ := strconv.Atoi(v)
			fileInfo.FilePartSize = filePartSize
		case "last_file_part_size":
			lastFilePartSize, _ := strconv.Atoi(v)
			fileInfo.LastFilePartSize = lastFilePartSize
		case "mtime":
			fileInfo.Mtime, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	if fileInfo.FileTotalParts == 0 {
		var (
			fileInfo2 = &model.DfsFileInfo{
				Creator: fileInfo.Creator,
				FileId:  fileInfo.FileId,
			}
		)

		fileInfo2.FileTotalParts, fileInfo2.LastFilePartSize, err = d.getFileTotalPartsByFile(ctx, fileInfo.Creator, fileInfo.FileId)
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HLEN %s) error(%v)", k, err)
			return
		}

		// save fileInfo
		d.SetFileInfo(ctx, fileInfo2)

		// refill to fileInfo
		fileInfo.FileTotalParts = fileInfo2.FileTotalParts
		fileInfo.LastFilePartSize = fileInfo2.LastFilePartSize
	}

	return fileInfo, nil
}

func (d *Dao) getFileTotalPartsByFile(ctx context.Context, ownerId, fileId int64) (fileTotalParts, lastFilePartSize int, err error) {
	var (
		k    = getFileKey(ownerId, fileId)
		bBuf string
	)
	if fileTotalParts, err = d.ssdb.HlenCtx(ctx, k); err != nil {
		logx.WithContext(ctx).Errorf("getFileTotalPartsByFile: conn.Do(HLEN %s) error(%v)", k, err)
		return
	} else if fileTotalParts == 0 {
		err = model.ErrorDfsFileNotFound
		logx.WithContext(ctx).Infof("getFileTotalPartsByFile: conn.Do(HLEN %s) error(%v)", k, err)
		return
	}

	if bBuf, err = d.ssdb.HgetCtx(ctx, k, strconv.Itoa(fileTotalParts-1)); err != nil {
		logx.WithContext(ctx).Errorf("getFileTotalPartsByFile: conn.Do(HGET %s %d) error(%v), fileTotalParts(%d)", k, err, fileTotalParts-1)
		return
	}

	lastFilePartSize = len([]byte(bBuf))
	return
}
