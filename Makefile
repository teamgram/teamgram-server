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

all: geoip idgen presence dfs media mediaprocessor authsession biz userupdates msg sync bff gateway

lint: lint-tg-primitives

lint-tg-primitives:
	@./scripts/lint-tg-primitives.sh

geoip:
	@echo "build geoip..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/geoip -tags=jsoniter app/infra/geoip/cmd/geoip/*.go

idgen:
	@echo "build idgen..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/idgen -tags=jsoniter app/service/idgen/cmd/idgen/*.go

presence:
	@echo "build presence..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/presence -tags=jsoniter app/service/presence/cmd/presence/*.go

dfs:
	@echo "build dfs..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/dfs -tags=jsoniter app/service/dfs/cmd/dfs/*.go

media:
	@echo "build media..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/media -tags=jsoniter app/service/media/cmd/media/*.go

mediaprocessor:
	@echo "build mediaprocessor..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/mediaprocessor -tags=jsoniter app/service/mediaprocessor/cmd/mediaprocessor/*.go

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

sync:
	@echo "build sync..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/sync -tags=jsoniter app/messenger/sync/cmd/sync/*.go

bff:
	@echo "build bff..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/bff -tags=jsoniter app/bff/bff/cmd/bff/*.go

gateway:
	@echo "build gateway..."
	@go build -ldflags ${ldflags} -o ${INSTALL}/bin/gateway -tags=jsoniter app/interface/gateway/cmd/gateway/*.go

clean:
	@rm -rf ${INSTALL}/bin/geoip
	@rm -rf ${INSTALL}/bin/idgen
	@rm -rf ${INSTALL}/bin/presence
	@rm -rf ${INSTALL}/bin/dfs
	@rm -rf ${INSTALL}/bin/media
	@rm -rf ${INSTALL}/bin/mediaprocessor
	@rm -rf ${INSTALL}/bin/authsession
	@rm -rf ${INSTALL}/bin/biz
	@rm -rf ${INSTALL}/bin/userupdates
	@rm -rf ${INSTALL}/bin/msg
	@rm -rf ${INSTALL}/bin/sync
	@rm -rf ${INSTALL}/bin/bff
	@rm -rf ${INSTALL}/bin/gateway
