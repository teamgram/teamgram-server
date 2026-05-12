package rpc

import (
	"context"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
)

type ChatClient struct {
	inner chatclient.ChatClient
}

func NewChatClient(inner chatclient.ChatClient) *ChatClient {
	return &ChatClient{inner: inner}
}

func (c *ChatClient) ChatCheckChatAccess(ctx context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
	return c.inner.ChatCheckChatAccess(ctx, in)
}

func (c *ChatClient) ChatCheckMessageAction(ctx context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
	return c.inner.ChatCheckMessageAction(ctx, in)
}

func (c *ChatClient) ChatGetChatParticipantIdList(ctx context.Context, in *chatpb.TLChatGetChatParticipantIdList) (*chatpb.VectorLong, error) {
	return c.inner.ChatGetChatParticipantIdList(ctx, in)
}
