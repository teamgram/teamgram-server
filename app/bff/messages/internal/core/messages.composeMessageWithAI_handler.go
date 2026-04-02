// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesComposeMessageWithAI
// messages.composeMessageWithAI#fd426afe flags:# proofread:flags.0?true emojify:flags.3?true text:TextWithEntities translate_to_lang:flags.1?string change_tone:flags.2?string = messages.ComposedMessageWithAI;
func (c *MessagesCore) MessagesComposeMessageWithAI(in *tg.TLMessagesComposeMessageWithAI) (*tg.MessagesComposedMessageWithAI, error) {
	text := ""
	if in.Text != nil {
		text = in.Text.Text
	}

	return &tg.MessagesComposedMessageWithAI{
		ResultText: tg.MakeTLTextWithEntities(&tg.TLTextWithEntities{
			Text:     text,
			Entities: []tg.MessageEntityClazz{},
		}),
	}, nil
}
