// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"bytes"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

// UserupdatesAppendDialogAuthSeqSideEffect
// userupdates.appendDialogAuthSeqSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserAuthSeqAppendResult;
func (c *UserupdatesCore) UserupdatesAppendDialogAuthSeqSideEffect(in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: missing append auth seq side effect request", userupdates.ErrOperationTerminal)
	}
	if err := validateDialogSideEffectAppend(dialogSideEffectAppendInput{
		UserID:              in.UserId,
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		TargetAuthPolicy:    in.TargetAuthPolicy,
		Payload:             in.Payload,
		PayloadHash:         in.PayloadHash,
	}); err != nil {
		return nil, err
	}

	return nil, userupdates.ErrAuthSeqLedgerUnavailable
}

type dialogSideEffectAppendInput struct {
	UserID              int64
	SourcePermAuthKeyID int64
	OperationID         string
	TargetAuthPolicy    string
	Payload             []byte
	PayloadHash         []byte
}

func validateDialogSideEffectAppend(in dialogSideEffectAppendInput) error {
	if in.UserID == 0 {
		return fmt.Errorf("%w: missing user_id", userupdates.ErrOperationTerminal)
	}
	if in.OperationID == "" {
		return fmt.Errorf("%w: missing operation_id", userupdates.ErrOperationTerminal)
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return userupdates.ErrOperationPayloadConflict
	}
	if in.TargetAuthPolicy == "not_source_perm_auth_key" && in.SourcePermAuthKeyID == 0 {
		return fmt.Errorf("%w: missing source_perm_auth_key_id", userupdates.ErrOperationTerminal)
	}
	return nil
}
