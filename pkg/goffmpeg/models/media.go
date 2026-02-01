package models

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Mediafile struct {
	aspect                string
	resolution            string
	videoBitRate          string
	videoBitRateTolerance int
	videoMaxBitRate       int
	videoMinBitrate       int
	videoCodec            string
	vframes               int
	frameRate             int
	audioRate             int
	maxKeyframe           int
	minKeyframe           int
	keyframeInterval      int
	audioCodec            string
	audioBitrate          string
	audioChannels         int
	audioVariableBitrate  bool
	bufferSize            int
	threadset             bool
	threads               int
	preset                string
	tune                  string
	audioProfile          string
	videoProfile          string
	target                string
	duration              string
	durationInput         string
	seekTime              string
	qscale                uint32
	crf                   uint32
	strict                int
	muxDelay              string
	seekUsingTsInput      bool
	seekTimeInput         string
	inputPath             string
	inputPipe             bool
	inputPipeReader       io.ReadCloser
	inputPipeWriter       io.Writer
	outputPipe            bool
	outputPipeReader      io.Reader
	outputPipeWriter      io.WriteCloser
	movFlags              string
	hideBanner            bool
	outputPath            string
	outputFormat          string
	copyTs                bool
	nativeFramerateInput  bool
	inputInitialOffset    string
	rtmpLive              string
	hlsPlaylistType       string
	hlsListSize           int
	hlsSegmentDuration    int
	hlsMasterPlaylistName string
	hlsSegmentFilename    string
	httpMethod            string
	httpKeepAlive         bool
	hwaccel               string
	streamIds             map[int]string
	metadata              Metadata
	videoFilter           string
	audioFilter           string
	skipVideo             bool
	skipAudio             bool
	compressionLevel      int
	mapMetadata           string
	tags                  map[string]string
	encryptionKey         string
	movflags              string
	bframe                int
	pixFmt                string
	rawInputArgs          []string
	rawOutputArgs         []string
	threadQueueSize       int
	map2                  string
	segmentTime           int
	resetTimestamps       int
	scThreshold           string
	forceKeyFrames        string
}

/*** SETTERS ***/

func (m *Mediafile) SetScThreshold(v string) {
	m.scThreshold = v
}

func (m *Mediafile) SetForceKeyFrames(v string) {
	m.forceKeyFrames = v
}

func (m *Mediafile) SetSegmentTime(v int) {
	m.segmentTime = v
}

func (m *Mediafile) SetResetTimestamps(v int) {
	m.resetTimestamps = v
}

func (m *Mediafile) SetAudioFilter(v string) {
	m.audioFilter = v
}

func (m *Mediafile) SetVideoFilter(v string) {
	m.videoFilter = v
}

// Deprecated: Use SetVideoFilter instead.
func (m *Mediafile) SetFilter(v string) {
	m.SetVideoFilter(v)
}

func (m *Mediafile) SetAspect(v string) {
	m.aspect = v
}

func (m *Mediafile) SetResolution(v string) {
	m.resolution = v
}

func (m *Mediafile) SetVideoBitRate(v string) {
	m.videoBitRate = v
}

func (m *Mediafile) SetVideoBitRateTolerance(v int) {
	m.videoBitRateTolerance = v
}

func (m *Mediafile) SetVideoMaxBitrate(v int) {
	m.videoMaxBitRate = v
}

func (m *Mediafile) SetVideoMinBitRate(v int) {
	m.videoMinBitrate = v
}

func (m *Mediafile) SetVideoCodec(v string) {
	m.videoCodec = v
}

func (m *Mediafile) SetVframes(v int) {
	m.vframes = v
}

func (m *Mediafile) SetFrameRate(v int) {
	m.frameRate = v
}

func (m *Mediafile) SetAudioRate(v int) {
	m.audioRate = v
}

func (m *Mediafile) SetAudioVariableBitrate() {
	m.audioVariableBitrate = true
}

func (m *Mediafile) SetMaxKeyFrame(v int) {
	m.maxKeyframe = v
}

func (m *Mediafile) SetMinKeyFrame(v int) {
	m.minKeyframe = v
}

func (m *Mediafile) SetKeyframeInterval(v int) {
	m.keyframeInterval = v
}

func (m *Mediafile) SetAudioCodec(v string) {
	m.audioCodec = v
}

func (m *Mediafile) SetAudioBitRate(v string) {
	m.audioBitrate = v
}

