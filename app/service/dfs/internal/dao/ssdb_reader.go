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
	"io"
	"strconv"

	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type SSDBReader struct {
	*model.DfsFileInfo
	ssdb kv.Store
	i    int64 // current reading index
}

func (d *Dao) NewSSDBReader(fileInfo *model.DfsFileInfo) *SSDBReader {
	return &SSDBReader{
		DfsFileInfo: fileInfo,
		ssdb:        d.ssdb,
		i:           0,
	}
}

func (r *SSDBReader) getPartIdx() int {
	if r.i < int64(r.FirstFilePartSize) {
		return 0
	}
	return 1 + (int(r.i)-r.FirstFilePartSize)/r.FilePartSize
}

func (r *SSDBReader) getPartOffset() int {
	if r.i < int64(r.FirstFilePartSize) {
		return int(r.i)
	} else {
		return (int(r.i) - r.FirstFilePartSize) % r.FilePartSize
	}
}

func (r *SSDBReader) Read(p []byte) (n int, err error) {
	if r.i >= r.DfsFileInfo.GetFileSize() {
		return 0, io.EOF
	}

	l := len(p)
	if l == 0 {
		return
	}

	var (
		b []byte
	)

	for idx := r.getPartIdx(); idx < r.DfsFileInfo.FileTotalParts; idx++ {
		b, err = r.readFile(context.Background(), idx)
		if err != nil {
			return 0, err
		}
		// l += len(b) - r.getPartOffset()
		offset := r.getPartOffset()
		if len(b)-offset >= l-n {
			// log.Debugf("endB - offset:%d, idx: %d, i: %d, l: %d, n: %d", offset, idx, r.i, l, n)
			copy(p[n:], b[offset:offset+l-n])
			r.i += int64(l - n)
			n = l
			// log.Debugf("endE - offset:%d, idx: %d, i: %d, l: %d, n: %d", offset, idx, r.i, l, n)
			return
		} else {
			// log.Debugf("b - offset:%d, idx: %d, i: %d, l: %d, n: %d", offset, idx, r.i, l, n)
			copy(p[n:], b[offset:])
			n += len(b) - offset
			r.i += int64(len(b) - offset)
			// log.Debugf("e - offset:%d, idx: %d, i: %d, l: %d, n: %d", offset, idx, r.i, l, n)
		}
	}

	return
}

// Seek implements the io.Seeker interface.
func (r *SSDBReader) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
	case io.SeekEnd:
		abs = r.DfsFileInfo.GetFileSize() + offset
	default:
		return 0, fmt.Errorf("ssdb.Reader.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, fmt.Errorf("ssdb.Reader.Seek: negative position")
	}
	r.i = abs
	return abs, nil
}

func (r *SSDBReader) readFile(ctx context.Context, filePart int) ([]byte, error) {
	var (
		err  error
		k    = getFileKey(r.Creator, r.FileId)
		bBuf string
	)

	bBuf, err = r.ssdb.HgetCtx(ctx, k, strconv.Itoa(int(filePart)))
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, filePart, err)
		return nil, err
	}

	return []byte(bBuf), nil
}

func (r *SSDBReader) ReadAll(ctx context.Context) ([]byte, error) {

	var (
		bytes []byte
		err   error
		bBuf  string
		b     []byte
		k     = getFileKey(r.Creator, r.FileId)
	)

	for i := 0; i < r.FileTotalParts; i++ {
		bBuf, err = r.ssdb.HgetCtx(ctx, k, strconv.Itoa(i))
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
			return nil, err
		}
		b = []byte(bBuf)
		if bytes == nil {
			bytes = make([]byte, 0, len(b)*r.FileTotalParts)
		}
		bytes = append(bytes, b...)
	}

	return bytes, nil
}
