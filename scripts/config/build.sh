#!/bin/bash

PWD2=`pwd`
echo $PWD2
chatengine="$GOPATH/src/github.com/nebula-chat/chatengine"

echo "build document ..."
cd ${chatengine}/service/document
go build

echo "build auth_session ..."
cd ${chatengine}/service/auth_session
go build

echo "build sync ..."
cd ${chatengine}/messenger/sync
go build

echo "build upload ..."
cd ${chatengine}/messenger/upload
go build


echo "build auth_key ..."
cd ${chatengine}/access/auth_key
go build

echo "build biz_server ..."
cd ${chatengine}/messenger/biz_server
go build

echo "build session ..."
cd ${chatengine}/access/session
go build

echo "build frontend ..."
cd ${chatengine}/access/frontend
go build

cd ${PWD2}