func (m *Mediafile) SetAudioChannels(v int) {
	m.audioChannels = v
}

func (m *Mediafile) SetPixFmt(v string) {
	m.pixFmt = v
}

func (m *Mediafile) SetBufferSize(v int) {
	m.bufferSize = v
}

func (m *Mediafile) SetThreads(v int) {
	m.threadset = true
	m.threads = v
}

func (m *Mediafile) SetPreset(v string) {
	m.preset = v
}

func (m *Mediafile) SetTune(v string) {
	m.tune = v
}

func (m *Mediafile) SetAudioProfile(v string) {
	m.audioProfile = v
}

func (m *Mediafile) SetVideoProfile(v string) {
	m.videoProfile = v
}

func (m *Mediafile) SetDuration(v string) {
	m.duration = v
}

func (m *Mediafile) SetDurationInput(v string) {
	m.durationInput = v
}

func (m *Mediafile) SetSeekTime(v string) {
	m.seekTime = v
}

func (m *Mediafile) SetSeekTimeInput(v string) {
	m.seekTimeInput = v
}

// Q Scale must be integer between 1 to 31 - https://trac.ffmpeg.org/wiki/Encode/MPEG-4
func (m *Mediafile) SetQScale(v uint32) {
	m.qscale = v
}

func (m *Mediafile) SetCRF(v uint32) {
	m.crf = v
}

func (m *Mediafile) SetStrict(v int) {
	m.strict = v
}

func (m *Mediafile) SetSeekUsingTsInput(val bool) {
	m.seekUsingTsInput = val
}

func (m *Mediafile) SetCopyTs(val bool) {
	m.copyTs = val
}

func (m *Mediafile) SetInputPath(val string) {
	m.inputPath = val
}

func (m *Mediafile) SetInputPipe(val bool) {
	m.inputPipe = val
}

func (m *Mediafile) SetInputPipeReader(r io.ReadCloser) {
	m.inputPipeReader = r
}

func (m *Mediafile) SetInputPipeWriter(w io.Writer) {
	m.inputPipeWriter = w
}

func (m *Mediafile) SetOutputPipe(val bool) {
	m.outputPipe = val
}

func (m *Mediafile) SetOutputPipeReader(r io.Reader) {
	m.outputPipeReader = r
}

func (m *Mediafile) SetOutputPipeWriter(w io.WriteCloser) {
	m.outputPipeWriter = w
}

func (m *Mediafile) SetMovFlags(val string) {
	m.movFlags = val
}

func (m *Mediafile) SetHideBanner(val bool) {
	m.hideBanner = val
}

func (m *Mediafile) SetMuxDelay(val string) {
	m.muxDelay = val
}

func (m *Mediafile) SetOutputPath(val string) {
	m.outputPath = val
}

func (m *Mediafile) SetOutputFormat(val string) {
	m.outputFormat = val
}

func (m *Mediafile) SetNativeFramerateInput(val bool) {
	m.nativeFramerateInput = val
}

func (m *Mediafile) SetRtmpLive(val string) {
	m.rtmpLive = val
}

func (m *Mediafile) SetHlsListSize(val int) {
	m.hlsListSize = val
}

func (m *Mediafile) SetHlsSegmentDuration(val int) {
	m.hlsSegmentDuration = val
}

func (m *Mediafile) SetHlsPlaylistType(val string) {
	m.hlsPlaylistType = val
}

func (m *Mediafile) SetHlsMasterPlaylistName(val string) {
	m.hlsMasterPlaylistName = val
}

func (m *Mediafile) SetHlsSegmentFilename(val string) {
	m.hlsSegmentFilename = val
}

func (m *Mediafile) SetHttpMethod(val string) {
	m.httpMethod = val
}

func (m *Mediafile) SetHttpKeepAlive(val bool) {
	m.httpKeepAlive = val
}

func (m *Mediafile) SetHardwareAcceleration(val string) {
	m.hwaccel = val
}

func (m *Mediafile) SetInputInitialOffset(val string) {
	m.inputInitialOffset = val
}

func (m *Mediafile) SetStreamIds(val map[int]string) {
	m.streamIds = val
}

func (m *Mediafile) SetSkipVideo(val bool) {
	m.skipVideo = val
}

func (m *Mediafile) SetSkipAudio(val bool) {
	m.skipAudio = val
}

