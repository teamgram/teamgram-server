// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package ip

import (
	"fmt"
	"testing"
)

// TestGetOutBoundIP
// @Description: Test function to get the outbound IP address
// @return ip string, err error
func TestGetOutBoundIP(t *testing.T) {
	ip, _ := GetOutBoundIP()
	fmt.Println(ip)
}

// TestGetIp
// @Description: Test function to get the IP address from environment variable
func TestGetIp(t *testing.T) {
	ip := GetIp("127.0.0.1")
	fmt.Println(ip)
}
