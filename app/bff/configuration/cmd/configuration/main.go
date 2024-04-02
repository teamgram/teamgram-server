/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package main

import (
	"github.com/teamgram/marmota/pkg/commands"

	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/server"
)

func main() {
	commands.Run(server.New())
}
