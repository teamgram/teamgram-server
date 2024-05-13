VERSION=v0.96.0-teamgram-server
BUILD=`date +%F`
SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir="github.com/teamgram/marmota/pkg/version"
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch}"

all: idgen status dfs media authsession biz msg sync bff session gateway gnetway

idgen:
	@echo "build idgen..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/idgen -tags=jsoniter app/service/idgen/cmd/idgen/*.go

status:
	@echo "build status..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/status -tags=jsoniter app/service/status/cmd/status/*.go

dfs:
	@echo "build dfs..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/dfs -tags=jsoniter app/service/dfs/cmd/dfs/*.go

media:
	@echo "build media..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/media -tags=jsoniter app/service/media/cmd/media/*.go

authsession:
	@echo "build authsession..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/authsession -tags=jsoniter app/service/authsession/cmd/authsession/*.go

biz:
	@echo "build biz..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/biz -tags=jsoniter app/service/biz/biz/cmd/biz/*.go

msg:
	@echo "build msg..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/msg -tags=jsoniter app/messenger/msg/cmd/msg/*.go

sync:
	@echo "build sync..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/sync -tags=jsoniter app/messenger/sync/cmd/sync/*.go

bff:
	@echo "build bff..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/bff -tags=jsoniter app/bff/bff/cmd/bff/*.go

session:
	@echo "build session..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/session -tags=jsoniter app/interface/session/cmd/session/*.go

gateway:
	@echo "build gateway..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/gateway -tags=jsoniter app/interface/gateway/cmd/gateway/*.go

gnetway:
	@echo "build gnetway..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/gnetway -tags=jsoniter app/interface/gnetway/cmd/gnetway/*.go

httpserver:
	@echo "build httpserver..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/httpserver -tags=jsoniter app/interface/httpserver/cmd/httpserver/*.go

clean:
	@rm -rf teamgramd/bin/idgen
	@rm -rf teamgramd/bin/status
	@rm -rf teamgramd/bin/dfs
	@rm -rf teamgramd/bin/media
	@rm -rf teamgramd/bin/authsession
	@rm -rf teamgramd/bin/biz
	@rm -rf teamgramd/bin/msg
	@rm -rf teamgramd/bin/sync
	@rm -rf teamgramd/bin/bff
	@rm -rf teamgramd/bin/session
	@rm -rf teamgramd/bin/gateway
	@rm -rf teamgramd/bin/gnetway
	@rm -rf teamgramd/bin/httpserver
