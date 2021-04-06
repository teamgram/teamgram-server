package dataobject

type ChannelMediaUnreadDO struct {
	Id               int64  `db:"id"`
	UserId           int32  `db:"user_id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	MediaUnread      int8   `db:"media_unread"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
