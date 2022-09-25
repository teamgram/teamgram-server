# [teamgram-ios](https://github.com/teamgram/teamgram-ios)

## Install

- Get *[teamgram-ios](https://github.com/teamgram/teamgram-ios)* source code
```
mkdir ~/Teamgram
cd ~/Teamgram
git clone --recursive https://github.com/teamgram/teamgram-ios.git
```

- build, see [build teamgram-ios](https://github.com/teamgram/teamgram-ios#compilation-guide), and google
```
cd ~/Teamgram/teamgram-ios
sh r.sh
```

## Patch

**Default connect to Teamgram Test Server.**

If you want to connect to your own server, you can modify the following code:

[Network.swift#L473](https://github.com/teamgram/teamgram-ios/blob/teamgram/submodules/TelegramCore/Sources/Network/Network.swift#L473)

```
https://github.com/teamgram/teamgram-ios/blob/teamgram/submodules/TelegramCore/Sources/Network/Network.swift#L473
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
