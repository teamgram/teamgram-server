// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/pp"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	// ph := "50524f585920544350342032302e36352e3139342e31362033392e3130312e38322e3535203435313832203434330d0a"
	// ph := "50524f585920544350342032302e31322e3234312e33203136332e3138312e34362e3431203336363936203434330d0a"
	// "0300002f 2ae00000 00000043 6f6f6b69 653a206d 73747368 6173683d 41646d69 6e697374 720d0a01 00080003 000000"
	ph := "50524f585920544350342032302e31322e3234312e33203136332e3138312e34362e3431203336363936203434330d0a"
	ppv1, _ := hex.DecodeString(ph)
	fmt.Printf("len: %d, \"%s\"\n", len(ppv1), ppv1)

	//ppv1, err := c.Peek(-1)
	//if err != nil {
	//	logx.Errorf("conn(%s) Peek fail: %v", c, err)
	//	return
	//}
	if len(ppv1) < len(pp.V1Identifier) {
		logx.Errorf("conn ppv1 < len(pp.V1Identifier), data: %s", hex.EncodeToString(ppv1))
		return
	}

	if bytes.HasPrefix(ppv1, pp.V1Identifier) {
		logx.Errorf("conn() ppv1 data: %s", hex.EncodeToString(ppv1))

		r := bytes.NewReader(ppv1)
		fmt.Printf("len(ppv1):%d - r.Len():%d, r.Size():%d = %d\n", len(ppv1), r.Len(), r.Size(), len(ppv1)-r.Len())
		h, err := pp.ReadHeader(r)
		if err != nil {
			logx.Errorf("conn() ReadHeader error: %v", err)
			if r.Len() > 107 {
				return
			} else {
				return
			}
		}

		fmt.Println(strings.Split(h.Source.String(), ":")[0])
		// _, _ = c.Discard(len(ppv1) - r.Len())
		fmt.Printf("len(ppv1):%d - r.Len():%d, r.Size():%d = %d\n", len(ppv1), r.Len(), r.Size(), len(ppv1)-r.Len())
		// ctx.ppv1 = false
	} else {
		// ctx.ppv1 = false
	}
}
