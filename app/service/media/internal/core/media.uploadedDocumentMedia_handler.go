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
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaUploadedDocumentMedia
// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
func (c *MediaCore) MediaUploadedDocumentMedia(in *media.TLMediaUploadedDocumentMedia) (*mtproto.MessageMedia, error) {
	var (
		media   = in.GetMedia()
		ownerId = in.GetOwnerId()

		isGif    = false
		document *mtproto.Document
		err      error
	)

	for _, attr := range in.GetMedia().GetAttributes() {
		if attr.PredicateName == mtproto.Predicate_documentAttributeAnimated {
			isGif = true
			break
		}
	}
	if isGif && media.GetMimeType() == "image/gif" {
		document, err = c.svcCtx.Dao.DfsClient.DfsUploadGifDocumentMedia(c.ctx, &dfs.TLDfsUploadGifDocumentMedia{
			Creator: ownerId,
			Media:   media,
		})
		if err != nil {
			c.Logger.Errorf("media.uploadedDocumentMedia - error: %v", err)
			return nil, err
		}

		if len(document.GetThumbs()) > 0 {
			c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
		}
		c.svcCtx.Dao.SaveDocumentV2(c.ctx, media.GetFile().GetName(), document)
	} else if in.GetMedia().GetMimeType() == "video/mp4" {
		document, err = c.svcCtx.Dao.DfsClient.DfsUploadMp4DocumentMedia(c.ctx, &dfs.TLDfsUploadMp4DocumentMedia{
			Creator: ownerId,
			Media:   media,
		})
		if err != nil {
			c.Logger.Errorf("media.uploadedDocumentMedia - error: %v", err)
			return nil, err
		}

		if len(document.GetThumbs()) > 0 {
			c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
		}
		c.svcCtx.Dao.SaveDocumentV2(c.ctx, media.GetFile().GetName(), document)
	} else {
		document, err = c.svcCtx.Dao.DfsClient.DfsUploadDocumentFileV2(c.ctx, &dfs.TLDfsUploadDocumentFileV2{
			Creator: ownerId,
			Media:   media,
		})
		if err != nil {
			c.Logger.Errorf("media.uploadedDocumentMedia - error: %v", err)
			return nil, err
		}

		if len(document.GetThumbs()) > 0 {
			c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
		}
		c.svcCtx.Dao.SaveDocumentV2(c.ctx, media.GetFile().GetName(), document)
	}

	// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
	return mtproto.MakeTLMessageMediaDocument(&mtproto.MessageMedia{
		Document:   document,
		TtlSeconds: in.GetMedia().GetTtlSeconds(),
	}).To_MessageMedia(), nil
}
