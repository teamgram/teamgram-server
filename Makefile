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

all: gnetway
gnetway:
	@echo "build gnetway..."
	@go build -ldflags ${ldflags} -o teamgramd/bin/gnetway -tags=jsoniter app/interface/gnetway/cmd/gnetway/*.go

clean:
	@rm -rf teamgramd/bin/gnetway
