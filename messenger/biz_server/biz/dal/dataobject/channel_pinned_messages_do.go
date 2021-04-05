package dataobject

type ChannelPinnedMessagesDO struct {
	Id               int64  `db:"id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	AdminId          int32  `db:"admin_id"`
	Pinned           int8   `db:"pinned"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
