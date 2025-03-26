#!/usr/bin/env bash

echo "run gnetway ..."
nohup ./gnetway -f=../etc/gnetway.yaml >> ../logs/gnetway.log  2>&1 &
sleep 1

#echo "run httpserver ..."
#nohup ./httpserver -f=../etc/httpserver.yaml >> ../logs/httpserver.log  2>&1 &
#sleep 1
