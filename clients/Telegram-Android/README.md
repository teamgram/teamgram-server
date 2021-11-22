# [Telegram-Android](https://github.com/DrKLO/Telegram) patch
> [Telegram-Android](https://github.com/DrKLO/Telegram) patched by [NebulaChat](https://nebula.chat)

## Install

- Get *[Telegram-Android](https://github.com/DrKLO/Telegram)* source code


- Switch to e222fded6cca5ace3649be6f18b55f526311bc79 

```
git checkout e222fded6cca5ace3649be6f18b55f526311bc79
```

- patch

- build, see [build Telegram-Android](https://github.com/DrKLO/Telegram/blob/master/README.md), and google

## Replace your server and port in ConnectionManager.cpp.diff file

**Default connect to NebulaChat test server.**

If you want to connect to your own server, you can modify the following code:

```
Telegram-Android/ConnectionManager.cpp.diff
L10

+    std::string _nebulaChatServer("47.103.102.219");

```