func (m *Mediafile) SetMetadata(v Metadata) {
	m.metadata = v
}

func (m *Mediafile) SetCompressionLevel(val int) {
	m.compressionLevel = val
}

func (m *Mediafile) SetMapMetadata(val string) {
	m.mapMetadata = val
}

func (m *Mediafile) SetTags(val map[string]string) {
	m.tags = val
}

func (m *Mediafile) SetBframe(v int) {
	m.bframe = v
}

func (m *Mediafile) SetRawInputArgs(args []string) {
	m.rawInputArgs = args
}

func (m *Mediafile) SetRawOutputArgs(args []string) {
	m.rawOutputArgs = args
}

func (m *Mediafile) SetThreadQueueSize(v int) {
	m.threadQueueSize = v
}

func (m *Mediafile) SetMap2(v string) {
	m.map2 = v
}

/*** GETTERS ***/

// Deprecated: Use VideoFilter instead.
func (m *Mediafile) Filter() string {
	return m.VideoFilter()
}

func (m *Mediafile) VideoFilter() string {
	return m.videoFilter
}

func (m *Mediafile) AudioFilter() string {
	return m.audioFilter
}

func (m *Mediafile) Aspect() string {
	return m.aspect
}

func (m *Mediafile) Resolution() string {
	return m.resolution
}

func (m *Mediafile) VideoBitrate() string {
	return m.videoBitRate
}

func (m *Mediafile) VideoBitRateTolerance() int {
	return m.videoBitRateTolerance
}

func (m *Mediafile) VideoMaxBitRate() int {
	return m.videoMaxBitRate
}

func (m *Mediafile) VideoMinBitRate() int {
	return m.videoMinBitrate
}

func (m *Mediafile) VideoCodec() string {
	return m.videoCodec
}

func (m *Mediafile) Vframes() int {
	return m.vframes
}

func (m *Mediafile) FrameRate() int {
	return m.frameRate
}

func (m *Mediafile) GetPixFmt() string {
	return m.pixFmt
}

func (m *Mediafile) AudioRate() int {
	return m.audioRate
}

func (m *Mediafile) MaxKeyFrame() int {
	return m.maxKeyframe
}

func (m *Mediafile) MinKeyFrame() int {
	return m.minKeyframe
}

func (m *Mediafile) KeyFrameInterval() int {
	return m.keyframeInterval
}

func (m *Mediafile) AudioCodec() string {
	return m.audioCodec
}

func (m *Mediafile) AudioBitrate() string {
	return m.audioBitrate
}

func (m *Mediafile) AudioChannels() int {
	return m.audioChannels
}

func (m *Mediafile) BufferSize() int {
	return m.bufferSize
}

func (m *Mediafile) Threads() int {
	return m.threads
}

func (m *Mediafile) Target() string {
	return m.target
}

func (m *Mediafile) Duration() string {
	return m.duration
}

func (m *Mediafile) DurationInput() string {
	return m.durationInput
}

func (m *Mediafile) SeekTime() string {
	return m.seekTime
}

func (m *Mediafile) Preset() string {
	return m.preset
}

func (m *Mediafile) AudioProfile() string {
	return m.audioProfile
}

func (m *Mediafile) VideoProfile() string {
	return m.videoProfile
}

func (m *Mediafile) Tune() string {
	return m.tune
}

func (m *Mediafile) SeekTimeInput() string {
	return m.seekTimeInput
}

func (m *Mediafile) QScale() uint32 {
	return m.qscale
}

func (m *Mediafile) CRF() uint32 {
	return m.crf
}

func (m *Mediafile) Strict() int {
	return m.strict
}

func (m *Mediafile) MuxDelay() string {
	return m.muxDelay
}

func (m *Mediafile) SeekUsingTsInput() bool {
	return m.seekUsingTsInput
}

func (m *Mediafile) CopyTs() bool {
	return m.copyTs
}

func (m *Mediafile) InputPath() string {
	return m.inputPath
}

func (m *Mediafile) InputPipe() bool {
	return m.inputPipe
}

func (m *Mediafile) InputPipeReader() io.ReadCloser {
	return m.inputPipeReader
}

func (m *Mediafile) InputPipeWriter() io.Writer {
	return m.inputPipeWriter
}

func (m *Mediafile) OutputPipe() bool {
	return m.outputPipe
}

func (m *Mediafile) OutputPipeReader() io.Reader {
	return m.outputPipeReader
}

