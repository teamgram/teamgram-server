#!/usr/bin/env bash

# todo(yumcoder) change dc ip
# sed -i '/ipAddress = /c\ipAddress = 127.0.0.1' a.txt
# todo(yumcoder) change folder path for nbfs

docker start mysql-docker redis-docker etcd-docker

chatengine="$GOPATH/src/github.com/nebula-chat/chatengine"


echo "build document ..."
cd ${chatengine}/service/document
go get
go build
./document &

echo "build auth_session ..."
cd ${chatengine}/service/auth_session
go get
go build
./auth_session &

echo "build sync ..."
cd ${chatengine}/messenger/sync
go get
go build
./sync &

echo "build upload ..."
cd ${chatengine}/messenger/upload
go get
go build
./upload &

echo "build biz_server ..."
cd ${chatengine}/messenger/biz_server
go get
go build
./biz_server &

echo "build session ..."
cd ${chatengine}/server/access/session
go get
go build
./session &

echo "build frontend ..."
cd ${chatengine}/server/access/frontend
go get
go build
./frontend &

echo "build auth_key ..."
cd ${chatengine}/server/access/auth_key
go get
go build
./auth_key &

echo "***** wait *****"
wait
