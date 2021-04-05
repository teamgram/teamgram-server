package dataobject

type ChannelMessageViewsDO struct {
	Id               int32  `db:"id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	ReaderUserId     int32  `db:"reader_user_id"`
	ViewAt           int32  `db:"view_at"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
