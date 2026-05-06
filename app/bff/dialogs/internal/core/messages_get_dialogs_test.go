package core

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	messageclient "github.com/teamgram/teamgram-server/v2/app/service/biz/message/client"
	messagepb "github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type dialogsFakeDialogClient struct {
	dialogclient.DialogClient
	getDialogsV2       func(context.Context, *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error)
	getPeerDialogsV2   func(context.Context, *dialogpb.TLDialogGetPeerDialogsV2) (*dialogpb.VectorDialogExtV2, error)
	getPinnedDialogsV2 func(context.Context, *dialogpb.TLDialogGetPinnedDialogsV2) (*dialogpb.VectorDialogExtV2, error)
	toggleDialogPin    func(context.Context, *dialogpb.TLDialogToggleDialogPin) (*tg.Int32, error)
	reorderPinned      func(context.Context, *dialogpb.TLDialogReorderPinnedDialogs) (*tg.Bool, error)
}

func (f *dialogsFakeDialogClient) DialogGetDialogsV2(ctx context.Context, in *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
	return f.getDialogsV2(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogGetPeerDialogsV2(ctx context.Context, in *dialogpb.TLDialogGetPeerDialogsV2) (*dialogpb.VectorDialogExtV2, error) {
	return f.getPeerDialogsV2(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogGetPinnedDialogsV2(ctx context.Context, in *dialogpb.TLDialogGetPinnedDialogsV2) (*dialogpb.VectorDialogExtV2, error) {
	return f.getPinnedDialogsV2(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogToggleDialogPin(ctx context.Context, in *dialogpb.TLDialogToggleDialogPin) (*tg.Int32, error) {
	return f.toggleDialogPin(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogReorderPinnedDialogs(ctx context.Context, in *dialogpb.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
	return f.reorderPinned(ctx, in)
}

type dialogsFakeUserClient struct {
	userclient.UserClient
	getMutableUsersV2     func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	getNotifySettingsList func(context.Context, *userpb.TLUserGetNotifySettingsList) (*userpb.VectorPeerPeerNotifySettings, error)
}

func (f *dialogsFakeUserClient) UserGetMutableUsersV2(ctx context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return f.getMutableUsersV2(ctx, in)
}

func (f *dialogsFakeUserClient) UserGetNotifySettingsList(ctx context.Context, in *userpb.TLUserGetNotifySettingsList) (*userpb.VectorPeerPeerNotifySettings, error) {
	if f.getNotifySettingsList == nil {
		return &userpb.VectorPeerPeerNotifySettings{}, nil
	}
	return f.getNotifySettingsList(ctx, in)
}

type dialogsFakeMessageClient struct {
	messageclient.MessageClient
	getUserMessageList func(context.Context, *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error)
}

func (f *dialogsFakeMessageClient) MessageGetUserMessageList(ctx context.Context, in *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
	return f.getUserMessageList(ctx, in)
}

type dialogsFakeUserupdatesClient struct {
	userupdatesclient.UserupdatesClient
	getState                  func(context.Context, *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error)
	getDifference             func(context.Context, *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error)
	getDialogsByPeers         func(context.Context, *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error)
	getMessageViewsByPeerSeqs func(context.Context, *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error)
	processOperation          func(context.Context, *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	return f.getState(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	return f.getDifference(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetDialogsByPeers(ctx context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	if f.getDialogsByPeers == nil {
		return &userupdates.VectorDialogProjection{Datas: []userupdates.DialogProjectionClazz{}}, nil
	}
	return f.getDialogsByPeers(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
	if f.getMessageViewsByPeerSeqs == nil {
		return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: []tg.MessageClazz{}}).ToMessageViewList(), nil
	}
	return f.getMessageViewsByPeerSeqs(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesProcessUserOperation(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	return f.processOperation(ctx, in)
}

func newDialogsGetDialogsCore(repo *repository.Repository, selfID int64) *DialogsCore {
	c := New(context.Background(), &svc.ServiceContext{Repo: repo})
	c.MD = &metadata.RpcMetadata{UserId: selfID, PermAuthKeyId: 9001}
	return c
}

func TestMessagesGetDialogsReturnsEmptySliceFromBizDialog(t *testing.T) {
	var got *dialogpb.TLDialogGetDialogsV2
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(_ context.Context, in *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				got = in
				return dialogpb.MakeTLDialogPage(&dialogpb.TLDialogPage{Dialogs: []dialogpb.DialogExtV2Clazz{}, Exhausted: tg.BoolTrueClazz}), nil
			},
		},
	}, 100)

	folderID := int32(1)
	r, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		ExcludePinned: true,
		FolderId:      &folderID,
		OffsetPeer:    tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:         50,
	})
	if err != nil {
		t.Fatalf("MessagesGetDialogs error = %v", err)
	}
	if got == nil {
		t.Fatal("DialogGetDialogsV2 was not called")
	}
	cursor := got.Cursor.ToDialogCursor()
	if got.UserId != 100 || got.ExcludePinned != tg.ToBoolClazz(true) || got.Limit != 50 || cursor == nil || cursor.FolderId != 1 {
		t.Fatalf("DialogGetDialogsV2 request = %+v cursor=%+v, want user_id=100 exclude_pinned=true folder_id=1 limit=50", got, cursor)
	}
	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 0 || len(slice.Dialogs) != 0 || len(slice.Messages) != 0 || len(slice.Chats) != 0 || len(slice.Users) != 0 {
		t.Fatalf("empty dialogs reply = %+v, want empty messages.dialogsSlice", slice)
	}
}

func TestMessagesGetDialogsNoDifferenceFallback(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(context.Context, *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				return dialogpb.MakeTLDialogPage(&dialogpb.TLDialogPage{Dialogs: []dialogpb.DialogExtV2Clazz{}, Exhausted: tg.BoolTrueClazz}), nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getDifference: func(context.Context, *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
				t.Fatal("UserupdatesGetDifference must not be called by messages.getDialogs")
				return nil, nil
			},
		},
	}, 100)

	r, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("MessagesGetDialogs error = %v", err)
	}
	slice, ok := r.ToMessagesDialogsSlice()
	if !ok || slice.Count != 0 || len(slice.Dialogs) != 0 {
		t.Fatalf("MessagesGetDialogs reply = %+v ok=%v, want empty facade slice", r, ok)
	}
}

func TestMessagesGetPeerDialogsEmptyVectorReturnsEmptyPeerDialogs(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{}, 100)

	r, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{})
	if err != nil {
		t.Fatalf("MessagesGetPeerDialogs error = %v", err)
	}
	if r == nil {
		t.Fatal("MessagesGetPeerDialogs returned nil")
	}
	if len(r.Dialogs) != 0 || len(r.Messages) != 0 || len(r.Chats) != 0 || len(r.Users) != 0 || r.State == nil {
		t.Fatalf("empty peer dialogs reply = %+v, want empty vectors with state", r)
	}
}

func TestMessagesGetPeerDialogsHydratesTopMessagesUsersChats(t *testing.T) {
	const selfID int64 = 100
	var gotPeerDialogs *dialogpb.TLDialogGetPeerDialogsV2
	var gotMessageViews *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs
	var gotUsers *userpb.TLUserGetMutableUsersV2
	var gotChats *chatpb.TLChatGetChatListByIdList

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getPeerDialogsV2: func(_ context.Context, in *dialogpb.TLDialogGetPeerDialogsV2) (*dialogpb.VectorDialogExtV2, error) {
				gotPeerDialogs = in
				return &dialogpb.VectorDialogExtV2{Datas: []dialogpb.DialogExtV2Clazz{
					dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
						PeerType:              dialogPeerTypeUser,
						PeerId:                200,
						TopPeerSeq:            7,
						TopCanonicalMessageId: 7,
					}),
					dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
						PeerType:              dialogPeerTypeChat,
						PeerId:                300,
						TopPeerSeq:            8,
						TopCanonicalMessageId: 8,
					}),
				}}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getMessageViewsByPeerSeqs: func(_ context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
				gotMessageViews = in
				return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: []tg.MessageClazz{
					tg.MakeTLMessage(&tg.TLMessage{Id: 7, PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}), Message: "u"}),
					tg.MakeTLMessage(&tg.TLMessage{Id: 8, PeerId: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 300}), Message: "c"}),
				}}).ToMessageViewList(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				gotUsers = in
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
					tg.MakeTLImmutableUser(&tg.TLImmutableUser{User: tg.MakeTLUserData(&tg.TLUserData{Id: 200, FirstName: "Alice"})}),
				}}).ToMutableUsers(), nil
			},
		},
		ChatClient: &dialogsFakeChatClient{
			getChatListByIDList: func(_ context.Context, in *chatpb.TLChatGetChatListByIdList) (*chatpb.VectorMutableChat, error) {
				gotChats = in
				return &chatpb.VectorMutableChat{Datas: []tg.MutableChatClazz{testDialogsMutableChat(300)}}, nil
			},
		},
	}, selfID)

	r, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{
		Peers: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200})}),
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 300})}),
		},
	})
	if err != nil {
		t.Fatalf("MessagesGetPeerDialogs error = %v", err)
	}
	if gotPeerDialogs == nil || len(gotPeerDialogs.Peers) != 2 {
		t.Fatalf("DialogGetPeerDialogsV2 request = %+v", gotPeerDialogs)
	}
	if gotMessageViews == nil || gotMessageViews.UserId != selfID || len(gotMessageViews.Peers) != 2 ||
		gotMessageViews.Peers[0].PeerType != dialogPeerTypeUser || gotMessageViews.Peers[0].PeerId != 200 || gotMessageViews.Peers[0].PeerSeq != 7 ||
		gotMessageViews.Peers[1].PeerType != dialogPeerTypeChat || gotMessageViews.Peers[1].PeerId != 300 || gotMessageViews.Peers[1].PeerSeq != 8 {
		t.Fatalf("UserupdatesGetMessageViewsByPeerSeqs request = %+v", gotMessageViews)
	}
	if gotUsers == nil || len(gotUsers.Id) != 1 || gotUsers.Id[0] != 200 {
		t.Fatalf("UserGetMutableUsersV2 request = %+v", gotUsers)
	}
	if gotChats == nil || len(gotChats.IdList) != 1 || gotChats.IdList[0] != 300 {
		t.Fatalf("ChatGetChatListByIdList request = %+v", gotChats)
	}
	if len(r.Dialogs) != 2 || len(r.Messages) != 2 || len(r.Users) != 1 || len(r.Chats) != 1 {
		t.Fatalf("reply lens = dialogs:%d messages:%d users:%d chats:%d", len(r.Dialogs), len(r.Messages), len(r.Users), len(r.Chats))
	}
}

