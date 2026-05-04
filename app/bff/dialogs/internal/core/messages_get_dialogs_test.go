package core

import (
	"context"
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
	getDialogs func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error)
}

func (f *dialogsFakeDialogClient) DialogGetDialogs(ctx context.Context, in *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
	return f.getDialogs(ctx, in)
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
	getDifference func(context.Context, *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error)
}

func (f *dialogsFakeUserupdatesClient) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	return f.getDifference(ctx, in)
}

func newDialogsGetDialogsCore(repo *repository.Repository, selfID int64) *DialogsCore {
	c := New(context.Background(), &svc.ServiceContext{Repo: repo})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
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
	var gotDifference *userupdates.TLUserupdatesGetDifference
	var gotUsers *userpb.TLUserGetMutableUsersV2

	c := newDialogsGetDialogsCore(&repository.Repository{
		DialogClient: &dialogsFakeDialogClient{
			getDialogs: func(context.Context, *dialogpb.TLDialogGetDialogs) (*dialogpb.VectorDialogExt, error) {
				return &dialogpb.VectorDialogExt{}, nil
			},
		},
		UserupdatesClient: &dialogsFakeUserupdatesClient{
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
	if gotDifference == nil || gotDifference.UserId != selfID || gotDifference.Pts != 0 || gotDifference.PtsTotalLimit == nil || *gotDifference.PtsTotalLimit != 20 {
		t.Fatalf("UserupdatesGetDifference request = %+v, want initial diff limit=20", gotDifference)
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
	c := newDialogsGetDialogsCore(&repository.Repository{}, 100)

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
