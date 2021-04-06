package dataobject

type ChannelMessageBoxesDO struct {
	Id                  int32  `db:"id"`
	SenderUserId        int32  `db:"sender_user_id"`
	ChannelId           int32  `db:"channel_id"`
	DialogId            int64  `db:"dialog_id"`
	DialogMessageId     int32  `db:"dialog_message_id"`
	ChannelMessageBoxId int32  `db:"channel_message_box_id"`
	MessageId           int64  `db:"message_id"`
	Date                int32  `db:"date"`
	Deleted             int8   `db:"deleted"`
	CreatedAt           string `db:"created_at"`
	UpdatedAt           string `db:"updated_at"`
}
