package core

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
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

type dialogsFakeMsgClient struct {
	msgclient.MsgClient
	resolveDialogCursorTopMessage func(context.Context, *msgpb.TLMsgResolveDialogCursorTopMessage) (*msgpb.ResolvedDialogCursor, error)
}

func (f *dialogsFakeMsgClient) MsgResolveDialogCursorTopMessage(ctx context.Context, in *msgpb.TLMsgResolveDialogCursorTopMessage) (*msgpb.ResolvedDialogCursor, error) {
	return f.resolveDialogCursorTopMessage(ctx, in)
}

type dialogsFakeUserClient struct {
	userclient.UserClient
	getMutableUsersV2     func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	getUserProjection     func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
	getNotifySettingsList func(context.Context, *userpb.TLUserGetNotifySettingsList) (*userpb.VectorPeerPeerNotifySettings, error)
}

func (f *dialogsFakeUserClient) UserGetMutableUsersV2(ctx context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return f.getMutableUsersV2(ctx, in)
}

func (f *dialogsFakeUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	return f.getUserProjection(ctx, in)
}

func (f *dialogsFakeUserClient) UserGetNotifySettingsList(ctx context.Context, in *userpb.TLUserGetNotifySettingsList) (*userpb.VectorPeerPeerNotifySettings, error) {
	if f.getNotifySettingsList == nil {
		return &userpb.VectorPeerPeerNotifySettings{}, nil
	}
	return f.getNotifySettingsList(ctx, in)
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

func TestMessagesGetDialogsResolvesPublicOffsetIDToInternalCursor(t *testing.T) {
	var gotResolve *msgpb.TLMsgResolveDialogCursorTopMessage
	var gotDialogs *dialogpb.TLDialogGetDialogsV2
	c := newDialogsGetDialogsCore(&repository.Repository{
		MsgClient: &dialogsFakeMsgClient{
			resolveDialogCursorTopMessage: func(_ context.Context, in *msgpb.TLMsgResolveDialogCursorTopMessage) (*msgpb.ResolvedDialogCursor, error) {
				gotResolve = in
				return msgpb.MakeTLResolvedDialogCursor(&msgpb.TLResolvedDialogCursor{
					Found:       tg.BoolTrueClazz,
					PeerType:    dialogPeerTypeUser,
					PeerId:      200,
					PeerSeq:     7,
					MessageDate: 123,
				}).ToResolvedDialogCursor(), nil
			},
		},
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(_ context.Context, in *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				gotDialogs = in
				return dialogpb.MakeTLDialogPage(&dialogpb.TLDialogPage{Dialogs: []dialogpb.DialogExtV2Clazz{}, Exhausted: tg.BoolTrueClazz}), nil
			},
		},
	}, 100)

	_, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetId:   42,
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("MessagesGetDialogs error = %v", err)
	}
	if gotResolve == nil || gotResolve.UserId != 100 || gotResolve.TopMessageId != 42 {
		t.Fatalf("resolver request = %+v, want user_id 100 top_message_id 42", gotResolve)
	}
	cursor := gotDialogs.Cursor.ToDialogCursor()
	if cursor.TopPeerSeq != 7 || cursor.TopMessageDate != 123 || cursor.PeerType != dialogPeerTypeUser || cursor.PeerId != 200 {
		t.Fatalf("dialog cursor = %+v, want resolved internal cursor", cursor)
	}
}

