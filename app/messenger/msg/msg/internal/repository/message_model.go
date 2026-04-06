// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessageBoxDO holds the persisted message box data.
type MessageBoxDO struct {
	UserId          int64
	MessageId       int32
	SenderUserId    int64
	PeerType        int32
	PeerId          int64
	RandomId        int64
	Message         *tg.TLMessage
	Pts             int32
	PtsCount        int32
}

// MessageModel manages message persistence.
type MessageModel interface {
	PutMessage(ctx context.Context, box *MessageBoxDO) error
	GetMessage(ctx context.Context, userId int64, messageId int32) (*MessageBoxDO, error)
	GetUserMessages(ctx context.Context, userId int64, peerType int32, peerId int64, limit int) ([]*MessageBoxDO, error)
}

type inMemoryMessageModel struct {
	mu  sync.RWMutex
	// Key: userId_messageId
	messages map[string]*MessageBoxDO
}

func NewMessageModel() MessageModel {
	return &inMemoryMessageModel{
		messages: make(map[string]*MessageBoxDO),
	}
}

func messageKey(userId int64, messageId int32) string {
	return fmt.Sprintf("%d_%d", userId, messageId)
}

func (m *inMemoryMessageModel) PutMessage(ctx context.Context, box *MessageBoxDO) error {
	m.mu.Lock()
	m.messages[messageKey(box.UserId, box.MessageId)] = box
	m.mu.Unlock()
	return nil
}

func (m *inMemoryMessageModel) GetMessage(ctx context.Context, userId int64, messageId int32) (*MessageBoxDO, error) {
	m.mu.RLock()
	box, ok := m.messages[messageKey(userId, messageId)]
	m.mu.RUnlock()
	if !ok {
		return nil, nil
	}
	return box, nil
}

func (m *inMemoryMessageModel) GetUserMessages(ctx context.Context, userId int64, peerType int32, peerId int64, limit int) ([]*MessageBoxDO, error) {
	m.mu.RLock()
	var result []*MessageBoxDO
	for _, box := range m.messages {
		if box.UserId == userId && box.PeerType == peerType && box.PeerId == peerId {
			result = append(result, box)
			if len(result) >= limit {
				break
			}
		}
	}
	m.mu.RUnlock()
	return result, nil
}
