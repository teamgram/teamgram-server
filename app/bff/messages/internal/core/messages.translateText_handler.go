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
	"github.com/teamgram/proto/mtproto"
)

// MessagesTranslateText
// messages.translateText#24ce6dee flags:# peer:flags.0?InputPeer msg_id:flags.0?int text:flags.1?string from_lang:flags.2?string to_lang:string = messages.TranslatedText;
func (c *MessagesCore) MessagesTranslateText(in *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.translateText blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
