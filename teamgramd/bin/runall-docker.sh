#!/usr/bin/env bash

echo "run idgen ..."
nohup ./idgen -f=../etc2/idgen.yaml >> ../logs/idgen.log  2>&1 &
sleep 1

echo "run status ..."
nohup ./status -f=../etc2/status.yaml >> ../logs/status.log  2>&1 &
sleep 1

echo "run authsession ..."
nohup ./authsession -f=../etc2/authsession.yaml >> ../logs/authsession.log  2>&1 &
sleep 1

echo "run dfs ..."
nohup ./dfs -f=../etc2/dfs.yaml >> ../logs/dfs.log  2>&1 &
sleep 1

echo "run media ..."
nohup ./media -f=../etc2/media.yaml >> ../logs/media.log  2>&1 &
sleep 1

echo "run biz ..."
nohup ./biz -f=../etc2/biz.yaml >> ../logs/biz.log  2>&1 &
sleep 1

echo "run msg ..."
nohup ./msg -f=../etc2/msg.yaml >> ../logs/msg.log  2>&1 &
sleep 1

echo "run sync ..."
nohup ./sync -f=../etc2/sync.yaml >> ../logs/sync.log  2>&1 &
sleep 1

echo "run bff ..."
nohup ./bff -f=../etc2/bff.yaml >> ../logs/bff.log  2>&1 &
sleep 5

echo "run session ..."
nohup ./session -f=../etc2/session.yaml >> ../logs/session.log  2>&1 &
sleep 1

#echo "run gateway ..."
#nohup ./gateway -f=../etc2/gateway.yaml >> ../logs/gateway.log  2>&1 &
#sleep 1

echo "run gnetway ..."
nohup ./gnetway -f=../etc/gnetway.yaml >> ../logs/gnetway.log  2>&1 &
sleep 1

#echo "run httpserver ..."
#nohup ./httpserver -f=../etc/httpserver.yaml >> ../logs/httpserver.log  2>&1 &
#sleep 1
