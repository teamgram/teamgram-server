// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
onlineSession user_id:long perm_auth_key_id:long auth_key_id:long auth_key_type:int session_id:long gateway_id:string gateway_generation:string gateway_rpc_addr:string layer:int client:string updated_at:long expires_at:long = OnlineSession;
userOnlineSessions user_id:long sessions:Vector<OnlineSession> = UserOnlineSessions;

---functions---
presence.setSessionOnline session:OnlineSession = Bool;
presence.setSessionOffline user_id:long auth_key_id:long session_id:long = Bool;
presence.getUserOnlineSessions user_id:long = UserOnlineSessions;
presence.getUsersOnlineSessions users:Vector<long> = Vector<UserOnlineSessions>;
presence.getGatewaySessions gateway_id:string = Vector<OnlineSession>;

// LAYER 0
