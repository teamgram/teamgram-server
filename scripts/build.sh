#!/usr/bin/env bash

teamgram="$GOPATH/src/github.com/teamgram/teamgram-server/app"
INSTALL="$GOPATH/src/github.com/teamgram/teamgram-server/teamgramd"

echo "build idgen ..."
cd ${teamgram}/service/idgen/cmd/idgen
go build -o ${INSTALL}/bin/idgen

echo "build status ..."
cd ${teamgram}/service/status/cmd/status
go build -o ${INSTALL}/bin/status

echo "build dfs ..."
cd ${teamgram}/service/dfs/cmd/dfs
go build -o ${INSTALL}/bin/dfs

echo "build media ..."
cd ${teamgram}/service/media/cmd/media
go build -o ${INSTALL}/bin/media

echo "build authsession ..."
cd ${teamgram}/service/authsession/cmd/authsession
go build -o ${INSTALL}/bin/authsession

#echo "build poll ..."
#cd ${teamgram}/service/poll/cmd/poll
#go build -o ${INSTALL}/bin/poll

#echo "build twofa ..."
#cd ${teamgram}/service/twofa/cmd/twofa
#go build -o ${INSTALL}/bin/twofa

echo "build biz ..."
cd ${teamgram}/service/biz/biz/cmd/biz
go build -o ${INSTALL}/bin/biz

echo "build msg ..."
cd ${teamgram}/messenger/msg/cmd/msg
go build -o ${INSTALL}/bin/msg

echo "build sync ..."
cd ${teamgram}/messenger/sync/cmd/sync
go build -o ${INSTALL}/bin/sync

#echo "build push ..."
#cd ${teamgram}/messenger/push/cmd/push
#go build -o ${INSTALL}/bin/push

#echo "build scheduled ..."
#cd ${teamgram}/job/scheduled/cmd/scheduled
#go build -o ${INSTALL}/bin/scheduled

echo "build bff ..."
cd ${teamgram}/bff/bff/cmd/bff
go build -o ${INSTALL}/bin/bff

echo "build session ..."
cd ${teamgram}/interface/session/cmd/session
go build -o ${INSTALL}/bin/session

echo "build gateway ..."
cd ${teamgram}/interface/gateway/cmd/gateway
go build -o ${INSTALL}/bin/gateway

#echo "build api_wallpaper ..."
#cd ${teamgram}/admin/wallpaper2/cmd/wallpaper2
#go build -o ${INSTALL}/bin/wallpaper2

#echo "build wsserver ..."
#cd ${teamgram}/interface/wsserver/cmd/wsserver
#go build -o ${INSTALL}/bin/wsserver
