package repository

import (
	"context"
	"errors"
	"time"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type ExportChatInviteArg struct {
	ChatID        int64
	AdminID       int64
	RequestNeeded bool
	ExpireDate    *int32
	UsageLimit    *int32
	Title         *string
}

type EditExportedChatInviteArg struct {
	ChatID        int64
	Link          string
	Revoked       bool
	ExpireDate    *int32
	UsageLimit    *int32
	RequestNeeded *bool
	Title         *string
}

type InviteParticipantArg struct {
	ChatID     int64
	Link       string
	UserID     int64
	Requested  bool
	ApprovedBy int64
}

type ChatInviteImporterQuery struct {
	ChatID     int64
	Link       string
	Requested  bool
	OffsetDate int32
	OffsetUser int64
	Limit      int32
}

type HideJoinRequestsArg struct {
	ChatID   int64
	UserID   int64
	Approver int64
	Approve  bool
}

type JoinRequest struct {
	ChatID int64
	Link   string
	UserID int64
	Date   int64
}

func (r *Repository) CreateExportedChatInvite(ctx context.Context, arg ExportChatInviteArg) (*tg.ExportedChatInvite, error) {
	now := time.Now().Unix()
	row := &model.ChatInvites{
		ChatId:        arg.ChatID,
		AdminId:       arg.AdminID,
		Link:          chatpb.GenChatInviteHash(),
		RequestNeeded: arg.RequestNeeded,
		Date2:         now,
	}
	if arg.ExpireDate != nil {
		row.ExpireDate = int64(*arg.ExpireDate)
	}
	if arg.UsageLimit != nil {
		row.UsageLimit = *arg.UsageLimit
	}
	if arg.Title != nil {
		row.Title = *arg.Title
	}
	if _, _, err := r.model.ChatInvitesModel.Insert(ctx, row); err != nil {
		return nil, wrapStorage("chat_invites.Insert", err)
	}
	return r.MakeChatInviteExported(ctx, row)
}

func (r *Repository) GetExportedChatInvite(ctx context.Context, chatID int64, link string) (*tg.ExportedChatInvite, error) {
	row, err := r.GetChatInviteByLink(ctx, link)
	if err != nil {
		return nil, err
	}
	if err := requireInviteRowForChat(row, chatID); err != nil {
		return nil, err
	}
	return r.MakeChatInviteExported(ctx, row)
}

func (r *Repository) GetExportedChatInvites(ctx context.Context, chatID, adminID int64, revoked bool, offsetDate *int32, offsetLink *string, limit int32) ([]tg.ExportedChatInviteClazz, error) {
	if limit == 0 {
		limit = 50
	}
	rows, err := r.model.ChatInvitesModel.SelectListByAdminId(ctx, chatID, adminID)
	if err != nil {
		return nil, wrapStorage("chat_invites.SelectListByAdminId", err)
	}

	all := make([]tg.ExportedChatInviteClazz, 0, len(rows))
	for i := range rows {
		if rows[i].Revoked != revoked {
			continue
		}
		invite, err := r.MakeChatInviteExported(ctx, &rows[i])
		if err != nil {
			return nil, err
		}
		all = append(all, invite.Clazz)
	}

	offset := 0
	if offsetDate != nil && offsetLink != nil && *offsetDate != 0 && *offsetLink != "" {
		offset = exportedInviteOffset(all, *offsetDate, *offsetLink)
	}
	if offset == -1 {
		return []tg.ExportedChatInviteClazz{}, nil
	}
	end := offset + int(limit)
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], nil
}

