/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sqlx.DB
var _ *sqlx.Tx

type CanonicalMessageRow struct {
	SendStateID        int64  `db:"send_state_id"`
	CanonicalMessageID int64  `db:"canonical_message_id"`
	PeerSeq            int64  `db:"peer_seq"`
	MessageDate        int64  `db:"message_date"`
	RequestPayloadHash []byte `db:"request_payload_hash"`
}

type HistoryMessageRow struct {
	CanonicalMessageID int64  `db:"canonical_message_id"`
	PeerSeq            int64  `db:"peer_seq"`
	UserMessageID      int64  `db:"user_message_id"`
	FromUserID         int64  `db:"from_user_id"`
	Outgoing           bool   `db:"outgoing"`
	PeerType           int32  `db:"peer_type"`
	PeerID             int64  `db:"peer_id"`
	MessageKind        int32  `db:"message_kind"`
	MessageText        string `db:"message_text"`
	MessageDate        int64  `db:"message_date"`
	ViewPayload        []byte `db:"view_payload"`
}

type ResolvedMessageIDRow struct {
	UserID             int64 `db:"user_id"`
	PeerType           int32 `db:"peer_type"`
	PeerID             int64 `db:"peer_id"`
	UserMessageID      int64 `db:"user_message_id"`
	PeerSeq            int64 `db:"peer_seq"`
	CanonicalMessageID int64 `db:"canonical_message_id"`
	MessageDate        int64 `db:"message_date"`
	Outgoing           bool  `db:"outgoing"`
}

type EditableMessageRow struct {
	CanonicalMessageID int64  `db:"canonical_message_id"`
	PeerSeq            int64  `db:"peer_seq"`
	FromUserID         int64  `db:"from_user_id"`
	PeerType           int32  `db:"peer_type"`
	PeerID             int64  `db:"peer_id"`
	MessageKind        int32  `db:"message_kind"`
	MessageText        string `db:"message_text"`
	MessageDate        int64  `db:"message_date"`
	EditDate           int64  `db:"edit_date"`
	EditVersion        int32  `db:"edit_version"`
}

type HistoryOffsetRow struct {
	OffsetFromID int64 `db:"offset_from_id"`
}

type PeerSeqFloorRow struct {
	PeerSeqFloor int64 `db:"peer_seq_floor"`
}

type CanonicalQueriesModel interface {
	SelectCanonicalByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error)
	SelectCanonicalByID(ctx context.Context, sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error)
	SelectUserMessageByID(ctx context.Context, userId int64, peerType int32, peerId int64, userMessageId int64, messageStatus int32) (*ResolvedMessageIDRow, error)
	SelectNearestLiveUserMessageByPeerSeq(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*ResolvedMessageIDRow, error)
	SelectHistoryMessages(ctx context.Context, userId int64, peerType int32, peerId int64, messageStatus int32, minPeerSeq int64, maxPeerSeq int64, limit int32) ([]HistoryMessageRow, error)
	SearchHashTagMessages(ctx context.Context, hashTag string, userId int64, peerType int32, peerId int64, messageStatus int32, offsetId int64, offsetIdLimit int64, likeTag string, limit int32) ([]HistoryMessageRow, error)
	SelectCanonicalByUserView(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*HistoryMessageRow, error)
	SelectEditableMessageForUpdate(ctx context.Context, actorUserId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*EditableMessageRow, error)
	CountHistoryOffset(ctx context.Context, curUserId int64, curPeerType int64, curPeerId int64, curPeerSeq int64, userId int64, peerType int32, peerId int64, messageStatus int32) (*HistoryOffsetRow, error)
	SelectHistoryMessagesPage(ctx context.Context, userId int64, peerType int32, peerId int64, messageStatus int32, offset int64, limit int32) ([]HistoryMessageRow, error)
	SelectConversationViewPeerSeqFloor(ctx context.Context, peerType int32, userId int64, peerId int64, otherUserId int64, otherPeerId int64) (*PeerSeqFloorRow, error)
}

