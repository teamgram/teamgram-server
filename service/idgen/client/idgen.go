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

package idgen

import "fmt"

///////////////////////////////////////////////////////////////////////////////////////////
type UUIDGen interface {
	Initialize(config string) error
	GetUUID() (int64, error)
}

type SeqIDGen interface {
	Initialize(config string) error
	GetCurrentSeqID(key string) (int64, error)
	GetNextSeqID(key string) (int64, error)
	GetNextNSeqID(key string, n int) (seq int64, err error)
}

type UUIDGenInstance func() UUIDGen

var uuidGenAdapters = make(map[string]UUIDGenInstance)

func UUIDGenRegister(name string, adapter UUIDGenInstance) {
	if adapter == nil {
		panic("uuidgen: Register adapter is nil")
	}
	if _, ok := uuidGenAdapters[name]; ok {
		panic("uuidgen: Register called twice for adapter " + name)
	}
	uuidGenAdapters[name] = adapter
}

func NewUUIDGen(adapterName, config string) (adapter UUIDGen, err error) {
	instanceFunc, ok := uuidGenAdapters[adapterName]
	if !ok {
		err = fmt.Errorf("uuidgen: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.Initialize(config)
	if err != nil {
		adapter = nil
	}
	return
}

///////////////////////////////////////////////////////////////////////////////////////////
type SeqIDGenInstance func() SeqIDGen

var seqIDGenAdapters = make(map[string]SeqIDGenInstance)

func SeqIDGenRegister(name string, adapter SeqIDGenInstance) {
	if adapter == nil {
		panic("seqidgen: Register adapter is nil")
	}
	if _, ok := seqIDGenAdapters[name]; ok {
		panic("seqidgen: Register called twice for adapter " + name)
	}
	seqIDGenAdapters[name] = adapter
}

func NewSeqIDGen(adapterName, config string) (adapter SeqIDGen, err error) {
	instanceFunc, ok := seqIDGenAdapters[adapterName]
	if !ok {
		err = fmt.Errorf("seqidgen: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.Initialize(config)
	if err != nil {
		adapter = nil
	}
	return
}
