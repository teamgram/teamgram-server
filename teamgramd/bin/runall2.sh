#!/usr/bin/env bash

echo "run idgen ..."
nohup ./idgen -f=../etc/idgen.yaml >> ../logs/idgen.log  2>&1 &
sleep 1

echo "run status ..."
nohup ./status -f=../etc/status.yaml >> ../logs/status.log  2>&1 &
sleep 1

echo "run authsession ..."
nohup ./authsession -f=../etc/authsession.yaml >> ../logs/authsession.log  2>&1 &
sleep 1

echo "run dfs ..."
nohup ./dfs -f=../etc/dfs.yaml >> ../logs/dfs.log  2>&1 &
sleep 1

echo "run media ..."
nohup ./media -f=../etc/media.yaml >> ../logs/media.log  2>&1 &
sleep 1

#echo "run poll ..."
#nohup ./poll -f=../etc/poll.yaml >> ../logs/poll.log  2>&1 &
#sleep 1
#
#echo "run twofa ..."
#nohup ./twofa -f=../etc/twofa.yaml >> ../logs/twofa.log  2>&1 &
#sleep 1

echo "run biz ..."
nohup ./biz -f=../etc/biz.yaml >> ../logs/biz.log  2>&1 &
sleep 1

echo "run msg ..."
nohup ./msg -f=../etc/msg.yaml >> ../logs/msg.log  2>&1 &
sleep 1

echo "run sync ..."
nohup ./sync -f=../etc/sync.yaml >> ../logs/sync.log  2>&1 &
sleep 1

#echo "run scheduled ..."
#nohup ./scheduled -f=../etc/scheduled.yaml >> ../logs/scheduled.log  2>&1 &
#sleep 1
#
#echo "run push ..."
#nohup ./push -f=../etc/push.yaml >> ../logs/push.log  2>&1 &
#sleep 1
#
#echo "run sfu ..."
#nohup ./sfu -f=../etc/sfu.yaml >> ../logs/sfu.log  2>&1 &
#sleep 5
#
#echo "run teamcall ..."
#nohup ./teamcall -f=../etc/teamcall.yaml >> ../logs/teamcall.log  2>&1 &
#sleep 1
#
#echo "run botfather ..."
#nohup ./botfather -f=../etc/botfather.yaml >> ../logs/botfather.log  2>&1 &
#sleep 5
#
#echo "run messenger.bot ..."
#nohup ./bots -f=../etc/bots.yaml >> ../logs/bots.log  2>&1 &
#sleep 1
#
#echo "run messenger.messeages.bot ..."
#nohup ./bot -f=../etc/bot.yaml >> ../logs/bot.log  2>&1 &
#sleep 1
#
#echo "run phone ..."
#nohup ./phone -f=../etc/phone.yaml >> ../logs/phone.log  2>&1 &
#sleep 1

echo "run bff ..."
nohup ./bff -f=../etc/bff.yaml >> ../logs/bff.log  2>&1 &
sleep 5

echo "run session ..."
nohup ./session -f=../etc/session.yaml >> ../logs/session.log  2>&1 &
sleep 1

echo "run gateway ..."
nohup ./gateway -f=../etc/gateway.yaml >> ../logs/gateway.log  2>&1 &
sleep 1

#echo "run api_wallpaper ..."
#nohup ./wallpaper2 -f=../etc/wallpaper2.yaml >> ../logs/wallpaper2.log  2>&1 &
#sleep 1
#
#echo "run botway ..."
#nohup ./botway -f=../etc/botway.yaml >> ../logs/botway.log  2>&1 &
#sleep 1
#
#echo "run wsserver ..."
#nohup ./wsserver  >> ../logs/wsserver.log  2>&1 &