func TestMessagesGetPinnedDialogsReturnsEmptyPeerDialogs(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getPinnedDialogsV2: func(context.Context, *dialogpb.TLDialogGetPinnedDialogsV2) (*dialogpb.VectorDialogExtV2, error) {
				return &dialogpb.VectorDialogExtV2{}, nil
			},
		},
	}, 100)

	r, err := c.MessagesGetPinnedDialogs(&tg.TLMessagesGetPinnedDialogs{FolderId: 0})
	if err != nil {
		t.Fatalf("MessagesGetPinnedDialogs error = %v", err)
	}
	if r == nil {
		t.Fatal("MessagesGetPinnedDialogs returned nil")
	}
	if len(r.Dialogs) != 0 || len(r.Messages) != 0 || len(r.Chats) != 0 || len(r.Users) != 0 {
		t.Fatalf("MessagesGetPinnedDialogs reply = %+v, want empty peerDialogs", r)
	}
	if r.State == nil {
		t.Fatal("MessagesGetPinnedDialogs state is nil")
	}
}

func TestMessagesGetPinnedDialogsHydratesFacadeProjection(t *testing.T) {
	var gotPinned *dialogpb.TLDialogGetPinnedDialogsV2
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getPinnedDialogsV2: func(_ context.Context, in *dialogpb.TLDialogGetPinnedDialogsV2) (*dialogpb.VectorDialogExtV2, error) {
				gotPinned = in
				return &dialogpb.VectorDialogExtV2{Datas: []dialogpb.DialogExtV2Clazz{
					dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
						PeerType:              dialogPeerTypeUser,
						PeerId:                200,
						TopPeerSeq:            7,
						TopCanonicalMessageId: 7,
						TopMessageDate:        123,
						UnreadCount:           2,
						MainPinnedOrder:       99,
					}),
				}}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getDialogsByPeers: func(_ context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
				if len(in.Peers) != 1 || in.Peers[0].PeerType != dialogPeerTypeUser || in.Peers[0].PeerId != 200 {
					t.Fatalf("UserupdatesGetDialogsByPeers request = %+v, want user/200", in)
				}
				return &userupdates.VectorDialogProjection{Datas: []userupdates.DialogProjectionClazz{
					userupdates.MakeTLDialogProjection(&userupdates.TLDialogProjection{
						PeerType:                 dialogPeerTypeUser,
						PeerId:                   200,
						ReadInboxMaxPeerSeq:      3,
						ReadOutboxMaxPeerSeq:     0,
						DialogSchemaVersion:      1,
						TopPeerSeq:               7,
						TopCanonicalMessageId:    7,
						TopMessageDate:           123,
						PinnedPeerSeq:            0,
						PinnedCanonicalMessageId: 0,
						LastPts:                  1,
						LastPtsAt:                123,
					}),
				}}, nil
			},
			getMessageViewsByPeerSeqs: func(context.Context, *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
				return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: []tg.MessageClazz{
					tg.MakeTLMessage(&tg.TLMessage{
						Id:      7,
						PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
						Message: "hello",
						Date:    123,
					}),
				}}).ToMessageViewList(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getMutableUsersV2: func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
					tg.MakeTLImmutableUser(&tg.TLImmutableUser{
						User: tg.MakeTLUserData(&tg.TLUserData{Id: 200, FirstName: "Peer"}),
					}),
				}}).ToMutableUsers(), nil
			},
		},
	}, 100)

	r, err := c.MessagesGetPinnedDialogs(&tg.TLMessagesGetPinnedDialogs{FolderId: 0})
	if err != nil {
		t.Fatalf("MessagesGetPinnedDialogs error = %v", err)
	}
	if gotPinned == nil || gotPinned.UserId != 100 || gotPinned.FolderId != 0 || gotPinned.Limit != 100 {
		t.Fatalf("DialogGetPinnedDialogsV2 request = %+v", gotPinned)
	}
	if len(r.Dialogs) != 1 || len(r.Messages) != 1 || len(r.Users) != 1 {
		t.Fatalf("reply lens = dialogs:%d messages:%d users:%d", len(r.Dialogs), len(r.Messages), len(r.Users))
	}
	dialog, ok := (&tg.Dialog{Clazz: r.Dialogs[0]}).ToDialog()
	if !ok || dialog.TopMessage != 7 || dialog.ReadInboxMaxId != 3 || dialog.ReadOutboxMaxId != 0 || dialog.UnreadCount != 2 {
		t.Fatalf("dialog = %+v ok=%v, want hydrated top/read/unread", dialog, ok)
	}
}

