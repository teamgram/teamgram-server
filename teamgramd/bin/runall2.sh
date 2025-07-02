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

echo "run bff ..."
nohup ./bff -f=../etc/bff.yaml >> ../logs/bff.log  2>&1 &
sleep 5

echo "run session ..."
nohup ./session -f=../etc/session.yaml >> ../logs/session.log  2>&1 &
sleep 1

echo "run gnetway ..."
nohup ./gnetway -f=../etc/gnetway.yaml >> ../logs/gnetway.log  2>&1 &
sleep 1

#echo "run httpserver ..."
#nohup ./httpserver -f=../etc/httpserver.yaml >> ../logs/httpserver.log  2>&1 &
#sleep 1
