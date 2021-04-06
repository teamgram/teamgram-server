package dataobject

type ChannelMessageLikesDO struct {
	Id          int64  `db:"id"`
	UserId      int32  `db:"user_id"`
	PostId      string `db:"post_id"`
	CommentId   string `db:"comment_id"`
	StatusLike  int8   `db:"status_like"`
	CreatedTime int32  `db:"created_time"`
	UpdatedTime int32  `db:"updated_time"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
