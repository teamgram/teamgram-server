# [teamgram-ios](https://github.com/teamgram/teamgram-ios)

## Install

- Get *[teamgram-ios](https://github.com/teamgram/teamgram-ios)* source code

- build, see [build teamgram-ios](https://github.com/teamgram/teamgram-ios), and google

## Patch

**Default connect to Teamgram Test Server.**

If you want to connect to your own server, you can modify the following code:

[Network.swift#L473](https://github.com/teamgram/teamgram-ios/blob/9076f397e48485b88cf0f3b9aa91deb0f311587c/submodules/TelegramCore/Sources/Network/Network.swift#L473)

```
https://github.com/teamgram/teamgram-ios/blob/9076f397e48485b88cf0f3b9aa91deb0f311587c/submodules/TelegramCore/Sources/Network/Network.swift#L473

if testingEnvironment {
    seedAddressList = [
        1: ["XXX.XXX.XXX.XXX"]
    ]
} else {
    seedAddressList = [
        1: ["XXX.XXX.XXX.XXX"]
    ]
}
```
