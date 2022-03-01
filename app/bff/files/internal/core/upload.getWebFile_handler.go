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

package core

import (
	"crypto/tls"
	"fmt"
	"github.com/teamgram/marmota/pkg/net/http/http_client"
	"github.com/teamgram/proto/mtproto"
	"github.com/zeromicro/go-zero/core/mathx"
	"time"
)

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (c *FilesCore) UploadGetWebFile(in *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	var (
		webfile *mtproto.Upload_WebFile
	)

	switch in.GetLocation().GetPredicateName() {
	case mtproto.Predicate_inputWebFileLocation:
		err := mtproto.ErrLocationInvalid
		c.Logger.Errorf("upload.getWebFile - error: %v", err)
		return nil, err
	case mtproto.Predicate_inputWebFileGeoPointLocation:
		/*
			Name	Type	Description
			geo_point	InputGeoPoint	Geolocation
			access_hash	long	Access hash
			w	int	Map width in pixels before applying scale; 16-1024
			h	int	Map height in pixels before applying scale; 16-1024
			zoom	int	Map zoom level; 13-20
			scale	int	Map scale; 1-3
		*/

		// https://static-maps.yandex.ru/1.x/?lang=en_US&ll={%f},{%f}&size=%d,%d&z=%d&l=map
		// https://osm-static-maps.herokuapp.com/?center=-73.998672,40.714728&zoom=12&height=400&width=400&attribution=Paytam%20Messenger
		webFilePath := fmt.Sprintf("https://static-maps.yandex.ru/1.x/?lang=en_US&ll=%f,%f&size=400,400&z=%d&l=map",
			in.GetLocation().GetGeoPoint().GetLong(),
			in.GetLocation().GetGeoPoint().GetLat(),
			// request.GetLocation().GetH()*request.GetLocation().GetScale(),
			// request.GetLocation().GetW()*request.GetLocation().GetScale(),
			in.GetLocation().GetZoom())

		// http://173.249.63.182:11443/maps/api/staticmap?center=40.714728,-73.998672&zoom=12&scale=2&size=160x120&key=AIzaSyA4Y5fAFpv1MTw1RjPIB6QYNdZ_ZHvhiG0

		//webFilePath := fmt.Sprintf("http://173.249.63.182:11443/maps/api/staticmap?center=%f,%f&zoom=%d&scale=%d&size=%dx%d&key=%s",
		//	request.GetLocation().GetGeoPoint().GetLat(),
		//	request.GetLocation().GetGeoPoint().GetLong(),
		//	request.GetLocation().GetZoom(),
		//	request.GetLocation().GetScale(),
		//	request.GetLocation().GetW(),
		//	request.GetLocation().GetH(),
		//	"AIzaSyA4Y5fAFpv1MTw1RjPIB6QYNdZ_ZHvhiG0")

		c.Logger.Infof("webfile: %s", webFilePath)
		bytes, err := http_client.Get(webFilePath).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(5*time.Second, 5*time.Second).
			Bytes()
		c.Logger.Infof("bytes: %d", len(bytes))
		size2 := int32(len(bytes))
		if err != nil {
			c.Logger.Errorf("upload.getWebFile - error: %v", err)
			err = mtproto.ErrLocationInvalid
			return nil, err
		}

		if in.GetOffset() > int32(len(bytes)) {
			bytes = bytes[0:0]
		} else {
			bytes = bytes[int(in.GetOffset()) : in.GetOffset()+int32(mathx.MinInt(int(in.GetLimit()), len(bytes))-int(in.GetOffset()))]
		}

		// size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes
		webfile = mtproto.MakeTLUploadWebFile(&mtproto.Upload_WebFile{
			Size2:    size2,
			MimeType: "img/png",
			FileType: mtproto.MakeTLStorageFileUnknown(nil).To_Storage_FileType(),
			Mtime:    int32(time.Now().Unix()),
			Bytes:    bytes,
		}).To_Upload_WebFile()
	default:
		err := mtproto.ErrLocationInvalid
		c.Logger.Errorf("upload.getWebFile - error: %v", err)
		return nil, err
	}

	return webfile, nil
}
