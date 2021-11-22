# [tdesktop](https://github.com/telegramdesktop/tdesktop) patch
> [tdesktop](https://github.com/telegramdesktop/tdesktop) patched by [NebulaChat](https://nebula.chat)

## Install

- Get *tdesktop* source code

- Switch to dc8abc74ed4d72a73315550b91283ff1f2e44199 

```
git checkout dc8abc74ed4d72a73315550b91283ff1f2e44199
```

- patch

- build, see [Build tdesktop](https://github.com/telegramdesktop/tdesktop/blob/dev/README.md#build-instructions)

## Edit your server and port in config.h.diff file

**Default connect to NebulaChat test server.**

If you want to connect to your own server, you can modify the following code:

```
tdesktop/config.h.diff
L24

+#define NEBULAIM_DC_IP4   "47.103.102.219"

```
