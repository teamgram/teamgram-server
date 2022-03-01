# Working with Different Data Centers
> Forwarded from [Working with Different Data Centers](https://core.telegram.org/api/datacenter)


The servers are divided into several data centers (hereinafter “DCs”) in different parts of the world.
A complete list of proxy access points for these DCs may be obtained using help.getConfig:

```
dcOption#2ec2a43c id:int hostname:string ip_address:string port:int = DcOption;
config#232d5905 date:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int = Config;

---functions---
help.getConfig#c4f9186b = Config;
```


In this context, this_dc is the number of the current DC, dc_options is a list of all DCs available at the moment, each of which has an id, ip, and port for establishing a connection. Please note that ip and port may change frequently, based on proxy server load and the user's current location.

To optimize client communication with the API, each client must use the connection to the closest access point for its main queries (sending messages, getting contacts, etc.). Therefore, knowing how to select a DC is required before communicating with the API.

## Registration/Authorization
The auth.sendCode method is the basic entry point when registering a new user or authorizing an existing user. 95% of all redirection cases to a different DC will occure when invoking this method.

The client does not yet know which DC it will be associated with; therefore, it establishes an encrypted connection to a random address and sends its query to that address.
Having received a phone_number from a client, we can find out whether or not it is registered in the system. If it is, then, if necessary, instead of sending a text message, we request that it establish a connection with a different DC first (PHONE_MIGRATE_X error).
If we do not yet have a user with this number, we examine its IP-address. We can use it to identify the closest DC. Again, if necessary, we redirect the user to a different DC (NETWORK_MIGRATE_X error).

### Testing Redirects
There are reserved phone number prefixes to test the correctness of the application’s handling of redirects between DCs. If you wish to emulate an application of a user associated with DC number X, it is sufficient to specify the phone number as 99966XYYYY, where YYYY are random numbers, when registering the user. A user like this would always get XXXXX as the confirmation code (the DC number, repeated five times).
Do not store any important or private information in such test users’ messages; anyone can make use of the simplified authorization mechanism.

## File Access
A file saved by a user with upload.saveFilePart will be available for direct download only from the DC where the query was executed. That is why each file has a dc_id parameter:

fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
To download the file, an encrypted connection to DC dc_id must be established and used to execute the upload.getFile query.
If an attempt is made to download the file over a wrong connection, the FILE_MIGRATE_X error will be returned.

Please note that encryption keys are not copied between DCs; therefore, the process of establishing an encrypted connection is started from the very beginning for each new DC. An issued auth_key can be associated with the current authorized user by using an authorization transfer.

## User Migration
During the process of working with the API, user information is accumulated in the DC with which the user is associated. This is the reason a user cannot be associated with a different DC by means of the client. However, in the future, during prolonged communication from an unusual location, we may decide that the user’s data must be moved to a different DC. After some time, the data will be copied and the association will be updated. Once this happens, when executing any query transmitted to the old DC, the API will return the USER_MIGRATE_X error. The client will then have to establish a connection with the new DC and repeat the query.

## Authorization Transfer
The following methods can be used to eliminate the need for users to enter the code from a text message every time:

```
auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
auth.authorization#f6b673a4 expires:int user:User = auth.Authorization;
---functions---
auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
```

auth.exportAuthorization must be executed in the current DC (the DC with which a connection has already been established), passing in dc_id as the value for the new DC. The method should return the user identifier and a long string of random data. An import operation can be performed at the new DC by sending it what was received. Queries requiring authorization can then be successfully executed in the new DC.
