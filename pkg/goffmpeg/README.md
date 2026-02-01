# Goffmpeg
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/93e018e5008b4439acbb30d715b22e7f)](https://www.codacy.com/app/francisco.romero/goffmpeg?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=xfrr/goffmpeg&amp;utm_campaign=Badge_Grade)

FFMPEG wrapper written in GO which allows to obtain the progress.

## V2
New implementation with an easy-to-use API and interfaces to extend the transcoding capabilities.
> https://github.com/floostack/transcoder

# Dependencies
- [FFmpeg](https://www.ffmpeg.org/)
- [FFProbe](https://www.ffmpeg.org/ffprobe.html)

# Supported platforms

 - Linux
 - OS X
 - Windows

# Getting started
## How to transcode a media file
```shell
go get github.com/xfrr/goffmpeg
```

```go
package main

import (
    "github.com/xfrr/goffmpeg/transcoder"
)

var inputPath = "/data/testmov"
var outputPath = "/data/testmp4.mp4"

func main() {

	// Create new instance of transcoder
    trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
    err := trans.Initialize( inputPath, outputPath )
    // Handle error...

	// Start transcoder process without checking progress
	done := trans.Run(false)

	// This channel is used to wait for the process to end
	err = <-done
	// Handle error...

}
```
## How to get the transcoding progress
```go
...
func main() {

	// Create new instance of transcoder
    trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
    err := trans.Initialize( inputPath, outputPath )
    // Handle error...

	// Start transcoder process with progress checking
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}

	// This channel is used to wait for the transcoding process to end
	err = <-done

}
```

## How to pipe in data using the [pipe protocol](https://ffmpeg.org/ffmpeg-protocols.html#pipe)
Creating an input pipe will return [\*io.PipeReader](https://golang.org/pkg/io/#PipeReader), and creating an output pipe will return [\*io.PipeWriter](https://golang.org/pkg/io/#PipeWriter). An example is shown which uses `cat` to pipe in data, and [ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) to read data as bytes from the pipe.
```go
func main() {

	// Create new instance of transcoder
    trans := new(transcoder.Transcoder)

	// Initialize an empty transcoder
    err := trans.InitializeEmptyTranscoder()
    // Handle error...

	// Create a command such that its output should be passed as stdin to ffmpeg
	cmd := exec.Command("cat", "/path/to/file")

	// Create an input pipe to write to, which will return *io.PipeWriter
	w, err := trans.CreateInputPipe()

	cmd.Stdout = w

	// Create an output pipe to read from, which will return *io.PipeReader.
	// Must also specify the output container format
	r, err := trans.CreateOutputPipe("mp4")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer r.Close()
		defer wg.Done()

		// Read data from output pipe
		data, err := ioutil.ReadAll(r)
		// Handle error and data...
	}()

	go func() {
		defer w.Close()
		err := cmd.Run()
		// Handle error...
	}()

	// Start transcoder process without checking progress
	done := trans.Run(false)

	// This channel is used to wait for the transcoding process to end
	err = <-done
	// Handle error...

	wg.Wait()
}
```

# Progress properties
```go
type Progress struct {
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
}
```
# Media setters
Those options can be set before starting the transcoding.
```js
SetAspect
SetResolution
SetVideoBitRate
SetVideoBitRateTolerance
SetVideoMaxBitrate
SetVideoMinBitRate
SetVideoCodec
SetVframes
SetFrameRate
SetAudioRate
SetSkipAudio
SetSkipVideo
SetMaxKeyFrame
SetMinKeyFrame
SetKeyframeInterval
SetAudioCodec
SetAudioBitRate
SetAudioChannels
SetBufferSize
SetThreads
SetPreset
SetTune
SetAudioProfile
SetVideoProfile
SetDuration
SetDurationInput
SetSeekTime
SetSeekTimeInput
SetSeekUsingTsInput
SetQuality
SetStrict
SetCopyTs
SetMuxDelay
SetHideBanner
SetInputPath
SetNativeFramerateInput
SetRtmpLive
SetHlsListSize
SetHlsSegmentDuration
SetHlsPlaylistType
SetHlsMasterPlaylistName
SetHlsSegmentFilename
SetHttpMethod
SetHttpKeepAlive
SetOutputPath
SetOutputFormat
SetAudioFilter
SetAudioVariableBitrate
SetCompressionLevel
SetFilter
SetInputInitialOffset
SetInputPipeCommand
SetMapMetadata
SetMetadata
SetStreamIds
SetTags
SetVideoFilter
```
Example
```golang
func main() {

	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( inputPath, outputPath )
	// Handle error...

	// SET FRAME RATE TO MEDIAFILE
	trans.MediaFile().SetFrameRate(70)
	// SET ULTRAFAST PRESET TO MEDIAFILE
	trans.MediaFile().SetPreset("ultrafast")

	// Start transcoder process to check progress
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}

	// This channel is used to wait for the transcoding process to end
	err = <-done

}
```

Example with AES encryption :

More information about [HLS encryption with FFMPEG](https://hlsbook.net/how-to-encrypt-hls-video-with-ffmpeg/)

```bash
# Generate key
openssl rand 16 > enc.key
```

Create key file info :

```enc.keyinfo
Key URI
Path to key file
```

```golang
func main() {

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)

	trans.MediaFile().SetVideoCodec("libx264")
	
	trans.MediaFile().SetHlsSegmentDuration(4)

	trans.MediaFile().SetEncryptionKey(keyinfoPath)

	progress := trans.Output()
	
	err = <-done
}
```
