#!/usr/bin/env bash

PWD=`pwd`
TEAMGRAMAPP=${PWD}"/app"
echo ${PWD}

cd ${TEAMGRAMAPP}/messenger/msg/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/messenger/sync/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/bff/authorization/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/chat/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/message/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/user/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/updates/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/dialog/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/biz/username/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/authsession/internal/dal/dalgen
./dalgen_all.sh

cd ${TEAMGRAMAPP}/service/media/internal/dal/dalgen
./dalgen_all.sh

