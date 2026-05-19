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

	MessageOperationSchemaVersion    = 2
	OperationResponseSchemaVersion   = 2
	MessageEventSchemaVersion        = 2
	OperationResponseSchemaVersionV3 = 3

	MediaRefSchemaVersionV1                    = 1
	MediaRefSchemaVersionV2                    = 2
	MessageAttrsSchemaVersionV1                = 1
	ForwardRefSchemaVersionV1                  = 1
	ServiceActionSchemaVersionV1         int32 = 1
	ServiceActionCodecTLBinary           int32 = 1
	ServiceActionLayer                   int32 = 224
	MessageOperationSchemaVersionV3            = 3
	MessageEventSchemaVersionV3                = 3
	MessageOperationSchemaVersionV4            = 4
	MessageEventSchemaVersionV4                = 4
	MessageOperationSchemaVersionBatchV1       = 5
	MessageEventSchemaVersionBatchV1           = 5

	ReplyEnvelopeCodecTLBinary int32 = 1
	ReplyEnvelopeSchemaV1      int32 = 1
)

const (
	PeerTypeUser    int32 = 1
	PeerTypeChat    int32 = 2
	PeerTypeChannel int32 = 3
)

const (
	OperationKindSendMessage         = "send_message"
	OperationKindSendMessageBatch    = "send_message_batch"
	OperationKindReadHistory         = "read_history"
	OperationKindDeleteMessages      = "delete_messages"
	OperationKindDeleteHistory       = "delete_history"
	OperationKindEditMessage         = "edit_message"
	OperationKindUpdatePinnedMessage = "update_pinned_message"
	OperationKindScheduledMarker     = "scheduled_marker"
	OperationKindMarkDialogUnread    = "mark_dialog_unread"
	EventKindNewMessage              = "new_message"
	EventKindChatParticipantsChanged = "chat_participants_changed"
	FactKindNewMessage               = "new_message"
	FactKindChatParticipantsChanged  = "chat_participants_changed"
	FactKindTLUpdate                 = "tl_update"
)