func (r *Repository) EditExportedChatInvite(ctx context.Context, arg EditExportedChatInviteArg) ([]tg.ExportedChatInviteClazz, error) {
	hash := chatpb.NormalizeInviteHash(arg.Link)
	if hash == "" {
		return nil, chatpb.ErrInviteHashInvalid
	}
	row, err := r.GetChatInviteByLink(ctx, hash)
	if err != nil {
		return nil, err
	}
	if err := requireInviteRowForChat(row, arg.ChatID); err != nil {
		return nil, err
	}

	out := make([]tg.ExportedChatInviteClazz, 0, 2)
	if arg.Revoked {
		usage, requested, err := r.chatInviteCounts(ctx, row.Link)
		if err != nil {
			return nil, err
		}
		oldRow := *row
		oldRow.Revoked = true
		if row.Revoked {
			return append(out, makeChatInviteExported(&oldRow, usage, requested).Clazz), nil
		}
		var newRow *model.ChatInvites
		if row.Permanent {
			newRow = &model.ChatInvites{
				ChatId:    arg.ChatID,
				AdminId:   row.AdminId,
				Link:      chatpb.GenChatInviteHash(),
				Permanent: true,
				Date2:     time.Now().Unix(),
			}
		}
		if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
			txModel := r.model.WithTx(tx)
			if _, err := updateChatInviteTx(r, tx, arg.ChatID, hash, chatInviteUpdate{Revoked: boolPtr(true)}); err != nil {
				return err
			}
			if newRow == nil {
				return nil
			}
			if _, _, err := txModel.ChatInvitesModel.Insert(newRow); err != nil {
				return err
			}
			_, err := txModel.ChatParticipantsModel.UpdateLink(newRow.Link, arg.ChatID, newRow.AdminId)
			return err
		}); err != nil {
			return nil, wrapStorage("chat_invites.Edit revoked transaction", err)
		}
		out = append(out, makeChatInviteExported(&oldRow, usage, requested).Clazz)
		if newRow != nil {
			out = append(out, makeChatInviteExported(newRow, 0, 0).Clazz)
		}
		return out, nil
	}

	update := chatInviteUpdate{
		ExpireDate:    arg.ExpireDate,
		UsageLimit:    arg.UsageLimit,
		RequestNeeded: arg.RequestNeeded,
		Title:         arg.Title,
	}
	if _, err := updateChatInvite(ctx, r, arg.ChatID, hash, update); err != nil {
		return nil, err
	}
	applyChatInviteUpdate(row, update)
	invite, err := r.MakeChatInviteExported(ctx, row)
	if err != nil {
		return nil, err
	}
	return append(out, invite.Clazz), nil
}

func (r *Repository) DeleteExportedChatInvite(ctx context.Context, chatID int64, link string) error {
	hash := chatpb.NormalizeInviteHash(link)
	if hash == "" {
		return chatpb.ErrInviteHashInvalid
	}
	row, err := r.GetChatInviteByLink(ctx, hash)
	if err != nil {
		return err
	}
	if err := requireInviteRowForChat(row, chatID); err != nil {
		return err
	}
	if _, err := r.model.ChatInvitesModel.DeleteByLink(ctx, chatID, hash); err != nil {
		return wrapStorage("chat_invites.DeleteByLink", err)
	}
	return nil
}

func (r *Repository) DeleteRevokedExportedChatInvites(ctx context.Context, chatID, adminID int64) error {
	if _, err := r.model.ChatInvitesModel.DeleteByRevoked(ctx, chatID, adminID); err != nil {
		return wrapStorage("chat_invites.DeleteByRevoked", err)
	}
	return nil
}

func (r *Repository) GetChatInviteByLink(ctx context.Context, link string) (*model.ChatInvites, error) {
	hash := chatpb.NormalizeInviteHash(link)
	if hash == "" {
		return nil, chatpb.ErrInviteHashInvalid
	}
	row, err := r.model.ChatInvitesModel.SelectByLink(ctx, hash)
	if err != nil {
		if isNotFound(err) {
			return nil, chatpb.ErrInviteHashInvalid
		}
		return nil, wrapStorage("chat_invites.SelectByLink", err)
	}
	if row == nil {
		return nil, chatpb.ErrInviteHashInvalid
	}
	return row, nil
}

func (r *Repository) GetAdminsWithInvites(ctx context.Context, chatID int64, adminIDs []int64) ([]tg.ChatAdminWithInvitesClazz, error) {
	allowed := make(map[int64]struct{}, len(adminIDs))
	outMap := make(map[int64]*tg.TLChatAdminWithInvites, len(adminIDs))
	for _, id := range adminIDs {
		allowed[id] = struct{}{}
	}
	rows, err := r.model.ChatInvitesModel.SelectListByChatId(ctx, chatID)
	if err != nil {
		return nil, wrapStorage("chat_invites.SelectListByChatId", err)
	}
	for i := range rows {
		if _, ok := allowed[rows[i].AdminId]; !ok {
			continue
		}
		admin := outMap[rows[i].AdminId]
		if admin == nil {
			admin = tg.MakeTLChatAdminWithInvites(&tg.TLChatAdminWithInvites{AdminId: rows[i].AdminId})
			outMap[rows[i].AdminId] = admin
		}
		if rows[i].Revoked {
			admin.RevokedInvitesCount++
		} else {
			admin.InvitesCount++
		}
	}
	out := make([]tg.ChatAdminWithInvitesClazz, 0, len(outMap))
	for _, id := range adminIDs {
		if admin := outMap[id]; admin != nil {
			out = append(out, admin)
		}
	}
	return out, nil
}

