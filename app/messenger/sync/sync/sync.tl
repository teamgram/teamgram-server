// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---

---functions---
sync.updatesMe flags:# user_id:long perm_auth_key_id:long auth_key_id:flags.0?long session_id:flags.1?long updates:Updates = Void;
sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
sync.pushUpdates user_id:long updates:Updates = Void;
sync.pushUpdatesIfNot flags:# user_id:long includes:flags.0?Vector<long> excludes:flags.1?Vector<long> updates:Updates = Void;
sync.pushRpcResult user_id:long perm_auth_key_id:long auth_key_id:long gateway_id:string gateway_generation:string gateway_rpc_addr:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;

// LAYER 0
