package dataobject

type ChannelParticipantsDO struct {
	Id                        int64  `db:"id"`
	ChannelId                 int32  `db:"channel_id"`
	UserId                    int32  `db:"user_id"`
	IsCreator                 int32  `db:"is_creator"`
	ParticipantType           int8   `db:"participant_type"`
	InviterUserId             int32  `db:"inviter_user_id"`
	InvitedAt                 int32  `db:"invited_at"`
	JoinedAt                  int32  `db:"joined_at"`
	PromotedAt                int32  `db:"promoted_at"`
	IsLeft                    int32  `db:"is_left"`
	HiddenPrehistory          int8   `db:"hidden_prehistory"`
	HiddenPrehistoryMessageId int32  `db:"hidden_prehistory_message_id"`
	LeftAt                    int32  `db:"left_at"`
	IsKicked                  int8   `db:"is_kicked"`
	KickedBy                  int32  `db:"kicked_by"`
	KickedAt                  int32  `db:"kicked_at"`
	BannedRights              int32  `db:"banned_rights"`
	BannedUntilDate           int32  `db:"banned_until_date"`
	BannedAt                  int32  `db:"banned_at"`
	ReadInboxMaxId            int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId           int32  `db:"read_outbox_max_id"`
	Date                      int32  `db:"date"`
	State                     int32  `db:"state"`
	CreatedAt                 string `db:"created_at"`
	UpdatedAt                 string `db:"updated_at"`
}
