# Protocol and compatibility

This document describes the MTProto sub-protocols, API Layer, and client compatibility for Teamgram Server.

## MTProto 2.0

Teamgram implements [MTProto 2.0](https://core.telegram.org/mtproto) with these transports:

- **Abridged**
- **Intermediate**
- **Padded intermediate**
- **Full**

Clients connect to gnetway over TCP or WebSocket and use one of the above for handshake and RPC.

## API Layer

- Supported **API Layer** is **222** (see main README; may be updated with proto).
- TL methods and types follow the project’s [teamgram/proto](https://github.com/teamgram/proto) dependency.

## Compatible clients

| Client | Notes |
|--------|--------|
| [Android (teamgram-android)](../clients/teamgram-android.md) | Patch server address/port in ConnectionsManager.cpp |
| [iOS (teamgram-ios)](../clients/teamgram-ios.md) | Configure to point to your server |
| [Desktop (teamgram-tdesktop)](../clients/teamgram-tdesktop.md) | Configure to point to your server |

**Important:** Default sign-in verification code is **12345** (development only; change for production).

## Deployment notes

- Clients must reach gnetway ports (default 10443, 11443, 5222). Adjust TLS or reverse proxy (e.g. Nginx for WebSocket) as needed.
