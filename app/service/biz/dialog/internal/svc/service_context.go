// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package svc

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

type DialogRepository interface {
	ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]repository.DialogRecord, error)
	ListPinnedDialogs(ctx context.Context, userID int64, folderID int32) ([]repository.DialogRecord, error)
	CountDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) (int32, error)
	GetDialogByPeer(ctx context.Context, userID int64, peerType int32, peerID int64) (*repository.DialogRecord, error)
	ListDialogsByPeerDialogIDs(ctx context.Context, userID int64, ids []int64) ([]repository.DialogRecord, error)
	SaveDraft(ctx context.Context, in repository.SaveDraftInput) (*repository.DraftMutationResult, error)
	ClearDraft(ctx context.Context, in repository.ClearDraftInput) (*repository.DraftMutationResult, error)
	ClearDraftAfterSend(ctx context.Context, in repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error)
	ClearAllDrafts(ctx context.Context, in repository.ClearAllDraftsInput) ([]repository.DraftMutationResult, error)
	ListActiveDrafts(ctx context.Context, userID int64) ([]repository.DraftRecord, error)
	ToggleDialogPin(ctx context.Context, in repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error)
	ReorderPinnedDialogs(ctx context.Context, in repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error)
	EditPeerFolders(ctx context.Context, in repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error)
}

type ServiceContext struct {
	Config config.Config
	Repo   DialogRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repository.NewRepository(c),
	}
}
