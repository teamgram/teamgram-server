package core

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountUpdateBirthdayDelegates(t *testing.T) {
	var got *userpb.TLUserUpdateBirthday
	year := int32(1815)
	birthday := tg.MakeTLBirthday(&tg.TLBirthday{Day: 10, Month: 12, Year: &year})
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		updateBirthday: func(_ context.Context, in *userpb.TLUserUpdateBirthday) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, nil, 1001)
	if _, err := core.AccountUpdateBirthday(&tg.TLAccountUpdateBirthday{Birthday: birthday}); err != nil {
		t.Fatalf("AccountUpdateBirthday returned error: %v", err)
	}
	if got == nil || got.UserId != 1001 || got.Birthday != birthday {
		t.Fatalf("birthday request = %+v", got)
	}
}

func TestAccountUpdatePersonalChannelDelegates(t *testing.T) {
	var got *userpb.TLUserUpdatePersonalChannel
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		updatePersonalChannel: func(_ context.Context, in *userpb.TLUserUpdatePersonalChannel) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, nil, 1001)
	if _, err := core.AccountUpdatePersonalChannel(&tg.TLAccountUpdatePersonalChannel{Channel: tg.MakeTLInputChannel(&tg.TLInputChannel{ChannelId: 7007})}); err != nil {
		t.Fatalf("AccountUpdatePersonalChannel returned error: %v", err)
	}
	if got == nil || got.UserId != 1001 || got.ChannelId != 7007 {
		t.Fatalf("personal channel request = %+v", got)
	}
}

func TestAccountSetMainProfileTabDelegates(t *testing.T) {
	var got *userpb.TLUserSetMainProfileTab
	tab := tg.MakeTLProfileTabPosts(&tg.TLProfileTabPosts{})
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		setMainProfileTab: func(_ context.Context, in *userpb.TLUserSetMainProfileTab) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, nil, 1001)
	if _, err := core.AccountSetMainProfileTab(&tg.TLAccountSetMainProfileTab{Tab: tab}); err != nil {
		t.Fatalf("AccountSetMainProfileTab returned error: %v", err)
	}
	if got == nil || got.UserId != 1001 || got.Tab != tab {
		t.Fatalf("main profile tab request = %+v", got)
	}
}

func TestAccountSaveMusicReturnsFalseNilOnServiceError(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		saveMusic: func(context.Context, *userpb.TLUserSaveMusic) (*tg.Bool, error) {
			return nil, errors.New("user service unavailable")
		},
	}, nil, 1001)
	got, err := core.AccountSaveMusic(&tg.TLAccountSaveMusic{Id: tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 9009})})
	if err != nil {
		t.Fatalf("AccountSaveMusic error = %v, want nil", err)
	}
	if got != tg.BoolFalse {
		t.Fatalf("AccountSaveMusic = %#v, want BoolFalse", got)
	}
}

func TestAccountGetSavedMusicIds(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getSavedMusicIDList: func(_ context.Context, in *userpb.TLUserGetSavedMusicIdList) (*userpb.VectorLong, error) {
			if in.UserId != 1001 {
				t.Fatalf("user id = %d, want 1001", in.UserId)
			}
			return &userpb.VectorLong{Datas: []int64{11, 22}}, nil
		},
	}, nil, 1001)
	got, err := core.AccountGetSavedMusicIds(&tg.TLAccountGetSavedMusicIds{})
	if err != nil {
		t.Fatalf("AccountGetSavedMusicIds returned error: %v", err)
	}
	ids, ok := got.Clazz.(*tg.TLAccountSavedMusicIds)
	if !ok || len(ids.Ids) != 2 || ids.Ids[0] != 11 || ids.Ids[1] != 22 {
		t.Fatalf("ids = %+v, want [11 22]", got)
	}
}
