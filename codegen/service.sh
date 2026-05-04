#!/bin/bash

# interface
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gateway -app=gateway -name=interface.gateway -rpc=Gateway -port=20110 -helper=true
# tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/session -app=session -name=interface.session -rpc=Session -port=20120 -helper=true
# tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/interface/gnetway -app=gnetway -name=interface.gnetway -rpc=Gnetway -port=20110 -helper=true

# messenger
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg -app=msg -name=messenger.msg -rpc=Msg -port=20380 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/userupdates -app=userupdates -name=messenger.userupdates -rpc=Userupdates -port=30670 -helper=true
#tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/msg -app=msg -name=messenger.msg.msg -rpc=Msg -port=20380 -helper=true
# tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/msg/inbox -app=inbox -name=messenger.msg.inbox -rpc=Inbox -port=20390 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/messenger/sync -app=sync -name=messenger.sync -rpc=Sync -port=29530 -helper=true

# service
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/authsession -app=authsession -name=service.authsession -rpc=Authsession -port=20450 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/dfs -app=dfs -name=service.dfs -rpc=Dfs -port=20640 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/media -app=media -name=service.media -rpc=Media -port=20650 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/idgen -app=idgen -name=service.idgen -rpc=Idgen -port=20660 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/status -app=status -name=service.status -rpc=Status -port=20670 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/presence -app=presence -name=service.presence -rpc=Presence -port=20680 -helper=true

# service/biz
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/chat -app=chat -name=service.biz_service.chat -rpc=Chat -port=20500 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/code -app=code -name=service.biz_service.code -rpc=Code -port=29500 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/dialog -app=dialog -name=service.biz_service.dialog -rpc=Dialog -port=20260 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/message -app=message -name=service.biz_service.message -rpc=Message -port=20530 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/updates -app=updates -name=service.biz_service.updates -rpc=Updates -port=29510 -helper=true
tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/user -app=user -name=service.biz_service.user -rpc=User -port=20610 -helper=true
# tgctl teamgooo rpc --out=/opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app/service/biz/username -app=username -name=service.biz_service.username -rpc=Username -port=29520 -helper=true

find /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/app -name "*.go" | xargs gofmt -w
