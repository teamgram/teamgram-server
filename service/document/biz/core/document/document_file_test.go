// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package document

import (
	"encoding/json"
	"fmt"
	"github.com/nebula-chat/chatengine/mtproto"
	"testing"
)

func TestDocumentAttributes(t *testing.T) {
	attributes := &mtproto.DocumentAttributeList{}
	imageSize := &mtproto.TLDocumentAttributeImageSize{Data2: &mtproto.DocumentAttribute_Data{
		W: 512,
		H: 512,
	}}
	attributes.Attributes = append(attributes.Attributes, imageSize.To_DocumentAttribute())

	sticker := &mtproto.TLDocumentAttributeSticker{Data2: &mtproto.DocumentAttribute_Data{
		Alt: "😂",
		Stickerset: &mtproto.InputStickerSet{
			Constructor: mtproto.TLConstructor_CRC32_inputStickerSetID,
			Data2: &mtproto.InputStickerSet_Data{
				Id:         835404231795015689,
				AccessHash: 987465871030319816,
			},
		},
	}}
	attributes.Attributes = append(attributes.Attributes, sticker.To_DocumentAttribute())

	fileName := &mtproto.TLDocumentAttributeFilename{Data2: &mtproto.DocumentAttribute_Data{
		FileName: "sticker.webp",
	}}

	attributes.Attributes = append(attributes.Attributes, fileName.To_DocumentAttribute())
	d, _ := json.Marshal(attributes)
	fmt.Println(string(d))
}
