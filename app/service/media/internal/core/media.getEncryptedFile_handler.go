/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaGetEncryptedFile
// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
func (c *MediaCore) MediaGetEncryptedFile(in *media.TLMediaGetEncryptedFile) (*mtproto.EncryptedFile, error) {
	// TODO: not impl
	c.Logger.Errorf("media.getEncryptedFile blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
