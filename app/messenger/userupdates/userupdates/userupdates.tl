// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
// op_type / op_source / payload_codec / result_status are stable Go constants.
// payload must be a Teamgram-owned DTO, not raw TL JSON.

userOperation flags:# user_id:long bucket_id:int partition_id:int operation_id:string op_type:int op_source:int actor_user_id:long auth_key_id:flags.0?long auth_key_id_exclude:flags.1?long peer_type:int peer_id:long canonical_message_id:flags.2?long canonical_peer_seq:flags.3?long canonical_date:flags.4?long dependency_pts:flags.5?long request_id:flags.6?string payload_schema_version:int payload_codec:int payload_hash:bytes payload:bytes = UserOperation;

userOperationResult flags:# user_id:long operation_id:string status:int pts:long pts_count:int current_pts:long response_schema_version:flags.0?int response_payload:flags.1?bytes response_payload_hash:flags.2?bytes = UserOperationResult;

userState pts:long qts:int date:int seq:int unread_count:int = UserState;

userDifferenceEmpty state:UserState = UserDifference;
userDifference new_messages:Vector<Message> other_updates:Vector<Update> state:UserState = UserDifference;
userDifferenceSlice new_messages:Vector<Message> other_updates:Vector<Update> intermediate_state:UserState = UserDifference;
userDifferenceTooLong pts:long = UserDifference;

dialogProjectionPeer peer_type:int peer_id:long = DialogProjectionPeer;
dialogProjection flags:# peer_type:int peer_id:long top_peer_seq:long top_canonical_message_id:long top_message_date:long top_message_status:int read_inbox_max_peer_seq:long read_outbox_max_peer_seq:long unread_count:int unread_mentions_count:int unread_reactions_count:int unread_mark:flags.0?true pinned_peer_seq:long pinned_canonical_message_id:long has_scheduled:flags.1?true available_min_peer_seq:long last_pts:long last_pts_at:long dialog_schema_version:int dialog_payload:bytes = DialogProjection;
dialogProjectionList projections:Vector<DialogProjection> next_top_message_date:long next_top_peer_seq:long next_peer_type:int next_peer_id:long exhausted:Bool = DialogProjectionList;
messageViewPeerSeq peer_type:int peer_id:long peer_seq:long = MessageViewPeerSeq;
messageViewList messages:Vector<Message> = MessageViewList;

userAuthSeqAppendResult user_id:long operation_id:string seq:long date:int already_applied:Bool = UserAuthSeqAppendResult;
userPtsAppendResult user_id:long operation_id:string pts:long pts_count:int already_applied:Bool = UserPtsAppendResult;

---functions---
userupdates.processUserOperation operation:UserOperation = UserOperationResult;

userupdates.getOperationResult user_id:long operation_id:string payload_hash:bytes = UserOperationResult;

userupdates.getState user_id:long auth_key_id:long = UserState;

userupdates.getDifference flags:# user_id:long auth_key_id:long pts:long pts_total_limit:flags.0?int date:flags.1?long = UserDifference;

userupdates.listDialogs user_id:long top_message_date:long top_peer_seq:long peer_type:int peer_id:long limit:int = DialogProjectionList;

userupdates.getDialogsByPeers user_id:long peers:Vector<DialogProjectionPeer> = Vector<DialogProjection>;

userupdates.getDialogCount user_id:long = Int32;

userupdates.getMessageViewsByPeerSeqs user_id:long peers:Vector<MessageViewPeerSeq> = MessageViewList;

userupdates.getOutboxReadDate user_id:long peer_type:int peer_id:long msg_id:int = OutboxReadDate;

userupdates.appendDialogAuthSeqSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserAuthSeqAppendResult;

userupdates.appendDialogPtsSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserPtsAppendResult;

// LAYER 0
