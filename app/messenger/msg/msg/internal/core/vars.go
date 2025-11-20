// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"flag"
)

var (
	kUseV3 bool
	kUseV4 bool
)

func init() {
	flag.BoolVar(&kUseV3, "usev3", false, "use v3")
	flag.BoolVar(&kUseV4, "usev4", false, "use v4")
}
