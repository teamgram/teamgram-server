#!/usr/bin/env bash

# todo(yumcoder) change dc ip
# sed -i '/ipAddress = /c\ipAddress = 127.0.0.1' a.txt
# todo(yumcoder) change folder path for nbfs

# docker start mysql-docker redis-docker etcd-docker

chatengine="$GOPATH/src/github.com/nebula-chat/chatengine"

echo "build document ..."
cd ${chatengine}/service/document
go build
#./document &
sleep 1

echo "build auth_session ..."
cd ${chatengine}/service/auth_session
go build
#./auth_session &
sleep 1

echo "build sync ..."
cd ${chatengine}/messenger/sync
go build
#./sync &
sleep 1

echo "build upload ..."
cd ${chatengine}/messenger/upload
go build
#./upload &
sleep 1


echo "build auth_key ..."
cd ${chatengine}/access/auth_key
go build
#./auth_key &

echo "build biz_server ..."
cd ${chatengine}/messenger/biz_server
go build
#./biz_server &
sleep 1

echo "build session ..."
cd ${chatengine}/access/session
go build
#./session &
sleep 1

echo "build frontend ..."
cd ${chatengine}/access/frontend
go build
#./frontend &
sleep 1

echo "***** wait *****"
wait
