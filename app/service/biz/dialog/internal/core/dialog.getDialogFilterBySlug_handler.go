// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogGetDialogFilterBySlug
// dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
func (c *DialogCore) DialogGetDialogFilterBySlug(in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error) {
	var (
		dialogFilter *dialog.DialogFilterExt
	)

	v, err := c.svcCtx.Dao.DialogFiltersDAO.SelectBySlug(c.ctx, in.UserId, in.Slug)
	if err != nil {
		c.Logger.Errorf("dialog.getDialogFilterBySlug - error: %v", err)
		return nil, err
	}

	if v == nil {
		c.Logger.Errorf("dialog.getDialogFilterBySlug - error: %v", err)
		return nil, mtproto.ErrInviteSlugEmpty
	} else {
		dialogFilter = &dialog.DialogFilterExt{
			Id:           v.DialogFilterId,
			JoinedBySlug: true,
			Slug:         in.Slug,
			DialogFilter: nil,
			Order:        v.OrderValue,
		}

		if err = jsonx.UnmarshalFromString(v.DialogFilter, &dialogFilter.DialogFilter); err != nil {
			c.Logger.Errorf("jsonx.UnmarshalFromString(%v) - error: %v", v, err)
			return nil, err
		}

		if dialogFilter.DialogFilter == nil {
			dialogFilter.DialogFilter = mtproto.MakeTLDialogFilter(nil).To_DialogFilter()
		}
	}

	return dialogFilter, nil
}
