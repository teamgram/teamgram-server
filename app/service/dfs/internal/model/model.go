// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package model

const (
	PhotoSZOriginalSize = 0 // client upload original photo
	PhotoSZSmallSize    = 90
	PhotoSZMediumSize   = 320
	PhotoSZXLargeSize   = 800
	PhotoSZYLargeSize   = 1280
	PhotoSZASize        = 160
	PhotoSZBSize        = 320
	PhotoSZCSize        = 640
	PhotoSZStrippedSize = 40
	PhotoSZUSize        = 320
	PhotoSZVSize        = 800
)

const (
	PhotoSZOriginalType = "0" // client upload original photo
	PhotoSZSmallType    = "s"
	PhotoSZMediumType   = "m"
	PhotoSZXLargeType   = "x"
	PhotoSZYLargeType   = "y"
	PhotoSZAType        = "a"
	PhotoSZBType        = "b"
	PhotoSZCType        = "c"
	PhotoSZStrippedType = "i"
	VideoSZUType        = "u"
	VideoSZVType        = "v"
)

const (
	PhotoSZOriginalLocalId = 0 // client upload original photo
	PhotoSZSmallLocalId    = 1
	PhotoSZMediumLocalId   = 2
	PhotoSZXLargeLocalId   = 3
	PhotoSZYLargeLocalId   = 4
	PhotoSZALocalId        = 5
	PhotoSZBLocalId        = 6
	PhotoSZCLocalId        = 7
	PhotoSZStrippedLocalId = 8
	VideoSZULocalId        = 9
	VideoSZVLocalId        = 10
)

func PhotoSizeIsVideo(t string) bool {
	switch t {
	case VideoSZUType:
		return true
	case VideoSZVType:
		return true
	}
	return false
}

func LocalIdIsVideo(localId int) bool {
	switch localId {
	case VideoSZULocalId:
		return true
	case VideoSZVLocalId:
		return true
	}
	return false
}

func GetSizeTypeByLocalId(localId int) string {
	switch localId {
	case PhotoSZOriginalLocalId:
		return PhotoSZOriginalType
	case PhotoSZSmallLocalId:
		return PhotoSZSmallType
	case PhotoSZMediumLocalId:
		return PhotoSZMediumType
	case PhotoSZXLargeLocalId:
		return PhotoSZXLargeType
	case PhotoSZYLargeLocalId:
		return PhotoSZYLargeType
	case PhotoSZALocalId:
		return PhotoSZAType
	case PhotoSZBLocalId:
		return PhotoSZBType
	case PhotoSZCLocalId:
		return PhotoSZCType
	case PhotoSZStrippedLocalId:
		return PhotoSZStrippedType
	case VideoSZULocalId:
		return VideoSZUType
	case VideoSZVLocalId:
		return VideoSZVType
	}

	return ""
}

func GetSizeLocalIdByType(t string) int {
	switch t {
	case PhotoSZOriginalType:
		return PhotoSZOriginalLocalId
	case PhotoSZSmallType:
		return PhotoSZSmallLocalId
	case PhotoSZMediumType:
		return PhotoSZMediumLocalId
	case PhotoSZXLargeType:
		return PhotoSZXLargeLocalId
	case PhotoSZYLargeType:
		return PhotoSZYLargeLocalId
	case PhotoSZAType:
		return PhotoSZALocalId
	case PhotoSZBType:
		return PhotoSZBLocalId
	case PhotoSZCType:
		return PhotoSZCLocalId
	case PhotoSZStrippedType:
		return PhotoSZStrippedLocalId
	case VideoSZUType:
		return VideoSZULocalId
	case VideoSZVType:
		return VideoSZVLocalId
	}

	return -1
}

type ReSizeInfo struct {
	LocalId int
	Type    string
	Size    int
}

var (
	ReSizeInfoABCList = []ReSizeInfo{
		{PhotoSZALocalId, PhotoSZAType, PhotoSZASize},
		{PhotoSZBLocalId, PhotoSZBType, PhotoSZBSize},
		{PhotoSZCLocalId, PhotoSZCType, PhotoSZCSize},
	}

	ReSizeInfoPhotoList = []ReSizeInfo{
		{PhotoSZSmallLocalId, PhotoSZSmallType, PhotoSZSmallSize},
		{PhotoSZMediumLocalId, PhotoSZMediumType, PhotoSZMediumSize},
		{PhotoSZXLargeLocalId, PhotoSZXLargeType, PhotoSZXLargeSize},
		{PhotoSZYLargeLocalId, PhotoSZYLargeType, PhotoSZYLargeSize},
	}
)

func GetMaxResizeInfo(isABC bool, w, h int) *ReSizeInfo {
	var (
		isW    = true
		szList []ReSizeInfo
		sz     ReSizeInfo
	)

	if w < h {
		isW = false
	}

	if isABC {
		szList = ReSizeInfoABCList
	} else {
		szList = ReSizeInfoPhotoList
	}

	for _, sz = range szList {
		if isW {
			if w <= sz.Size {
				break
			}
		} else {
			if h <= sz.Size {
				break
			}
		}
	}

	return &sz
}
