# Security Guidelines for Client Developers
**See also:**

**Perfect Forward Secrecy**

**Secret chats, end-to-end encryption**

**Perfect Forward Secrecy in Secret Chats**

**MTProto 2.0, Detailed Description**

While MTProto is designed to be a reasonably fast and secure protocol, its advantages can be easily negated by careless implementation. We collected some security guidelines for client software developers on this page. All Telegram clients are required to comply.

> Note that as of version 4.6, major Telegram clients are using MTProto 2.0. MTProto v.1.0 is deprecated and is currently being phased out.

## Diffie-Hellman key exchange
We use DH key exchange in two cases:

- Creating an authorization key
- Establishing Secret Chats with end-to-end encryption

In both cases, there are some verifications to be done whenever DH is used:

### Validation of DH parameters
Client is expected to check whether p = dh_prime is a safe 2048-bit prime (meaning that both p and (p-1)/2 are prime, and that 2^2047 < p < 2^2048), and that g generates a cyclic subgroup of prime order (p-1)/2, i.e. is a quadratic residue mod p. Since g is always equal to 2, 3, 4, 5, 6 or 7, this is easily done using quadratic reciprocity law, yielding a simple condition on p mod 4g -- namely, p mod 8 = 7 for g = 2; p mod 3 = 2 for g = 3; no extra condition for g = 4; p mod 5 = 1 or 4 for g = 5; p mod 24 = 19 or 23 for g = 6; and p mod 7 = 3, 5 or 6 for g = 7. After g and p have been checked by the client, it makes sense to cache the result, so as not to repeat lengthy computations in future.

If the verification takes too long (which is the case for older mobile devices), one might initially run only 15 Miller--Rabin iterations (use parameter 30 in Java) for verifying primeness of p and (p - 1)/2 with error probability not exceeding one billionth, and do more iterations in the background later.

Another way to optimize this is to embed into the client application code a small table with some known "good" couples (g,p) (or just known safe primes p, since the condition on g is easily verified during execution), checked during code generation phase, so as to avoid doing such verification during runtime altogether. The server rarely changes these values, thus one usually needs to put the current value of server's dh_prime into such a table. For example, the current value of dh_prime equals (in big-endian byte order)

```
C7 1C AE B9 C6 B1 C9 04 8E 6C 52 2F 70 F1 3F 73 98 0D 40 23 8E 3E 21 C1 49 34 D0 37 56 3D 93 0F 48 19 8A 0A A7 C1 40 58 22 94 93 D2 25 30 F4 DB FA 33 6F 6E 0A C9 25 13 95 43 AE D4 4C CE 7C 37 20 FD 51 F6 94 58 70 5A C6 8C D4 FE 6B 6B 13 AB DC 97 46 51 29 69 32 84 54 F1 8F AF 8C 59 5F 64 24 77 FE 96 BB 2A 94 1D 5B CD 1D 4A C8 CC 49 88 07 08 FA 9B 37 8E 3C 4F 3A 90 60 BE E6 7C F9 A4 A4 A6 95 81 10 51 90 7E 16 27 53 B5 6B 0F 6B 41 0D BA 74 D8 A8 4B 2A 14 B3 14 4E 0E F1 28 47 54 FD 17 ED 95 0D 59 65 B4 B9 DD 46 58 2D B1 17 8D 16 9C 6B C4 65 B0 D6 FF 9C A3 92 8F EF 5B 9A E4 E4 18 FC 15 E8 3E BE A0 F8 7F A9 FF 5E ED 70 05 0D ED 28 49 F4 7B F9 59 D9 56 85 0C E9 29 85 1F 0D 81 15 F6 35 B1 05 EE 2E 4E 15 D0 4B 24 54 BF 6F 4F AD F0 34 B1 04 03 11 9C D8 E3 B9 2F CC 5B
```

### g_a and g_b validation
Apart from the conditions on the Diffie-Hellman prime dh_prime and generator g, both sides are to check that g, g_a and g_b are greater than 1 and less than dh_prime - 1. We recommend checking that g_a and g_b are between 2^{2048-64} and dh_prime - 2^{2048-64} as well.

