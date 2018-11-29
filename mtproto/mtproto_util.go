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

package mtproto

////////////////////////////////////////////////////////////////////////////////
func ToBool(b bool) *Bool {
	if b {
		return NewTLBoolTrue().To_Bool()
	} else {
		return NewTLBoolFalse().To_Bool()
	}
}

func FromBool(b *Bool) bool {
	return TLConstructor_CRC32_boolTrue == b.GetConstructor()
}

////////////////////////////////////////////////////////////////////////////////
//func ToInt32(v int32) *Int32 {
//	i := &mtproto.TLInt32{Data2: &mtproto.Int32_Data{
//		V: v,
//	}}
//	return i.To_Int32()
//}
//
//func FromInt32(i *Int32) int32 {
//	return i.GetData2().GetV()
//}

/*
//////////////////////////////////////////////////////////////////////////////////
// 太麻烦了
func GetUserIdListByChatParticipants(participants *TLChatParticipants) []int32 {
	chatUserIdList := []int32{}

	// TODO(@benqi):  nil check
	for _, participant := range participants.GetParticipants() {
		switch participant.Payload.(type) {
		case *ChatParticipant_ChatParticipant:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipant().GetUserId())
		case *ChatParticipant_ChatParticipantAdmin:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantAdmin().GetUserId())
		case *ChatParticipant_ChatParticipantCreator:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantCreator().GetUserId())
		}
	}
	return chatUserIdList
}

func (this *InputMedia) ToMessageMedia() (*MessageMedia) {
	switch this.Payload.(type) {
	case InputMedia_InputMediaUploadedPhoto:
		imedia := this.GetInputMediaUploadedPhoto()
		_ = TLInputMediaUploadedPhoto{}
		media := &TLMessageMediaPhoto{
			TtlSeconds: imedia.TtlSeconds,
			Caption: imedia.Caption,
		}

		p := &TLPhoto{
			HasStickers: len(imedia.Stickers) > 0,
			Date:int32(time.Now().Unix()),
		}

		switch imedia.GetFile().Payload.(type) {
		case *InputFile_InputFile:
			f := imedia.GetFile().GetInputFile()
			p.Id = f.Id
			p.AccessHash = 1 // f.Md5Checksum
			photoSize := &TLPhotoSize{
			}
			p.Sizes = append(p.Sizes, photoSize.ToPhotoSize())
		case *InputFile_InputFileBig:
			f := imedia.GetFile().GetInputFileBig()
			p.Id = f.Id
			p.AccessHash = 1 // f.Md5Checksum
			_ = TLInputFile{}
		}

		media.Photo = p.ToPhoto()
		return media.ToMessageMedia()
	}

	return nil
}
*/
