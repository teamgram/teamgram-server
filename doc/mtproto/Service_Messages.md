# Service Messages
> Forwarded from [Service Messages](https://core.telegram.org/mtproto/service_messages)


## Response to an RPC query
A response to an RPC query is normally wrapped as follows:

```
rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult;
```

Here req_msg_id is the identifier of the message sent by the other party and containing an RPC query. This way, the recipient knows that the result is a response to the specific RPC query in question.
At the same time, this response serves as acknowledgment of the other party’s receipt of the req_msg_id message.

Note that the response to an RPC query must also be acknowledged. Most frequently, this coincides with the transmission of the next message (which may have a container attached to it carrying a service message with the acknowledgment).

## RPC Error
The result field returned in response to any RPC query may also contain an error message in the following format:

```
rpc_error#2144ca19 error_code:int error_message:string = RpcError;
```

## Cancellation of an RPC Query
In certain situations, the client does not want to receive a response to an already transmitted RPC query, for example because the response turns out to be long and the client has decided to do without it because of insufficient link capacity. Simply interrupting the TCP connection will not have any effect because the server would re-send the missing response at the first opportunity. Therefore, the client needs a way to cancel receipt of the RPC response message, actually acknowledging its receipt prior to it being in fact received, which will settle the server down and prevent it from re-sending the response. However, the client does not know the RPC response’s msg_id prior to receiving the response; the only thing it knows is the req_msg_id. i. e. the msg_id of the relevant RPC query. Therefore, a special query is used:

```
rpc_drop_answer#58e4a740 req_msg_id:long = RpcDropAnswer;
```

The response to this query returns as one of the following messages wrapped in rpc_result and requiring an acknowledgment:

```
rpc_answer_unknown#5e2ad36e = RpcDropAnswer;
rpc_answer_dropped_running#cd78e586 = RpcDropAnswer;
rpc_answer_dropped#a43ad8b7 msg_id:long seq_no:int bytes:int = RpcDropAnswer;
```

The first version of the response is used if the server remembers nothing of the incoming req_msg_id (if it has already been responded to, for example). The second version is used if the response was canceled while the RPC query was being processed (where the RPC query itself was still fully processed); in this case, the same rpc_answer_dropped_running is also returned in response to the original query, and both of these responses require an acknowledgment from the client. The final version means that the RPC response was removed from the server’s outgoing queue, and its msg_id, seq_no, and length in bytes are transmitted to the client.

Note that rpc_answer_dropped_running and rpc_answer_dropped serve as acknowledgments of the server’s receipt of the original query (the same one, the response to which we wish to forget). In addition, same as for any RPC queries, any response to rpc_drop_answer is an acknowledgment for rpc_drop_answer itself.

As an alternative to using rpc_drop_answer, a new session may be created after the connection is reset and the old session is removed through destroy_session.

### Messages associated with querying, changing, and receiving the status of other messages
See Mobile Protocol: Service Messages about Messages

## Request for several future salts
The client may at any time request from the server several (between 1 and 64) future server salts together with their validity periods. Having stored them in persistent memory, the client may use them to send messages in the future even if he changes sessions (a server salt is attached to the authorization key rather than being session-specific).

```
get_future_salts#b921bd04 num:int = FutureSalts;
future_salt#0949d9dc valid_since:int valid_until:int salt:long = FutureSalt;
future_salts#ae500895 req_msg_id:long now:int salts:vector future_salt = FutureSalts;
```

The client must check to see that the response’s req_msg_id in fact coincides with msg_id of the query for get_future_salts. The server returns a maximum of num future server salts (may return fewer). The response serves as the acknowledgment of the query and does not require an acknowledgment itself.

## Ping Messages (PING/PONG)
```
ping#7abe77ec ping_id:long = Pong;
```
A response is usually returned to the same connection:

```
pong#347773c5 msg_id:long ping_id:long = Pong;
```

These messages do not require acknowledgments. A pong is transmitted only in response to a ping while a ping can be initiated by either side.

## Deferred Connection Closure + PING
```
ping_delay_disconnect#f3427b8c ping_id:long disconnect_delay:int = Pong;
```

Works like ping. In addition, after this is received, the server starts a timer which will close the current connection disconnect_delay seconds later unless it receives a new message of the same type which automatically resets all previous timers. If the client sends these pings once every 60 seconds, for example, it may set disconnect_delay equal to 75 seconds.

## Request to Destroy Session
Used by the client to notify the server that it may forget the data from a different session belonging to the same user (i. e. with the same auth_key_id). The result of this being applied to the current session is undefined.

```
destroy_session#e7512126 session_id:long = DestroySessionRes;
destroy_session_ok#e22045fc session_id:long = DestroySessionRes;
destroy_session_none#62d350c9 session_id:long = DestroySessionRes;
```

## New Session Creation Notification
The server notifies the client that a new session (from the server’s standpoint) had to be created to handle a client message. If, after this, the server receives a message with an even smaller msg_id within the same session, a similar notification will be generated for this msg_id as well. No such notifications are generated for high msg_id values.

```
new_session_created#9ec20908 first_msg_id:long unique_id:long server_salt:long = NewSession
```

The unique_id parameter is generated by the server every time a session is (re-)created.

This notification must be acknowledged by the client. It is necessary, for instance, for the client to understand that there is, in fact, a “gap” in the stream of long poll notifications received from the server (the user may have failed to receive notifications during some period of time).

Notice that the server may unilaterally destroy (close) any existing client sessions with all pending messages and notifications, without sending any notifications. This happens, for example, if the session is inactive for a long time, and the server runs out of memory. If the client at some point decides to send new messages to the server using the old session, already forgotten by the server, such a “new session created” notification will be generated. The client is expected to handle such situations gracefully.

## Containers
Containers are messages containing several other messages. Used for the ability to transmit several RPC queries and/or service messages at the same time, using HTTP or even TCP or UDP protocol. A container may only be accepted or rejected by the other party as a whole.

### Simple Container
A simple container carries several messages as follows:

```
msg_container#73f1f8dc messages:vector message = MessageContainer;
```

Here message refers to any message together with its length and msg_id:

```
message msg_id:long seqno:int bytes:int body:Object = Message;
```
bytes is the number of bytes in the body serialization.
All messages in a container must have msg_id lower than that of the container itself. A container does not require an acknowledgment and may not carry other simple containers. When messages are re-sent, they may be combined into a container in a different manner or sent individually.

Empty containers are also allowed. They are used by the server, for example, to respond to an HTTP request when the timeout specified in http_wait expires, and there are no messages to transmit.

## Message Copies
In some situations, an old message with a msg_id that is no longer valid needs to be re-sent. Then, it is wrapped in a copy container:

```
msg_copy#e06046b2 orig_message:Message = MessageCopy;
```

Once received, the message is processed as if the wrapper were not there. However, if it is known for certain that the message orig_message.msg_id was received, then the new message is not processed (while at the same time, it and orig_message.msg_id are acknowledged). The value of orig_message.msg_id must be lower than the container’s msg_id.

This is not used at this time, because an old message can be wrapped in a simple container with the same result.

## Packed Object
Used to replace any other object (or rather, a serialization thereof) with its archived (gzipped) representation:

```
gzip_packed#3072cfa1 packed_data:string = Object;
```

At the present time, it is supported in the body of an RPC response (i.e., as result in rpc_result) and generated by the server for a limited number of high-level queries. In addition, in the future it may be used to transmit non-service messages (i. e. RPC queries) from client to server.

## HTTP Wait/Long Poll
The following special service query not requiring an acknowledgement (which must be transmitted only through an HTTP connection) is used to enable the server to send messages in the future to the client using HTTP protocol:

```
http_wait#9299359f max_delay:int wait_after:int max_wait:int = HttpWait;
```

When such a message (or a container carrying such a message) is received, the server either waits max_delay milliseconds, whereupon it forwards all the messages that it is holding on to the client if there is at least one message queued in session (if needed, by placing them into a container to which acknowledgments may also be added); or else waits no more than max_wait milliseconds until such a message is available. If a message never appears, an empty container is transmitted.

The max_delay parameter denotes the maximum number of milliseconds that has elapsed between the first message for this session and the transmission of an HTTP response. The wait_after parameter works as follows: after the receipt of the latest message for a particular session, the server waits another wait_after milliseconds in case there are more messages. If there are no additional messages, the result is transmitted (a container with all the messages). If more messages appear, the wait_after timer is reset.

At the same time, the max_delay parameter has higher priority than wait_after, and max_wait has higher priority than max_delay.

This message does not require a response or an acknowledgement. If the container transmitted over HTTP carries several such messages, the behavior is undefined (in fact, the latest parameter will be used).

If no http_wait is present in container, default values max_delay=0 (milliseconds), wait_after=0 (milliseconds), and max_wait=25000 (milliseconds) are used.

If the client’s ping of the server takes a long time, it may make sense to set max_delay to a value that is comparable in magnitude to ping time.
