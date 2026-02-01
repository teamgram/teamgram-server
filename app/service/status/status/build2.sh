#!/bin/sh

SRC_DIR=.
DST_DIR=$GOPATH/src/

#GOGOPROTO_PATH=$GOPATH/src/github.com/gogo/protobuf/protobuf
MTPROTO_PATH=$GOPATH/src/github.com/teamgram/proto/mtproto

protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/src:./ --go_out=$DST_DIR --go-grpc_out=require_unimplemented_servers=false:$DST_DIR $SRC_DIR/*.proto

#protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/src:$GOGOPROTO_PATH:./ \
#    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
#    $SRC_DIR/*.proto
#protoc -I=$SRC_DIR --proto_path=$GOPATH/src:$GOPATH/src/nebula.chat/vendor:$GOGOPROTO_PATH:./ \
#    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
#    $SRC_DIR/rpc_error_codes.proto

gofmt -w *.go
