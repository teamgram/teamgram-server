package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMasterUnimplementedMethodsStayUnimplemented(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(nil, nil, 1001)
	if _, err := core.PhotosUploadContactProfilePhoto(&tg.TLPhotosUploadContactProfilePhoto{}); err != tg.ErrMethodNotImpl {
		t.Fatalf("PhotosUploadContactProfilePhoto error = %v, want METHOD_NOT_IMPL", err)
	}
}

func TestEnterpriseBlockedMethodsStayBlocked(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(nil, nil, 1001)
	if _, err := core.UsersSuggestBirthday(&tg.TLUsersSuggestBirthday{}); err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("UsersSuggestBirthday error = %v, want ENTERPRISE_IS_BLOCKED", err)
	}
	if _, err := core.ChannelsSetMainProfileTab(&tg.TLChannelsSetMainProfileTab{}); err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("ChannelsSetMainProfileTab error = %v, want ENTERPRISE_IS_BLOCKED", err)
	}
	if _, err := core.AccountUpdateVerified(&tg.TLAccountUpdateVerified{}); err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("AccountUpdateVerified error = %v, want ENTERPRISE_IS_BLOCKED", err)
	}
}