func TestMessagesToggleDialogPinPassesSourceAuthAndOutbox(t *testing.T) {
	var got *dialogpb.TLDialogToggleDialogPin
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			toggleDialogPin: func(_ context.Context, in *dialogpb.TLDialogToggleDialogPin) (*tg.Int32, error) {
				got = in
				return tg.MakeInt32(3), nil
			},
		},
	}, 100)

	r, err := c.MessagesToggleDialogPin(&tg.TLMessagesToggleDialogPin{
		Pinned: true,
		Peer: tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
			Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		}),
	})
	if err != nil {
		t.Fatalf("MessagesToggleDialogPin error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("reply = %v, want boolTrue", r)
	}
	if got == nil {
		t.Fatal("DialogToggleDialogPin was not called")
	}
	if got.UserId != 100 || got.PeerType != dialogPeerTypeUser || got.PeerId != 200 || got.Pinned != tg.BoolTrueClazz {
		t.Fatalf("DialogToggleDialogPin request = %+v", got)
	}
	if got.SourcePermAuthKeyId != 9001 || got.OperationId == "" || got.OutboxId == 0 {
		t.Fatalf("source auth/outbox fields = auth:%d op:%q outbox:%d", got.SourcePermAuthKeyId, got.OperationId, got.OutboxId)
	}
}

