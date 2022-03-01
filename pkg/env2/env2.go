// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package env2

import "flag"

var (
	SmsCodeName    = ""
	MyAppName      = "Teamgram"
	MyWebSite      = "nebula.chat"
	TDotMe         = "t.me"
	PredefinedUser = false

	// PredefinedUser2
	// predefined2 - auto register
	PredefinedUser2 = false

	Server2 = false
)

func init() {
	flag.StringVar(&SmsCodeName, "code", "", "code")
	flag.StringVar(&MyAppName, "app_name", "Teamgram", "app_name")
	flag.StringVar(&MyWebSite, "site_name", "nebula.chat", "site_name")
	flag.StringVar(&TDotMe, "t.me", "t.me", "t.me")
	flag.BoolVar(&PredefinedUser, "predefined", false, "predefined")
	flag.BoolVar(&PredefinedUser2, "predefined2", false, "predefined2")
	flag.BoolVar(&Server2, "server2", false, "server2")
}
