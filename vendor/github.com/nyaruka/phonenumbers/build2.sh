#!/bin/sh

PWD2=$PWD

SRC_DIR=.
DST_DIR=.

GOGOPROTO_PATH=$GOPATH/src/nebula.chat/vendor/github.com/gogo/protobuf/protobuf

protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/src:$GOPATH/src/nebula.chat/vendor:$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
    $SRC_DIR/*.proto

#gofmt -w codec_schema.tl.pb.go