func TestMessagesReorderPinnedDialogsPassesPeerDialogIDs(t *testing.T) {
	var got *dialogpb.TLDialogReorderPinnedDialogs
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			reorderPinned: func(_ context.Context, in *dialogpb.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
				got = in
				return tg.BoolTrue, nil
			},
		},
	}, 100)

	_, err := c.MessagesReorderPinnedDialogs(&tg.TLMessagesReorderPinnedDialogs{
		Force:    true,
		FolderId: 1,
		Order: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
			}),
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 300}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("MessagesReorderPinnedDialogs error = %v", err)
	}
	if got == nil {
		t.Fatal("DialogReorderPinnedDialogs was not called")
	}
	wantIDs := []int64{200*16 + int64(dialogPeerTypeUser), 300*16 + int64(dialogPeerTypeChat)}
	if got.UserId != 100 || got.Force != tg.BoolTrueClazz || got.FolderId != 1 || len(got.IdList) != 2 || got.IdList[0] != wantIDs[0] || got.IdList[1] != wantIDs[1] {
		t.Fatalf("DialogReorderPinnedDialogs request = %+v, want ids %v", got, wantIDs)
	}
	if got.SourcePermAuthKeyId != 9001 || got.OperationId == "" || got.OutboxId == 0 {
		t.Fatalf("source auth/outbox fields = auth:%d op:%q outbox:%d", got.SourcePermAuthKeyId, got.OperationId, got.OutboxId)
	}
}