func (r *Repository) MakeChatInviteExported(ctx context.Context, row *model.ChatInvites) (*tg.ExportedChatInvite, error) {
	if row == nil {
		return nil, chatpb.ErrInviteHashInvalid
	}
	usage, requested, err := r.chatInviteCounts(ctx, row.Link)
	if err != nil {
		return nil, err
	}
	return makeChatInviteExported(row, usage, requested), nil
}

func (r *Repository) chatInviteCounts(ctx context.Context, link string) (int32, int32, error) {
	usage, err := r.model.ChatInviteParticipantsModel.SelectCountByLink(ctx, link, 0)
	if err != nil {
		return 0, 0, wrapStorage("chat_invite_participants.SelectCountByLink usage", err)
	}
	requested, err := r.model.ChatInviteParticipantsModel.SelectCountByLink(ctx, link, 1)
	if err != nil {
		return 0, 0, wrapStorage("chat_invite_participants.SelectCountByLink requested", err)
	}
	return usage, requested, nil
}

func (r *Repository) RecordInviteParticipant(ctx context.Context, arg InviteParticipantArg) error {
	row := &model.ChatInviteParticipants{
		ChatId:     arg.ChatID,
		Link:       chatpb.NormalizeInviteHash(arg.Link),
		UserId:     arg.UserID,
		Requested:  arg.Requested,
		ApprovedBy: arg.ApprovedBy,
		Date2:      time.Now().Unix(),
	}
	if _, _, err := r.model.ChatInviteParticipantsModel.Insert(ctx, row); err != nil {
		return wrapStorage("chat_invite_participants.Insert", err)
	}
	return nil
}

func (r *Repository) GetChatInviteImporters(ctx context.Context, q ChatInviteImporterQuery) ([]tg.ChatInviteImporterClazz, error) {
	if q.Limit == 0 {
		q.Limit = 50
	}
	var (
		rows []model.ChatInviteParticipants
		err  error
	)
	linkRequested, useRecent := inviteImporterLinkQuery(q)
	if useRecent {
		rows, err = r.model.ChatInviteParticipantsModel.SelectRecentRequestedList(ctx, q.ChatID)
	} else {
		hash, err := r.validateInviteLinkForChat(ctx, q.ChatID, q.Link)
		if err != nil {
			return nil, err
		}
		rows, err = r.model.ChatInviteParticipantsModel.SelectListByLink(ctx, hash, linkRequested)
	}
	if err != nil {
		return nil, wrapStorage("chat_invite_participants.SelectList", err)
	}

	all := make([]tg.ChatInviteImporterClazz, 0, len(rows))
	for i := range rows {
		all = append(all, makeChatInviteImporter(&rows[i]))
	}
	offset := 0
	for i, importer := range all {
		if importer.UserId == q.OffsetUser && importer.Date == q.OffsetDate {
			offset = i + 1
			break
		}
	}
	end := offset + int(q.Limit)
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], nil
}

func (r *Repository) GetPendingJoinRequests(ctx context.Context, chatID int64, link *string) ([]JoinRequest, error) {
	var (
		rows []model.ChatInviteParticipants
		err  error
	)
	if link != nil {
		hash, err := r.validateInviteLinkForChat(ctx, chatID, *link)
		if err != nil {
			return nil, err
		}
		rows, err = r.model.ChatInviteParticipantsModel.SelectListByLink(ctx, hash, 1)
	} else {
		rows, err = r.model.ChatInviteParticipantsModel.SelectRecentRequestedList(ctx, chatID)
	}
	if err != nil {
		return nil, wrapStorage("chat_invite_participants.SelectPendingJoinRequests", err)
	}
	out := make([]JoinRequest, 0, len(rows))
	for i := range rows {
		if rows[i].ChatId != chatID {
			continue
		}
		out = append(out, JoinRequest{
			ChatID: rows[i].ChatId,
			Link:   rows[i].Link,
			UserID: rows[i].UserId,
			Date:   rows[i].Date2,
		})
	}
	return out, nil
}

