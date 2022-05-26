package utils

import (
	"bytes"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/teamgram/teamgram-server/pkg/goffmpeg/models"
)

func DurToSec(dur string) (sec float64) {
	durAry := strings.Split(dur, ":")
	var secs float64
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.ParseFloat(durAry[0], 64)
	secs = hr * (60 * 60)
	min, _ := strconv.ParseFloat(durAry[1], 64)
	secs += min * (60)
	second, _ := strconv.ParseFloat(durAry[2], 64)
	secs += second
	return secs
}

func GetFFmpegExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffmpeg"}

	switch platform {
	case "windows":
		command[0] = "where"
		break
	default:
		command[0] = "which"
		break
	}

	return command
}

func GetFFprobeExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffprobe"}

	switch platform {
	case "windows":
		command[0] = "where"
		break
	default:
		command[0] = "which"
		break
	}
	return command
}

func CheckFileType(streams []models.Streams) string {
	for i := 0; i < len(streams); i++ {
		st := streams[i]
		if st.CodecType == "video" {
			return "video"
		}
	}

	return "audio"
}

func LineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

// TestCmd ...
func TestCmd(command string, args string) (bytes.Buffer, error) {
	var out bytes.Buffer

	cmd := exec.Command(command, args)

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return out, err
	}

	return out, nil
}