func TestMessagesMarkDialogUnreadWritesUserupdatesOperation(t *testing.T) {
	var got *userupdates.TLUserOperation
	c := newDialogsGetDialogsCore(&repository.Repository{
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			processOperation: func(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
				got = in.Operation
				return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{UserId: 100, OperationId: in.Operation.OperationId, Status: 1, Pts: 5, PtsCount: 1}), nil
			},
		},
	}, 100)

	r, err := c.MessagesMarkDialogUnread(&tg.TLMessagesMarkDialogUnread{
		Unread: true,
		Peer: tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
			Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		}),
	})
	if err != nil {
		t.Fatalf("MessagesMarkDialogUnread error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("reply = %v, want boolTrue", r)
	}
	if got == nil || got.UserId != 100 || got.PeerType != dialogPeerTypeUser || got.PeerId != 200 {
		t.Fatalf("operation = %+v", got)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(got.Payload, &op); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindMarkDialogUnread || op.UnreadMark == nil || !*op.UnreadMark {
		t.Fatalf("payload = %+v", op)
	}
}

func TestMessagesGetPeerSettingsReturnsDefaultSettings(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{}, 100)

	r, err := c.MessagesGetPeerSettings(&tg.TLMessagesGetPeerSettings{
		Peer: tg.MakeTLInputPeerSelf(&tg.TLInputPeerSelf{}),
	})
	if err != nil {
		t.Fatalf("MessagesGetPeerSettings error = %v", err)
	}
	if r == nil || r.Settings == nil {
		t.Fatalf("MessagesGetPeerSettings reply = %+v, want default settings", r)
	}
	if len(r.Chats) != 0 || len(r.Users) != 0 {
		t.Fatalf("MessagesGetPeerSettings reply = %+v, want empty peers", r)
	}
}

