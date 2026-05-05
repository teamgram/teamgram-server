package payload

const (
	PayloadCodecJSON  int32 = 1
	OpTypeSendMessage int32 = 1

	BucketCount            = 4096
	ReceiverPartitionCount = 256
	PushPartitionCount     = 256

	MaxOperationIDLength = 160

	MessageOperationSchemaVersion  = 1
	OperationResponseSchemaVersion = 1
	MessageEventSchemaVersion      = 1
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
