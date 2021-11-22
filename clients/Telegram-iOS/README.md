# [Telegram-iOS](https://github.com/peter-iakovlev/Telegram) patch
> [Telegram-iOS](https://github.com/peter-iakovlev/Telegram) patched by [NebulaChat](https://nebula.chat)

## Install

- Get *[Telegram-iOS](https://github.com/peter-iakovlev/Telegram)* source code

- patch
```
MTDatacenterAuthMessageService.m.diff ==> submodules/MtProtoKit/MTProtoKit/MTDatacenterAuthMessageService.m

TGShareContextSignal.m.diff ==> LegacyDatabase/TGShareContextSignal.m
TGTelegramNetworking.m.diff ==> Telegraph/TGTelegramNetworking.m
```

- build, see [build Telegram-iOS](https://github.com/peter-iakovlev/Telegram), and google

## Replace your server and port in *TGTelegramNetworking.m.diff* and *TGShareContextSignal.m.diff*

**Default connect to NebulaChat test server.**

