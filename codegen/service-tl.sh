#!/bin/bash

# interface
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gateway/gateway --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gateway/gateway
# tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gnetway/gnetway --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gnetway/gnetway
# tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/session/session --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/session/session

# messenger
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/userupdates/userupdates --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/userupdates/userupdates
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/msg --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/msg
# tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/inbox/inbox --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/inbox/inbox
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/sync/sync --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/sync/sync

# service
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/idgen/idgen --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/idgen/idgen
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/status/status --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/status/status
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/authsession/authsession --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/authsession/authsession
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/media/media --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/media/media
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/dfs/dfs --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/dfs/dfs
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/presence/presence --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/presence/presence

# service/biz
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/chat/chat --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/chat/chat
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/code/code --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/code/code
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/dialog/dialog --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/dialog/dialog
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/message/message --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/message/message
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/updates/updates --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/updates/updates
tgctl teamgooo mtprotoc service --schemas=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/user/user --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/user/user

find /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app -name "*.go" | xargs gofmt -w
