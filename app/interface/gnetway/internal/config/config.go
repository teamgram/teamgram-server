/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package config

import (
	"github.com/teamgram/marmota/pkg/container2"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// MaxProc        int
	KeyFile        string
	KeyFingerprint string
	Gnetway        *GnetwayConfig
	Session        zrpc.RpcClientConf
}

type GnetwayServer struct {
	Proto     string `json:",default=tcp,options=tcp|websocket|http"`
	Addresses []string
}

type GnetwayConfig struct {
	Server     []GnetwayServer
	Multicore  bool
	SendBuf    int
	ReceiveBuf int
}

func (c GnetwayConfig) IsWebsocket(addr string) bool {
	for _, server := range c.Server {
		if server.Proto == "websocket" {
			for _, address := range server.Addresses {
				if address == addr {
					return true
				}
			}
		}
	}
	return false
}

func (c GnetwayConfig) IsHttp(addr string) bool {
	for _, server := range c.Server {
		if server.Proto == "http" {
			for _, address := range server.Addresses {
				if address == addr {
					return true
				}
			}
		}
	}
	return false
}

func (c GnetwayConfig) IsTcp(addr string) bool {
	for _, server := range c.Server {
		if server.Proto == "tcp" {
			for _, address := range server.Addresses {
				if address == addr {
					return true
				}
			}
		}
	}
	return false
}

func (c GnetwayConfig) ToAddresses() []string {
	var addresses []string
	for _, server := range c.Server {
		for _, address := range server.Addresses {
			if ok := container2.ContainsString(addresses, address); !ok {
				addresses = append(addresses, "tcp://"+address)
			}
		}
	}
	return addresses
}
