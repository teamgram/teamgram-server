package chat

const (
	ChatAccessGetHistory  = "get_history"
	ChatAccessReadHistory = "read_history"
	ChatAccessGetDialog   = "get_dialog"
	ChatAccessSearch      = "search"

	MessageActionSendText       = "send_text"
	MessageActionSendMediaPhoto = "send_media_photo"
	MessageActionSendMediaDoc   = "send_media_document"
	MessageActionSendAlbum      = "send_album"
	MessageActionForwardToChat  = "forward_to_chat"
	MessageActionEditOwnMessage = "edit_own_message"
	MessageActionDeleteLocal    = "delete_local"
	MessageActionDeleteRevoke   = "delete_revoke"
	MessageActionPinMessage     = "pin_message"
	MessageActionUnpinAll       = "unpin_all"
	MessageActionSendPoll       = "send_poll"
	MessageActionSendInline     = "send_inline"
	MessageActionSendGame       = "send_game"
)