func TestMessagesGetDialogsMapsUserDialogAndTopMessage(t *testing.T) {
	const selfID int64 = 100
	var gotMessageViews *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs
	var gotUsers *userpb.TLUserGetMutableUsersV2

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(context.Context, *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				draftPayload, err := json.Marshal(tg.MakeTLDraftMessage(&tg.TLDraftMessage{
					NoWebpage: true,
					Message:   "draft text",
					Date:      456,
				}))
				if err != nil {
					t.Fatalf("marshal draft payload: %v", err)
				}
				return dialogpb.MakeTLDialogPage(&dialogpb.TLDialogPage{Dialogs: []dialogpb.DialogExtV2Clazz{
					dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
						PeerType:              dialogPeerTypeUser,
						PeerId:                200,
						TopPeerSeq:            7,
						TopCanonicalMessageId: 2051345931852316672,
						TopMessageDate:        123,
						UnreadCount:           1,
						Extras: dialogpb.MakeTLDialogExtras(&dialogpb.TLDialogExtras{
							PeerType:     dialogPeerTypeUser,
							PeerId:       200,
							DraftPayload: draftPayload,
						}),
					}),
				}, Exhausted: tg.BoolTrueClazz}), nil
			},
		},
		MessageClient: &dialogsFakeMessageClient{
			getUserMessageList: func(_ context.Context, in *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
				t.Fatalf("legacy MessageGetUserMessageList must not be called: %+v", in)
				return nil, tg.ErrMethodNotImpl
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getMessageViewsByPeerSeqs: func(_ context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
				gotMessageViews = in
				return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: []tg.MessageClazz{
					tg.MakeTLMessage(&tg.TLMessage{
						Out:     true,
						Id:      7,
						PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
						Message: "hello",
						Date:    123,
					}),
				}}).ToMessageViewList(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getNotifySettingsList: func(context.Context, *userpb.TLUserGetNotifySettingsList) (*userpb.VectorPeerPeerNotifySettings, error) {
				return &userpb.VectorPeerPeerNotifySettings{Datas: []userpb.PeerPeerNotifySettingsClazz{
					userpb.MakeTLPeerPeerNotifySettings(&userpb.TLPeerPeerNotifySettings{
						PeerType: tg.PEER_USER,
						PeerId:   200,
						Settings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{MuteUntil: int32Ptr(99)}),
					}),
				}}, nil
			},
			getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				gotUsers = in
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
					tg.MakeTLImmutableUser(&tg.TLImmutableUser{
						User: tg.MakeTLUserData(&tg.TLUserData{
							Id:        200,
							FirstName: "Alice",
						}),
						KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
					}),
				}}).ToMutableUsers(), nil
			},
		},
	}, selfID)

	r, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("MessagesGetDialogs error = %v", err)
	}
	if gotMessageViews == nil || gotMessageViews.UserId != selfID || len(gotMessageViews.Peers) != 1 ||
		gotMessageViews.Peers[0].PeerType != dialogPeerTypeUser || gotMessageViews.Peers[0].PeerId != 200 || gotMessageViews.Peers[0].PeerSeq != 7 {
		t.Fatalf("UserupdatesGetMessageViewsByPeerSeqs request = %+v, want self/user/200 seq 7", gotMessageViews)
	}
	if gotUsers == nil || len(gotUsers.Id) != 1 || gotUsers.Id[0] != 200 || !gotUsers.Privacy || !gotUsers.HasTo || len(gotUsers.To) != 1 || gotUsers.To[0] != selfID {
		t.Fatalf("UserGetMutableUsersV2 request = %+v, want peer user 200 with privacy to self", gotUsers)
	}

	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 1 || len(slice.Dialogs) != 1 || len(slice.Messages) != 1 || len(slice.Users) != 1 {
		t.Fatalf("reply lens = count:%d dialogs:%d messages:%d users:%d", slice.Count, len(slice.Dialogs), len(slice.Messages), len(slice.Users))
	}
	if dialog, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog(); !ok || dialog.TopMessage != 7 || dialog.ReadInboxMaxId != 0 || dialog.ReadOutboxMaxId != 0 {
		t.Fatalf("dialog = %+v, ok=%v, want top_message=7 and unread read state", dialog, ok)
	} else if settings := dialog.NotifySettings; settings == nil || settings.MuteUntil == nil || *settings.MuteUntil != 99 {
		t.Fatalf("notify settings = %+v, want batched mute_until=99", dialog.NotifySettings)
	} else if draft, ok := (&tg.DraftMessage{Clazz: dialog.Draft}).ToDraftMessage(); !ok || !draft.NoWebpage || draft.Message != "draft text" || draft.Date != 456 {
		t.Fatalf("draft = %+v, ok=%v, want saved draft payload", draft, ok)
	}
	if user, ok := (&tg.User{Clazz: slice.Users[0]}).ToUser(); !ok || user.Id != 200 || user.FirstName == nil || *user.FirstName != "Alice" {
		t.Fatalf("user = %+v, ok=%v, want Alice/200", user, ok)
	}
}

func TestMessagesGetDialogsMapsBizDialogErrorToInternal(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(context.Context, *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				return nil, errors.New("biz unavailable")
			},
		},
	}, 100)

	_, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      10,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesGetDialogs error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func int32Ptr(v int32) *int32 {
	return &v
}
