package ffmpeg

import (
	"bytes"
	"strings"

	"github.com/teamgram/teamgram-server/pkg/goffmpeg/utils"
)

// Configuration ...
type Configuration struct {
	FfmpegBin  string
	FfprobeBin string
}

// Configure Get and set FFmpeg and FFprobe bin paths
func Configure() (Configuration, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer

	execFFmpegCommand := utils.GetFFmpegExec()
	execFFprobeCommand := utils.GetFFprobeExec()

	outFFmpeg, err := utils.TestCmd(execFFmpegCommand[0], execFFmpegCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	outProbe, err = utils.TestCmd(execFFprobeCommand[0], execFFprobeCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	ffmpeg := strings.Replace(strings.Split(outFFmpeg.String(), "\n")[0], utils.LineSeparator(), "", -1)
	ffprobe := strings.Replace(strings.Split(outProbe.String(), "\n")[0], utils.LineSeparator(), "", -1)

	cnf := Configuration{ffmpeg, ffprobe}
	return cnf, nil
}
