// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
updateDraftMessage peer:Peer draft:DraftMessage = PeerWithDraftMessage;

dialogExt flags:# order:long dialog:Dialog available_min_id:int date:long theme_emoticon:string ttl_period:int wallpaper_id:long wallpaper_overridden:flags.0?true = DialogExt;
dialogPinnedExt order:long peer_type:int peer_id:long = DialogPinnedExt;

dialogFilterExt flags:# id:int joined_by_slug:flags.0?true slug:string dialog_filter:DialogFilter order:long = DialogFilterExt;

simpleDialogsData users:Vector<long> chats:Vector<long> channels:Vector<long> = DialogsData;

savedDialogList count:int dialogs:Vector<SavedDialog> = SavedDialogList;

---functions---
dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
dialog.clearDraftMessage user_id:long peer_type:int peer_id:long source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
dialog.clearDraftAfterSend user_id:long peer_type:int peer_id:long clear_before_date:int source_perm_auth_key_id:long source_operation_id:string outbox_id:long = Bool;
dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
dialog.clearAllDrafts user_id:long source_perm_auth_key_id:long operation_id:string outbox_ids:Vector<long> = Vector<PeerWithDraftMessage>;

dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool source_perm_auth_key_id:long operation_id:string outbox_id:long = Int32;
dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;

dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;

dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long pinned_msg_id:flags.6?int = Bool;
dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;

dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;

// folders
dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
dialog.deleteDialogFilter user_id:long id:int = Bool;
dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;

dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int source_perm_auth_key_id:long operation_id:string outbox_ids:Vector<long> = Vector<DialogPinnedExt>;

dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;

// chatTheme
dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;

// ttl
dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;

dialog.getMyDialogsData flags:# user_id:long user:flags.0?true chat:flags.1?true channel:flags.2?true = DialogsData;

// saved_dialogs
dialog.getSavedDialogs user_id:long exclude_pinned:Bool offset_date:int offset_id:int offset_peer:PeerUtil limit:int = SavedDialogList;
dialog.getPinnedSavedDialogs user_id:long = SavedDialogList;
dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;

dialog.getDialogFilter user_id:long id:int = DialogFilterExt;
dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;

dialog.updateUnreadCount flags:# user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;

dialog.toggleDialogFilterTags user_id:long enabled:Bool = Bool;
dialog.getDialogFilterTags user_id:long = Bool;

// chatTheme
dialog.setChatWallpaper flags:# user_id:long peer_type:int peer_id:long wallpaper_id:long wallpaper_overridden:flags.0?true = Bool;

// LAYER 0
