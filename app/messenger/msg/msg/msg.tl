// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
sender user_id:long type:int auth_key_id:long = Sender;
outboxMessage flags:# no_webpage:flags.0?true background:flags.1?true random_id:long message:Message schedule_date:flags.2?int = OutboxMessage;

contentMessage flags:# id:int dialog_message_id:long mentioned:flags.0?true media_unread:flags.1?true reaction:flags.2?true send_user_id:long = ContentMessage;

---functions---
msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;

msg.sendMessageV2 flags:# clear_draft:flags.0?true user_id:long auth_key_id:long source_perm_auth_key_id:flags.1?long clear_draft_before_date:flags.2?int peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;

msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;

msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;

msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
msg.getHistory user_id:long auth_key_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
msg.searchHashtag user_id:long auth_key_id:long peer_type:int peer_id:long hash_tag:string offset_id:int limit:int = messages.Messages;

msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;

// LAYER 0
