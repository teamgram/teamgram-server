// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package dao

import (
	"context"

	"github.com/teamgooo/teamgooo-server/pkg/net/kitex/metadata"
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"

	_ "github.com/teamgooo/teamgooo-server/app/bff/account/account/accountservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/authorization/authorization/authorizationservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/autodownload/autodownload/autodownloadservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/chatinvites/chatinvites/chatinvitesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/chats/chats/chatsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/configuration/configuration/configurationservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/contacts/contacts/contactsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/dialogs/dialogs/dialogsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/drafts/drafts/draftsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/files/files/filesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/messages/messages/messagesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/notification/notification/notificationservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/nsfw/nsfw/nsfwservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/passport/passport/passportservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/premium/premium/premiumservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/privacysettings/privacysettings/privacysettingsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/qrcode/qrcode/qrcodeservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/tos/tos/tosservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/updates/updates/updatesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/userchannelprofiles/userchannelprofiles/userchannelprofilesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/usernames/usernames/usernamesservice"
	_ "github.com/teamgooo/teamgooo-server/app/bff/users/users/usersservice"
)

func (d *Dao) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object iface.TLObject) (iface.TLObject, error) {
	return d.BFFProxyClient2.InvokeContext(ctx, rpcMetaData, object)
}