type CanonicalQueriesTxModel interface {
	SelectCanonicalByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error)
	SelectCanonicalByID(sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error)
	SelectUserMessageByID(userId int64, peerType int32, peerId int64, userMessageId int64, messageStatus int32) (*ResolvedMessageIDRow, error)
	SelectNearestLiveUserMessageByPeerSeq(userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*ResolvedMessageIDRow, error)
	SelectHistoryMessages(userId int64, peerType int32, peerId int64, messageStatus int32, minPeerSeq int64, maxPeerSeq int64, limit int32) ([]HistoryMessageRow, error)
	SearchHashTagMessages(hashTag string, userId int64, peerType int32, peerId int64, messageStatus int32, offsetId int64, offsetIdLimit int64, likeTag string, limit int32) ([]HistoryMessageRow, error)
	SelectCanonicalByUserView(userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*HistoryMessageRow, error)
	SelectEditableMessageForUpdate(actorUserId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*EditableMessageRow, error)
	CountHistoryOffset(curUserId int64, curPeerType int64, curPeerId int64, curPeerSeq int64, userId int64, peerType int32, peerId int64, messageStatus int32) (*HistoryOffsetRow, error)
	SelectHistoryMessagesPage(userId int64, peerType int32, peerId int64, messageStatus int32, offset int64, limit int32) ([]HistoryMessageRow, error)
	SelectConversationViewPeerSeqFloor(peerType int32, userId int64, peerId int64, otherUserId int64, otherPeerId int64) (*PeerSeqFloorRow, error)
}

type defaultCanonicalQueriesModel struct {
	db *sqlx.DB
}

func NewCanonicalQueriesModel(db *sqlx.DB) CanonicalQueriesModel {
	return &defaultCanonicalQueriesModel{db: db}
}

type defaultCanonicalQueriesTxModel struct {
	tx *sqlx.Tx
}

func NewCanonicalQueriesTxModel(tx *sqlx.Tx) CanonicalQueriesTxModel {
	return &defaultCanonicalQueriesTxModel{tx: tx}
}

