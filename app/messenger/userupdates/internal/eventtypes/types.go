package eventtypes

const (
	PayloadCodecJSON int32 = 1

	EventTypeNewMessage              int32 = 1
	EventTypeReadHistory             int32 = 2
	EventTypeUpdatePinnedMessage     int32 = 3
	EventTypeMarkDialogUnread        int32 = 4
	EventTypeScheduledMarker         int32 = 5
	EventTypeDeleteMessages          int32 = 6
	EventTypeDeleteHistory           int32 = 7
	EventTypeEditMessage             int32 = 8
	EventTypeChatParticipantsChanged int32 = 9
	EventTypeDialogPublicUpdate      int32 = 100
)

type UserEvent struct {
	UserID             int64
	Pts                int64
	PtsCount           int32
	OperationID        string
	OpType             int32
	EventType          int32
	PeerType           int32
	PeerID             int64
	CanonicalMessageID int64
	PeerSeq            int64
	ActorUserID        int64
	EventSchemaVersion int32
	EventCodec         int32
	EventPayload       []byte
	EventPayloadHash   []byte
}
