VERSION=v0.224.0-teamgooo-server
BUILD=`date +%F`
SHELL := /bin/bash
BASEDIR = $(shell pwd)
### "teamgoood"
INSTALL="./teamgooo"

# build with verison infos
versionDir="github.com/teamgram/marmota/pkg/version"
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch}"

all: geoip idgen status dfs media authsession biz userupdates msg bff gateway

lint: lint-tg-primitives

lint-tg-primitives:
	@./scripts/lint-tg-primitives.sh

geoip:
	@echo "build geoip..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/geoip -tags=jsoniter app/infra/geoip/cmd/geoip/*.go

idgen:
	@echo "build idgen..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/idgen -tags=jsoniter app/service/idgen/cmd/idgen/*.go

status:
	@echo "build status..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/status -tags=jsoniter app/service/status/cmd/status/*.go

dfs:
	@echo "build dfs..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/dfs -tags=jsoniter app/service/dfs/cmd/dfs/*.go

media:
	@echo "build media..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/media -tags=jsoniter app/service/media/cmd/media/*.go

authsession:
	@echo "build authsession..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/authsession -tags=jsoniter app/service/authsession/cmd/authsession/*.go

biz:
	@echo "build biz..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/biz -tags=jsoniter app/service/biz/biz/cmd/biz/*.go

userupdates:
	@echo "build userupdates..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/userupdates -tags=jsoniter app/messenger/userupdates/cmd/userupdates/*.go

msg:
	@echo "build msg..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/msg -tags=jsoniter app/messenger/msg/cmd/msg/*.go

bff:
	@echo "build bff..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/bff -tags=jsoniter app/bff/bff/cmd/bff/*.go

gateway:
	@echo "build gateway..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/gateway -tags=jsoniter app/interface/gateway/cmd/gateway/*.go

clean:
	@rm -rf ${INSTALL}/bin/geoip
	@rm -rf ${INSTALL}/bin/idgen
	@rm -rf ${INSTALL}/bin/status
	@rm -rf ${INSTALL}/bin/dfs
	@rm -rf ${INSTALL}/bin/media
	@rm -rf ${INSTALL}/bin/authsession
	@rm -rf ${INSTALL}/bin/biz
	@rm -rf ${INSTALL}/bin/userupdates
	@rm -rf ${INSTALL}/bin/msg
	@rm -rf ${INSTALL}/bin/bff
	@rm -rf ${INSTALL}/bin/gateway