func (m *Mediafile) OutputPipeWriter() io.WriteCloser {
	return m.outputPipeWriter
}

func (m *Mediafile) MovFlags() string {
	return m.movFlags
}

func (m *Mediafile) HideBanner() bool {
	return m.hideBanner
}

func (m *Mediafile) OutputPath() string {
	return m.outputPath
}

func (m *Mediafile) OutputFormat() string {
	return m.outputFormat
}

func (m *Mediafile) NativeFramerateInput() bool {
	return m.nativeFramerateInput
}

func (m *Mediafile) RtmpLive() string {
	return m.rtmpLive
}

func (m *Mediafile) HlsListSize() int {
	return m.hlsListSize
}

func (m *Mediafile) HlsSegmentDuration() int {
	return m.hlsSegmentDuration
}

func (m *Mediafile) HlsMasterPlaylistName() string {
	return m.hlsMasterPlaylistName
}

func (m *Mediafile) HlsSegmentFilename() string {
	return m.hlsSegmentFilename
}

func (m *Mediafile) HlsPlaylistType() string {
	return m.hlsPlaylistType
}

func (m *Mediafile) InputInitialOffset() string {
	return m.inputInitialOffset
}

func (m *Mediafile) HttpMethod() string {
	return m.httpMethod
}

func (m *Mediafile) HttpKeepAlive() bool {
	return m.httpKeepAlive
}

func (m *Mediafile) HardwareAcceleration() string {
	return m.hwaccel
}

func (m *Mediafile) StreamIds() map[int]string {
	return m.streamIds
}

func (m *Mediafile) SkipVideo() bool {
	return m.skipVideo
}

func (m *Mediafile) SkipAudio() bool {
	return m.skipAudio
}

func (m *Mediafile) Metadata() Metadata {
	return m.metadata
}

func (m *Mediafile) GetMetadata() *Metadata {
	return &m.metadata
}

func (m *Mediafile) CompressionLevel() int {
	return m.compressionLevel
}

func (m *Mediafile) MapMetadata() string {
	return m.mapMetadata
}

func (m *Mediafile) Tags() map[string]string {
	return m.tags
}

func (m *Mediafile) SetEncryptionKey(v string) {
	m.encryptionKey = v
}

func (m *Mediafile) EncryptionKey() string {
	return m.encryptionKey
}

func (m *Mediafile) RawInputArgs() []string {
	return m.rawInputArgs
}

func (m *Mediafile) RawOutputArgs() []string {
	return m.rawOutputArgs
}

func (m *Mediafile) ThreadQueueSize() int {
	return m.threadQueueSize
}

/** OPTS **/
func (m *Mediafile) ToStrCommand() []string {
	var strCommand []string

	opts := []string{
		"ThreadQueueSize",
		"SeekTimeInput",
		"SeekUsingTsInput",
		"NativeFramerateInput",
		"DurationInput",
		"RtmpLive",
		"InputInitialOffset",
		"HardwareAcceleration",
		"RawInputArgs",
		"InputPath",
		"InputPipe",
		"HideBanner",
		"Aspect",
		"Resolution",
		"FrameRate",
		"AudioRate",
		"VideoCodec",
		"Vframes",
		"VideoBitRate",
		"VideoBitRateTolerance",
		"VideoMaxBitRate",
		"VideoMinBitRate",
		"VideoProfile",
		"Map2",
		"SkipVideo",
		"AudioCodec",
		"AudioBitRate",
		"AudioChannels",
		"AudioProfile",
		"SkipAudio",
		"CRF",
		"QScale",
		"Strict",
		"BufferSize",
		"MuxDelay",
		"Threads",
		"KeyframeInterval",
		"Preset",
		"PixFmt",
		"Tune",
		"Target",
		"SeekTime",
		"Duration",
		"CopyTs",
		"StreamIds",
		"MovFlags",
		"RawOutputArgs",
		"SegmentTime",
		"ResetTimestamps",
		"HlsListSize",
		"HlsSegmentDuration",
		"HlsPlaylistType",
		"HlsMasterPlaylistName",
		"HlsSegmentFilename",
		"ScThreshold",
		"ForceKeyFrames",
		"AudioFilter",
		"VideoFilter",
		"HttpMethod",
		"HttpKeepAlive",
		"CompressionLevel",
		"MapMetadata",
		"Tags",
		"EncryptionKey",
		"OutputFormat",
		"OutputPath",
		"Bframe",
		"MovFlags",
		"OutputPipe",
	}

	for _, name := range opts {
		opt := reflect.ValueOf(m).MethodByName(fmt.Sprintf("Obtain%s", name))
		if (opt != reflect.Value{}) {
			result := opt.Call([]reflect.Value{})

			if val, ok := result[0].Interface().([]string); ok {
				strCommand = append(strCommand, val...)
			}
		}
	}

	return strCommand
}

