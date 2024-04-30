/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package media

import (
	"reflect"

	"github.com/teamgram/proto/mtproto"
)

var _ *mtproto.Bool

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLMediaUploadPhotoFile":        RPCContextTuple{"/mtproto.RPCMedia/media_uploadPhotoFile", func() interface{} { return new(mtproto.Photo) }},
	"TLMediaUploadProfilePhotoFile": RPCContextTuple{"/mtproto.RPCMedia/media_uploadProfilePhotoFile", func() interface{} { return new(mtproto.Photo) }},
	"TLMediaGetPhoto":               RPCContextTuple{"/mtproto.RPCMedia/media_getPhoto", func() interface{} { return new(mtproto.Photo) }},
	"TLMediaGetPhotoSizeList":       RPCContextTuple{"/mtproto.RPCMedia/media_getPhotoSizeList", func() interface{} { return new(PhotoSizeList) }},
	"TLMediaGetPhotoSizeListList":   RPCContextTuple{"/mtproto.RPCMedia/media_getPhotoSizeListList", func() interface{} { return new(Vector_PhotoSizeList) }},
	"TLMediaGetVideoSizeList":       RPCContextTuple{"/mtproto.RPCMedia/media_getVideoSizeList", func() interface{} { return new(VideoSizeList) }},
	"TLMediaUploadedDocumentMedia":  RPCContextTuple{"/mtproto.RPCMedia/media_uploadedDocumentMedia", func() interface{} { return new(mtproto.MessageMedia) }},
	"TLMediaGetDocument":            RPCContextTuple{"/mtproto.RPCMedia/media_getDocument", func() interface{} { return new(mtproto.Document) }},
	"TLMediaGetDocumentList":        RPCContextTuple{"/mtproto.RPCMedia/media_getDocumentList", func() interface{} { return new(Vector_Document) }},
	"TLMediaUploadEncryptedFile":    RPCContextTuple{"/mtproto.RPCMedia/media_uploadEncryptedFile", func() interface{} { return new(mtproto.EncryptedFile) }},
	"TLMediaGetEncryptedFile":       RPCContextTuple{"/mtproto.RPCMedia/media_getEncryptedFile", func() interface{} { return new(mtproto.EncryptedFile) }},
	"TLMediaUploadWallPaperFile":    RPCContextTuple{"/mtproto.RPCMedia/media_uploadWallPaperFile", func() interface{} { return new(mtproto.Document) }},
	"TLMediaUploadThemeFile":        RPCContextTuple{"/mtproto.RPCMedia/media_uploadThemeFile", func() interface{} { return new(mtproto.Document) }},
	"TLMediaUploadStickerFile":      RPCContextTuple{"/mtproto.RPCMedia/media_uploadStickerFile", func() interface{} { return new(mtproto.Document) }},
	"TLMediaUploadRingtoneFile":     RPCContextTuple{"/mtproto.RPCMedia/media_uploadRingtoneFile", func() interface{} { return new(mtproto.Document) }},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		// log.Errorf("Can't find name: %s", rt.Name())
		return nil
	}
	return &m
}

func GetRPCContextRegisters() map[string]RPCContextTuple {
	return rpcContextRegisters
}
