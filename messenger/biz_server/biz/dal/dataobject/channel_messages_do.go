package dataobject

type ChannelMessagesDO struct {
	Id               int32  `db:"id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	SenderUserId     int32  `db:"sender_user_id"`
	RandomId         int64  `db:"random_id"`
	PostInfoId       int64  `db:"post_info_id"`
	MessageDataId    int64  `db:"message_data_id"`
	MessageType      int8   `db:"message_type"`
	MessageData      string `db:"message_data"`
	HasMediaUnread   int8   `db:"has_media_unread"`
	EditMessage      string `db:"edit_message"`
	EditDate         int32  `db:"edit_date"`
	Views            int32  `db:"views"`
	Date             int32  `db:"date"`
	Deleted          int8   `db:"deleted"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
