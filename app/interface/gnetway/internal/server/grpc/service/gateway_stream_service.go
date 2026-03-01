// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package service

import (
	"github.com/teamgram/teamgram-server/app/interface/gnetway/gateway"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func (s *Service) GatewayDataStream(stream grpc.BidiStreamingServer[gateway.GatewayStreamRequest, gateway.GatewayStreamResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			logx.Errorf("GatewayDataStream Recv error: %v", err)
			return err
		}

		sendData := req.GetSendData()
		if sendData == nil {
			if err := stream.Send(&gateway.GatewayStreamResponse{
				RequestId: req.GetRequestId(),
				Success:   false,
			}); err != nil {
				logx.Errorf("GatewayDataStream Send error: %v", err)
				return err
			}
			continue
		}

		// delegate to existing unary handler
		_, rpcErr := s.RPCGatewayServer.GatewaySendDataToGateway(stream.Context(), sendData)

		resp := &gateway.GatewayStreamResponse{
			RequestId: req.GetRequestId(),
			Success:   rpcErr == nil,
		}
		if err := stream.Send(resp); err != nil {
			logx.Errorf("GatewayDataStream Send error: %v", err)
			return err
		}
	}
}
