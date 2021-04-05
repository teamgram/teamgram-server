package dataobject

type MessageReactDataDO struct {
	Id            int32  `db:"id"`
	ReactDataId   int64  `db:"react_data_id"`
	ReactId       int32  `db:"react_id"`
	MessageDataID int64  `db:"message_data_id"`
	SenderUserId  int32  `db:"sender_user_id"`
	Date3         int32  `db:"date3"`
	EditDate      int32  `db:"edit_date"`
	Deleted       int8   `db:"deleted"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
