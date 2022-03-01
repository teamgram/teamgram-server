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

package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/teamgram/marmota/pkg/commands"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	c config.Config
)

func main() {
	commands.Run(New())
}

type Server struct {
	dao *dao.Dao
}

func New() *Server {
	s := new(Server)
	return s
}

func (s *Server) Initialize() error {
	s.dao = dao.New(c)
	return nil
}

func (s *Server) RunLoop() {
	defer func() {
		s.Destroy()
		os.Exit(0)
	}()

	logx.Infof("aaaa", "aa")
	// buf, err := ioutil.ReadFile("./test001.jpeg")
	// ../../../../../../../tools/gif2mp4/safe_image.gif
	buf, err := ioutil.ReadFile("../../../../../../../tools/gif2mp4/safe_image.gif")
	if err != nil {
		logx.Error("open error: %v", err)
		return
	}

	logx.Infof("aaaa", "bb")

	szParts := len(buf) / 10240
	lastSize := len(buf) % 10240
	_ = lastSize
	// idx := sz/10240
	for i := 0; i < szParts; i++ {
		s.dao.WriteFilePartData(context.Background(), 100, 1, int32(i), buf[i*10240:(i+1)*10240])
	}
	logx.Infof("aaaa", "cc")
	if lastSize > 0 {
		logx.Infof("aaaa: lastSize = %d, bufSize = %d", lastSize, len(buf))
		// szParts++
		s.dao.WriteFilePartData(context.Background(), 100, 1, int32(szParts), buf[szParts*10240:])
	}

	logx.Infof("aaaa", "dd")

	if err = s.dao.SetFileInfo(context.Background(), &model.DfsFileInfo{
		Creator:        100,
		FileId:         1,
		Big:            false,
		FileName:       "safe_image.gif",
		FileTotalParts: szParts + 1,
		FilePartSize:   10240,
		// FileSize:       int64(len(buf)),
	}); err != nil {
		logx.Errorf("%v", err)
	}
}

func (s *Server) Destroy() {
	// s.dao.Close()
}
