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

package ffmpegutil

import (
	"io/ioutil"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestConvertGifToMp4(t *testing.T) {
	ffmpegUtil := NewFFmpegUtil()
	data, duration, err := ffmpegUtil.ConvertToMp4ByPipe("./gsmarena_003.gif", -1, 320)
	if err != nil {
		logx.Errorf("%v", err)
		return
	}
	logx.Infof("duration = %d", duration)
	ioutil.WriteFile("./gsmarena_003.gif.mp4", data, 0644)

	oData, err := ffmpegUtil.GetFirstFrameByPipe(data)
	if err != nil {
		logx.Errorf("%v", err)
		return
	}

	ioutil.WriteFile("./gsmarena_003.gif.jpg", oData, 0644)

	md, _ := ffmpegUtil.GetMetadataByPipe(data)
	if md != nil {
		logx.Infof("%#v", md)
		w, h := GetWidthHeight(md)
		logx.Infof("duration: %d, (w, h): (%d,%d)", GetDuration(md), w, h)
	}
}
