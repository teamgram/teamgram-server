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
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"sync"

	"github.com/teamgram/teamgram-server/pkg/goffmpeg/models"
	"github.com/teamgram/teamgram-server/pkg/goffmpeg/transcoder"
	"github.com/teamgram/teamgram-server/pkg/goffmpeg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type FFmpegUtil struct {
	*transcoder.Transcoder
}

func NewFFmpegUtil() *FFmpegUtil {
	return &FFmpegUtil{
		Transcoder: new(transcoder.Transcoder),
	}
}

/*
	// ffmpeg -i safe_image.gif -f mp4 -movflags +faststart -pix_fmt yuv420p -vf scale=320:-1 -c:v libx264 -strict experimental -b:v 218k -bufsize 218k safe_image.gif.mp4
	// os.Setenv("PATH", "/usr/bin:/sbin:/usr/local/Cellar/ffmpeg/4.2.1_2/bin")
	o, err := exec.Command(
		"ffmpeg",
		"-f", "gif",
		"-i", gifPath,
		"-movflags", "+faststart",
		"-pix_fmt", "yuv420p",
		// "-s", "320x268",
		"-vf", "scale=320:-1",
		"-c:v", "libx264",
		// "-profile:v", "baseline",
		// "-x264opts", "cabac=0:bframes=0:ref=1:weightp=0:level=30:bitrate=700:vbv_maxrate=768:vbv_bufsize=1400",
		// "-pass", "1",
		"-strict", "experimental",
		"-b:v", "218k",
		"-bufsize", "218k",
		videoPath,
	).CombinedOutput()
*/

func (trans *FFmpegUtil) ConvertToMp4(gifPath string) (err error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)
	videoPath := gifPath + ".mp4"

	// Initialize transcoder passing the input file path and output file path
	err = trans.Initialize(gifPath, videoPath)
	if err != nil {
		logx.Errorf("convertGifToMp4 - error: %v", err)
		return
	}
	// Handle error...

	// "-f gif"
	trans.MediaFile().SetOutputFormat("mp4")
	trans.MediaFile().SetMovFlags("+faststart")
	trans.MediaFile().SetPixFmt("yuv420p")

	w, h := GetWidthHeight(trans.MediaFile().GetMetadata())
	if w >= h {
		trans.MediaFile().SetVideoFilter("scale=320:-2")
	} else {
		trans.MediaFile().SetVideoFilter("scale=-2:320")
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetStrict(-2)
	trans.MediaFile().SetVideoBitRate("218k")
	trans.MediaFile().SetBufferSize(218)

	// Start transcoder process without checking progress
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}

	// This channel is used to wait for the process to end
	err = <-done
	if err != nil {
		logx.Errorf("convertGifToMp4 - error: %v", err)
	}

	//data, err = ioutil.ReadFile(videoPath)
	//if err != nil {
	//	log.Errorf("convertGifToMp4 - error: %v", err)
	//}
	return
}

func (trans *FFmpegUtil) ConvertToMp4ByPipe(gifPath string, dstW, dstH int) (bytes []byte, duration int32, err error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err = trans.InitializeEmptyTranscoder()
	if err != nil {
		logx.Errorf("InitializeEmptyTranscoder - error: %v")
		return
	}

	// ffmpeg -i safe_image.gif -f mp4 -movflags +faststart -pix_fmt yuv420p -vf scale=320:-1 -c:v libx264 -strict experimental -b:v 218k -bufsize 218k safe_image.gif.mp4
	trans.MediaFile().SetInputPath(gifPath)

	trans.MediaFile().SetOutputFormat("mp4")
	// trans.MediaFile().SetMovFlags("+faststart")
	trans.MediaFile().SetPixFmt("yuv420p")
	trans.MediaFile().SetVideoFilter(fmt.Sprintf("scale=%d:%d", dstW, dstH))
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetStrict(-2)
	// trans.MediaFile().SetVideoBitRate("218k")
	// trans.MediaFile().SetBufferSize(218)
	// trans.MediaFile().SetFrameRate(30)

	r, err := trans.CreateOutputPipe("mp4")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer r.Close()
		defer wg.Done()

		// Read data from output pipe
		bytes, err = ioutil.ReadAll(r)
		// Handle error and data...
		if err != nil {
			logx.Errorf("readPipe error: %v", err)
			return
		}
	}()

	// Start transcoder process without checking progress
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()
	// Example of printing transcoding progress
	for msg := range progress {
		duration = int32(math.Round(utils.DurToSec(msg.CurrentTime)))
		fmt.Println(msg)
	}

	// This channel is used to wait for the transcoding process to end
	err = <-done
	// Handle error...
	if err != nil {
		logx.Errorf("transcoding error: %v", err)
		return
	}

	wg.Wait()

	return
}

