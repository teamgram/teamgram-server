# Working with Updates
> Forwarded from [Working with Updates](https://core.telegram.org/api/updates)

When a client is being actively used, events will occur that affect the current user and that he or she must learn about as soon as possible. For example, when someone sends him or her a message. To eliminate the need for the client itself to periodically download these events, there is an update delivery mechanism in which the server sends the user notifications over one of its available connections with the client.

## Subscribing to Updates
When any high-level API query is executed on behalf of an authorized user, the session identifier (session_id) will be stored in association with the user identifier and the current authorization key identifier (auth_key_id). This applies to all methods except these file handling methods:

- upload.saveFilePart
- upload.getFile

Thereafter, whenever there is an event that the user needs to be notified of, the server will find a list of the 10 most recent active sessions and send messages to those sessions. Note that a session is also associated with an authorization key, and only one notification will be sent at a time, per key, to the most recent session. Other sessions will be ignored. Therefore, you should not start several sessions for regular API queries, but if you must, all messages with updates must be processed in each of them.

An exception has been made for the file handling methods; therefore, separate sessions may be established for uploading/downloading files.

## Update Status
Before processing updates, the current status must be obtained: updates.State. This could be done using a call to updates.getState or updates.getDifference.

The status must always be stored on the client and kept up to date. Each message received from the server may change this status; hence, it is important to process the pts, date, and seq fields on received updates, and when calling certain methods.

## Processing Incoming Messages
Each update message is of the type Updates. It can be seen from the schema below that this type has several constructors.

Messages with updateShort constructors, normally, have lower priority and are broadcast to a large number of users, i.e. one of the chat participants started entering text in a big conversation (updateChatUserTyping).

The updates constructor may contain the sequence number of the current user’s update (seq > 0). Once received, it is useful to compare it with the value on the client. If the difference is greater than one, then the client is out of sync (normally, as a result of the client disconnecting or the server rebooting).

The updateShortMessage and updateShortChatMessage constructors are redundant but help significantly reduce the transmitted message size for 90% of the updates.


```
updatesTooLong#e317af7e = Updates;
updateShortMessage#d3f45784 id:int from_id:int message:string pts:int date:int seq:int = Updates;
updateShortChatMessage#2b2fbd4e id:int from_id:int chat_id:int message:string pts:int date:int seq:int = Updates;
updateShort#78d4dec1 update:Update date:int = Updates;
updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
```

### Combined Updates
The server can also regroup updates. When a client does not maintain an active connection and does not receive updates, events may be regrouped before the final delivery in order to combine them into a single message. In doing so, obsolete updates, such as updateUserTyping will be deleted, while updates such as updateUserStatus will be regrouped so that only the most recent updates from each user_id remain. It may be that one message will contain events from updates with different sequential values (**seq**). In this case, the updatesCombined constructor will be used. It defines the range of seq. When such a constructor is received, it should be checked that seq_start is exactly one greater than the client’s value for seq, and the current value of seq should be overwritten with the value in the update’s seq field.

If the final message becomes too large (the current limit is 200 KB), the updatesTooLong constructor will be used. If a client receives this constructor, it means the difference must be obtained manually.

### Compressing Updates
All updates can also be compressed using gzip in the same way as responses to queries.

## Obtaining Differences
Manually obtaining updates is required in the following situations:

- Loss of sync: an update contains seq / seq_start (> 0) and it is not equal to client_seq + 1. It may be useful to wait up to 0.5 seconds in this situation and abort the sync in case a new update arrives, that fixes the ‘hole’.
	
- Session loss on the server: the client receives a new session created notification. This can be caused by garbage collection on the MTProto server or a server reboot.

- Incorrect update: the client cannot deserialize the received data.

- Incomplete update: the client is missing data about a chat/user from one of the shortened constructors, such as updateShortChatMessage, etc.

- Long period without updates: no updates for 15 minutes or longer.

To manually obtain the difference between the client’s status and the updated status, a call to updates.getDifference is made, passing in the parameters of the client’s current state. If the updates.differenceSlice constructor is returned in the response, the full difference was too large to be received by the client all at once. The intermediate status, intermediate_state, must be saved on the client and the query must be repeated, using the intermediate status as the current status.


## PUSH Notifications about Updates
If a client does not have an active connection at the time of an event, PUSH Notifications will also be useful.

