package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func makeDefaultAutoDownloadSettings() *tg.AccountAutoDownloadSettings {
	return tg.MakeTLAccountAutoDownloadSettings(&tg.TLAccountAutoDownloadSettings{
		Low: makeAutoDownloadSettings(true, 512000, 512000),
		Medium: makeAutoDownloadSettings(
			false,
			10485760,
			1048576,
		),
		High: makeAutoDownloadSettings(
			false,
			15728640,
			3145728,
		),
	}).ToAccountAutoDownloadSettings()
}

func makeAutoDownloadSettings(phonecallsLessData bool, videoSizeMax int64, fileSizeMax int64) *tg.AutoDownloadSettings {
	return tg.MakeTLAutoDownloadSettings(&tg.TLAutoDownloadSettings{
		Disabled:           false,
		VideoPreloadLarge:  true,
		AudioPreloadNext:   true,
		PhonecallsLessData: phonecallsLessData,
		PhotoSizeMax:       1048576,
		VideoSizeMax:       videoSizeMax,
		FileSizeMax:        fileSizeMax,
	}).ToAutoDownloadSettings()
}
