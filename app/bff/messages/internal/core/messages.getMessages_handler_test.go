package core

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesGetMessagesReturnsMessagesInRequestOrder(t *testing.T) {
	var got *msg.TLMsgGetUserMessageList
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			getUserMessageList: func(_ context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
				got = in
				return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
					tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:       100,
						MessageId:    7,
						PeerType:     payload.PeerTypeUser,
						PeerId:       300,
						SenderUserId: 300,
						Message: tg.MakeTLMessage(&tg.TLMessage{
							Id:      7,
							FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 100}),
							Date:    1,
							Message: "seven",
						}),
					}),
					tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:       100,
						MessageId:    9,
						PeerType:     payload.PeerTypeUser,
						PeerId:       400,
						SenderUserId: 400,
						Message: tg.MakeTLMessage(&tg.TLMessage{
							Id:      9,
							FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 400}),
							PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 100}),
							Date:    2,
							Message: "nine",
						}),
					}),
				}}, nil
			},
		},
		UserClient: &messagesFakeUserClient{
			projectUsers: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 100, Self: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 300, Contact: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 400, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100, 200)

	r, err := c.MessagesGetMessages(&tg.TLMessagesGetMessages{
		Id_VECTORINPUTMESSAGE: []tg.InputMessageClazz{
			tg.MakeTLInputMessageID(&tg.TLInputMessageID{Id: 9}),
			tg.MakeTLInputMessageReplyTo(&tg.TLInputMessageReplyTo{Id: 7}),
		},
		Id_VECTORINT32: []int32{11},
	})
	if err != nil {
		t.Fatalf("MessagesGetMessages() error = %v", err)
	}
	if got == nil || got.UserId != 100 || !reflect.DeepEqual(got.IdList, []int32{9, 7, 11}) {
		t.Fatalf("msg request = %+v, want user 100 ids [9 7 11]", got)
	}
	out, ok := r.ToMessagesMessages()
	if !ok {
		t.Fatalf("result type = %T, want messages.messages", r.Clazz)
	}
	if len(out.Messages) != 3 {
		t.Fatalf("messages len = %d, want 3", len(out.Messages))
	}
	if m, ok := out.Messages[0].(*tg.TLMessage); !ok || m.Id != 9 || m.Message != "nine" {
		t.Fatalf("message[0] = %#v, want id 9 text nine", out.Messages[0])
	}
	if m, ok := out.Messages[1].(*tg.TLMessage); !ok || m.Id != 7 || m.Message != "seven" {
		t.Fatalf("message[1] = %#v, want id 7 text seven", out.Messages[1])
	}
	if m, ok := out.Messages[2].(*tg.TLMessageEmpty); !ok || m.Id != 11 {
		t.Fatalf("message[2] = %#v, want messageEmpty id 11", out.Messages[2])
	}
	if gotProjection == nil || !reflect.DeepEqual(gotProjection.TargetUserIds, []int64{400, 100, 300}) {
		t.Fatalf("projection target ids = %+v, want [400 100 300]", gotProjection)
	}
	if len(out.Users) != 3 {
		t.Fatalf("users len = %d, want 3", len(out.Users))
	}
}

func TestMessagesGetMessagesFallsBackWhenBatchHasMissingMessage(t *testing.T) {
	var singleIDs []int32
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			getUserMessageList: func(_ context.Context, _ *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
				return nil, msg.ErrMsgIdInvalid
			},
			getUserMessage: func(_ context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
				singleIDs = append(singleIDs, in.Id)
				if in.Id == 7 {
					return tg.MakeTLMessageBox(&tg.TLMessageBox{
						UserId:    100,
						MessageId: 7,
						Message:   tg.MakeTLMessage(&tg.TLMessage{Id: 7, Message: "seven"}),
					}).ToMessageBox(), nil
				}
				return nil, msg.ErrMsgIdInvalid
			},
		},
	}, 100, 200)

	r, err := c.MessagesGetMessages(&tg.TLMessagesGetMessages{
		Id_VECTORINPUTMESSAGE: []tg.InputMessageClazz{
			tg.MakeTLInputMessageID(&tg.TLInputMessageID{Id: 7}),
			tg.MakeTLInputMessageID(&tg.TLInputMessageID{Id: 99}),
		},
	})
	if err != nil {
		t.Fatalf("MessagesGetMessages() error = %v", err)
	}
	if !reflect.DeepEqual(singleIDs, []int32{7, 99}) {
		t.Fatalf("single fetch ids = %v, want [7 99]", singleIDs)
	}
	out, ok := r.ToMessagesMessages()
	if !ok || len(out.Messages) != 2 {
		t.Fatalf("messages = %#v, ok=%v, want two messages", r, ok)
	}
	if m, ok := out.Messages[0].(*tg.TLMessage); !ok || m.Id != 7 {
		t.Fatalf("message[0] = %#v, want id 7", out.Messages[0])
	}
	if m, ok := out.Messages[1].(*tg.TLMessageEmpty); !ok || m.Id != 99 {
		t.Fatalf("message[1] = %#v, want messageEmpty id 99", out.Messages[1])
	}
}

func TestMessagesGetMessagesMapsStorageError(t *testing.T) {
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			getUserMessageList: func(_ context.Context, _ *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
				return nil, msg.ErrMsgStorage
			},
		},
	}, 100, 200)

	_, err := c.MessagesGetMessages(&tg.TLMessagesGetMessages{
		Id_VECTORINPUTMESSAGE: []tg.InputMessageClazz{
			tg.MakeTLInputMessageID(&tg.TLInputMessageID{Id: 7}),
		},
	})
	if !errors.Is(err, tg.ErrInternalServerError) {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}
