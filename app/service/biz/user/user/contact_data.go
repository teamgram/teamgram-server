// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package user

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *Vector_ContactData) ToContacts() []*mtproto.Contact {
	contacts := make([]*mtproto.Contact, 0, len(m.GetDatas()))
	for _, c := range m.GetDatas() {
		contacts = append(contacts, mtproto.MakeTLContact(&mtproto.Contact{
			UserId: c.ContactUserId,
			Mutual: mtproto.ToBool(c.MutualContact),
		}).To_Contact())
	}

	return contacts
}