func (m *Mediafile) ObtainAudioFilter() []string {
	if m.audioFilter != "" {
		return []string{"-af", m.audioFilter}
	}
	return nil
}

func (m *Mediafile) ObtainVideoFilter() []string {
	if m.videoFilter != "" {
		return []string{"-vf", m.videoFilter}
	}
	return nil
}

func (m *Mediafile) ObtainAspect() []string {
	// Set aspect
	if m.resolution != "" {
		resolution := strings.Split(m.resolution, "x")
		if len(resolution) != 0 {
			width, _ := strconv.ParseFloat(resolution[0], 64)
			height, _ := strconv.ParseFloat(resolution[1], 64)
			return []string{"-aspect", fmt.Sprintf("%f", width/height)}
		}
	}

	if m.aspect != "" {
		return []string{"-aspect", m.aspect}
	}
	return nil
}

func (m *Mediafile) ObtainHardwareAcceleration() []string {
	if m.hwaccel != "" {
		return []string{"-hwaccel", m.hwaccel}
	}
	return nil
}

func (m *Mediafile) ObtainInputPath() []string {
	if m.inputPath != "" {
		return []string{"-i", m.inputPath}
	}
	return nil
}

func (m *Mediafile) ObtainInputPipe() []string {
	if m.inputPipe {
		return []string{"-i", "pipe:0"}
	}
	return nil
}

func (m *Mediafile) ObtainOutputPipe() []string {
	if m.outputPipe {
		return []string{"pipe:1"}
	}
	return nil
}

func (m *Mediafile) ObtainMovFlags() []string {
	if m.movFlags != "" {
		return []string{"-movflags", m.movFlags}
	}
	return nil
}

func (m *Mediafile) ObtainHideBanner() []string {
	if m.hideBanner {
		return []string{"-hide_banner"}
	}
	return nil
}

func (m *Mediafile) ObtainNativeFramerateInput() []string {
	if m.nativeFramerateInput {
		return []string{"-re"}
	}
	return nil
}

func (m *Mediafile) ObtainOutputPath() []string {
	if m.outputPath != "" {
		return []string{m.outputPath}
	}
	return nil
}

func (m *Mediafile) ObtainVideoCodec() []string {
	if m.videoCodec != "" {
		return []string{"-c:v", m.videoCodec}
	}
	return nil
}

func (m *Mediafile) ObtainVframes() []string {
	if m.vframes != 0 {
		return []string{"-vframes", fmt.Sprintf("%d", m.vframes)}
	}
	return nil
}

func (m *Mediafile) ObtainFrameRate() []string {
	if m.frameRate != 0 {
		return []string{"-r", fmt.Sprintf("%d", m.frameRate)}
	}
	return nil
}

func (m *Mediafile) ObtainAudioRate() []string {
	if m.audioRate != 0 {
		return []string{"-ar", fmt.Sprintf("%d", m.audioRate)}
	}
	return nil
}

func (m *Mediafile) ObtainResolution() []string {
	if m.resolution != "" {
		return []string{"-s", m.resolution}
	}
	return nil
}

func (m *Mediafile) ObtainVideoBitRate() []string {
	if m.videoBitRate != "" {
		return []string{"-b:v", m.videoBitRate}
	}
	return nil
}

func (m *Mediafile) ObtainAudioCodec() []string {
	if m.audioCodec != "" {
		return []string{"-c:a", m.audioCodec}
	}
	return nil
}

func (m *Mediafile) ObtainAudioBitRate() []string {
	switch {
	case !m.audioVariableBitrate && m.audioBitrate != "":
		return []string{"-b:a", m.audioBitrate}
	case m.audioVariableBitrate && m.audioBitrate != "":
		return []string{"-q:a", m.audioBitrate}
	case m.audioVariableBitrate:
		return []string{"-q:a", "0"}
	default:
		return nil
	}
}

