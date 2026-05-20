package cursor

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

type DeliveryClass string

const (
	DeliveryClassUserPTS           DeliveryClass = "user_pts"
	DeliveryClassChannelPTS        DeliveryClass = "channel_pts"
	DeliveryClassAuthSeq           DeliveryClass = "auth_seq"
	DeliveryClassRealtimeOnly      DeliveryClass = "realtime_only"
	DeliveryClassAttachedAuxiliary DeliveryClass = "attached_auxiliary"
)

func ClassifyUpdate(update tg.UpdateClazz) DeliveryClass {
	switch update.(type) {
	case *tg.TLUpdateNewMessage,
		*tg.TLUpdateDeleteMessages,
		*tg.TLUpdateReadHistoryInbox,
		*tg.TLUpdateReadHistoryOutbox,
		*tg.TLUpdateWebPage,
		*tg.TLUpdateReadMessagesContents,
		*tg.TLUpdateEditMessage,
		*tg.TLUpdateFolderPeers,
		*tg.TLUpdatePinnedMessages:
		return DeliveryClassUserPTS
	case *tg.TLUpdateChannelTooLong,
		*tg.TLUpdateNewChannelMessage,
		*tg.TLUpdateReadChannelInbox,
		*tg.TLUpdateDeleteChannelMessages,
		*tg.TLUpdateEditChannelMessage,
		*tg.TLUpdateChannelWebPage,
		*tg.TLUpdatePinnedChannelMessages:
		return DeliveryClassChannelPTS
	default:
		return DeliveryClassRealtimeOnly
	}
}