func (r *Repository) GetRecentChatInviteRequesters(ctx context.Context, chatID int64) (*chatpb.RecentChatInviteRequesters, error) {
	rows, err := r.model.ChatInviteParticipantsModel.SelectRecentRequestedList(ctx, chatID)
	if err != nil {
		return nil, wrapStorage("chat_invite_participants.SelectRecentRequestedList", err)
	}
	requesters := make([]int64, 0, len(rows))
	for i := range rows {
		requesters = append(requesters, rows[i].UserId)
	}
	return chatpb.MakeTLRecentChatInviteRequesters(&chatpb.TLRecentChatInviteRequesters{
		RequestsPending:  int32(len(rows)),
		RecentRequesters: requesters,
	}).ToRecentChatInviteRequesters(), nil
}

func (r *Repository) HideChatJoinRequest(ctx context.Context, arg HideJoinRequestsArg) error {
	if arg.Approve {
		rowsAffected, err := r.model.ChatInviteParticipantsModel.UpdateApprovedBy(ctx, arg.Approver, arg.ChatID, arg.UserID)
		if err != nil {
			return wrapStorage("chat_invite_participants.UpdateApprovedBy", err)
		}
		if err := requireRowsAffected(rowsAffected); err != nil {
			return err
		}
		return nil
	}
	rowsAffected, err := r.model.ChatInviteParticipantsModel.Delete(ctx, arg.ChatID, arg.UserID)
	if err != nil {
		return wrapStorage("chat_invite_participants.Delete", err)
	}
	if err := requireRowsAffected(rowsAffected); err != nil {
		return err
	}
	return nil
}

func requireRowsAffected(rowsAffected int64) error {
	if rowsAffected == 0 {
		return chatpb.ErrUserNotParticipant
	}
	return nil
}

func wrapMutationError(op string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, chatpb.ErrUserNotParticipant) {
		return err
	}
	return wrapStorage(op, err)
}

func (r *Repository) validateInviteLinkForChat(ctx context.Context, chatID int64, link string) (string, error) {
	hash := chatpb.NormalizeInviteHash(link)
	if hash == "" {
		return "", chatpb.ErrInviteHashInvalid
	}
	row, err := r.GetChatInviteByLink(ctx, hash)
	if err != nil {
		return "", err
	}
	if err := requireInviteRowForChat(row, chatID); err != nil {
		return "", err
	}
	return hash, nil
}

func requireInviteRowForChat(row *model.ChatInvites, chatID int64) error {
	if row == nil {
		return chatpb.ErrInviteHashInvalid
	}
	if row.ChatId != chatID {
		return chatpb.ErrChatLinkExists
	}
	return nil
}

func inviteImporterLinkQuery(q ChatInviteImporterQuery) (requested int32, useRecent bool) {
	if q.Requested {
		if q.Link == "" {
			return 1, true
		}
		return 1, false
	}
	return 0, false
}

func (r *Repository) CountChatInviteParticipants(ctx context.Context, link string, requested bool) (int32, error) {
	v := int32(0)
	if requested {
		v = 1
	}
	count, err := r.model.ChatInviteParticipantsModel.SelectCountByLink(ctx, chatpb.NormalizeInviteHash(link), v)
	if err != nil {
		return 0, wrapStorage("chat_invite_participants.SelectCountByLink", err)
	}
	return count, nil
}

type chatInviteUpdate struct {
	Revoked       *bool
	ExpireDate    *int32
	UsageLimit    *int32
	RequestNeeded *bool
	Title         *string
}

func updateChatInvite(ctx context.Context, r *Repository, chatID int64, link string, update chatInviteUpdate) (int64, error) {
	cMap := makeChatInviteUpdateMap(update)
	if len(cMap) == 0 {
		return 0, nil
	}
	rows, err := r.model.ChatInvitesModel.Update(ctx, cMap, chatID, link)
	if err != nil {
		return 0, wrapStorage("chat_invites.Update", err)
	}
	return rows, nil
}

