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

type SeqsvrClient struct {
}

func seqsvrClientInstance() SeqIDGen {
	cli := &SeqsvrClient{}
	return cli
}

func (c *SeqsvrClient) Initialize(config string) error {
	// TODO(@benqi): Impl Initialize logic
	return fmt.Errorf("not impl seqsvr.Initialize")
}

func (c *SeqsvrClient) GetCurrentSeqID(key string) (int64, error) {
	// TODO(@benqi): Impl GetCurrentSeqID logic
	return 0, fmt.Errorf("not impl seqsvr.GetCurrentSeqID")
}

func (c *SeqsvrClient) GetNextSeqID(key string) (int64, error) {
	// TODO(@benqi): Impl GetNextSeqID logic
	return 0, fmt.Errorf("not impl seqsvr.GetNextSeqID")
}


func (c *SeqsvrClient) GetNextNSeqID(key string, n int) (int64, error) {
	return 0, fmt.Errorf("not impl seqsvr.GetNextNSeqID")
}

func init() {
	SeqIDGenRegister("seqsvr", seqsvrClientInstance)
}
