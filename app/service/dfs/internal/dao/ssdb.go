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
	"fmt"
	"strconv"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	_fileKeyPrefix     = "file_%d_%d"
	_fileInfoKeyPrefix = "file_info_%d_%d"
)

func getFileKey(ownerId, fileId int64) string {
	return fmt.Sprintf(_fileKeyPrefix, ownerId, fileId)
}

func getFileInfoKey(ownerId, fileId int64) string {
	return fmt.Sprintf(_fileInfoKeyPrefix, ownerId, fileId)
}

func (d *Dao) WriteFilePartData(ctx context.Context, ownerId, fileId int64, filePart int32, bytes []byte) (err error) {
	var (
		k = getFileKey(ownerId, fileId)
	)

	err = d.ssdb.HsetCtx(ctx, k, strconv.Itoa(int(filePart)), string(bytes))
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(HSET %d, %d, %s) error(%v)", ownerId, fileId, filePart, err)
		return
	}

	_, err = d.ssdb.ExpireCtx(ctx, k, ssdbExpire)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(EXPIRE %d,%d,%s) error(%v)", ownerId, fileId, filePart, err)
		return
	}

	return
}

func (d *Dao) ReadFile(ctx context.Context, ownerId, fileId int64, parts int32) (partLength int32, bytes []byte, err error) {
	var (
		k = getFileKey(ownerId, fileId)
	)

	for i := int32(0); i < parts; i++ {
		var (
			bBuf string
			b    []byte
		)
		bBuf, err = d.ssdb.HgetCtx(ctx, k, strconv.Itoa(int(i)))
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
			return 0, nil, err
		}
		b = []byte(bBuf)
		if bytes == nil {
			bytes = make([]byte, 0, len(b)*int(parts))
		}
		if i == 0 {
			partLength = int32(len(b))
		}
		bytes = append(bytes, b...)
	}

	return
}

func (d *Dao) ReadFileCB(ctx context.Context, ownerId, fileId int64, parts int32, cb func(part int32, bytes []byte) error) (err error) {
	var (
		k    = getFileKey(ownerId, fileId)
		bBuf string
	)

	for i := int32(0); i < parts; i++ {
		bBuf, err = d.ssdb.HgetCtx(ctx, k, strconv.Itoa(int(i)))
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HGET %s, %d) error(%v)", k, i, err)
			return
		}
		if err = cb(i, []byte(bBuf)); err != nil {
			return
		}
	}

	return
}

func (d *Dao) ReadOffsetLimit(ctx context.Context, fileInfo *model.DfsFileInfo, offset, limit int32) (bytes []byte, err error) {
	if limit == 0 && offset != 0 {
		return
	}

	var (
		k    = getFileKey(fileInfo.Creator, fileInfo.FileId)
		bBuf string
		b    []byte
	)

	if limit == 0 && offset == 0 {
		for i := 0; i < fileInfo.FileTotalParts; i++ {
			bBuf, err = d.ssdb.HgetCtx(ctx, k, strconv.Itoa(int(i)))
			if err != nil {
				logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
				return
			}
			b = []byte(bBuf)
			if bytes == nil {
				bytes = make([]byte, 0, len(b)*fileInfo.FileTotalParts)
			}
			bytes = append(bytes, b...)
		}
	} else {
		var (
			bPart = int(offset) / fileInfo.FilePartSize
			ePart = int(offset+limit) / fileInfo.FilePartSize
			bP    = int(offset) % fileInfo.FilePartSize
			eP    = int(offset+limit) % fileInfo.FilePartSize
		)

		bytes = make([]byte, 0, limit)
		for i := bPart; i <= ePart; i++ {
			bBuf, err = d.ssdb.HgetCtx(ctx, k, strconv.Itoa(i))
			if err != nil {
				logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
				return
			}
			b = []byte(bBuf)
			if i == bPart {
				if i == ePart {
					bytes = append(bytes, b[bP:eP]...)
				} else {
					bytes = append(bytes, b[bP:]...)
				}
			} else if i == ePart {
				bytes = append(bytes, b[:eP]...)
			} else {
				bytes = append(bytes, b...)
			}
		}
	}

	return
}

func (d *Dao) OpenFile(ctx context.Context, ownerId, fileId int64, parts int32) (*SSDBReader, error) {
	fileInfo, err := d.GetFileInfo(ctx, ownerId, fileId)
	if err != nil {
		return nil, err
	}
	if parts > 0 {
		// TODO(@benqi): check fileInfo.FileTotalParts == parts
	}
	return d.NewSSDBReader(fileInfo), nil
}
