// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package authsession

const (
	AuthStateUnknown       = 0
	AuthStateNew           = 1
	AuthStatePermBound     = 2
	AuthStateInited        = 3
	AuthStateAuthorization = 4
	AuthStateNeedPassword  = 5
	AuthStateNormal        = 6
	AuthStateLogout        = 7
	AuthStateDeleted       = 8
)
