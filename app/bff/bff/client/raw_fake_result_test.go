package bffproxyclient

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestRawFakeReturnsEncodedTLBytes(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, &tg.TLHelpGetTermsOfServiceUpdate{}))
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	if _, ok := obj.(*tg.TLHelpTermsOfServiceUpdateEmpty); !ok {
		t.Fatalf("DecodeObject() = %T, want *tg.TLHelpTermsOfServiceUpdateEmpty", obj)
	}
}

func TestRawFakeDecodesRequestFields(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, &tg.TLLangpackGetDifference{
		LangPack:    "tdesktop",
		LangCode:    "zh-hans",
		FromVersion: 7,
	}))
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	diff, ok := obj.(*tg.TLLangPackDifference)
	if !ok {
		t.Fatalf("DecodeObject() = %T, want *tg.TLLangPackDifference", obj)
	}
	if diff.LangCode != "zh-hans" || diff.FromVersion != 7 || diff.Version != 7 {
		t.Fatalf("difference = %#v", diff)
	}
}

func TestRawFakeMainScreenStartupMethods(t *testing.T) {
	tests := []struct {
		name string
		req  interface {
			Encode(*bin.Encoder, int32) error
		}
		want func(t *testing.T, obj iface.TLObject)
	}{
		{
			name: "account.updateStatus",
			req:  &tg.TLAccountUpdateStatus{Offline: tg.BoolTrueClazz},
			want: wantType[*tg.TLBoolTrue],
		},
		{
			name: "updates.getState",
			req:  &tg.TLUpdatesGetState{},
			want: wantType[*tg.TLUpdatesState],
		},
		{
			name: "messages.getDialogs",
			req: &tg.TLMessagesGetDialogs{
				OffsetPeer: tg.InputPeerEmptyClazz,
			},
			want: wantType[*tg.TLMessagesDialogs],
		},
		{
			name: "messages.getPinnedDialogs",
			req:  &tg.TLMessagesGetPinnedDialogs{FolderId: 0},
			want: wantType[*tg.TLMessagesPeerDialogs],
		},
		{
			name: "messages.getDialogFilters",
			req:  &tg.TLMessagesGetDialogFilters{},
			want: wantType[*tg.TLMessagesDialogFilters],
		},
		{
			name: "help.getPeerColors",
			req:  &tg.TLHelpGetPeerColors{Hash: 0},
			want: wantType[*tg.TLHelpPeerColors],
		},
		{
			name: "messages.getAvailableReactions",
			req:  &tg.TLMessagesGetAvailableReactions{Hash: 0},
			want: wantType[*tg.TLMessagesAvailableReactions],
		},
		{
			name: "messages.getTopReactions",
			req:  &tg.TLMessagesGetTopReactions{Limit: 10, Hash: 0},
			want: wantType[*tg.TLMessagesReactions],
		},
		{
			name: "messages.getRecentReactions",
			req:  &tg.TLMessagesGetRecentReactions{Limit: 10, Hash: 0},
			want: wantType[*tg.TLMessagesReactions],
		},
		{
			name: "messages.getSavedReactionTags",
			req:  &tg.TLMessagesGetSavedReactionTags{Hash: 0},
			want: wantType[*tg.TLMessagesSavedReactionTags],
		},
		{
			name: "messages.getScheduledHistory",
			req:  &tg.TLMessagesGetScheduledHistory{Peer: tg.InputPeerSelfClazz, Hash: 0},
			want: wantType[*tg.TLMessagesMessages],
		},
		{
			name: "messages.getDefaultTagReactions",
			req:  &tg.TLMessagesGetDefaultTagReactions{Hash: 0},
			want: wantType[*tg.TLMessagesReactions],
		},
		{
			name: "messages.getAvailableEffects",
			req:  &tg.TLMessagesGetAvailableEffects{Hash: 0},
			want: wantType[*tg.TLMessagesAvailableEffects],
		},
		{
			name: "messages.getStickerSet",
			req:  &tg.TLMessagesGetStickerSet{Stickerset: tg.InputStickerSetEmptyClazz, Hash: 0},
			want: wantType[*tg.TLMessagesStickerSetNotModified],
		},
		{
			name: "account.getDefaultEmojiStatuses",
			req:  &tg.TLAccountGetDefaultEmojiStatuses{Hash: 0},
			want: wantType[*tg.TLAccountEmojiStatuses],
		},
		{
			name: "users.getFullUser",
			req:  &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz},
			want: wantType[*tg.TLUsersUserFull],
		},
		{
			name: "account.getNotifySettings",
			req:  &tg.TLAccountGetNotifySettings{Peer: tg.InputNotifyUsersClazz},
			want: wantType[*tg.TLPeerNotifySettings],
		},
		{
			name: "messages.getEmojiGroups",
			req:  &tg.TLMessagesGetEmojiGroups{Hash: 0},
			want: wantType[*tg.TLMessagesEmojiGroups],
		},
		{
			name: "messages.getAttachMenuBots",
			req:  &tg.TLMessagesGetAttachMenuBots{Hash: 0},
			want: wantType[*tg.TLAttachMenuBots],
		},
		{
			name: "messages.getQuickReplies",
			req:  &tg.TLMessagesGetQuickReplies{Hash: 0},
			want: wantType[*tg.TLMessagesQuickReplies],
		},
		{
			name: "contacts.getContacts",
			req:  &tg.TLContactsGetContacts{Hash: 0},
			want: wantType[*tg.TLContactsContacts],
		},
		{
			name: "contacts.getTopPeers",
			req:  &tg.TLContactsGetTopPeers{Correspondents: true, Limit: 64, Hash: 0},
			want: wantType[*tg.TLContactsTopPeers],
		},
		{
			name: "payments.getStarGiftActiveAuctions",
			req:  &tg.TLPaymentsGetStarGiftActiveAuctions{Hash: 0},
			want: wantType[*tg.TLPaymentsStarGiftActiveAuctions],
		},
		{
			name: "stories.getAllStories",
			req:  &tg.TLStoriesGetAllStories{Next: false},
			want: wantType[*tg.TLStoriesAllStories],
		},
		{
			name: "stories.getStoriesArchive",
			req:  &tg.TLStoriesGetStoriesArchive{Peer: tg.InputPeerSelfClazz, Limit: 100},
			want: wantType[*tg.TLStoriesStories],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, tt.req))
			if err != nil {
				t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
			}
			if !ok {
				t.Fatal("TryReturnRawFakeRpcResult() ok = false")
			}
			obj, err := iface.DecodeObject(bin.NewDecoder(payload))
			if err != nil {
				t.Fatalf("DecodeObject() error = %v", err)
			}
			tt.want(t, obj)
		})
	}
}

