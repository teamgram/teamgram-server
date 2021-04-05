package dataobject

type ChannelsDO struct {
	Id               int32  `db:"id"`
	CreatorUserId    int32  `db:"creator_user_id"`
	AccessHash       int64  `db:"access_hash"`
	RandomId         int64  `db:"random_id"`
	Type             int32  `db:"type"`
	TopMessage       int32  `db:"top_message"`
	ParticipantCount int32  `db:"participant_count"`
	Title            string `db:"title"`
	About            string `db:"about"`
	PhotoId          int64  `db:"photo_id"`
	Public           int32  `db:"public"`
	Link             string `db:"link"`
	BroadCast        int8   `db:"broad_cast"`
	Verified         int8   `db:"verified"`
	MegaGroup        int8   `db:"mega_group"`
	DemoCracy        int8   `db:"demo_cracy"`
	Signatures       int8   `db:"signatures"`
	AdminsEnabled    int8   `db:"admins_enabled"`
	Deactivated      int8   `db:"deactivated"`
	Version          int32  `db:"version"`
	Date             int32  `db:"date"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_ad"`
}
