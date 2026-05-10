package payload

const (
	PayloadCodecJSON  int32 = 1
	OpTypeSendMessage int32 = 1

	BucketCount            = 4096
	ReceiverPartitionCount = 256
	PushPartitionCount     = 256

	MaxOperationIDLength = 160

	MessageOperationSchemaVersionV1  = 1
	OperationResponseSchemaVersionV1 = 1
	MessageEventSchemaVersionV1      = 1

	MessageOperationSchemaVersion  = 2
	OperationResponseSchemaVersion = 2
	MessageEventSchemaVersion      = 2

	MediaRefSchemaVersionV1         = 1
	MediaRefSchemaVersionV2         = 2
	MessageAttrsSchemaVersionV1     = 1
	ForwardRefSchemaVersionV1       = 1
	MessageOperationSchemaVersionV3 = 3
	MessageEventSchemaVersionV3     = 3
)

const (
	PeerTypeUser    int32 = 1
	PeerTypeChat    int32 = 2
	PeerTypeChannel int32 = 3
)

const (
	OperationKindSendMessage         = "send_message"
	OperationKindReadHistory         = "read_history"
	OperationKindDeleteMessages      = "delete_messages"
	OperationKindDeleteHistory       = "delete_history"
	OperationKindEditMessage         = "edit_message"
	OperationKindUpdatePinnedMessage = "update_pinned_message"
	OperationKindScheduledMarker     = "scheduled_marker"
	OperationKindMarkDialogUnread    = "mark_dialog_unread"
	EventKindNewMessage              = "new_message"
)
