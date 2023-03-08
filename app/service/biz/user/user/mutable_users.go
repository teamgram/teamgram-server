// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package user

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *Vector_ImmutableUser) GetUnsafeUser(selfId, id int64) (*mtproto.User, error) {
	if len(m.GetDatas()) == 0 {
		return nil, mtproto.ErrUserIdInvalid
	}

	if selfId == id {
		return m.GetUnsafeUserSelf(selfId)
	}

	me := m.findImmutableUser(selfId)
	to := m.findImmutableUser(id)

	if me == nil || to == nil {
		return nil, mtproto.ErrUserIdInvalid
	}

	return to.ToUnsafeUser(me), nil
}

func (m *Vector_ImmutableUser) GetUnsafeUserSelf(id int64) (*mtproto.User, error) {
	if len(m.GetDatas()) == 0 {
		return nil, mtproto.ErrUserIdInvalid
	}

	me := m.findImmutableUser(id)
	if me == nil {
		return nil, mtproto.ErrUserIdInvalid
	}

	return me.ToSelfUser(), nil
}

func (m *Vector_ImmutableUser) findImmutableUser(id int64) *mtproto.ImmutableUser {
	for _, v := range m.GetDatas() {
		if v.Id() == id {
			return v
		}
	}

	return nil
}

func (m *Vector_ImmutableUser) GetImmutableUser(id int64) (u *mtproto.ImmutableUser, ok bool) {
	for _, v := range m.GetDatas() {
		if v.Id() == id {
			return v, true
		}
	}

	return nil, false
}

func (m *Vector_ImmutableUser) CheckExistUser(id ...int64) bool {
	for _, id2 := range id {
		if _, ok := m.GetImmutableUser(id2); !ok {
			return false
		}
	}

	return true
}

func (m *Vector_ImmutableUser) GetUserListByIdList(selfId int64, id ...int64) []*mtproto.User {
	users := make([]*mtproto.User, 0, len(id))

	if len(m.GetDatas()) == 0 {
		return users
	}

	me := m.findImmutableUser(selfId)
	if me == nil {
		return users
	}

	for _, id2 := range id {
		to := m.findImmutableUser(id2)
		if to != nil {
			users = append(users, to.ToUnsafeUser(me))
		}
	}

	return users
}

func (m *Vector_ImmutableUser) Visit(cb func(it *mtproto.ImmutableUser)) {
	for _, v := range m.GetDatas() {
		cb(v)
	}
}

func (m *Vector_ImmutableUser) VisitByMe(meId int64, cb func(me, it *mtproto.ImmutableUser)) {
	var (
		me *mtproto.ImmutableUser
	)

	for _, v := range m.GetDatas() {
		if v.Id() == meId {
			me = v
			break
		}
	}

	if me == nil {
		return
	}

	for _, v := range m.GetDatas() {
		cb(me, v)
	}
}

func (m *Vector_ImmutableUser) Length() int {
	return len(m.GetDatas())
}

func (m *Vector_ImmutableUser) Empty() bool {
	return m == nil || len(m.GetDatas()) == 0
}
