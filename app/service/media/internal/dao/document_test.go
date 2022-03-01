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

package dao

import (
	"testing"
)

func TestGetDocumentList(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	//mysqlConfig1 := mysql_client.MySQLConfig{
	//	Name:   "immaster",
	//	DSN:    "root:@/enterprise?charset=utf8",
	//	Active: 5,
	//	Idle:   2,
	//}
	//mysqlConfig2 := mysql_client.MySQLConfig{
	//	Name:   "imslave",
	//	DSN:    "root:@/enterprise?charset=utf8",
	//	Active: 5,
	//	Idle:   2,
	//}
	//mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig1, mysqlConfig2})
	//
	//mdidia, _ := InstallMediaCore(&conf.ServiceConfig{MysqlNames: []string{"immaster"}})
	//// photoModel := photo.NewPhotoModel(1, "immaster", "")
	//// documentModel := NewDocumentModel(1, "immaster", "", photoModel)
	//documents := mdidia.GetDocumentList([]int64{1108902269895577600})
	//fmt.Printf("%s\n", log.DebugString(documents))
}

func TestDocumentAttributes(t *testing.T) {
	//attributes := &mtproto.DocumentAttributeList{}
	//imageSize := &mtproto.TLDocumentAttributeImageSize{Data2: &mtproto.DocumentAttribute_Data{
	//	W: 512,
	//	H: 512,
	//}}
	//attributes.Attributes = append(attributes.Attributes, imageSize.To_DocumentAttribute())
	//
	//sticker := &mtproto.TLDocumentAttributeSticker{Data2: &mtproto.DocumentAttribute_Data{
	//	Alt: "😂",
	//	Stickerset: &mtproto.InputStickerSet{
	//		Constructor: mtproto.TLConstructor_CRC32_inputStickerSetID,
	//		Data2: &mtproto.InputStickerSet_Data{
	//			Id:         835404231795015689,
	//			AccessHash: 987465871030319816,
	//		},
	//	},
	//}}
	//attributes.Attributes = append(attributes.Attributes, sticker.To_DocumentAttribute())
	//
	//fileName := &mtproto.TLDocumentAttributeFilename{Data2: &mtproto.DocumentAttribute_Data{
	//	FileName: "sticker.webp",
	//}}
	//
	//attributes.Attributes = append(attributes.Attributes, fileName.To_DocumentAttribute())
	//d, _ := json.Marshal(attributes)
	//fmt.Println(string(d))
}
