package core

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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
	getDialogs       func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error)
	getPinnedDialogs func(context.Context, *dialogpb.TLDialogGetPinnedDialogs) (*dialogpb.VectorDialogExt, error)
	toggleDialogPin  func(context.Context, *dialogpb.TLDialogToggleDialogPin) (*tg.Int32, error)
	reorderPinned    func(context.Context, *dialogpb.TLDialogReorderPinnedDialogs) (*tg.Bool, error)
}

func (f *dialogsFakeDialogClient) DialogGetDialogs(ctx context.Context, in *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
	return f.getDialogs(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogGetPinnedDialogs(ctx context.Context, in *dialogpb.TLDialogGetPinnedDialogs) (*dialogpb.VectorDialogExt, error) {
	return f.getPinnedDialogs(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogToggleDialogPin(ctx context.Context, in *dialogpb.TLDialogToggleDialogPin) (*tg.Int32, error) {
	return f.toggleDialogPin(ctx, in)
}

func (f *dialogsFakeDialogClient) DialogReorderPinnedDialogs(ctx context.Context, in *dialogpb.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
	return f.reorderPinned(ctx, in)
}

type dialogsFakeUserClient struct {
	userclient.UserClient
	getMutableUsersV2 func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
}

func (f *dialogsFakeUserClient) UserGetMutableUsersV2(ctx context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return f.getMutableUsersV2(ctx, in)
}

type dialogsFakeMessageClient struct {
	messageclient.MessageClient
	getUserMessageList func(context.Context, *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error)
}

func (f *dialogsFakeMessageClient) MessageGetUserMessageList(ctx context.Context, in *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
	return f.getUserMessageList(ctx, in)
}

type dialogsFakeMsgClient struct {
	msgclient.MsgClient
	getHistory func(context.Context, *msg.TLMsgGetHistory) (*tg.MessagesMessages, error)
}

func (f *dialogsFakeMsgClient) MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
	return f.getHistory(ctx, in)
}

type dialogsFakeUserupdatesClient struct {
	userupdatesclient.UserupdatesClient
	getState          func(context.Context, *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error)
	getDifference     func(context.Context, *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error)
	getDialogsByPeers func(context.Context, *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error)
	processOperation  func(context.Context, *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	return f.getState(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	return f.getDifference(ctx, in)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetDialogsByPeers(ctx context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	return f.getDialogsByPeers(ctx, in)
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
	var got *dialogpb.TLDialogGetDialogs
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(_ context.Context, in *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				got = in
				return &dialogpb.VectorDialogExt{}, nil
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
		t.Fatal("DialogGetDialogs was not called")
	}
	if got.UserId != 100 || got.ExcludePinned != tg.ToBoolClazz(true) || got.FolderId != 1 {
		t.Fatalf("DialogGetDialogs request = %+v, want user_id=100 exclude_pinned=true folder_id=1", got)
	}
	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 0 || len(slice.Dialogs) != 0 || len(slice.Messages) != 0 || len(slice.Chats) != 0 || len(slice.Users) != 0 {
		t.Fatalf("empty dialogs reply = %+v, want empty messages.dialogsSlice", slice)
	}
}

func TestMessagesGetDialogsFallsBackToCanonicalSelfHistory(t *testing.T) {
	const selfID int64 = 100
	var gotHistory *msg.TLMsgGetHistory
	var gotUsers *userpb.TLUserGetMutableUsersV2

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{}, nil
			},
		},
		MsgClient: &dialogsFakeMsgClient{
			getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
				gotHistory = in
				return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
					Messages: []tg.MessageClazz{
						tg.MakeTLMessage(&tg.TLMessage{
							Out:     true,
							Id:      2,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: selfID}),
							Message: "self message",
							Date:    123,
						}),
					},
					Chats: []tg.ChatClazz{},
					Users: []tg.UserClazz{},
				}).ToMessagesMessages(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				gotUsers = in
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
					tg.MakeTLImmutableUser(&tg.TLImmutableUser{
						User: tg.MakeTLUserData(&tg.TLUserData{
							Id:        selfID,
							FirstName: "Self",
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
	if gotHistory == nil || gotHistory.UserId != selfID || gotHistory.PeerType != payload.PeerTypeUser || gotHistory.PeerId != selfID || gotHistory.Limit != 1 {
		t.Fatalf("MsgGetHistory request = %+v, want canonical self history limit=1", gotHistory)
	}
	if gotUsers == nil || len(gotUsers.Id) != 1 || gotUsers.Id[0] != selfID {
		t.Fatalf("UserGetMutableUsersV2 request = %+v, want self user", gotUsers)
	}

	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 1 || len(slice.Dialogs) != 1 || len(slice.Messages) != 1 || len(slice.Users) != 1 {
		t.Fatalf("reply lens = count:%d dialogs:%d messages:%d users:%d", slice.Count, len(slice.Dialogs), len(slice.Messages), len(slice.Users))
	}
	dialog, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog()
	if !ok || dialog.TopMessage != 2 || dialog.ReadInboxMaxId != 0 || dialog.ReadOutboxMaxId != 0 {
		t.Fatalf("dialog = %+v, ok=%v, want top_message=2 and unread fallback read state", dialog, ok)
	}
}

func TestMessagesGetDialogsFallsBackToCanonicalUserupdatesDialogs(t *testing.T) {
	const selfID int64 = 100
	var gotState *userupdates.TLUserupdatesGetState
	var gotDifference *userupdates.TLUserupdatesGetDifference
	var gotUsers *userpb.TLUserGetMutableUsersV2

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getState: func(_ context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
				gotState = in
				return userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 260}).ToUserState(), nil
			},
			getDifference: func(_ context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
				gotDifference = in
				return userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
					NewMessages: []tg.MessageClazz{
						tg.MakeTLMessage(&tg.TLMessage{
							Id:      2,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							Message: "incoming",
							Date:    123,
						}),
					},
					OtherUpdates: []tg.UpdateClazz{},
					State: userupdates.MakeTLUserState(&userupdates.TLUserState{
						Pts: 2,
					}),
				}).ToUserDifference(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
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
	if gotState == nil || gotState.UserId != selfID {
		t.Fatalf("UserupdatesGetState request = %+v, want user_id=%d", gotState, selfID)
	}
	if gotDifference == nil || gotDifference.UserId != selfID || gotDifference.Pts != 60 || gotDifference.PtsTotalLimit == nil || *gotDifference.PtsTotalLimit != 200 {
		t.Fatalf("UserupdatesGetDifference request = %+v, want latest-window diff pts=60 limit=200", gotDifference)
	}
	if gotUsers == nil || len(gotUsers.Id) != 1 || gotUsers.Id[0] != 200 {
		t.Fatalf("UserGetMutableUsersV2 request = %+v, want peer user 200", gotUsers)
	}
	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 1 || len(slice.Dialogs) != 1 || len(slice.Messages) != 1 || len(slice.Users) != 1 {
		t.Fatalf("reply lens = count:%d dialogs:%d messages:%d users:%d", slice.Count, len(slice.Dialogs), len(slice.Messages), len(slice.Users))
	}
	dialog, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog()
	if !ok || dialog.TopMessage != 2 || dialog.UnreadCount != 1 {
		t.Fatalf("dialog = %+v, ok=%v, want top_message=2 unread=1", dialog, ok)
	}
}

func TestMessagesGetDialogsMergesCanonicalUserupdatesDialogs(t *testing.T) {
	const selfID int64 = 100

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{Datas: []dialogpb.DialogExtClazz{
					dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
						Order: 1,
						Date:  100,
						Dialog: tg.MakeTLDialog(&tg.TLDialog{
							Peer:           tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
							TopMessage:     7,
							NotifySettings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
						}),
					}),
				}}, nil
			},
		},
		MessageClient: &dialogsFakeMessageClient{
			getUserMessageList: func(context.Context, *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
				return &messagepb.VectorMessageBox{Datas: []tg.MessageBoxClazz{
					tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:    selfID,
						MessageId: 7,
						Message: tg.MakeTLMessage(&tg.TLMessage{
							Out:     true,
							Id:      7,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
							Message: "legacy",
							Date:    100,
						}),
					}),
				}}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getState: func(context.Context, *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
				return userupdates.MakeTLUserState(&userupdates.TLUserState{Pts: 260}).ToUserState(), nil
			},
			getDifference: func(context.Context, *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
				return userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
					NewMessages: []tg.MessageClazz{
						tg.MakeTLMessage(&tg.TLMessage{
							Id:      2,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							Message: "canonical",
							Date:    123,
						}),
					},
					OtherUpdates: []tg.UpdateClazz{},
					State: userupdates.MakeTLUserState(&userupdates.TLUserState{
						Pts: 2,
					}),
				}).ToUserDifference(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				users := make([]tg.ImmutableUserClazz, 0, len(in.Id))
				for _, id := range in.Id {
					users = append(users, tg.MakeTLImmutableUser(&tg.TLImmutableUser{
						User: tg.MakeTLUserData(&tg.TLUserData{
							Id:        id,
							FirstName: "User",
						}),
						KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
					}))
				}
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: users}).ToMutableUsers(), nil
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

	slice, ok := r.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("MessagesGetDialogs returned %s, want messages.dialogsSlice", r.ClazzName())
	}
	if slice.Count != 2 || len(slice.Dialogs) != 2 || len(slice.Messages) != 2 || len(slice.Users) != 2 {
		t.Fatalf("reply lens = count:%d dialogs:%d messages:%d users:%d", slice.Count, len(slice.Dialogs), len(slice.Messages), len(slice.Users))
	}
	first, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog()
	if !ok || first.TopMessage != 2 {
		t.Fatalf("first dialog = %+v, ok=%v, want canonical top_message=2", first, ok)
	}
}