func (m *Mediafile) ObtainAudioChannels() []string {
	if m.audioChannels != 0 {
		return []string{"-ac", fmt.Sprintf("%d", m.audioChannels)}
	}
	return nil
}

func (m *Mediafile) ObtainVideoMaxBitRate() []string {
	if m.videoMaxBitRate != 0 {
		return []string{"-maxrate", fmt.Sprintf("%dk", m.videoMaxBitRate)}
	}
	return nil
}

func (m *Mediafile) ObtainVideoMinBitRate() []string {
	if m.videoMinBitrate != 0 {
		return []string{"-minrate", fmt.Sprintf("%dk", m.videoMinBitrate)}
	}
	return nil
}

func (m *Mediafile) ObtainBufferSize() []string {
	if m.bufferSize != 0 {
		return []string{"-bufsize", fmt.Sprintf("%dk", m.bufferSize)}
	}
	return nil
}

func (m *Mediafile) ObtainVideoBitRateTolerance() []string {
	if m.videoBitRateTolerance != 0 {
		return []string{"-bt", fmt.Sprintf("%dk", m.videoBitRateTolerance)}
	}
	return nil
}

func (m *Mediafile) ObtainThreads() []string {
	if m.threadset {
		return []string{"-threads", fmt.Sprintf("%d", m.threads)}
	}
	return nil
}

func (m *Mediafile) ObtainTarget() []string {
	if m.target != "" {
		return []string{"-target", m.target}
	}
	return nil
}

func (m *Mediafile) ObtainDuration() []string {
	if m.duration != "" {
		return []string{"-t", m.duration}
	}
	return nil
}

func (m *Mediafile) ObtainDurationInput() []string {
	if m.durationInput != "" {
		return []string{"-t", m.durationInput}
	}
	return nil
}

func (m *Mediafile) ObtainKeyframeInterval() []string {
	if m.keyframeInterval != 0 {
		return []string{"-g", fmt.Sprintf("%d", m.keyframeInterval)}
	}
	return nil
}

func (m *Mediafile) ObtainSeekTime() []string {
	if m.seekTime != "" {
		return []string{"-ss", m.seekTime}
	}
	return nil
}

func (m *Mediafile) ObtainSeekTimeInput() []string {
	if m.seekTimeInput != "" {
		return []string{"-ss", m.seekTimeInput}
	}
	return nil
}

func (m *Mediafile) ObtainPreset() []string {
	if m.preset != "" {
		return []string{"-preset", m.preset}
	}
	return nil
}

func (m *Mediafile) ObtainTune() []string {
	if m.tune != "" {
		return []string{"-tune", m.tune}
	}
	return nil
}

func (m *Mediafile) ObtainCRF() []string {
	if m.crf != 0 {
		return []string{"-crf", fmt.Sprintf("%d", m.crf)}
	}
	return nil
}

func (m *Mediafile) ObtainQScale() []string {
	if m.qscale != 0 {
		return []string{"-qscale", fmt.Sprintf("%d", m.qscale)}
	}
	return nil
}

func (m *Mediafile) ObtainStrict() []string {
	if m.strict != 0 {
		return []string{"-strict", fmt.Sprintf("%d", m.strict)}
	}
	return nil
}

func (m *Mediafile) ObtainVideoProfile() []string {
	if m.videoProfile != "" {
		return []string{"-profile:v", m.videoProfile}
	}
	return nil
}

func (m *Mediafile) ObtainAudioProfile() []string {
	if m.audioProfile != "" {
		return []string{"-profile:a", m.audioProfile}
	}
	return nil
}

func (m *Mediafile) ObtainCopyTs() []string {
	if m.copyTs {
		return []string{"-copyts"}
	}
	return nil
}

func (m *Mediafile) ObtainOutputFormat() []string {
	if m.outputFormat != "" {
		return []string{"-f", m.outputFormat}
	}
	return nil
}

func (m *Mediafile) ObtainMuxDelay() []string {
	if m.muxDelay != "" {
		return []string{"-muxdelay", m.muxDelay}
	}
	return nil
}

func (m *Mediafile) ObtainSeekUsingTsInput() []string {
	if m.seekUsingTsInput {
		return []string{"-seek_timestamp", "1"}
	}
	return nil
}

