# teamgram-android

## Install

- Get *[teamgram-android](https://github.com/teamgram/teamgram-android)* source code

- build, see [build teamgram-android](https://github.com/teamgram/teamgram-android/blob/master/README.md), and google

## patch

**Default connect to Teamgram Test Server.**

If you want to connect to your own server, you can modify the following code:

[ConnectionsManager.cpp#L1684](https://github.com/teamgram/teamgram-android/blob/5b790e0fd27aab272ed5933968e9187ce7c61899/TMessagesProj/jni/tgnet/ConnectionsManager.cpp#L1684)

```
https://github.com/teamgram/teamgram-android/blob/5b790e0fd27aab272ed5933968e9187ce7c61899/TMessagesProj/jni/tgnet/ConnectionsManager.cpp#L1684

void ConnectionsManager::initDatacenters() {
    Datacenter *datacenter;
    if (!testBackend) {
        if (datacenters.find(1) == datacenters.end()) {
            datacenter = new Datacenter(instanceNum, 1);
            datacenter->addAddressAndPort("XXX.XXX.XXX.XXX", 10443, 0, "");
            datacenters[1] = datacenter;
        }
    } else {
        if (datacenters.find(1) == datacenters.end()) {
            datacenter = new Datacenter(instanceNum, 1);
            datacenter->addAddressAndPort("XXX.XXX.XXX.XXX", 10443, 0, "");
            datacenters[1] = datacenter;
        }
    }
}


```
