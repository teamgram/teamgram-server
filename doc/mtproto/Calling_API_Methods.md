# Calling API Methods
> Forwarded from [Calling API Methods](https://core.telegram.org/api/invoking)

## Layers
Versioning in the API is supported by so-called TL layers.

The need to add a new object constructor or to add/remove a field in a constructor creates a backwards compatibility problem for previous versions of API clients. After all, simply changing a constructor in a schema also changes its number. To address this problem, each schema update is separated into a layer.
A layer is a collection of updated methods or constructors in a TL schema. Each layer is numbered with sequentially increasing numbers starting with 2. The first layer is the base layer — the TL schema without any changes.

There are helper methods to let the API know that a client supports Layer N:
invokeWithLayerN {X:Type} query:!X = X;

If a client supports Layer 2, then the following constructor must be used:

```
invokeWithLayer2#289dd1f6 {X:Type} query:!X = X;
```

In practice, this means that before every API call, an int with the value 0x289dd1f6 must be added before the method number.

Then the API can return constructors that appeared in Layer 2.

Update: Starting with Layer 9, helper methods invokeWithLayerN can be used only together with initConnection: the present layer will be saved with all other parameters of the client and future calls will be using this saved value. See more below.

List of Available Layers

## Saving Client Info
Starting with Layer 9, it is possible to save information about the current client on the server in conjunction with an authorization key. This may help eliminate client-side problems with certain releases on certain devices or with certain localizations, as well as eliminate the need for sending layer information in each call.

The new helper method initConnection accepts client parameters. This method must be called when first calling the API after the application has restarted or in case the value of one of the parameters could have changed.
UPDATE: initConnection must also be called after each auth.bindTempAuthKey.

When calling this method, the current layer used by the client is also saved (the layer in which initConnection was wrapped is used). After a successful call to initConnection it is no longer necessary to wrap each API call in invokeWithLayerN.

## Sequential Calls
Sometimes a client needs to transmit several send message method calls to the server all at once in a single message or in several consecutive messages. However, there is a chance that the server may execute these requests out of order (queries are handled by different servers to improve performance, which introduces a degree of randomness to the process).

There are helper methods for making several consecutive API calls without wasting time waiting for a response:

```
invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;
invokeAfterMsgs#3dc4b4f0 {X:Type} msg_ids:Vector<long> query:!X = X;
```

They may be used, for example, if a client attempts to send accumulated messages after the Internet connection has been restored after being absent for a long time. In this case, the 32-bit number 0xcb9f372d must be added before the method number in each call, followed by a 64-bit message identifier, msg_id, which contains the previous call in the queue.
The second method is similar, except it takes several messages that must be waited for.

If the waiting period exceeds 0.5 seconds (this value may change in the future) and no result has appeared, the method will be executed just the same. If any of the queries returns an error, all its dependent queries will also return the 400 MSG_WAIT_FAILED error.

**Helper Method Sequence**

Important: if the helper methods invokeAfterMsg / invokeAfterMsgs are used together with invokeWithLayerN or other helper methods, invokeAfterMsg / invokeAfterMsgs must always be the outermost wrapper.

## invokeWithoutUpdates
```
invokeWithoutUpdates#bf9459b7 {X:Type} query:!X = X;
```
Invoke with method without returning updates in the socket

## Data Compression
We recommend using gzip compression when making method calls in order to reduce the amount of network traffic.

The schema and constructor information are given in the protocol documentation.

### Data Compression when Making a Call
Before transmitting a query, the string containing the entire body of the serialized high-level query (starting with the method number) must be compressed using gzip. If the resulting string is smaller than the original, it makes sense to transmit the gzip_packed constructor.

There is no point in doing the above when transmitting binary multimedia data (photos, videos) or small messages (up to 255 bytes).

### Uncompressing Data
By default, the server compresses the response to any call as well as updates, in accordance with the rules stated above. If the gzip_packed constructor is received as a response in rpc_result, then the string that follows must be extracted and uncompressed. Processing then continues on the resulting new string.