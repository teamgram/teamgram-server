package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetBotInfo(ctx context.Context, botID int64) (*tg.BotInfo, error) {
	botDO, err := r.model.BotsModel.FindOneByBotId(ctx, botID)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrBotNotFound
		}
		return nil, fmt.Errorf("%w: get bot info %d: %w", userpb.ErrUserStorage, botID, err)
	}
	return r.makeBotInfo(ctx, botDO)
}

func (r *Repository) GetBotInfoData(ctx context.Context, botID int64) (*userpb.BotInfoData, error) {
	botDO, err := r.model.BotsModel.FindOneByBotId(ctx, botID)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrBotNotFound
		}
		return nil, fmt.Errorf("%w: get bot info data %d: %w", userpb.ErrUserStorage, botID, err)
	}
	botInfo, err := r.makeBotInfo(ctx, botDO)
	if err != nil {
		return nil, err
	}
	return userpb.MakeTLBotInfoData(&userpb.TLBotInfoData{
		BotInfo:    botInfo,
		MainAppUrl: tg.MakeFlagsString(botDO.MainAppUrl),
		BotInline:  botDO.BotInlinePlaceholder != "",
		Token:      botDO.Token,
		BotId:      botDO.BotId,
	}).ToBotInfoData(), nil
}

func (r *Repository) SetBotCommands(ctx context.Context, botID int64, commands []tg.BotCommandClazz) error {
	commandDOList := make([]*model.BotCommands, 0, len(commands))
	for _, command := range commands {
		if command == nil {
			continue
		}
		commandDOList = append(commandDOList, &model.BotCommands{
			BotId:       botID,
			Command:     command.Command,
			Description: command.Description,
		})
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		if _, err := txModel.BotCommandsModel.Delete(botID); err != nil {
			return fmt.Errorf("delete bot commands: %w", err)
		}
		if len(commandDOList) > 0 {
			if _, _, err := txModel.BotCommandsModel.InsertBulk(commandDOList); err != nil {
				return fmt.Errorf("insert bot commands: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("%w: set bot commands %d: %w", userpb.ErrUserStorage, botID, err)
	}
	return nil
}

func (r *Repository) UpdateBotData(ctx context.Context, botID int64, botChatHistory, botNoChats, botInlineGeo, botAttachMenu, botHasMainApp tg.BoolClazz, botInlinePlaceholder *string) error {
	botDO, err := r.model.BotsModel.FindOneByBotId(ctx, botID)
	if err != nil {
		if isNotFound(err) {
			return userpb.ErrBotNotFound
		}
		return fmt.Errorf("%w: update bot data %d: %w", userpb.ErrUserStorage, botID, err)
	}
	if botChatHistory != nil {
		botDO.BotChatHistory = tg.FromBoolClazz(botChatHistory)
	}
	if botNoChats != nil {
		botDO.BotNochats = tg.FromBoolClazz(botNoChats)
	}
	if botInlineGeo != nil {
		botDO.BotInlineGeo = tg.FromBoolClazz(botInlineGeo)
	}
	if botAttachMenu != nil {
		botDO.BotAttachMenu = tg.FromBoolClazz(botAttachMenu)
	}
	if botHasMainApp != nil {
		botDO.BotHasMainApp = tg.FromBoolClazz(botHasMainApp)
	}
	if botInlinePlaceholder != nil {
		botDO.BotInlinePlaceholder = *botInlinePlaceholder
	}
	botDO.BotInfoVersion++
	if err := r.model.BotsModel.Update2(ctx, botDO); err != nil {
		return fmt.Errorf("%w: update bot data row %d: %w", userpb.ErrUserStorage, botID, err)
	}
	r.invalidateProjectionFactCache(ctx, botID)
	return nil
}

func (r *Repository) makeBotInfo(ctx context.Context, botDO *model.Bots) (*tg.BotInfo, error) {
	description := tg.MakeFlagsString(botDO.Description)
	botInfo := tg.MakeTLBotInfo(&tg.TLBotInfo{
		HasPreviewMedias: botDO.HasPreviewMedias,
		UserId:           tg.MakeFlagsInt64(botDO.BotId),
		Description:      description,
		Commands:         []tg.BotCommandClazz{},
		PrivacyPolicyUrl: tg.MakeFlagsString(botDO.PrivacyPolicyUrl),
	}).ToBotInfo()

	commandDOList, err := r.model.BotCommandsModel.SelectList(ctx, botDO.BotId)
	if err != nil {
		return nil, fmt.Errorf("%w: get bot commands %d: %w", userpb.ErrUserStorage, botDO.BotId, err)
	}
	for i := range commandDOList {
		botInfo.Commands = append(botInfo.Commands, tg.MakeTLBotCommand(&tg.TLBotCommand{
			Command:     commandDOList[i].Command,
			Description: commandDOList[i].Description,
		}).ToBotCommand())
	}

	if botDO.HasMenuButton {
		botInfo.MenuButton = tg.MakeTLBotMenuButton(&tg.TLBotMenuButton{
			Text: botDO.MenuButtonText,
			Url:  botDO.MenuButtonUrl,
		})
	}
	if botDO.DescriptionPhotoId != 0 && r.mediaReader != nil {
		photo, err := r.mediaReader.GetPhoto(ctx, botDO.DescriptionPhotoId)
		if err != nil {
			return nil, err
		}
		botInfo.DescriptionPhoto = photo.Clazz
	}
	if botDO.DescriptionDocumentId != 0 && r.mediaReader != nil {
		document, err := r.mediaReader.GetDocument(ctx, botDO.DescriptionDocumentId)
		if err != nil {
			return nil, err
		}
		botInfo.DescriptionDocument = document.Clazz
	}
	if botDO.HasAppSettings {
		botInfo.AppSettings = tg.MakeTLBotAppSettings(&tg.TLBotAppSettings{
			PlaceholderPath:     []byte(botDO.PlaceholderPath),
			BackgroundColor:     tg.MakeFlagsInt32(botDO.BackgroundColor),
			BackgroundDarkColor: tg.MakeFlagsInt32(botDO.BackgroundDarkColor),
			HeaderColor:         tg.MakeFlagsInt32(botDO.HeaderColor),
			HeaderDarkColor:     tg.MakeFlagsInt32(botDO.HeaderDarkColor),
		})
	}

	return botInfo, nil
}