func (m *defaultCanonicalQueriesModel) SelectCanonicalByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select r.send_state_id, r.canonical_message_id, c.peer_seq, c.`date` as message_date, r.request_payload_hash from message_client_randoms as r join canonical_messages as c on c.canonical_message_id = r.canonical_message_id where r.sender_user_id = ? and r.peer_type = ? and r.peer_id = ? and r.client_random_id = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, senderUserId, peerType, peerId, clientRandomId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectCanonicalByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select r.send_state_id, r.canonical_message_id, c.peer_seq, c.`date` as message_date, r.request_payload_hash from message_client_randoms as r join canonical_messages as c on c.canonical_message_id = r.canonical_message_id where r.sender_user_id = ? and r.peer_type = ? and r.peer_id = ? and r.client_random_id = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, senderUserId, peerType, peerId, clientRandomId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectCanonicalByID(ctx context.Context, sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select ? as send_state_id, canonical_message_id, peer_seq, `date` as message_date, ? as request_payload_hash from canonical_messages where canonical_message_id = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, sendStateId, requestPayloadHash, canonicalMessageId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectCanonicalByID(sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select ? as send_state_id, canonical_message_id, peer_seq, `date` as message_date, ? as request_payload_hash from canonical_messages where canonical_message_id = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, sendStateId, requestPayloadHash, canonicalMessageId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectUserMessageByID(ctx context.Context, userId int64, peerType int32, peerId int64, userMessageId int64, messageStatus int32) (*ResolvedMessageIDRow, error) {
	var rValue ResolvedMessageIDRow
	query := "select user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, `date` as message_date, outgoing from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and user_message_id = ? and message_status = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId, peerType, peerId, userMessageId, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectUserMessageByID(userId int64, peerType int32, peerId int64, userMessageId int64, messageStatus int32) (*ResolvedMessageIDRow, error) {
	var rValue ResolvedMessageIDRow
	query := "select user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, `date` as message_date, outgoing from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and user_message_id = ? and message_status = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, userId, peerType, peerId, userMessageId, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectNearestLiveUserMessageByPeerSeq(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*ResolvedMessageIDRow, error) {
	var rValue ResolvedMessageIDRow
	query := "select user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, `date` as message_date, outgoing from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and peer_seq <= ? and message_status = ? order by peer_seq desc limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectNearestLiveUserMessageByPeerSeq(userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*ResolvedMessageIDRow, error) {
	var rValue ResolvedMessageIDRow
	query := "select user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, `date` as message_date, outgoing from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and peer_seq <= ? and message_status = ? order by peer_seq desc limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, userId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectHistoryMessages(ctx context.Context, userId int64, peerType int32, peerId int64, messageStatus int32, minPeerSeq int64, maxPeerSeq int64, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and v.peer_seq > ? and v.peer_seq < ? order by v.`date` desc, v.peer_seq desc limit ?"

	err := m.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, messageStatus, minPeerSeq, maxPeerSeq, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectHistoryMessages(userId int64, peerType int32, peerId int64, messageStatus int32, minPeerSeq int64, maxPeerSeq int64, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and v.peer_seq > ? and v.peer_seq < ? order by v.`date` desc, v.peer_seq desc limit ?"

	err := m.tx.QueryRowsPartial(&rList, query, userId, peerType, peerId, messageStatus, minPeerSeq, maxPeerSeq, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesModel) SearchHashTagMessages(ctx context.Context, hashTag string, userId int64, peerType int32, peerId int64, messageStatus int32, offsetId int64, offsetIdLimit int64, likeTag string, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select distinct v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id left join hash_tags as h on h.user_id = v.user_id and h.peer_type = v.peer_type and h.peer_id = v.peer_id and h.hash_tag_user_message_id = v.user_message_id and h.hash_tag = ? and h.deleted = 0 where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and (? <= 0 or v.peer_seq < ?) and (h.hash_tag_user_message_id is not null or c.message_text like ?) order by v.peer_seq desc limit ?"

	err := m.db.QueryRowsPartial(ctx, &rList, query, hashTag, userId, peerType, peerId, messageStatus, offsetId, offsetIdLimit, likeTag, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesTxModel) SearchHashTagMessages(hashTag string, userId int64, peerType int32, peerId int64, messageStatus int32, offsetId int64, offsetIdLimit int64, likeTag string, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select distinct v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id left join hash_tags as h on h.user_id = v.user_id and h.peer_type = v.peer_type and h.peer_id = v.peer_id and h.hash_tag_user_message_id = v.user_message_id and h.hash_tag = ? and h.deleted = 0 where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and (? <= 0 or v.peer_seq < ?) and (h.hash_tag_user_message_id is not null or c.message_text like ?) order by v.peer_seq desc limit ?"

	err := m.tx.QueryRowsPartial(&rList, query, hashTag, userId, peerType, peerId, messageStatus, offsetId, offsetIdLimit, likeTag, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesModel) SelectCanonicalByUserView(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*HistoryMessageRow, error) {
	var rValue HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.peer_seq = ? and v.message_status = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectCanonicalByUserView(userId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*HistoryMessageRow, error) {
	var rValue HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.peer_seq = ? and v.message_status = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, userId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectEditableMessageForUpdate(ctx context.Context, actorUserId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*EditableMessageRow, error) {
	var rValue EditableMessageRow
	query := "select c.canonical_message_id, c.peer_seq, c.from_user_id, c.peer_type, c.peer_id, c.message_kind, c.message_text, c.`date` as message_date, COALESCE(c.edit_date, c.`date`) as edit_date, c.edit_version from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.peer_seq = ? and v.message_status = ? limit 1 for update"

	err := m.db.QueryRowPartial(ctx, &rValue, query, actorUserId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectEditableMessageForUpdate(actorUserId int64, peerType int32, peerId int64, peerSeq int64, messageStatus int32) (*EditableMessageRow, error) {
	var rValue EditableMessageRow
	query := "select c.canonical_message_id, c.peer_seq, c.from_user_id, c.peer_type, c.peer_id, c.message_kind, c.message_text, c.`date` as message_date, COALESCE(c.edit_date, c.`date`) as edit_date, c.edit_version from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.peer_seq = ? and v.message_status = ? limit 1 for update"

	err := m.tx.QueryRowPartial(&rValue, query, actorUserId, peerType, peerId, peerSeq, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) CountHistoryOffset(ctx context.Context, curUserId int64, curPeerType int64, curPeerId int64, curPeerSeq int64, userId int64, peerType int32, peerId int64, messageStatus int32) (*HistoryOffsetRow, error) {
	var rValue HistoryOffsetRow
	query := "select COUNT(*) as offset_from_id from user_message_views as v join user_message_views as cur on cur.user_id = ? and cur.peer_type = ? and cur.peer_id = ? and cur.peer_seq = ? where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and (v.`date` > cur.`date` or (v.`date` = cur.`date` and v.peer_seq >= cur.peer_seq))"

	err := m.db.QueryRowPartial(ctx, &rValue, query, curUserId, curPeerType, curPeerId, curPeerSeq, userId, peerType, peerId, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) CountHistoryOffset(curUserId int64, curPeerType int64, curPeerId int64, curPeerSeq int64, userId int64, peerType int32, peerId int64, messageStatus int32) (*HistoryOffsetRow, error) {
	var rValue HistoryOffsetRow
	query := "select COUNT(*) as offset_from_id from user_message_views as v join user_message_views as cur on cur.user_id = ? and cur.peer_type = ? and cur.peer_id = ? and cur.peer_seq = ? where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? and (v.`date` > cur.`date` or (v.`date` = cur.`date` and v.peer_seq >= cur.peer_seq))"

	err := m.tx.QueryRowPartial(&rValue, query, curUserId, curPeerType, curPeerId, curPeerSeq, userId, peerType, peerId, messageStatus)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectHistoryMessagesPage(ctx context.Context, userId int64, peerType int32, peerId int64, messageStatus int32, offset int64, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? order by v.`date` desc, v.peer_seq desc limit ?, ?"

	err := m.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, messageStatus, offset, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectHistoryMessagesPage(userId int64, peerType int32, peerId int64, messageStatus int32, offset int64, limit int32) ([]HistoryMessageRow, error) {
	var rList []HistoryMessageRow
	query := "select v.canonical_message_id, v.peer_seq, v.user_message_id, c.from_user_id, v.outgoing, v.peer_type, v.peer_id, v.message_kind, c.message_text, v.`date` as message_date, v.view_payload from user_message_views as v join canonical_messages as c on c.canonical_message_id = v.canonical_message_id where v.user_id = ? and v.peer_type = ? and v.peer_id = ? and v.message_status = ? order by v.`date` desc, v.peer_seq desc limit ?, ?"

	err := m.tx.QueryRowsPartial(&rList, query, userId, peerType, peerId, messageStatus, offset, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultCanonicalQueriesModel) SelectConversationViewPeerSeqFloor(ctx context.Context, peerType int32, userId int64, peerId int64, otherUserId int64, otherPeerId int64) (*PeerSeqFloorRow, error) {
	var rValue PeerSeqFloorRow
	query := "select COALESCE(MAX(peer_seq), 0) + 1 as peer_seq_floor from user_message_views where peer_type = ? and ((user_id = ? and peer_id = ?) or (user_id = ? and peer_id = ?))"

	err := m.db.QueryRowPartial(ctx, &rValue, query, peerType, userId, peerId, otherUserId, otherPeerId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectConversationViewPeerSeqFloor(peerType int32, userId int64, peerId int64, otherUserId int64, otherPeerId int64) (*PeerSeqFloorRow, error) {
	var rValue PeerSeqFloorRow
	query := "select COALESCE(MAX(peer_seq), 0) + 1 as peer_seq_floor from user_message_views where peer_type = ? and ((user_id = ? and peer_id = ?) or (user_id = ? and peer_id = ?))"

	err := m.tx.QueryRowPartial(&rValue, query, peerType, userId, peerId, otherUserId, otherPeerId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}