func TestMessagesGetPeerDialogsSelfFromCanonicalHistory(t *testing.T) {
	const selfID int64 = 100
	var gotHistory *msg.TLMsgGetHistory

	c := newDialogsGetDialogsCore(&repository.Repository{
		MsgClient: &dialogsFakeMsgClient{
			getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
				gotHistory = in
				return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
					Messages: []tg.MessageClazz{
						tg.MakeTLMessage(&tg.TLMessage{
							Out:     true,
							Id:      2,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: selfID}),
							Message: "self message",
							Date:    123,
						}),
					},
					Chats: []tg.ChatClazz{},
					Users: []tg.UserClazz{},
				}).ToMessagesMessages(), nil
			},
		},
		UserClient: &dialogsFakeUserClient{
			getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
				return tg.MakeTLMutableUsers(&tg.TLMutableUsers{Users: []tg.ImmutableUserClazz{
					tg.MakeTLImmutableUser(&tg.TLImmutableUser{
						User: tg.MakeTLUserData(&tg.TLUserData{
							Id:        selfID,
							FirstName: "Self",
						}),
						KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
					}),
				}}).ToMutableUsers(), nil
			},
		},
	}, selfID)

	r, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{
		Peers: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerSelf(&tg.TLInputPeerSelf{}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("MessagesGetPeerDialogs error = %v", err)
	}
	if gotHistory == nil || gotHistory.UserId != selfID || gotHistory.PeerId != selfID || gotHistory.Limit != 1 {
		t.Fatalf("MsgGetHistory request = %+v, want canonical self history limit=1", gotHistory)
	}
	if len(r.Dialogs) != 1 || len(r.Messages) != 1 || len(r.Users) != 1 {
		t.Fatalf("reply lens = dialogs:%d messages:%d users:%d", len(r.Dialogs), len(r.Messages), len(r.Users))
	}
	dialog, ok := (&tg.Dialog{Clazz: r.Dialogs[0]}).ToDialog()
	if !ok || dialog.TopMessage != 2 || dialog.ReadInboxMaxId != 0 || dialog.ReadOutboxMaxId != 0 {
		t.Fatalf("dialog = %+v, ok=%v, want top_message=2 and unread fallback read state", dialog, ok)
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

func TestMessagesGetPinnedDialogsReturnsEmptyPeerDialogs(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getPinnedDialogs: func(context.Context, *dialogpb.TLDialogGetPinnedDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{}, nil
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

func TestMessagesGetPinnedDialogsHydratesUserupdatesProjection(t *testing.T) {
	var gotProjection *userupdates.TLUserupdatesGetDialogsByPeers
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getPinnedDialogs: func(context.Context, *dialogpb.TLDialogGetPinnedDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{Datas: []dialogpb.DialogExtClazz{
					dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
						Order: 99,
						Dialog: tg.MakeTLDialog(&tg.TLDialog{
							Peer: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
						}),
					}),
				}}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
			getDialogsByPeers: func(_ context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
				gotProjection = in
				return &userupdates.VectorDialogProjection{Datas: []userupdates.DialogProjectionClazz{
					userupdates.MakeTLDialogProjection(&userupdates.TLDialogProjection{
						PeerType:              dialogPeerTypeUser,
						PeerId:                200,
						TopCanonicalMessageId: 7,
						TopMessageDate:        123,
						UnreadCount:           2,
					}),
				}}, nil
			},
		},
		MessageClient: &dialogsFakeMessageClient{
			getUserMessageList: func(context.Context, *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
				return &messagepb.VectorMessageBox{Datas: []tg.MessageBoxClazz{
					tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:    100,
						MessageId: 7,
						Message: tg.MakeTLMessage(&tg.TLMessage{
							Id:      7,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							Message: "hello",
							Date:    123,
						}),
					}),
				}}, nil
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
	if gotProjection == nil || len(gotProjection.Peers) != 1 || gotProjection.Peers[0].PeerType != dialogPeerTypeUser || gotProjection.Peers[0].PeerId != 200 {
		t.Fatalf("UserupdatesGetDialogsByPeers request = %+v", gotProjection)
	}
	if len(r.Dialogs) != 1 || len(r.Messages) != 1 || len(r.Users) != 1 {
		t.Fatalf("reply lens = dialogs:%d messages:%d users:%d", len(r.Dialogs), len(r.Messages), len(r.Users))
	}
	dialog, ok := (&tg.Dialog{Clazz: r.Dialogs[0]}).ToDialog()
	if !ok || dialog.TopMessage != 7 || dialog.UnreadCount != 2 {
		t.Fatalf("dialog = %+v ok=%v, want hydrated top/unread", dialog, ok)
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
	var gotMessages *messagepb.TLMessageGetUserMessageList
	var gotUsers *userpb.TLUserGetMutableUsersV2

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{Datas: []dialogpb.DialogExtClazz{
					dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
						Order: 10,
						Dialog: tg.MakeTLDialog(&tg.TLDialog{
							Peer:           tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							TopMessage:     7,
							ReadInboxMaxId: 7,
							NotifySettings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
						}),
						Date:          123,
						ThemeEmoticon: "",
					}),
				}}, nil
			},
		},
		MessageClient: &dialogsFakeMessageClient{
			getUserMessageList: func(_ context.Context, in *messagepb.TLMessageGetUserMessageList) (*messagepb.VectorMessageBox, error) {
				gotMessages = in
				return &messagepb.VectorMessageBox{Datas: []tg.MessageBoxClazz{
					tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:    selfID,
						MessageId: 7,
						Message: tg.MakeTLMessage(&tg.TLMessage{
							Out:     true,
							Id:      7,
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
							Message: "hello",
							Date:    123,
						}),
						Reaction: "",
					}),
				}}, nil
			},
		},
		UserClient: &dialogsFakeUserClient{
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
	if gotMessages == nil || gotMessages.UserId != selfID || len(gotMessages.IdList) != 1 || gotMessages.IdList[0] != 7 {
		t.Fatalf("MessageGetUserMessageList request = %+v, want self top message 7", gotMessages)
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
	if dialog, ok := (&tg.Dialog{Clazz: slice.Dialogs[0]}).ToDialog(); !ok || dialog.TopMessage != 7 {
		t.Fatalf("dialog = %+v, ok=%v, want top_message=7", dialog, ok)
	}
	if user, ok := (&tg.User{Clazz: slice.Users[0]}).ToUser(); !ok || user.Id != 200 || user.FirstName == nil || *user.FirstName != "Alice" {
		t.Fatalf("user = %+v, ok=%v, want Alice/200", user, ok)
	}
}

func TestMessagesGetDialogsMapsBizDialogErrorToInternal(t *testing.T) {
	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
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

func TestOffsetDialogExtsMatchesOffsetPeerByID(t *testing.T) {
	dialogs := []dialogpb.DialogExtClazz{
		dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
			Dialog: tg.MakeTLDialog(&tg.TLDialog{
				Peer:       tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
				TopMessage: 10,
			}),
		}),
		dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
			Dialog: tg.MakeTLDialog(&tg.TLDialog{
				Peer:       tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 201}),
				TopMessage: 9,
			}),
		}),
	}

	got := offsetDialogExts(dialogs, &tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 201}),
	}, 100)
	if len(got) != 0 {
		t.Fatalf("offset by second peer returned %d dialogs, want 0", len(got))
	}

	got = offsetDialogExts(dialogs, &tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 999}),
	}, 100)
	if len(got) != 2 {
		t.Fatalf("offset by missing peer returned %d dialogs, want original 2", len(got))
	}
}