// GetFirstFrame
// cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
func (trans *FFmpegUtil) GetFirstFrame(iFilePath string) (bytes []byte, err error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err = trans.InitializeEmptyTranscoder()
	if err != nil {
		logx.Errorf("InitializeEmptyTranscoder - error: %v", err)
		return
	}

	trans.MediaFile().SetInputPath(iFilePath)

	trans.MediaFile().SetOutputFormat("image2")
	trans.MediaFile().SetVframes(1)

	r, err := trans.CreateOutputPipe("image2")
	if err != nil {
		logx.Errorf("createOutputPipe error: %v", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer r.Close()
		defer wg.Done()

		// Read data from output pipe
		bytes, err = ioutil.ReadAll(r)
		// log.Debugf("data: %d", data)
		// Handle error and data...
		if err != nil {
			if errors.Is(err, io.ErrClosedPipe) {
				//
			} else {
				logx.Errorf("readPipe error: %v", err)
			}
		}
	}()

	// Start transcoder process without checking progress
	done := trans.Run(true)

	// This channel is used to wait for the transcoding process to end
	err = <-done
	// Handle error...
	if err != nil {
		if errors.Is(err, io.ErrClosedPipe) {
			//
		} else {
			logx.Errorf("getFirstFrameByPipe error: %v", err)
			return
		}
	}

	wg.Wait()

	err = nil
	return
}

func (trans *FFmpegUtil) GetFirstFrameByPipe(iPipeData []byte) (bytes []byte, err error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err = trans.InitializeEmptyTranscoder()
	if err != nil {
		logx.Errorf("InitializeEmptyTranscoder - error: %v")
		return
	}

	trans.MediaFile().SetOutputFormat("image2")
	trans.MediaFile().SetVframes(1)

	// Create an input pipe to write to, which will return *io.PipeWriter
	w, err := trans.CreateInputPipe()
	if err != nil {
		logx.Errorf("createInputPipe error: %v", err)
		return
	}
	// log.Println(err)

	r, err := trans.CreateOutputPipe("image2")
	if err != nil {
		logx.Errorf("createOutputPipe error: %v", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer r.Close()
		defer wg.Done()

		// Read data from output pipe
		bytes, err = ioutil.ReadAll(r)
		// log.Debugf("data: %d", data)
		// Handle error and data...
		if err != nil {
			if errors.Is(err, io.ErrClosedPipe) {
				//
			} else {
				logx.Errorf("readPipe error: %v", err)
			}
		}
	}()

	go func() {
		var (
			n int
		)

		defer w.Close()
		n, err = w.Write(iPipeData)
		_ = n
		// log.Debugf("data: %d", data)
		if err != nil {
			if errors.Is(err, io.ErrClosedPipe) {
				//
			} else {
				logx.Errorf("writePipe error: %v", err)
				return
			}
		}

	}()

	// Start transcoder process without checking progress
	done := trans.Run(true)

	// This channel is used to wait for the transcoding process to end
	err = <-done
	// Handle error...
	if err != nil {
		if errors.Is(err, io.ErrClosedPipe) {
			//
		} else {
			logx.Errorf("getFirstFrameByPipe error: %v", err)
			return
		}
	}

	wg.Wait()

	err = nil
	return
}

func (trans *FFmpegUtil) GetMetadataByPipe(iPipeData []byte) (*models.Metadata, error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)
	if err := trans.Transcoder.GetMetadataByPipe(bytes.NewBuffer(iPipeData)); err != nil {
		return nil, err
	}
	// fmt.Println(trans.MediaFile().Metadata())
	return trans.MediaFile().GetMetadata(), nil
}

func (trans *FFmpegUtil) GetMetadata(iFilePath string) (*models.Metadata, error) {
	// Create new instance of transcoder
	// trans := new(transcoder.Transcoder)
	if err := trans.Transcoder.GetMetadata(iFilePath); err != nil {
		return nil, err
	}
	// fmt.Println(trans.MediaFile().Metadata())
	return trans.MediaFile().GetMetadata(), nil
}

func GetDuration(md *models.Metadata) int {
	return int(math.Round(utils.DurToSec(md.Format.Duration)))
}

func GetWidthHeight(md *models.Metadata) (w, h int) {
	for i := 0; i < len(md.Streams); i++ {
		if md.Streams[i].CodecType == "video" {
			w = md.Streams[i].Width
			h = md.Streams[i].Height
			break
		}
	}
	return
}

func (trans *FFmpegUtil) GetDuration() int {
	return int(math.Round(utils.DurToSec(trans.MediaFile().Metadata().Format.Duration)))
}

func (trans *FFmpegUtil) GetWidth() int {
	streams := trans.MediaFile().Metadata().Streams
	for i := 0; i < len(streams); i++ {
		if streams[i].CodecType == "video" {
			return streams[i].Width
		}
	}
	return 0
}

func (trans *FFmpegUtil) GetHeight() int {
	streams := trans.MediaFile().Metadata().Streams
	for i := 0; i < len(streams); i++ {
		if streams[i].CodecType == "video" {
			return streams[i].Height
		}
	}
	return 0
}