func TestMessagesGetDialogsUnresolvedPositiveOffsetIDReturnsEmptySlice(t *testing.T) {
	dialogCalled := false
	c := newDialogsGetDialogsCore(&repository.Repository{
		MsgClient: &dialogsFakeMsgClient{
			resolveDialogCursorTopMessage: func(context.Context, *msgpb.TLMsgResolveDialogCursorTopMessage) (*msgpb.ResolvedDialogCursor, error) {
				return msgpb.MakeTLResolvedDialogCursor(&msgpb.TLResolvedDialogCursor{Found: tg.BoolFalseClazz}).ToResolvedDialogCursor(), nil
			},
		},
		DialogClient: &dialogsFakeDialogClient{
			getDialogsV2: func(context.Context, *dialogpb.TLDialogGetDialogsV2) (*dialogpb.DialogPage, error) {
				dialogCalled = true
				return nil, nil
			},
		},
	}, 100)

	r, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetId:   404,
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("MessagesGetDialogs error = %v", err)
	}
	if dialogCalled {
		t.Fatal("DialogGetDialogsV2 was called for unresolved positive offset_id")
	}
	slice, ok := r.ToMessagesDialogsSlice()
	if !ok || slice.Count != 0 || len(slice.Dialogs) != 0 {
		t.Fatalf("reply = %+v ok=%v, want empty slice", r, ok)
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
	var gotUsers *userpb.TLUserGetUserProjectionBundle
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
						TopUserMessageId:      7,
						TopCanonicalMessageId: 7,
					}),
					dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
						PeerType:              dialogPeerTypeChat,
						PeerId:                300,
						TopPeerSeq:            8,
						TopUserMessageId:      8,
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
			getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotUsers = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: selfID, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 200, FirstName: nonEmptyStringPtr("Alice")}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
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
	if gotUsers == nil || len(gotUsers.ViewerUserIds) != 1 || gotUsers.ViewerUserIds[0] != selfID || len(gotUsers.TargetUserIds) != 1 || gotUsers.TargetUserIds[0] != 200 {
		t.Fatalf("UserGetUserProjectionBundle request = %+v", gotUsers)
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
						TopUserMessageId:      7,
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
						PeerType:                  dialogPeerTypeUser,
						PeerId:                    200,
						ReadInboxMaxPeerSeq:       3,
						ReadInboxMaxUserMessageId: 3,
						ReadOutboxMaxPeerSeq:      0,
						DialogSchemaVersion:       1,
						TopPeerSeq:                7,
						TopUserMessageId:          7,
						TopCanonicalMessageId:     7,
						TopMessageDate:            123,
						PinnedPeerSeq:             0,
						PinnedCanonicalMessageId:  0,
						LastPts:                   1,
						LastPtsAt:                 123,
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
			getUserProjection: func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 200, FirstName: nonEmptyStringPtr("Peer")}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
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
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	c := newDialogsGetDialogsCore(&repository.Repository{
		UserClient: &dialogsFakeUserClient{
			getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 100, Self: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100)

	r, err := c.MessagesGetPeerSettings(&tg.TLMessagesGetPeerSettings{
		Peer: tg.MakeTLInputPeerSelf(&tg.TLInputPeerSelf{}),
	})
	if err != nil {
		t.Fatalf("MessagesGetPeerSettings error = %v", err)
	}
	if r == nil || r.Settings == nil {
		t.Fatalf("MessagesGetPeerSettings reply = %+v, want default settings", r)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 100 ||
		len(gotProjection.TargetUserIds) != 1 || gotProjection.TargetUserIds[0] != 100 {
		t.Fatalf("projection request = %+v, want viewer/target 100", gotProjection)
	}
	if len(r.Chats) != 0 || len(r.Users) != 1 {
		t.Fatalf("MessagesGetPeerSettings reply = %+v, want projected self user", r)
	}
}

func TestMessagesGetDialogsMapsUserDialogAndTopMessage(t *testing.T) {
	const selfID int64 = 100
	var gotMessageViews *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs
	var gotUsers *userpb.TLUserGetUserProjectionBundle

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
						PeerType:                   dialogPeerTypeUser,
						PeerId:                     200,
						TopPeerSeq:                 7,
						TopUserMessageId:           42,
						TopCanonicalMessageId:      2051345931852316672,
						TopMessageDate:             123,
						ReadInboxMaxUserMessageId:  43,
						ReadOutboxMaxUserMessageId: 44,
						UnreadCount:                1,
						Extras: dialogpb.MakeTLDialogExtras(&dialogpb.TLDialogExtras{
							PeerType:     dialogPeerTypeUser,
							PeerId:       200,
							DraftPayload: draftPayload,
						}),
					}),
				}, Exhausted: tg.BoolTrueClazz}), nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getMessageViewsByPeerSeqs: func(_ context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
				gotMessageViews = in
				return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: []tg.MessageClazz{
					tg.MakeTLMessage(&tg.TLMessage{
						Out:     false,
						Id:      42,
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
			getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotUsers = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: selfID, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 200, FirstName: nonEmptyStringPtr("Alice")}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
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
	if gotUsers == nil || len(gotUsers.ViewerUserIds) != 1 || gotUsers.ViewerUserIds[0] != selfID || len(gotUsers.TargetUserIds) != 1 || gotUsers.TargetUserIds[0] != 200 {
		t.Fatalf("UserGetUserProjectionBundle request = %+v, want viewer self and peer user 200", gotUsers)
	}

	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 1 || len(slice.Dialogs) != 1 || len(slice.Messages) != 1 || len(slice.Users) != 1 {
		t.Fatalf("reply lens = count:%d dialogs:%d messages:%d users:%d", slice.Count, len(slice.Dialogs), len(slice.Messages), len(slice.Users))
	}
	topMessage, ok := slice.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("top message = %T, want *tg.TLMessage", slice.Messages[0])
	}
	if topMessage.Out || topMessage.FromId != nil {
		t.Fatalf("incoming top message = out:%t from_id:%#v, want out=false from_id=nil", topMessage.Out, topMessage.FromId)
	}
	topPeer, ok := topMessage.PeerId.(*tg.TLPeerUser)
	if !ok || topPeer.UserId != 200 {
		t.Fatalf("top message peer_id = %#v, want peerUser(200)", topMessage.PeerId)
	}
	if dialog, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog(); !ok || dialog.TopMessage != 42 || dialog.ReadInboxMaxId != 43 || dialog.ReadOutboxMaxId != 44 {
		t.Fatalf("dialog = %+v, ok=%v, want public top/read ids", dialog, ok)
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

func nonEmptyStringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
