package dataobject

type ChannelAdminLogsDO struct {
	Id          int64  `db:"id"`
	ChannelId   int32  `db:"channel_id"`
	AdminUserId int32  `db:"admin_user_id"`
	Event       string `db:"event"`
	EventType   int32  `db:"event_type"`
	Date        int32  `db:"date"`
	CreatedAt   string `db:"created_at"`
}
