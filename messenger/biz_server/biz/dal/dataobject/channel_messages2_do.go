package dataobject

type ChannelMessage2DO struct {
	Id                  int32  `db:"id"`
	ChannelId           int32  `db:"channel_id"`
	SenderUserId        int32  `db:"sender_user_id"`
	ChannelMessageBoxId int32  `db:"channel_message_box_id"`
	DialogMessageId     int64  `db:"dialog_message_id"`
	RandomId            int64  `db:"random_id"`
	MessageType         int8   `db:"message_type"`
	MessageData         string `db:"message_data"`
	Date2               int32  `db:"date2"`
	Deleted             int8   `db:"deleted"`
	CreatedAt           string `db:"created_at"`
	UpdatedAt           string `db:"updated_at"`
}
