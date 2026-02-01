/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mt

import (
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"
)

const (
	ClazzName_resPQ                      = "resPQ"
	ClazzName_p_q_inner_data             = "p_q_inner_data"
	ClazzName_p_q_inner_data_dc          = "p_q_inner_data_dc"
	ClazzName_p_q_inner_data_temp        = "p_q_inner_data_temp"
	ClazzName_p_q_inner_data_temp_dc     = "p_q_inner_data_temp_dc"
	ClazzName_bind_auth_key_inner        = "bind_auth_key_inner"
	ClazzName_server_DH_params_fail      = "server_DH_params_fail"
	ClazzName_server_DH_params_ok        = "server_DH_params_ok"
	ClazzName_server_DH_inner_data       = "server_DH_inner_data"
	ClazzName_client_DH_inner_data       = "client_DH_inner_data"
	ClazzName_dh_gen_ok                  = "dh_gen_ok"
	ClazzName_dh_gen_retry               = "dh_gen_retry"
	ClazzName_dh_gen_fail                = "dh_gen_fail"
	ClazzName_destroy_auth_key_ok        = "destroy_auth_key_ok"
	ClazzName_destroy_auth_key_none      = "destroy_auth_key_none"
	ClazzName_destroy_auth_key_fail      = "destroy_auth_key_fail"
	ClazzName_req_pq                     = "req_pq"
	ClazzName_req_pq_multi               = "req_pq_multi"
	ClazzName_req_DH_params              = "req_DH_params"
	ClazzName_set_client_DH_params       = "set_client_DH_params"
	ClazzName_destroy_auth_key           = "destroy_auth_key"
	ClazzName_msgs_ack                   = "msgs_ack"
	ClazzName_bad_msg_notification       = "bad_msg_notification"
	ClazzName_bad_server_salt            = "bad_server_salt"
	ClazzName_msgs_state_req             = "msgs_state_req"
	ClazzName_msgs_state_info            = "msgs_state_info"
	ClazzName_msgs_all_info              = "msgs_all_info"
	ClazzName_msg_detailed_info          = "msg_detailed_info"
	ClazzName_msg_new_detailed_info      = "msg_new_detailed_info"
	ClazzName_msg_resend_req             = "msg_resend_req"
	ClazzName_rpc_error                  = "rpc_error"
	ClazzName_rpc_answer_unknown         = "rpc_answer_unknown"
	ClazzName_rpc_answer_dropped_running = "rpc_answer_dropped_running"
	ClazzName_rpc_answer_dropped         = "rpc_answer_dropped"
	ClazzName_future_salt                = "future_salt"
	ClazzName_future_salts               = "future_salts"
	ClazzName_pong                       = "pong"
	ClazzName_destroy_session_ok         = "destroy_session_ok"
	ClazzName_destroy_session_none       = "destroy_session_none"
	ClazzName_new_session_created        = "new_session_created"
	ClazzName_http_wait                  = "http_wait"
	ClazzName_ipPort                     = "ipPort"
	ClazzName_ipPortSecret               = "ipPortSecret"
	ClazzName_accessPointRule            = "accessPointRule"
	ClazzName_help_configSimple          = "help_configSimple"
	ClazzName_tlsClientHello             = "tlsClientHello"
	ClazzName_tlsBlockString             = "tlsBlockString"
	ClazzName_tlsBlockRandom             = "tlsBlockRandom"
	ClazzName_tlsBlockZero               = "tlsBlockZero"
	ClazzName_tlsBlockDomain             = "tlsBlockDomain"
	ClazzName_tlsBlockGrease             = "tlsBlockGrease"
	ClazzName_tlsBlockPublicKey          = "tlsBlockPublicKey"
	ClazzName_tlsBlockScope              = "tlsBlockScope"
	ClazzName_help_test                  = "help_test"
	ClazzName_test_useError              = "test_useError"
	ClazzName_test_useConfigSimple       = "test_useConfigSimple"
	ClazzName_rpc_drop_answer            = "rpc_drop_answer"
	ClazzName_get_future_salts           = "get_future_salts"
	ClazzName_ping                       = "ping"
	ClazzName_ping_delay_disconnect      = "ping_delay_disconnect"
	ClazzName_destroy_session            = "destroy_session"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_resPQ, 0, 0x5162463)                       // 5162463
	iface.RegisterClazzName(ClazzName_p_q_inner_data, 0, 0x83c95aec)             // 83c95aec
	iface.RegisterClazzName(ClazzName_p_q_inner_data_dc, 0, 0xa9f55f95)          // a9f55f95
	iface.RegisterClazzName(ClazzName_p_q_inner_data_temp, 0, 0x3c6a84d4)        // 3c6a84d4
	iface.RegisterClazzName(ClazzName_p_q_inner_data_temp_dc, 0, 0x56fddf88)     // 56fddf88
	iface.RegisterClazzName(ClazzName_bind_auth_key_inner, 0, 0x75a3f765)        // 75a3f765
	iface.RegisterClazzName(ClazzName_server_DH_params_fail, 0, 0x79cb045d)      // 79cb045d
	iface.RegisterClazzName(ClazzName_server_DH_params_ok, 0, 0xd0e8075c)        // d0e8075c
	iface.RegisterClazzName(ClazzName_server_DH_inner_data, 0, 0xb5890dba)       // b5890dba
	iface.RegisterClazzName(ClazzName_client_DH_inner_data, 0, 0x6643b654)       // 6643b654
	iface.RegisterClazzName(ClazzName_dh_gen_ok, 0, 0x3bcbf734)                  // 3bcbf734
	iface.RegisterClazzName(ClazzName_dh_gen_retry, 0, 0x46dc1fb9)               // 46dc1fb9
	iface.RegisterClazzName(ClazzName_dh_gen_fail, 0, 0xa69dae02)                // a69dae02
	iface.RegisterClazzName(ClazzName_destroy_auth_key_ok, 0, 0xf660e1d4)        // f660e1d4
	iface.RegisterClazzName(ClazzName_destroy_auth_key_none, 0, 0xa9f2259)       // a9f2259
	iface.RegisterClazzName(ClazzName_destroy_auth_key_fail, 0, 0xea109b13)      // ea109b13
	iface.RegisterClazzName(ClazzName_req_pq, 0, 0x60469778)                     // 60469778
	iface.RegisterClazzName(ClazzName_req_pq_multi, 0, 0xbe7e8ef1)               // be7e8ef1
	iface.RegisterClazzName(ClazzName_req_DH_params, 0, 0xd712e4be)              // d712e4be
	iface.RegisterClazzName(ClazzName_set_client_DH_params, 0, 0xf5045f1f)       // f5045f1f
	iface.RegisterClazzName(ClazzName_destroy_auth_key, 0, 0xd1435160)           // d1435160
	iface.RegisterClazzName(ClazzName_msgs_ack, 0, 0x62d6b459)                   // 62d6b459
	iface.RegisterClazzName(ClazzName_bad_msg_notification, 0, 0xa7eff811)       // a7eff811
	iface.RegisterClazzName(ClazzName_bad_server_salt, 0, 0xedab447b)            // edab447b
	iface.RegisterClazzName(ClazzName_msgs_state_req, 0, 0xda69fb52)             // da69fb52
	iface.RegisterClazzName(ClazzName_msgs_state_info, 0, 0x4deb57d)             // 4deb57d
	iface.RegisterClazzName(ClazzName_msgs_all_info, 0, 0x8cc0d131)              // 8cc0d131
	iface.RegisterClazzName(ClazzName_msg_detailed_info, 0, 0x276d3ec6)          // 276d3ec6
	iface.RegisterClazzName(ClazzName_msg_new_detailed_info, 0, 0x809db6df)      // 809db6df
	iface.RegisterClazzName(ClazzName_msg_resend_req, 0, 0x7d861a08)             // 7d861a08
	iface.RegisterClazzName(ClazzName_rpc_error, 0, 0x2144ca19)                  // 2144ca19
	iface.RegisterClazzName(ClazzName_rpc_answer_unknown, 0, 0x5e2ad36e)         // 5e2ad36e
	iface.RegisterClazzName(ClazzName_rpc_answer_dropped_running, 0, 0xcd78e586) // cd78e586
	iface.RegisterClazzName(ClazzName_rpc_answer_dropped, 0, 0xa43ad8b7)         // a43ad8b7
	iface.RegisterClazzName(ClazzName_future_salt, 0, 0x949d9dc)                 // 949d9dc
	iface.RegisterClazzName(ClazzName_future_salts, 0, 0xae500895)               // ae500895
	iface.RegisterClazzName(ClazzName_pong, 0, 0x347773c5)                       // 347773c5
	iface.RegisterClazzName(ClazzName_destroy_session_ok, 0, 0xe22045fc)         // e22045fc
	iface.RegisterClazzName(ClazzName_destroy_session_none, 0, 0x62d350c9)       // 62d350c9
	iface.RegisterClazzName(ClazzName_new_session_created, 0, 0x9ec20908)        // 9ec20908
	iface.RegisterClazzName(ClazzName_http_wait, 0, 0x9299359f)                  // 9299359f
	iface.RegisterClazzName(ClazzName_ipPort, 0, 0xd433ad73)                     // d433ad73
	iface.RegisterClazzName(ClazzName_ipPortSecret, 0, 0x37982646)               // 37982646
	iface.RegisterClazzName(ClazzName_accessPointRule, 0, 0x4679b65f)            // 4679b65f
	iface.RegisterClazzName(ClazzName_help_configSimple, 0, 0x5a592a6c)          // 5a592a6c
	iface.RegisterClazzName(ClazzName_tlsClientHello, 0, 0x6c52c484)             // 6c52c484
	iface.RegisterClazzName(ClazzName_tlsBlockString, 0, 0x4218a164)             // 4218a164
	iface.RegisterClazzName(ClazzName_tlsBlockRandom, 0, 0x4d4dc41e)             // 4d4dc41e
	iface.RegisterClazzName(ClazzName_tlsBlockZero, 0, 0x9333afb)                // 9333afb
	iface.RegisterClazzName(ClazzName_tlsBlockDomain, 0, 0x10e8636f)             // 10e8636f
	iface.RegisterClazzName(ClazzName_tlsBlockGrease, 0, 0xe675a1c1)             // e675a1c1
	iface.RegisterClazzName(ClazzName_tlsBlockPublicKey, 0, 0x9eb95b5c)          // 9eb95b5c
	iface.RegisterClazzName(ClazzName_tlsBlockScope, 0, 0xe725d44f)              // e725d44f
	iface.RegisterClazzName(ClazzName_help_test, 0, 0xc0e202f7)                  // c0e202f7
	iface.RegisterClazzName(ClazzName_test_useError, 0, 0xee75af01)              // ee75af01
	iface.RegisterClazzName(ClazzName_test_useConfigSimple, 0, 0xf9b7b23d)       // f9b7b23d
	iface.RegisterClazzName(ClazzName_rpc_drop_answer, 0, 0x58e4a740)            // 58e4a740
	iface.RegisterClazzName(ClazzName_get_future_salts, 0, 0xb921bd04)           // b921bd04
	iface.RegisterClazzName(ClazzName_ping, 0, 0x7abe77ec)                       // 7abe77ec
	iface.RegisterClazzName(ClazzName_ping_delay_disconnect, 0, 0xf3427b8c)      // f3427b8c
	iface.RegisterClazzName(ClazzName_destroy_session, 0, 0xe7512126)            // e7512126

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_resPQ, 0x5162463)                       // 5162463
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data, 0x83c95aec)             // 83c95aec
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_dc, 0xa9f55f95)          // a9f55f95
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_temp, 0x3c6a84d4)        // 3c6a84d4
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_temp_dc, 0x56fddf88)     // 56fddf88
	iface.RegisterClazzIDName(ClazzName_bind_auth_key_inner, 0x75a3f765)        // 75a3f765
	iface.RegisterClazzIDName(ClazzName_server_DH_params_fail, 0x79cb045d)      // 79cb045d
	iface.RegisterClazzIDName(ClazzName_server_DH_params_ok, 0xd0e8075c)        // d0e8075c
	iface.RegisterClazzIDName(ClazzName_server_DH_inner_data, 0xb5890dba)       // b5890dba
	iface.RegisterClazzIDName(ClazzName_client_DH_inner_data, 0x6643b654)       // 6643b654
	iface.RegisterClazzIDName(ClazzName_dh_gen_ok, 0x3bcbf734)                  // 3bcbf734
	iface.RegisterClazzIDName(ClazzName_dh_gen_retry, 0x46dc1fb9)               // 46dc1fb9
	iface.RegisterClazzIDName(ClazzName_dh_gen_fail, 0xa69dae02)                // a69dae02
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_ok, 0xf660e1d4)        // f660e1d4
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_none, 0xa9f2259)       // a9f2259
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_fail, 0xea109b13)      // ea109b13
	iface.RegisterClazzIDName(ClazzName_req_pq, 0x60469778)                     // 60469778
	iface.RegisterClazzIDName(ClazzName_req_pq_multi, 0xbe7e8ef1)               // be7e8ef1
	iface.RegisterClazzIDName(ClazzName_req_DH_params, 0xd712e4be)              // d712e4be
	iface.RegisterClazzIDName(ClazzName_set_client_DH_params, 0xf5045f1f)       // f5045f1f
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key, 0xd1435160)           // d1435160
	iface.RegisterClazzIDName(ClazzName_msgs_ack, 0x62d6b459)                   // 62d6b459
	iface.RegisterClazzIDName(ClazzName_bad_msg_notification, 0xa7eff811)       // a7eff811
	iface.RegisterClazzIDName(ClazzName_bad_server_salt, 0xedab447b)            // edab447b
	iface.RegisterClazzIDName(ClazzName_msgs_state_req, 0xda69fb52)             // da69fb52
	iface.RegisterClazzIDName(ClazzName_msgs_state_info, 0x4deb57d)             // 4deb57d
	iface.RegisterClazzIDName(ClazzName_msgs_all_info, 0x8cc0d131)              // 8cc0d131
	iface.RegisterClazzIDName(ClazzName_msg_detailed_info, 0x276d3ec6)          // 276d3ec6
	iface.RegisterClazzIDName(ClazzName_msg_new_detailed_info, 0x809db6df)      // 809db6df
	iface.RegisterClazzIDName(ClazzName_msg_resend_req, 0x7d861a08)             // 7d861a08
	iface.RegisterClazzIDName(ClazzName_rpc_error, 0x2144ca19)                  // 2144ca19
	iface.RegisterClazzIDName(ClazzName_rpc_answer_unknown, 0x5e2ad36e)         // 5e2ad36e
	iface.RegisterClazzIDName(ClazzName_rpc_answer_dropped_running, 0xcd78e586) // cd78e586
	iface.RegisterClazzIDName(ClazzName_rpc_answer_dropped, 0xa43ad8b7)         // a43ad8b7
	iface.RegisterClazzIDName(ClazzName_future_salt, 0x949d9dc)                 // 949d9dc
	iface.RegisterClazzIDName(ClazzName_future_salts, 0xae500895)               // ae500895
	iface.RegisterClazzIDName(ClazzName_pong, 0x347773c5)                       // 347773c5
	iface.RegisterClazzIDName(ClazzName_destroy_session_ok, 0xe22045fc)         // e22045fc
	iface.RegisterClazzIDName(ClazzName_destroy_session_none, 0x62d350c9)       // 62d350c9
	iface.RegisterClazzIDName(ClazzName_new_session_created, 0x9ec20908)        // 9ec20908
	iface.RegisterClazzIDName(ClazzName_http_wait, 0x9299359f)                  // 9299359f
	iface.RegisterClazzIDName(ClazzName_ipPort, 0xd433ad73)                     // d433ad73
	iface.RegisterClazzIDName(ClazzName_ipPortSecret, 0x37982646)               // 37982646
	iface.RegisterClazzIDName(ClazzName_accessPointRule, 0x4679b65f)            // 4679b65f
	iface.RegisterClazzIDName(ClazzName_help_configSimple, 0x5a592a6c)          // 5a592a6c
	iface.RegisterClazzIDName(ClazzName_tlsClientHello, 0x6c52c484)             // 6c52c484
	iface.RegisterClazzIDName(ClazzName_tlsBlockString, 0x4218a164)             // 4218a164
	iface.RegisterClazzIDName(ClazzName_tlsBlockRandom, 0x4d4dc41e)             // 4d4dc41e
	iface.RegisterClazzIDName(ClazzName_tlsBlockZero, 0x9333afb)                // 9333afb
	iface.RegisterClazzIDName(ClazzName_tlsBlockDomain, 0x10e8636f)             // 10e8636f
	iface.RegisterClazzIDName(ClazzName_tlsBlockGrease, 0xe675a1c1)             // e675a1c1
	iface.RegisterClazzIDName(ClazzName_tlsBlockPublicKey, 0x9eb95b5c)          // 9eb95b5c
	iface.RegisterClazzIDName(ClazzName_tlsBlockScope, 0xe725d44f)              // e725d44f
	iface.RegisterClazzIDName(ClazzName_help_test, 0xc0e202f7)                  // c0e202f7
	iface.RegisterClazzIDName(ClazzName_test_useError, 0xee75af01)              // ee75af01
	iface.RegisterClazzIDName(ClazzName_test_useConfigSimple, 0xf9b7b23d)       // f9b7b23d
	iface.RegisterClazzIDName(ClazzName_rpc_drop_answer, 0x58e4a740)            // 58e4a740
	iface.RegisterClazzIDName(ClazzName_get_future_salts, 0xb921bd04)           // b921bd04
	iface.RegisterClazzIDName(ClazzName_ping, 0x7abe77ec)                       // 7abe77ec
	iface.RegisterClazzIDName(ClazzName_ping_delay_disconnect, 0xf3427b8c)      // f3427b8c
	iface.RegisterClazzIDName(ClazzName_destroy_session, 0xe7512126)            // e7512126
}
