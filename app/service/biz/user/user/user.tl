// LAYER 0
// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
//
//privacyKeyRules key:int rules:Vector<PrivacyRule> = PrivacyKeyRules;
//
//contactData flags:# user_id:long contact_user_id:long first_name:flags.0?string last_name:flags.1?string mutual_contact:flags.2?true = ContactData;
//
//botData flags:# id:long bot_type:int creator:long token:string description:string bot_chat_history:flags.1?true bot_nochats:flags.2?true bot_inline_geo:flags.3?true bot_info_version:int bot_inline_placeholder:flags.4?string= BotData;
//userData flags:# id:long access_hash:long user_type:int sceret_key_id:long first_name:string last_name:string username:string phone:string profile_photo:flags.0?Photo bot:flags.8?BotData country_code:string verified:flags.1?true support:flags.2?true scam:flags.3?true fake:flags.4?true about:flags.5?string restricted:flags.7?true restriction_reason:flags.7?Vector<RestrictionReason> contacts_version:int privacies_version:int deleted:flags.9?true bot_attach_menu:flags.10?true premium:flags.11?true emoji_status:flags.12?EmojiStatus = UserData;
//
//immutableUser flags:# user:UserData last_seen_at:long contacts:flags.0?Vector<ContactData> reverse_contacts:flags.1?Vector<ContactData> keys_privacy_rules:Vector<PrivacyKeyRules> = ImmutableUser;

viewerUsers viewer_user_id:long users:Vector<User> = ViewerUsers;
userProjectionBundle flags:# facts:flags.0?Vector<ImmutableUser> viewer_users:Vector<ViewerUsers> missing_user_ids:Vector<long> = UserProjectionBundle;

// userProfilePhotos photo:Photo id_list:Vector<long> =  UserProfilePhotos;
// simpleUserData id:long access_hash:long first_name:string last_name:string about:string username:= SimpleUserData;
userImportedContacts imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> update_id_list:Vector<long> = UserImportedContacts;

usersDataFound count:int users:Vector<UserData> next_offset:string = UsersFound;
usersIdFound id_list:Vector<long> = UsersFound;

peerPeerNotifySettings peer_type:int peer_id:long settings:PeerNotifySettings = PeerPeerNotifySettings;
lastSeenData user_id:long last_seen_at:long expires:int = LastSeenData;
botInfoData flags:# bot_info:BotInfo main_app_url:flags.0?string bot_inline:flags.1?true token:string bot_id:long = BotInfoData;

usernameNotExisted = UsernameExist;
usernameExisted = UsernameExist;
usernameExistedNotMe = UsernameExist;
usernameExistedIsMe = UsernameExist;

usernameData flags:# username:string peer:flags.0?Peer editable:flags.1?true active:flags.2?true = UsernameData;

---functions---
user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
user.getLastSeen id:long = LastSeenData;

user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
user.getImmutableUserByPhone phone:string = ImmutableUser;
user.getImmutableUserByToken token:string = ImmutableUser;
user.setAccountDaysTTL user_id:long ttl:int = Bool;
user.getAccountDaysTTL user_id:long = AccountDaysTTL;

user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
user.resetNotifySettings user_id:long = Bool;
user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;

user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;

user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;

user.changePhone user_id:long phone:string = Bool;

user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
user.deleteUser user_id:long reason:string phone:string = Bool;
//user.getUserSelf id:long = User;
//user.getUserById self_user_id:long user_id:long = User;
//user.getUserListByIdList self_user_id:long user_id_list:Vector<long> = Vector<User>;
//user.getUserName user_id:long = String;
//user.getUserByToken token:string = User;

user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
user.blockedByUser user_id:long peer_user_id:long = Bool;
user.isBlockedByUser user_id:long peer_user_id:long = Bool;
user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;

user.getContactSignUpNotification user_id:long = Bool;
user.setContactSignUpNotification user_id:long silent:Bool = Bool;

user.getContentSettings user_id:long = account.ContentSettings;
user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;

user.deleteContact user_id:long id:long = Bool;
user.getContactList user_id:long = Vector<ContactData>;
user.getContactIdList user_id:long = Vector<long>;
user.getContact user_id:long id:long = ContactData;
user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
user.checkContact user_id:long id:long = Bool;

//
user.getImportersByPhone phone:string = Vector<InputContact>;
user.deleteImportersByPhone phone:string = Bool;
user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;

user.getCountryCode user_id:long = String;

user.updateAbout user_id:long about:string = Bool;
user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
user.updateVerified user_id:long verified:Bool = Bool;
user.updateUsername user_id:long username:string = Bool;
user.updateProfilePhoto user_id:long id:long = Int64;
user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
user.getProfilePhotos user_id:long = Vector<long>;
// user.putCacheUserPhotos user_id:long photos:UserPhotos = Bool;

user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
user.isBot id:long = Bool;
user.getBotInfo bot_id:long = BotInfo;
user.checkBots id:Vector<long> = Vector<long>;

user.getFullUser self_user_id:long id:long = users.UserFull;

user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;

user.getUserDataById user_id:long = UserData;
user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
user.getUserDataByToken token:string = UserData;

user.search q:string excluded_contacts:Vector<long> offset:long limit:int = UsersFound;

user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?string bot_has_main_app:flags.13?Bool = Bool;

user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
user.getUserProjectionBundle flags:# with_facts:flags.0?true viewer_user_ids:Vector<long> target_user_ids:Vector<long> = UserProjectionBundle;

// 10,000,000
user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;

// close friends list
user.editCloseFriends user_id:long id:Vector<long> = Bool;

// stories_max_id
user.setStoriesMaxId user_id:long id:int = Bool;

user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;

user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
user.getBirthdays user_id:long = Vector<ContactBirthday>;

user.setStoriesHidden user_id:long id:long hidden:Bool = Bool;

user.updatePersonalChannel user_id:long channel_id:long = Bool;

user.getUserIdByPhone phone:string = Int64;

user.setAuthorizationTTL user_id:long ttl:int = Bool;
user.getAuthorizationTTL user_id:long = AccountDaysTTL;

user.updatePremium flags:# user_id:long premium:Bool months:flags.1?int = Bool;

user.getBotInfoV2 bot_id:long = BotInfoData;

user.saveMusic flags:# unsave:flags.0?true user_id:long id:long after_id:flags.15?long = Bool;
user.getSavedMusicIdList user_id:long = Vector<long>;

user.setMainProfileTab user_id:long tab:ProfileTab = Bool;

user.setDefaultHistoryTTL user_id:long ttl:int = Bool;
user.getDefaultHistoryTTL user_id:long = DefaultHistoryTTL;

user.getAccountUsername user_id:long = UsernameData;
user.checkAccountUsername user_id:long username:string = UsernameExist;
user.getChannelUsername channel_id:long = UsernameData;
user.checkChannelUsername channel_id:long username:string = UsernameExist;
user.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
user.checkUsername username:string = UsernameExist;
user.updateUsernameByUsername peer_type:int peer_id:long username:string = Bool;
user.deleteUsername username:string = Bool;
user.resolveUsername username:string = Peer;
user.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
user.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
user.searchUsername q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
user.toggleUsername peer_type:int peer_id:long username:string active:Bool = Bool;
user.reorderUsernames peer_type:int peer_id:long username_list:Vector<string> = Bool;
user.deactivateAllChannelUsernames channel_id:long = Bool;
