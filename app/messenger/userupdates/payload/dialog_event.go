package payload

const (
	DialogEventSchemaVersion = 1

	DialogEventDraftSaved                = "dialog.draftSaved"
	DialogEventDraftCleared              = "dialog.draftCleared"
	DialogEventDraftClearedAfterSend     = "dialog.draftClearedAfterSend"
	DialogEventPinToggled                = "dialog.pinToggled"
	DialogEventPinnedDialogsReordered    = "dialog.pinnedDialogsReordered"
	DialogEventFolderPeersChanged        = "updateFolderPeers"
	DialogEventFilterUpdated             = "dialog.filterUpdated"
	DialogEventFilterDeleted             = "dialog.filterDeleted"
	DialogEventFiltersOrderUpdated       = "dialog.filtersOrderUpdated"
	DialogEventWallpaperChanged          = "dialog.wallpaperChanged"
	DialogEventPrivatePeerHistoryTTL     = "updatePeerHistoryTTL"
	DialogEventPrivateThemeChanged       = "messageActionSetChatTheme"
	DialogEventSavedDialogPinned         = "dialog.savedDialogPinned"
	DialogEventPinnedSavedDialogsChanged = "dialog.pinnedSavedDialogs"
)

type DialogEventV1 struct {
	SchemaVersion    int    `json:"schema_version"`
	EventKind        string `json:"event_kind"`
	PublicUpdateType string `json:"public_update_type,omitempty"`
	PeerType         int32  `json:"peer_type,omitempty"`
	PeerID           int64  `json:"peer_id,omitempty"`
	FolderID         *int32 `json:"folder_id,omitempty"`
	Pinned           *bool  `json:"pinned,omitempty"`
	TTLPeriod        *int32 `json:"ttl_period,omitempty"`
}
