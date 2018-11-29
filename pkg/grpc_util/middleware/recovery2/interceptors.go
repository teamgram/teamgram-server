// Copyright 2017 David Ackroyd. All Rights Reserved.
// See LICENSE for licensing terms.

// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

// helper recovery, RecoveryHandlerFunc spit UnaryRecoveryHandlerFunc and StreamRecoveryHandlerFunc
//
package grpc_recovery2

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type UnaryRecoveryHandlerFunc func(ctx context.Context, p interface{}) (err error)
type StreamRecoveryHandlerFunc func(stream grpc.ServerStream, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = unaryRecoverFrom(ctx, r, o.unaryRecoveryHandlerFunc)
			}
		}()

		// TODO(@benqi): 加一层
		r, err2 := handler(ctx, req)
		if err2 != nil {
			err = unaryRecoverFrom(ctx, err2, o.unaryRecoveryHandlerFunc2)
			return r, err
		} else {
			return r, nil
		}
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = streamRecoverFrom(stream, r, o.streamRecoveryHandlerFunc)
			}
		}()

		return handler(srv, stream)
	}
}

func unaryRecoverFrom(ctx context.Context, p interface{}, f UnaryRecoveryHandlerFunc) error {
	if f == nil {
		return status.Errorf(codes.Internal, "%s", p)
	}
	return f(ctx, p)
}

//func unaryRecoverFrom2(ctx context.Context, p interface{}, f UnaryRecoveryHandlerFunc) error {
//	if f == nil {
//		return status.Errorf(codes.Internal, "%s", p)
//	}
//	return f(ctx, p)
//}
//

func streamRecoverFrom(stream grpc.ServerStream, p interface{}, f StreamRecoveryHandlerFunc) error {
	if f == nil {
		return status.Errorf(codes.Internal, "%s", p)
	}
	return f(stream, p)
}
