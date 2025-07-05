// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package dao

import (
	"context"

	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"

	_ "github.com/teamgram/teamgram-server/v2/app/bff/account/account/accountservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/authorization/authorization/authorizationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/autodownload/autodownload/autodownloadservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/chatinvites/chatinvitesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/chats/chats/chatsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/contacts/contacts/contactsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/dialogs/dialogs/dialogsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/drafts/drafts/draftsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/files/files/filesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/messages/messages/messagesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/notification/notification/notificationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/nsfw/nsfw/nsfwservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/passport/passport/passportservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/premium/premium/premiumservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/privacysettings/privacysettingsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/tos/tos/tosservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/updates/updates/updatesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/usernames/usernames/usernamesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/userprofile/userprofile/userprofileservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/users/users/usersservice"
)

func (d *Dao) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object iface.TLObject) (iface.TLObject, error) {
	return d.BFFProxyClient2.InvokeContext(ctx, rpcMetaData, object)
}