func (m *Mediafile) ObtainRtmpLive() []string {
	if m.rtmpLive != "" {
		return []string{"-rtmp_live", m.rtmpLive}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHlsPlaylistType() []string {
	if m.hlsPlaylistType != "" {
		return []string{"-hls_playlist_type", m.hlsPlaylistType}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainInputInitialOffset() []string {
	if m.inputInitialOffset != "" {
		return []string{"-itsoffset", m.inputInitialOffset}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHlsListSize() []string {
	return []string{"-hls_list_size", fmt.Sprintf("%d", m.hlsListSize)}
}

func (m *Mediafile) ObtainHlsSegmentDuration() []string {
	if m.hlsSegmentDuration != 0 {
		return []string{"-hls_time", fmt.Sprintf("%d", m.hlsSegmentDuration)}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHlsMasterPlaylistName() []string {
	if m.hlsMasterPlaylistName != "" {
		return []string{"-master_pl_name", m.hlsMasterPlaylistName}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHlsSegmentFilename() []string {
	if m.hlsSegmentFilename != "" {
		return []string{"-hls_segment_filename", m.hlsSegmentFilename}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHttpMethod() []string {
	if m.httpMethod != "" {
		return []string{"-method", m.httpMethod}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainPixFmt() []string {
	if m.pixFmt != "" {
		return []string{"-pix_fmt", m.pixFmt}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainHttpKeepAlive() []string {
	if m.httpKeepAlive {
		return []string{"-multiple_requests", "1"}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainSkipVideo() []string {
	if m.skipVideo {
		return []string{"-vn"}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainSkipAudio() []string {
	if m.skipAudio {
		return []string{"-an"}
	} else {
		return nil
	}
}

func (m *Mediafile) ObtainStreamIds() []string {
	if m.streamIds != nil && len(m.streamIds) != 0 {
		result := []string{}
		for i, val := range m.streamIds {
			result = append(result, []string{"-streamid", fmt.Sprintf("%d:%s", i, val)}...)
		}
		return result
	}
	return nil
}

func (m *Mediafile) ObtainCompressionLevel() []string {
	if m.compressionLevel != 0 {
		return []string{"-compression_level", fmt.Sprintf("%d", m.compressionLevel)}
	}
	return nil
}

func (m *Mediafile) ObtainMapMetadata() []string {
	if m.mapMetadata != "" {
		return []string{"-map_metadata", m.mapMetadata}
	}
	return nil
}

func (m *Mediafile) ObtainEncryptionKey() []string {
	if m.encryptionKey != "" {
		return []string{"-hls_key_info_file", m.encryptionKey}
	}

	return nil
}

func (m *Mediafile) ObtainBframe() []string {
	if m.bframe != 0 {
		return []string{"-bf", fmt.Sprintf("%d", m.bframe)}
	}
	return nil
}

func (m *Mediafile) ObtainTags() []string {
	if m.tags != nil && len(m.tags) != 0 {
		result := []string{}
		for key, val := range m.tags {
			result = append(result, []string{"-metadata", fmt.Sprintf("%s=%s", key, val)}...)
		}
		return result
	}
	return nil
}

func (m *Mediafile) ObtainRawInputArgs() []string {
	return m.rawInputArgs
}

func (m *Mediafile) ObtainRawOutputArgs() []string {
	return m.rawOutputArgs
}

func (m *Mediafile) ObtainThreadQueueSize() []string {
	if m.threadQueueSize != 0 {
		return []string{"-thread_queue_size", fmt.Sprintf("%d", m.threadQueueSize)}
	}
	return nil
}

func (m *Mediafile) ObtainMap2() []string {
	if m.map2 != "" {
		return []string{"-map", m.map2}
	}
	return nil
}

func (m *Mediafile) ObtainSegmentTime() []string {
	if m.segmentTime != 0 {
		return []string{"-segment_time", fmt.Sprintf("%d", m.segmentTime)}
	}
	return nil
}

func (m *Mediafile) ObtainResetTimestamps() []string {
	if m.resetTimestamps != 0 {
		return []string{"-reset_timestamps", fmt.Sprintf("%d", m.resetTimestamps)}
	}
	return nil
}

func (m *Mediafile) ObtainScThreshold() []string {
	if m.scThreshold != "" {
		return []string{"-sc_threshold", m.scThreshold}
	}
	return nil
}

func (m *Mediafile) ObtainForceKeyFrames() []string {
	if m.forceKeyFrames != "" {
		return []string{"-force_key_frames", m.forceKeyFrames}
	}
	return nil
}
