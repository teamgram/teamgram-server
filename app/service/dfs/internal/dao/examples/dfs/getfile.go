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
	"flag"
	"os"

	"github.com/teamgram/marmota/pkg/commands"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"
	"github.com/zeromicro/go-zero/core/conf"

	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/dfs.yaml", "the config file")

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
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	s.dao = dao.New(c)
	return nil
}

func (s *Server) RunLoop() {
	defer func() {
		s.Destroy()
		os.Exit(0)
	}()

	fileInfo, err := s.dao.GetFileInfo(context.Background(), 7997959636588162716, -8695032368284712706)
	if err != nil {
		logx.Errorf("open error: %v", err)
		// panic(err)
		return
	} else {
		logx.Debugf("fileInfo: %v", fileInfo)
	}

}

func (s *Server) Destroy() {
	// s.dao.Close()
}
