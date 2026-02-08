#!/usr/bin/env bash
# Logs go to /app/logs/*.log (volume) for filebeat; works on Docker Desktop Mac/Win where
# /var/lib/docker/containers is not visible to filebeat container.

#mkdir -p ../logs

echo "run idgen ..."
./idgen -f=../etc2/idgen.yaml &
sleep 1

echo "run status ..."
./status -f=../etc2/status.yaml &
sleep 1

echo "run authsession ..."
./authsession -f=../etc2/authsession.yaml &
sleep 1

echo "run dfs ..."
./dfs -f=../etc2/dfs.yaml &
sleep 1

echo "run media ..."
./media -f=../etc2/media.yaml &
sleep 1

echo "run biz ..."
./biz -f=../etc2/biz.yaml &
sleep 1

echo "run msg ..."
./msg -f=../etc2/msg.yaml &
sleep 1

echo "run sync ..."
./sync -f=../etc2/sync.yaml &
sleep 1

echo "run bff ..."
./bff -f=../etc2/bff.yaml &
sleep 5

echo "run session ..."
./session -f=../etc2/session.yaml &
sleep 1

echo "run gnetway ..."
./gnetway -f=../etc2/gnetway.yaml &
sleep 1

# echo "run httpserver ..."
# ./httpserver -f=../etc/httpserver.yaml &
# sleep 1
