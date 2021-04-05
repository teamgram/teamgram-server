package dataobject

type MessageReactDO struct {
	Id         int32  `db:"id"`
	ReactionId int64  `db:"reaction_id"`
	Text       string `db:"text"`
	FileId     int64  `db:"file_id"`
	FileHash   string `db:"file_hash"`
	FileSize   int32  `db:"file_size"`
	Width      int32  `db:"width"`
	Height     int32  `db:"height"`
}
