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

package grpc_recovery2

var (
	defaultOptions = &options{
		unaryRecoveryHandlerFunc:  nil,
		unaryRecoveryHandlerFunc2: nil,
		streamRecoveryHandlerFunc: nil,
	}
)

type options struct {
	unaryRecoveryHandlerFunc  UnaryRecoveryHandlerFunc
	unaryRecoveryHandlerFunc2 UnaryRecoveryHandlerFunc
	streamRecoveryHandlerFunc StreamRecoveryHandlerFunc
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithUnaryRecoveryHandler customizes the function for recovering from a panic.
func WithUnaryRecoveryHandler(f UnaryRecoveryHandlerFunc) Option {
	return func(o *options) {
		o.unaryRecoveryHandlerFunc = f
	}
}

func WithUnaryRecoveryHandler2(f UnaryRecoveryHandlerFunc) Option {
	return func(o *options) {
		o.unaryRecoveryHandlerFunc2 = f
	}
}

// WithStreamRecoveryHandler customizes the function for recovering from a panic.
func WithStreamRecoveryHandler(f StreamRecoveryHandlerFunc) Option {
	return func(o *options) {
		o.streamRecoveryHandlerFunc = f
	}
}
