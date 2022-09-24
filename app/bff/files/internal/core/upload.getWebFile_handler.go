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
	"google.golang.org/grpc/status"
)

var (
	ErrWebfileNotAvailable = status.Error(mtproto.ErrBadRequest, "WEBFILE_NOT_AVAILABLE")
)

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (c *FilesCore) UploadGetWebFile(in *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	switch in.GetLocation().GetPredicateName() {
	case mtproto.Predicate_inputWebFileAudioAlbumThumbLocation:
		err := ErrWebfileNotAvailable
		c.Logger.Errorf("upload.getWebFile - error: %v", err)

		return nil, err
	default:
		c.Logger.Errorf("upload.getWebFile blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}
}
