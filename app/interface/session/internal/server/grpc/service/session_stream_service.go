// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package service

import (
	"context"
	"strings"

	"github.com/teamgram/teamgram-server/app/interface/session/internal/core"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func (s *Service) SessionDataStream(stream grpc.BidiStreamingServer[session.SessionStreamRequest, session.SessionStreamResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			logx.Errorf("SessionDataStream Recv error: %v", err)
			return err
		}

		resp := s.handleStreamRequest(stream.Context(), req)
		if err := stream.Send(resp); err != nil {
			logx.Errorf("SessionDataStream Send error: %v", err)
			return err
		}
	}
}

func (s *Service) handleStreamRequest(ctx context.Context, req *session.SessionStreamRequest) *session.SessionStreamResponse {
	reqId := req.GetRequestId()

	switch p := req.GetPayload().(type) {
	case *session.SessionStreamRequest_SendData:
		return s.handleSendData(ctx, reqId, p.SendData)
	case *session.SessionStreamRequest_CreateSession:
		return s.handleCreateSession(ctx, reqId, p.CreateSession)
	case *session.SessionStreamRequest_CloseSession:
		return s.handleCloseSession(ctx, reqId, p.CloseSession)
	case *session.SessionStreamRequest_QueryAuthKey:
		return s.handleQueryAuthKey(ctx, reqId, p.QueryAuthKey)
	case *session.SessionStreamRequest_SetAuthKey:
		return s.handleSetAuthKey(ctx, reqId, p.SetAuthKey)
	case *session.SessionStreamRequest_PushUpdates:
		return s.handlePushUpdates(ctx, reqId, p.PushUpdates)
	case *session.SessionStreamRequest_PushSessionUpdates:
		return s.handlePushSessionUpdates(ctx, reqId, p.PushSessionUpdates)
	case *session.SessionStreamRequest_PushRpcResult:
		return s.handlePushRpcResult(ctx, reqId, p.PushRpcResult)
	default:
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    400,
					Message: "unknown payload type",
				},
			},
		}
	}
}

func (s *Service) makeShardingErrorResp(reqId string, err error) *session.SessionStreamResponse {
	// check if it's a redirect error (code 700)
	if target, ok := parseRedirectTarget(err); ok {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    700,
					Message: "REDIRECT_TO_" + target,
				},
			},
		}
	}
	return &session.SessionStreamResponse{
		RequestId: reqId,
		Payload: &session.SessionStreamResponse_Error{
			Error: &session.SessionStreamError{
				Code:    500,
				Message: err.Error(),
			},
		},
	}
}

func (s *Service) makeAckResp(reqId string) *session.SessionStreamResponse {
	return &session.SessionStreamResponse{
		RequestId: reqId,
		Payload: &session.SessionStreamResponse_Ack{
			Ack: &session.SessionStreamAck{Success: true},
		},
	}
}

func (s *Service) handleSendData(ctx context.Context, reqId string, in *session.TLSessionSendDataToSession) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetData().GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionSendDataToSession(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handleCreateSession(ctx context.Context, reqId string, in *session.TLSessionCreateSession) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetClient().GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionCreateSession(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handleCloseSession(ctx context.Context, reqId string, in *session.TLSessionCloseSession) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetClient().GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionCloseSession(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handleQueryAuthKey(ctx context.Context, reqId string, in *session.TLSessionQueryAuthKey) *session.SessionStreamResponse {
	c := core.New(ctx, s.svcCtx)
	authKey, err := c.SessionQueryAuthKey(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return &session.SessionStreamResponse{
		RequestId: reqId,
		Payload: &session.SessionStreamResponse_AuthKey{
			AuthKey: authKey,
		},
	}
}

func (s *Service) handleSetAuthKey(ctx context.Context, reqId string, in *session.TLSessionSetAuthKey) *session.SessionStreamResponse {
	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionSetAuthKey(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handlePushUpdates(ctx context.Context, reqId string, in *session.TLSessionPushUpdatesData) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionPushUpdatesData(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handlePushSessionUpdates(ctx context.Context, reqId string, in *session.TLSessionPushSessionUpdatesData) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionPushSessionUpdatesData(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

func (s *Service) handlePushRpcResult(ctx context.Context, reqId string, in *session.TLSessionPushRpcResultData) *session.SessionStreamResponse {
	if err := s.checkShardingV(ctx, in.GetPermAuthKeyId()); err != nil {
		return s.makeShardingErrorResp(reqId, err)
	}

	c := core.New(ctx, s.svcCtx)
	_, err := c.SessionPushRpcResultData(in)
	if err != nil {
		return &session.SessionStreamResponse{
			RequestId: reqId,
			Payload: &session.SessionStreamResponse_Error{
				Error: &session.SessionStreamError{
					Code:    500,
					Message: err.Error(),
				},
			},
		}
	}
	return s.makeAckResp(reqId)
}

// parseRedirectTarget extracts the target address from a redirect error.
// The error format follows mtproto.NewErrRedirectToX: gRPC code 700, message "REDIRECT_TO_<addr>"
func parseRedirectTarget(err error) (string, bool) {
	if err == nil {
		return "", false
	}
	s := err.Error()
	const prefix = "REDIRECT_TO_"
	idx := strings.Index(s, prefix)
	if idx < 0 {
		return "", false
	}
	target := s[idx+len(prefix):]
	if target == "" {
		return "", false
	}
	return target, true
}