### Checking SHA1 hash values during key generation
Once the client receives a server_DH_params_ok answer in step 5) of the Authorization Key generation protocol and decrypts it obtaining answer_with_hash, it MUST check that

```
answer_with_hash := SHA1(answer) + answer + (0-15 random bytes)
```

In other words, the first 20 bytes of answer_with_hash must be equal to SHA1 of the remainder of the decrypted message without the padding random bytes.

### Checking nonce, server_nonce and new_nonce fields
When the client receives and/or decrypts server messages during creation of Authorization Key, and these messages contain some nonce fields already known to the client from messages previously obtained during the same run of the protocol, the client is to check that these fields indeed contain the values previosly known.

### Using secure pseudorandom number generator to create DH secret parameters a and b
Client must use a cryptographically secure PRNG to generate secret exponents a or b for DH key exchange. For secret chats, the client might request some entropy (random bytes) from the server while invoking messages.getDhConfig and feed these random bytes into its PRNG (for example, by PRNG_seed if OpenSSL library is used), but never using these "random" bytes by themselves or replacing by them the local PRNG seed. One should mix bytes received from server into local PRNG seed.

## MTProto Encrypted Messages
Some important checks are to be done while sending and especially receiving encrypted MTProto messages.

### Checking SHA256 hash value of msg_key
msg_key is used not only to compute the AES key and IV to decrypt the received message. After decryption, the client MUST check that msg_key is indeed equal to SHA256 of the plaintext obtained as the result of decryption (including the final 12...1024 padding bytes), prepended with 32 bytes taken from the auth_key, as explained in MTProto 2.0 Description.

If an error is encountered before this check could be performed, the client must perform the msg_key check anyway before returning any result. Note that the response to any error encountered before the msg_key check must be the same as the response to a failed msg_key check.

### Checking message length
The client must check that the length of the message or container obtained from the decrypted message (computed from its length field) does not exceed the total size of the plaintext, and that the difference (i.e. the length of the random padding) lies in the range from 12 to 1024 bytes.

The length should be always divisible by 4 and non-negative. On no account the client is to access data past the end of the decryption buffer containing the plaintext message.

### Checking session_id
The client is to check that the session_id field in the decrypted message indeed equals to that of an active session created by the client.

### Checking msg_id
The client must check that msg_id has even parity for messages from client to server, and odd parity for messages from server to client.

In addition, the identifiers (msg_id) of the last N messages received from the other side must be stored, and if a message comes in with an msg_id lower than all or equal to any of the stored values, that message is to be ignored. Otherwise, the new message msg_id is added to the set, and, if the number of stored msg_id values is greater than N, the oldest (i. e. the lowest) is discarded.

In addition, msg_id values that belong over 30 seconds in the future or over 300 seconds in the past are to be ignored (recall that msg_id approximately equals unixtime * 2^32). This is especially important for the server. The client would also find this useful (to protect from a replay attack), but only if it is certain of its time (for example, if its time has been synchronized with that of the server).

Certain client-to-server service messages containing data sent by the client to the server (for example, msg_id of a recent client query) may, nonetheless, be processed on the client even if the time appears to be “incorrect”. This is especially true of messages to change server_salt and notifications about invalid time on the client. See Mobile Protocol: Service Messages.

## Behavior in case of mismatch
If one of the checks listed above fails, the client is to completely discard the message obtained from server. We also recommend closing and reestablishing the TCP connection to the server, then retrying the operation or the whole key generation protocol.

No information from incorrect messages can be used. Even if the application throws an exception and dies, this is much better than continuing with invalid data.

Notice that invalid messages will infrequently appear during normal work even if no malicious tampering is being done. This is due to network transmission errors. We recommend ignoring the invalid message and closing the TCP connection, then creating a new TCP connection to the server and retrying the original query.

> The previous version of security recommendations relevant for MTProto 1.0 clients is available here.