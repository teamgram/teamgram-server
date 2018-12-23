#!/bin/bash

PWD2=`pwd`
echo $PWD2
chatengine="$GOPATH/src/github.com/nebula-chat/chatengine"

echo "build document ..."
cd ${chatengine}/service/document
# go build
nohup ./document >> ${PWD2}/document.log 2>&1 &
sleep 1

echo "build auth_session ..."
cd ${chatengine}/service/auth_session
# go build
nohup ./auth_session >> ${PWD2}/auth_session.log 2>&1 &
sleep 1

echo "build sync ..."
cd ${chatengine}/messenger/sync
# go build
nohup ./sync >> ${PWD2}/sync.log 2>&1 &
sleep 1

echo "build upload ..."
cd ${chatengine}/messenger/upload
# go build
nohup ./upload >> ${PWD2}/upload.log 2>&1 &
sleep 1


echo "build auth_key ..."
cd ${chatengine}/access/auth_key
# go build
nohup ./auth_key >> ${PWD2}/auth_key.log 2>&1 &
sleep 1

echo "build biz_server ..."
cd ${chatengine}/messenger/biz_server
# go build
nohup ./biz_server >> ${PWD2}/biz_server.log 2>&1 &
sleep 1

echo "build session ..."
cd ${chatengine}/access/session
# go build
nohup ./session >> ${PWD2}/session.log 2>&1 &
sleep 1

echo "build frontend ..."
cd ${chatengine}/access/frontend
# go build
nohup ./frontend >> ${PWD2}/frontend.log 2>&1 &
sleep 1

cd $PWD2

