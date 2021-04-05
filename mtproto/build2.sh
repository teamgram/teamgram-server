#!/bin/sh
SRC_DIR=/home/anhttn/go/src/github.com/nebula-chat/chatengine/mtproto/
DST_DIR=/home/anhttn/Desktop/mtproto/

#./codegen.sh
#protoc -I=$SRC_DIR --go_out=$DST_DIR/ $SRC_DIR/*.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/*.proto

gofmt -w *.go
