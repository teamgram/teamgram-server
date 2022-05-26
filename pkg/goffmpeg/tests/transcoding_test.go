package test

import (
	"io/ioutil"
	"os/exec"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/teamgram/teamgram-server/pkg/goffmpeg/transcoder"
)

func TestInputNotFound(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/nf"
	var outputPath = "/tmp/ffmpeg/out/nf.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.NotNil(t, err)
}

func TestTranscoding3GP(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/3gp"
	var outputPath = "/tmp/ffmpeg/out/3gp.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingAVI(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/avi"
	var outputPath = "/tmp/ffmpeg/out/avi.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingFLV(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/flv"
	var outputPath = "/tmp/ffmpeg/out/flv.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingMKV(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/mkv"
	var outputPath = "/tmp/ffmpeg/out/mkv.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingMOV(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/mov"
	var outputPath = "/tmp/ffmpeg/out/mov.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingMPEG(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/mpeg"
	var outputPath = "/tmp/ffmpeg/out/mpeg.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingOGG(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/ogg"
	var outputPath = "/tmp/ffmpeg/out/ogg.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingWAV(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/wav"
	var outputPath = "/tmp/ffmpeg/out/wav.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingWEBM(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/webm"
	var outputPath = "/tmp/ffmpeg/out/webm.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingWMV(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/wmv"
	var outputPath = "/tmp/ffmpeg/out/wmv.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)
}

func TestTranscodingProgress(t *testing.T) {

	var inputPath = "/tmp/ffmpeg/avi"
	var outputPath = "/tmp/ffmpeg/out/avi.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	assert.Nil(t, err)

	done := trans.Run(true)
	for val := range trans.Output() {
		if &val != nil {
			break
		}
	}

	err = <-done
	assert.Nil(t, err)
}

func TestTranscodePipes(t *testing.T) {
	c1 := exec.Command("cat", "/tmp/ffmpeg/mkv")

	trans := new(transcoder.Transcoder)

	err := trans.InitializeEmptyTranscoder()
	assert.Nil(t, err)

	w, err := trans.CreateInputPipe()
	assert.Nil(t, err)
	c1.Stdout = w

	r, err := trans.CreateOutputPipe("mp4")
	assert.Nil(t, err)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_, err := ioutil.ReadAll(r)
		assert.Nil(t, err)

		r.Close()
		wg.Done()
	}()

	go func() {
		err := c1.Run()
		assert.Nil(t, err)
		w.Close()
	}()
	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)

	wg.Wait()
}
