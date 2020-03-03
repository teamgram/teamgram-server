FROM golang:1.12.12 AS builder

ENV CGO_ENABLED 0
ENV TARGET_DIR $GOPATH/src/github.com/nebula-chat/chatengine

RUN echo $GOPATH
RUN mkdir -p $TARGET_DIR
RUN cd $TARGET_DIR
COPY . $TARGET_DIR/

# build biz_server
RUN cd ${TARGET_DIR}/messenger/biz_server && go build -ldflags='-s -w'
# build document
RUN cd ${TARGET_DIR}/service/document && go build -ldflags='-s -w'
# build auth_session
RUN cd ${TARGET_DIR}/service/auth_session && go build -ldflags='-s -w'
# build sync
RUN cd ${TARGET_DIR}/messenger/sync && go build -ldflags='-s -w'
# build upload
RUN cd ${TARGET_DIR}/messenger/upload && go build -ldflags='-s -w'
# build auth_key
RUN cd ${TARGET_DIR}/access/auth_key && go build -ldflags='-s -w'
# build session
RUN cd ${TARGET_DIR}/access/session && go build -ldflags='-s -w'
# build frontend
RUN cd ${TARGET_DIR}/access/frontend && go build -ldflags='-s -w'



FROM alpine:3.10.3
RUN apk add --no-cache ca-certificates tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV TARGET_DIR /go/src/github.com/nebula-chat/chatengine
WORKDIR /app/

COPY ./docker/entrypont.sh /app/

RUN mkdir -p /app/config-templates

# build document
COPY --from=builder ${TARGET_DIR}/service/document/document ./
# build auth_session
COPY --from=builder ${TARGET_DIR}/service/auth_session/auth_session ./
# build sync
COPY --from=builder ${TARGET_DIR}/messenger/sync/sync ./
# build upload
COPY --from=builder ${TARGET_DIR}/messenger/upload/upload ./
# build biz_server
COPY --from=builder ${TARGET_DIR}/messenger/biz_server/biz_server ./
# build auth_key
COPY --from=builder ${TARGET_DIR}/access/auth_key/auth_key ./
# build session
COPY --from=builder ${TARGET_DIR}/access/session/session ./
# build frontend
COPY --from=builder ${TARGET_DIR}/access/frontend/frontend ./

# copy configs
COPY --from=builder ${TARGET_DIR}/scripts/config/*.toml ./config-templates/
COPY --from=builder ${TARGET_DIR}/scripts/config/*.json ./config-templates/
COPY --from=builder ${TARGET_DIR}/scripts/config/*.key ./config-templates/

ENTRYPOINT sh /app/entrypont.sh
