#!/usr/bin/env bash

echo "run idgen ..."
nohup ./idgen -f=../app/service/idgen/etc/idgen.yaml >> ../logs/idgen.log  2>&1 &
sleep 1

echo "run status ..."
nohup ./status -f=../app/service/status/etc/status.yaml >> ../logs/status.log  2>&1 &
sleep 1

echo "run authsession ..."
nohup ./authsession -f=../app/service/authsession/etc/authsession.yaml >> ../logs/authsession.log  2>&1 &
sleep 1

echo "run dfs ..."
nohup ./dfs -f=../app/service/dfs/etc/dfs.yaml >> ../logs/dfs.log  2>&1 &
sleep 1

echo "run media ..."
nohup ./media -f=../app/service/dfs/media/media.yaml >> ../logs/media.log  2>&1 &
sleep 1

echo "run biz ..."
nohup ./biz -f=../app/service/biz/etc/biz.yaml >> ../logs/biz.log  2>&1 &
sleep 1

echo "run userupdates ..."
nohup ./userupdates -f=../app/messenger/userupdates/etc/userupdates.yaml >> ../logs/userupdates.log  2>&1 &
sleep 1

echo "run msg ..."
nohup ./msg -f=../app/messenger/msg/etc/msg.yaml >> ../logs/msg.log  2>&1 &
sleep 1

echo "run bff ..."
nohup ./bff -f=../app/bff/bff/etc/bff.yaml >> ../logs/bff.log  2>&1 &
sleep 5

echo "run gateway ..."
nohup ./gateway -f=../app/interface/gateway/etc/gateway.yaml >> ../logs/gateway.log  2>&1 &
sleep 1
