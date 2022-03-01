#!/bin/bash

dalgen2 --xml=$1 --db=teamgram --go2=github.com/teamgram/teamgram-server/app/messenger/biz_server/auth/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
