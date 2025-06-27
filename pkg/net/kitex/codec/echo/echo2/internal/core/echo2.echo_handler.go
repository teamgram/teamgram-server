// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2"
)

var _ *tg.Bool

// Echo2Echo
// echo2.echo message:string = Echo;
func (c *Echo2Core) Echo2Echo(in *echo2.TLEcho2Echo) (*echo2.Echo, error) {
	// TODO: not impl
	// c.Logger.Errorf("echo2.echo blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return echo2.MakeEcho(&echo2.TLEcho{
		Message: in.Message,
	}), nil
}
