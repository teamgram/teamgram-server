/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfs

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
	"TLDfsWriteFilePartData":        RPCContextTuple{"/mtproto.RPCDfs/dfs_writeFilePartData", func() interface{} { return new(mtproto.Bool) }},
	"TLDfsUploadPhotoFileV2":        RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadPhotoFileV2", func() interface{} { return new(mtproto.Photo) }},
	"TLDfsUploadProfilePhotoFileV2": RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadProfilePhotoFileV2", func() interface{} { return new(mtproto.Photo) }},
	"TLDfsUploadEncryptedFileV2":    RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadEncryptedFileV2", func() interface{} { return new(mtproto.EncryptedFile) }},
	"TLDfsDownloadFile":             RPCContextTuple{"/mtproto.RPCDfs/dfs_downloadFile", func() interface{} { return new(mtproto.Upload_File) }},
	"TLDfsUploadDocumentFileV2":     RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadDocumentFileV2", func() interface{} { return new(mtproto.Document) }},
	"TLDfsUploadGifDocumentMedia":   RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadGifDocumentMedia", func() interface{} { return new(mtproto.Document) }},
	"TLDfsUploadMp4DocumentMedia":   RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadMp4DocumentMedia", func() interface{} { return new(mtproto.Document) }},
	"TLDfsUploadWallPaperFile":      RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadWallPaperFile", func() interface{} { return new(mtproto.Document) }},
	"TLDfsUploadThemeFile":          RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadThemeFile", func() interface{} { return new(mtproto.Document) }},
	"TLDfsUploadRingtoneFile":       RPCContextTuple{"/mtproto.RPCDfs/dfs_uploadRingtoneFile", func() interface{} { return new(mtproto.Document) }},
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
