# [teamgram-tdesktop](https://github.com/teamgram/teamgram-tdesktop)

## Install

- Get *[teamgram-tdesktop](https://github.com/teamgram/teamgram-tdesktop)* source code

- build, see [Build teamgram-tdesktop](https://github.com/teamgram/teamgram-tdesktop#build-instructions)

## Patch

**Default connect to Teamgram Test Server.**

If you want to connect to your own server, you can modify the following code:

[mtproto_dc_options.cpp#L31](https://github.com/teamgram/teamgram-tdesktop/blob/teamgram2/Telegram/SourceFiles/mtproto/mtproto_dc_options.cpp#L31)

```
https://github.com/teamgram/teamgram-tdesktop/blob/teamgram2/Telegram/SourceFiles/mtproto/mtproto_dc_options.cpp#L31

const BuiltInDc kBuiltInDcs[] = {
    { 1, "XXX.XXX.XXX.XXX" , 10443 },
};

const BuiltInDc kBuiltInDcsTest[] = {
    { 1, "XXX.XXX.XXX.XXX" , 10443 },
};

```
