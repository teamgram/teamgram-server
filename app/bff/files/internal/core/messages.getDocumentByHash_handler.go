// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
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
	"errors"

	"github.com/teamgram/proto/v2/tg"
)

// MessagesGetDocumentByHash
// messages.getDocumentByHash#b1f2061f sha256:bytes size:long mime_type:string = Document;
func (c *FilesCore) MessagesGetDocumentByHash(in *tg.TLMessagesGetDocumentByHash) (*tg.Document, error) {
	// TODO: not impl
	// c.Logger.Errorf("messages.getDocumentByHash blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, errors.New("messages.getDocumentByHash not implemented")
}
