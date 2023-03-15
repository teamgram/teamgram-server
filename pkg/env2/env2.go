// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package env2

import (
	"flag"
)

var (
	MyAppName      = "Teamgram"
	MyWebSite      = "teamgram.net"
	TDotMe         = "t.me"
	PredefinedUser = false

	// PredefinedUser2
	// predefined2 - auto register
	PredefinedUser2 = false
)

func init() {
	flag.StringVar(&MyAppName, "app_name", "Teamgram", "app_name")
	flag.StringVar(&MyWebSite, "site_name", "teamgram.net", "site_name")
	flag.StringVar(&TDotMe, "t.me", "t.me", "t.me")
	flag.BoolVar(&PredefinedUser, "predefined", false, "predefined")
	flag.BoolVar(&PredefinedUser2, "predefined2", false, "predefined2")
}

func IsTDotMe(me string) bool {
	switch me {
	case "teamgram.me":
		return true
	case "t.me":
		return true
	case TDotMe:
		return true
	}

	return false
}
