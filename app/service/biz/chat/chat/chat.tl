// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
//immutableChatParticipant id:long chat_id:long user_id:long state:int participant_type:int link:string useage:int inviter_user_id:long invited_at:long kicked_at:long left_at:long admin_rights:ChatAdminRights date:long = ImmutableChatParticipant;
//immutableChat flags:# id:long creator:long title:string photo:Photo deactivated:flags.0?true call_active:flags.1?true call_not_empty:flags.2?true noforwards:flags.25?true participants_count:int date:long version:int migrated_to:flags.3?InputChannel default_banned_rights:ChatBannedRights can_set_username:flags.4?true about:string exported_invite:flags.5?ExportedChatInvite bot_info:flags.6?Vector<BotInfo> call:flags.7?InputGroupCall available_reactions_type:int available_reactions:Vector<string> ttl_period:int = ImmutableChat;
//
//mutableChat chat:ImmutableChat chat_participants:Vector<ImmutableChatParticipant> = MutableChat;

// ChatInviteExt;
chatInviteAlready chat:MutableChat = ChatInviteExt;
chatInvite flags:# request_needed:flags.6?true title:string about:flags.5?string photo:Photo participants_count:int participants:flags.4?Vector<long> = ChatInviteExt;
chatInvitePeek chat:MutableChat expires:int = ChatInviteExt;

userChatIdList user_id:long chat_id_list:Vector<long> = UserChatIdList;

recentChatInviteRequesters requests_pending:int recent_requesters:Vector<long> = RecentChatInviteRequesters;
chatAccessCheckResult self_id:long chat_id:long access_kind:string = ChatAccessCheckResult;
messageActionCheckResult self_id:long chat_id:long action:string media_kind:string = MessageActionCheckResult;
chatInviteImported flags:# chat:MutableChat requesters:flags.0?RecentChatInviteRequesters = ChatInviteImported;

---functions---
chat.getMutableChat chat_id:long = MutableChat;
chat.getChatListByIdList self_id:long id_list:Vector<long> = Vector<MutableChat>;
chat.getChatBySelfId self_id:long chat_id:long = MutableChat;
chat.checkChatAccess self_id:long chat_id:long access_kind:string = ChatAccessCheckResult;
chat.checkMessageAction self_id:long chat_id:long action:string media_kind:string = MessageActionCheckResult;

chat.createChat2 flags:# creator_id:long user_id_list:Vector<long> title:string bots:flags.0?Vector<long> ttl_period:flags.1?int client_msg_id:flags.2?long operation_id:flags.3?string = MutableChat;
chat.deleteChat chat_id:long operator_id:long = MutableChat;
chat.deleteChatUser chat_id:long operator_id:long delete_user_id:long = MutableChat;
chat.editChatTitle chat_id:long edit_user_id:long title:string = MutableChat;
chat.editChatAbout chat_id:long edit_user_id:long about:string = MutableChat;
chat.editChatPhoto chat_id:long edit_user_id:long chat_photo:Photo = MutableChat;
chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
chat.addChatUser flags:# chat_id:long inviter_id:long user_id:long is_bot:flags.0?true = MutableChat;
chat.getMutableChatByLink link:string = MutableChat;
chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;

chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;

chat.getChatParticipantIdList chat_id:long = Vector<long>;

chat.getUsersChatIdList id:Vector<long> = Vector<UserChatIdList>;
chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;

chat.exportChatInvite flags:# chat_id:long admin_id:long legacy_revoke_permanent:flags.2?true request_needed:flags.3?true expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
chat.checkChatInvite self_id:long hash:string = ChatInviteExt;
chat.importChatInvite self_id:long hash:string = MutableChat;
chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
chat.deleteExportedChatInvite self_id:long chat_id:long link:string = Bool;
chat.deleteRevokedExportedChatInvites self_id:long chat_id:long admin_id:long = Bool;
chat.editExportedChatInvite flags:# self_id:long chat_id:long revoked:flags.2?true link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = Vector<ExportedChatInvite>;

chat.setChatAvailableReactions self_id:long chat_id:long available_reactions_type:int available_reactions:Vector<string> = MutableChat;

// ttl
chat.setHistoryTTL self_id:long chat_id:long ttl_period:int = MutableChat;

chat.search self_id:long q:string offset:long limit:int = Vector<MutableChat>;

chat.getRecentChatInviteRequesters self_id:long chat_id:long = RecentChatInviteRequesters;
chat.hideChatJoinRequests flags:# self_id:long chat_id:long approved:flags.0?true link:flags.1?string user_id:flags.2?long = RecentChatInviteRequesters;
chat.importChatInvite2 self_id:long hash:string = ChatInviteImported;

// LAYER 0