func updateChatInviteTx(r *Repository, tx *sqlx.Tx, chatID int64, link string, update chatInviteUpdate) (int64, error) {
	cMap := makeChatInviteUpdateMap(update)
	if len(cMap) == 0 {
		return 0, nil
	}
	return r.model.WithTx(tx).ChatInvitesModel.Update(cMap, chatID, link)
}

func makeChatInviteUpdateMap(update chatInviteUpdate) map[string]interface{} {
	cMap := make(map[string]interface{}, 5)
	if update.Revoked != nil {
		cMap["revoked"] = *update.Revoked
	}
	if update.ExpireDate != nil {
		cMap["expire_date"] = int64(*update.ExpireDate)
	}
	if update.UsageLimit != nil {
		cMap["usage_limit"] = *update.UsageLimit
	}
	if update.RequestNeeded != nil {
		cMap["request_needed"] = *update.RequestNeeded
	}
	if update.Title != nil {
		cMap["title"] = *update.Title
	}
	return cMap
}

func applyChatInviteUpdate(row *model.ChatInvites, update chatInviteUpdate) {
	if row == nil {
		return
	}
	if update.Revoked != nil {
		row.Revoked = *update.Revoked
	}
	if update.ExpireDate != nil {
		row.ExpireDate = int64(*update.ExpireDate)
	}
	if update.UsageLimit != nil {
		row.UsageLimit = *update.UsageLimit
	}
	if update.RequestNeeded != nil {
		row.RequestNeeded = *update.RequestNeeded
	}
	if update.Title != nil {
		row.Title = *update.Title
	}
}

func makeChatInviteExported(row *model.ChatInvites, usage, requested int32) *tg.ExportedChatInvite {
	if row == nil {
		return nil
	}
	date := int32(row.Date2)
	invite := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{
		Revoked:       row.Revoked,
		Permanent:     row.Permanent,
		RequestNeeded: row.RequestNeeded,
		Link:          chatpb.BuildInviteLink(row.Link),
		AdminId:       row.AdminId,
		Date:          date,
		Usage:         &usage,
		Requested:     &requested,
	}).ToExportedChatInvite()
	v := invite.Clazz.(*tg.TLChatInviteExported)
	if row.StartDate != 0 {
		start := int32(row.StartDate)
		v.StartDate = &start
	}
	if row.ExpireDate != 0 {
		expire := int32(row.ExpireDate)
		v.ExpireDate = &expire
	}
	if row.UsageLimit != 0 {
		limit := row.UsageLimit
		v.UsageLimit = &limit
	}
	if row.Title != "" {
		title := row.Title
		v.Title = &title
	}
	return invite
}

func makeChatInviteImporter(row *model.ChatInviteParticipants) tg.ChatInviteImporterClazz {
	if row == nil {
		return nil
	}
	importer := tg.MakeTLChatInviteImporter(&tg.TLChatInviteImporter{
		Requested: row.Requested,
		UserId:    row.UserId,
		Date:      int32(row.Date2),
	}).ToChatInviteImporter()
	if !row.Requested && row.ApprovedBy != 0 {
		importer.ApprovedBy = &row.ApprovedBy
	}
	return importer
}

func exportedInviteLink(invite tg.ExportedChatInviteClazz) string {
	if v, ok := invite.(*tg.TLChatInviteExported); ok {
		return v.Link
	}
	return ""
}

func exportedInviteDate(invite tg.ExportedChatInviteClazz) int32 {
	if v, ok := invite.(*tg.TLChatInviteExported); ok {
		return v.Date
	}
	return 0
}

func exportedInviteOffset(invites []tg.ExportedChatInviteClazz, offsetDate int32, offsetLink string) int {
	hash := chatpb.NormalizeInviteHash(offsetLink)
	for i, invite := range invites {
		if exportedInviteDate(invite) != offsetDate {
			continue
		}
		if chatpb.NormalizeInviteHash(exportedInviteLink(invite)) == hash {
			return i + 1
		}
	}
	return -1
}

func boolPtr(v bool) *bool {
	return &v
}

func IsInviteExpired(row *model.ChatInvites, usage int32, now int64) bool {
	if row == nil {
		return true
	}
	if row.Revoked {
		return true
	}
	if row.ExpireDate != 0 && now > row.ExpireDate {
		return true
	}
	return row.UsageLimit > 0 && usage >= row.UsageLimit
}

func IsInviteNotFound(err error) bool {
	return errors.Is(err, chatpb.ErrInviteHashInvalid)
}
