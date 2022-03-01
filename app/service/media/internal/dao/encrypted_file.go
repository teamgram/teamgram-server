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
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

func (m *Dao) SaveEncryptedFileV2(ctx context.Context, eF *mtproto.EncryptedFile) error {
	do := &dataobject.EncryptedFilesDO{
		EncryptedFileId: eF.Id,
		AccessHash:      eF.AccessHash,
		DcId:            eF.DcId,
		FilePath:        "",
		FileSize:        eF.Size2,
		KeyFingerprint:  eF.KeyFingerprint,
		Md5Checksum:     "",
	}

	_, _, err := m.EncryptedFilesDAO.Insert(ctx, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("error - %v", err)
	}

	return err
}

func (m *Dao) GetEncryptedFile(ctx context.Context, id, accessHash int64) (*mtproto.EncryptedFile, error) {
	do, err := m.EncryptedFilesDAO.SelectByFileLocation(ctx, id, accessHash)
	if err != nil {
		logx.WithContext(ctx).Errorf("error - %v", err)
		return nil, err
	}

	if do == nil {
		return mtproto.MakeTLEncryptedFileEmpty(nil).To_EncryptedFile(), nil
	} else {
		encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
			Id:             do.EncryptedFileId,
			AccessHash:     do.AccessHash,
			Size2:          do.FileSize,
			DcId:           do.DcId,
			KeyFingerprint: do.KeyFingerprint,
		})
		return encryptedFile.To_EncryptedFile(), nil
	}
}
