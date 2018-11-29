# Handling PUSH-notifications
> Forwarded from [Handling PUSH-notifications](https://core.telegram.org/api/push-updates)

PUSH notifications for various events related to the current user are useful when the client is not running or a server connection has not been established.

## Configuring the Application
To be able to send APNS notifications to Apple servers or GCM notifications to Google servers, application certificates (APNS) or an application key (GCM) must be specified in the application settings.

## Subscribing to Notifications
To subscribe to notifications, the client must invoke the account.registerDevice query, passing in token_type and token as parameters that identify the current device. It is useful to repeat this query at least once every 24 hours or when restarting the application.
Use account.unregisterDevice to unsubscribe.

## Notification Structure
Each notification has several parameters that describe it.

### Notification Type
A string literal in the form /[A-Z_0-9]+/, which summarizes the notification. For example, CHAT_MESSAGE_TEXT.

### Notification Text Arguments
A list or arguments which, when inserted into a template, produce a readable notification.

### Opening Parameters
Parameters which are be passed into the application when a notification is opened.

### Sound
The name of an audio file to be played.

### Badge
New value for the unread notification counter.

## Processing Notifications
GCM
A GCM notification is provided as a JSON object in the following format:

```
{
  "collapse_key": "CHAT_MESSAGE_CONTACT",
  "data": {
    "loc_key": "CHAT_MESSAGE_CONTACT",
    "loc_args": ["John Doe", "Contact Exchange"],
    "text": "John Doe shared a contact in the group Contact Exchange",
    "custom": {
      "chat_id": 241233,
      "msg_id": 123
    },
    "badge": 1,
    "sound": "sound1.mp3",
    "mute": true
  }
}
```

In principle, data.loc_key, data.custom, and an Internet connection are sufficient to generate a notification. Obviously, if possible, when generating a visual notification you need not use all of the transmitted data and may rely on the information already stored on the client. But if a user or a chat is not cached locally, the values passed in loc_args may also be used.

## Possible Notifications
|Type	|Template Example|	Text Arguments|	Opening Parameters|
|:-:|:-:|:-:|:-:|
|MESSAGE_TEXT|	{1}: {2}	|1. Message author<br>2. Message body	|from_id: author identifier|
|MESSAGE_NOTEXT|	{1} sent you a message	|1. Message author|	from_id: author identifier|
|MESSAGE_PHOTO	|{1} sent you a photo	|1. Message author|	from_id: author identifier|
|MESSAGE_VIDEO	|{1} sent you a video	|1. Message author|	from_id: author identifier|
|MESSAGE_DOC	|{1} sent you a document	|1. Message author|	from_id - author identifier|
|MESSAGE_AUDIO	|{1} sent you a voice message|	1. Message author	|from_id - author identifier|
|MESSAGE_CONTACT	|{1} shared a contact with you|	1. Message author|	from_id: author identifier|
|MESSAGE_GEO	|{1} sent you a map	|1. Message author|	from_id: author identifier|
|CHAT_MESSAGE_TEXT	|{1}@{2}: {3}	|1. Message author<br>2. Chat name<br>3. Message body	|from_id: author identifier|
|CHAT_MESSAGE_NOTEXT|	{1} sent a message to the group {2}	|1. Message author<br>2. Chat name	|chat_id: chat identifier|
|CHAT_MESSAGE_PHOTO	|{1} sent a photo to the group {2}| 1. Message author<br>2. Chat name	|chat_id: chat identifier|
|CHAT_MESSAGE_VIDEO|	{1} sent a video to the group {2}|1. Message author<br>2. Chat name	|chat_id: chat identifier|
|CHAT_MESSAGE_DOC|	{1} sent a document to the group {2}| 1. Message author <br>2. Chat name	|chat_id - chat identifier|
|CHAT_MESSAGE_AUDIO	|{1} sent a voice message to the group {2}|1. Message author<br>2. Chat name|	chat_id - chat identifier|
|CHAT_MESSAGE_CONTACT	|{1} shared a contact in the group {2}|1. Message author<br>2. Chat name	|chat_id: chat identifier|
|`CHAT_MESSAGE_GEO`|	{1} sent a map to the group {2}|1. Message author<br>2. Chat name	|chat_id: chat identifier|
|CHAT_CREATED	|{1} invited you to the group {2}|1. Message author<br>2. Chat name	|chat_id: chat identifier|
|CHAT_TITLE_EDITED	|{1} edited the group's {2} name,|1. Message author <br> 2. New chat name	|chat_id: chat identifier|
|CHAT_PHOTO_EDITED	|{1} edited the group's {2} photo|1. Message author <br> 2. Chat name	|chat_id: chat identifier|
|`CHAT_ADD_MEMBER`	|{1} invited {3} to the group {2}|1. Message author<br>2. Chat name<br>3. New participant name | chat_id: chat identifier|
|`CHAT_ADD_YOU`	 | {1} invited you to the group {2} | 1. Message author<br>2. Chat name	|chat_id: chat identifier|
|`CHAT_DELETE_MEMBER`	|{1} kicked {3} from the group {2}|1. Message author <br> 2. Chat name<br>3. Dropped participant name|chat_id: chat identifier |
|`CHAT_DELETE_YOU`	|{1} kicked you from the group {2}|1. Message author <br>2. Chat name	|chat_id: chat identifier|
|CHAT_LEFT|	{1} has left the group {2} | 1. Message author <br>2. Chat name |	chat_id: chat identifier|
|CHAT_RETURNED|	{1} has returned to the group {2} | 1. Message author<br>2. Chat name	|chat_id: chat identifier|
|GEOCHAT_CHECKIN	|{1} has checked-in at {2}|1. Message author<br>2. Chat name	|geochat_id: GeoChat identifier|
|CONTACT_JOINED	|{1} joined the App!	|1. Contact name | contact_id: contact identifier|
|AUTH_UNKNOWN	|New login from unrecognized device {1}| 1. Device name||
|AUTH_REGION	|New login from unrecognized device {1}, location: {2}|1. Device name<br>2. Location||	
|CONTACT_PHOTO|{1} updated profile photo|1. Contact name |from_id: contact identifier|
|ENCRYPTION_REQUEST	|You have a new message|||
|ENCRYPTION_ACCEPT	|You have a new message|||
|ENCRYPTED_MESSAGE	|You have a new message|||		
## Service notifications
The following notifications can be used to update app settings.

|Type	| Text |	Opening Parameters |	Description|
|:-:|:-:|:-:|:-:|
|DC_UPDATE|	Open this notification to update app settings|dc - number of the data-center<br>addr - server address with port number in the format 111.112.113.114:443| In case the client gets this notification, it is necessary to add the received server address to the list of possible addresses. In case the address of the first DC was passed (dc=1), it is recommended to call it immediately using help.getConfig to update dc-configuration.|


