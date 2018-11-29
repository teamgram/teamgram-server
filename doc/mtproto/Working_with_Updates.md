# Working with Updates

When a client is being actively used, events will occur that affect the current user and that he or she must learn about as soon as possible. For example, when someone sends him or her a message. To eliminate the need for the client itself to periodically download these events, there is an update delivery mechanism in which the server sends the user notifications over one of its available connections with the client.

当客户端被主动使用时，将发生影响当前用户的事件，并且他或她必须尽快了解。例如，当有人向他或她发送消息时。为了消除客户端本身定期下载这些事件的需要，有一种更新传递机制，其中服务器通过其与客户端的一个可用连接发送用户通知。

## Subscribing to Updates
When any high-level API query is executed on behalf of an authorized user, the session identifier (session_id) will be stored in association with the user identifier and the current authorization key identifier (auth_key_id). This applies to all methods except these file handling methods:

当代表授权用户执行任何高级API查询时，会话标识符（session_id）将与用户标识符和当前授权密钥标识符（auth_key_id）相关联地存储。这适用于除这些文件处理方法之外的所有方法：

- upload.saveFilePart
- upload.getFile

Thereafter, whenever there is an event that the user needs to be notified of, the server will find a list of the 10 most recent active sessions and send messages to those sessions. Note that a session is also associated with an authorization key, and only one notification will be sent at a time, per key, to the most recent session. Other sessions will be ignored. Therefore, you should not start several sessions for regular API queries, but if you must, all messages with updates must be processed in each of them.

此后，每当存在需要通知用户的事件时，服务器将找到10个最近活动会话的列表并向这些会话发送消息。请注意，会话还与授权密钥相关联，并且每个密钥一次仅向最近的会话发送一个通知。其他会话将被忽略。因此，您不应为常规API查询启动多个会话，但如果必须，则必须在每个会话中处理所有包含更新的消息。

An exception has been made for the file handling methods; therefore, separate sessions may be established for uploading/downloading files.

文件处理方法已经例外; 因此，可以建立单独的会话来上载/下载文件。


## Update Status
Before processing updates, the current status must be obtained: updates.State. This could be done using a call to updates.getState or updates.getDifference.

在处理更新之前，必须获取当前状态：updates.State。这可以通过调用updates.getState或updates.getDifference来完成。

The status must always be stored on the client and kept up to date. Each message received from the server may change this status; hence, it is important to process the pts, date, and seq fields on received updates, and when calling certain methods.

状态必须始终存储在客户端上并保持最新。从服务器收到的每条消息都可能改变这种状态; 因此，重要的是处理收到的更新上的pts，date和seq字段，以及调用某些方法时。

## Processing Incoming Messages
Each update message is of the type Updates. It can be seen from the schema below that this type has several constructors.

每个更新消息都是Updates类型。从下面的模式可以看出，这种类型有几个构造函数。

Messages with updateShort constructors, normally, have lower priority and are broadcast to a large number of users, i.e. one of the chat participants started entering text in a big conversation (updateChatUserTyping).

具有updateShort构造函数的消息通常具有较低的优先级并且被广播给大量用户，即其中一个聊天参与者开始在大对话中输入文本（updateChatUserTyping）。

The updates constructor may contain the sequence number of the current user’s update (seq > 0). Once received, it is useful to compare it with the value on the client. If the difference is greater than one, then the client is out of sync (normally, as a result of the client disconnecting or the server rebooting).

该更新的构造可能包含当前用户的更新序列号（seq > 0）。收到后，将其与客户端上的值进行比较是有用的。如果差异大于1，则客户端不同步（通常，由于客户端断开连接或服务器重新启动）。

The updateShortMessage and updateShortChatMessage constructors are redundant but help significantly reduce the transmitted message size for 90% of the updates.

该updateShortMessage和updateShortChatMessage构造函数是冗余的，但有助于显着减少90％的更新传输的消息大小。


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

服务器还可以重新组合更新。当客户端不维护活动连接并且未接收更新时，可以在最终传递之前重新分组事件，以便将它们组合成单个消息。在这样做时，将删除过时的更新，例如updateUserTyping，而updateUserStatus等更新将重新分组，以便仅保留每个user_id的最新更新。可能是一条消息将包含来自具有不同顺序值的更新的事件（** seq **）。在这种情况下，将使用updatesCombined构造函数。它定义了seq的范围。收到这样的构造函数时，应检查seq_start正好是一个比客户端对值序列，和当前值序列应与更新的价值被覆盖序列字段。

If the final message becomes too large (the current limit is 200 KB), the updatesTooLong constructor will be used. If a client receives this constructor, it means the difference must be obtained manually.

如果最终消息变得太大（当前限制为200 KB），则将使用updatesTooLong构造函数。如果客户端收到此构造函数，则表示必须手动获取差异。

### Compressing Updates
All updates can also be compressed using gzip in the same way as responses to queries.

所有更新也可以使用gzip压缩，其方式与对查询的响应相同。

## Obtaining Differences
Manually obtaining updates is required in the following situations:

在以下情况下需要手动获取更新：

- Loss of sync: an update contains seq / seq_start (> 0) and it is not equal to client_seq + 1. It may be useful to wait up to 0.5 seconds in this situation and abort the sync in case a new update arrives, that fixes the ‘hole’.

	同步丢失：更新包含seq / seq_start（> 0），它不等于client_seq + 1。在这种情况下等待最多0.5秒并在新的更新到达时中止同步可能是有用的，这可以修复“漏洞”。
	
- Session loss on the server: the client receives a new session created notification. This can be caused by garbage collection on the MTProto server or a server reboot.

	服务器上的会话丢失：客户端收到新会话创建的通知。这可能是由MTProto服务器上的垃圾收集或服务器重启引起的。
- Incorrect update: the client cannot deserialize the received data.

	更新不正确：客户端无法反序列化接收的数据。

- Incomplete update: the client is missing data about a chat/user from one of the shortened constructors, such as updateShortChatMessage, etc.

	不完整的更新：客户端缺少来自其中一个缩短的构造函数的聊天/用户数据，例如updateShortChatMessage等。

- Long period without updates: no updates for 15 minutes or longer.

	没有更新的长时间：15分钟或更长时间内没有更新。

To manually obtain the difference between the client’s status and the updated status, a call to updates.getDifference is made, passing in the parameters of the client’s current state. If the updates.differenceSlice constructor is returned in the response, the full difference was too large to be received by the client all at once. The intermediate status, intermediate_state, must be saved on the client and the query must be repeated, using the intermediate status as the current status.

要手动获取客户端状态和更新状态之间的差异，需要调用updates.getDifference，传入客户端当前状态的参数。如果在响应中返回了updates.differenceSlice构造函数，则完全差异太大而无法立即被客户端接收。中间状态intermediate_state必须保存在客户端上，并且必须使用中间状态作为当前状态重复查询。


## PUSH Notifications about Updates
If a client does not have an active connection at the time of an event, PUSH Notifications will also be useful.

如果客户端在事件发生时没有活动连接，则PUSH通知也很有用。