func TestRawFakeReturnsEmptyEmojiKeywordsLanguages(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, &tg.TLMessagesGetEmojiKeywordsLanguages{}))
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}

	var got tg.VectorEmojiLanguage
	if err := got.Decode(bin.NewDecoder(payload)); err != nil {
		t.Fatalf("VectorEmojiLanguage.Decode() error = %v", err)
	}
	if len(got.Datas) != 0 {
		t.Fatalf("len(VectorEmojiLanguage.Datas) = %d, want 0", len(got.Datas))
	}
}

func TestRawFakeEncodesResponseWithClientLayer(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(
		context.Background(),
		&metadata.RpcMetadata{Layer: 223},
		encodeRawFakeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}),
	)
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}

	d := bin.NewDecoder(payload)
	if id, err := d.ClazzID(); err != nil || id != tg.ClazzID_users_userFull {
		t.Fatalf("response constructor = %#x, %v; want users.userFull", id, err)
	}
	if id, err := d.ClazzID(); err != nil || id != 0xa02bc13e {
		t.Fatalf("full_user constructor = %#x, %v; want userFull#a02bc13e for layer 223", id, err)
	}
}

func TestRawFakeUsersGetFullUserUsesMetadataSelfUser(t *testing.T) {
	const selfUserID int64 = 1571766986

	payload, ok, err := TryReturnRawFakeRpcResult(
		context.Background(),
		&metadata.RpcMetadata{Layer: 223, UserId: selfUserID},
		encodeRawFakeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}),
	)
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}

	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	userFull, ok := obj.(*tg.TLUsersUserFull)
	if !ok {
		t.Fatalf("DecodeObject() = %T, want *tg.TLUsersUserFull", obj)
	}
	if userFull.FullUser.Id != selfUserID {
		t.Fatalf("FullUser.Id = %d, want %d", userFull.FullUser.Id, selfUserID)
	}
	if len(userFull.Users) != 1 {
		t.Fatalf("len(Users) = %d, want 1", len(userFull.Users))
	}
	self, ok := userFull.Users[0].(*tg.TLUser)
	if !ok {
		t.Fatalf("Users[0] = %T, want *tg.TLUser", userFull.Users[0])
	}
	if self.Id != selfUserID || !self.Self {
		t.Fatalf("Users[0] = {Id:%d Self:%t}, want {Id:%d Self:true}", self.Id, self.Self, selfUserID)
	}
}

func TestRawFakeUnknownConstructor(t *testing.T) {
	x := bin.NewEncoder()
	x.PutClazzID(0xfeed9999)
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, append([]byte(nil), x.Bytes()...))
	x.End()
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if ok || payload != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() = %x, %v; want nil, false", payload, ok)
	}
}

func wantType[T iface.TLObject](t *testing.T, obj iface.TLObject) {
	t.Helper()
	if _, ok := obj.(T); !ok {
		t.Fatalf("DecodeObject() = %T, want %T", obj, *new(T))
	}
}

func encodeRawFakeTL(t *testing.T, obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	t.Helper()
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 224); err != nil {
		x.Reset()
		if err2 := obj.Encode(x, 0); err2 != nil {
			t.Fatalf("Encode(%T) error = %v", obj, err)
		}
	}
	return append([]byte(nil), x.Bytes()...)
}
