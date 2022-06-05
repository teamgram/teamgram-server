# 高级 FAQ
> 关于MTProto的FAQ适用于高级用户。您可能还想查看我们的基础 FAQ。
> 请注意，客户端开发人员必须遵守安全指南。
> 

### General

Why did you use a custom protocol?

How does it work?

Server-client encryption

End-to-end encryption

Why didn't you use a different solution?

Why are you mostly relying on classical crypto algorithms?

I'm a security expert and I have comments about your setup
### Encryption

How are MTProto messages authenticated?

Are you using Encrypt-and-MAC?

Why not go for Encrypt-then-MAC?

Do you still use SHA-1?

Do you use IGE? IGE is broken!
### Authentication

How is the server authenticated during DH key exchange?

How are clients authenticated?

How are secret chats authenticated?

How are voice calls authentication?

Do you have Forward Secrecy?

### Protection against known attacks

Known-plaintext attacks

Chosen-plaintext attacks

Chosen-ciphertext attacks

What about IND CCA?

Replay attacks

Man-in-the-middle attacks

Hash collisions for DH keys

Length extension attacks

### Encrypted CDNs

Why do you use CDNs?

Can CDNs decipher any files?

Can CDNs substitute files with their own versions?

Can CDNs be used for censorship?

Can I verify this?

Does this affect private data?

Is this connected to government requests to move servers to their territory?
Does this give countries any influence over Telegram?

## General questions
### Q: Why did you go for a custom protocol?
In order to achieve reliability on weak mobile connections as well as speed when dealing with large files (such as photos, large videos and files up to 2 GB each), MTProto uses an original approach. This document is intended to clarify certain details of our setup, as well as address some important points that might be overlooked at first glance.

### Q: Where can I read more about the protocol?
Detailed protocol documentation is available here. Please note that MTProto supports two layers: client-server encryption that is used in Telegram cloud chats and end-to-end encryption that is used in Telegram Secret Chats. See below for more information.

If you have any comments, feel free to reach out to security@telegram.org

### Q: How does server-client encryption work in MTProto?
Server-client encryption is used in Telegram Cloud Chats. Here's a brief overview of the setup:

MTProto 2.0, Part I. Cloud chats (server-client encryption)
#### Note 1
Each plaintext message to be encrypted in MTProto always contains the following data to be checked upon decryption in order to make the system robust against known problems with the components:

- server salt (64-Bit)
- session id
- message sequence number
- message length
- time

#### Note 2
See additional comments on our use of IGE and message authentication.

#### Note 3
Telegram's End-to-end encrypted Secret Chats are using an additional layer of encryption on top of the described above.

### Q: How does end-to-end encryption work in MTProto?
End-to-end encryption is used in Telegram Secret Chats, as well as voice and video calls. You can read more about it here: Secret Chats, End-to-End encryption. Here's a brief overview of the setup:

End-to-end encryption in MTProto 2.0 (Secret Chats)
Please see these articles for details:

- Secret Chats, End-to-End encryption
- End-to-End TL Schema
- Sequence numbers in secret chats
- Perfect Forward Secrecy
- End-to-end encrypted Voice Calls

### Q: Why are you not using X? (insert solution)
While other ways of achieving the same cryptographic goals, undoubtedly, exist, we feel that the present solution is both robust and also sucсeeds at our secondary task of beating unencrypted messengers in terms of delivery time and stability.

### Q: Why are you mostly relying on classical crypto algorithms?
We prefer to use well-known algorithms, created in the days when bandwidth and processing power were both a much rarer commodity. This has valuable side-effects for modern-day mobile development and sending large files, provided one takes care of the known drawbacks.

The weakspots of such algorithms are also well-known, and have been exploited for decades. We use these algorithms in such a combination that, to the best of our knowledge, prevents any known attacks.

### Q: I'm a security expert and I have comments about your setup.
Any comments on Telegram's security are welcome at security@telegram.org. All submissions which result in a change of code or configuration are eligible for bounties, ranging from $100 to $100,000 or more, depending on the severity of the issue.

Please note that we can not offer bounties for issues that are disclosed to the public before they are addressed.

## Encryption
### Q: How are MTProto messages authenticated?
All Telegram apps ensure that msg_key is equal to SHA-256 of a fragment of the auth_key concatenated with the decrypted message (including 12…1024 bytes of random padding). It is important that the plaintext always contains message length, server salt, session_id and other data not known to the attacker.

It is crucial that AES decryption keys depend both on msg_key, and on auth_key, known only to the parties involved in the exchange.

### Q: Are you doing Encrypt-then-MAC, MAC-then-Encrypt or MAC-and-Encrypt?
We do none of the above, strictly speaking. For message authentication, we compute SHA-256(auth_key_fragment + AES_decrypt(…,encrypted_message)) upon message receipt and compare this value to the msg_key received with the encrypted message.

See also: Why not Encrypt-then-MAC?

### Q: Why don't you go for a standard encrypt-then-MAC approach?
Using encrypt-then-MAC, e.g. involving GCM (Galois Counter Mode), would enable the receiving party to detect unauthorized or modified ciphertexts, thus eliminating the need to decrypt them in case of tampering.

In MTProto, the clients and the server authenticate messages by ensuring that SHA-256(auth_key_fragment + plaintext + padding) = msg_key and that the plaintext always contains message length, server salt, session_id and other data not known to a potential attacker before accepting any message. These security checks performed on the client before any message is accepted ensure that invalid or tampered with messages will always be safely (and silently) discarded.

This way we arrive at the same result. The difference is that the security check is performed before decryption in Encrypt-then-MAC and after decryption in MTProto – but in either case before a message is accepted. AES encryption / decryption on devices currently in use is comparable in speed with the additional HMAC computation required for the encrypt-then-MAC approach.

### Q: Do you still use SHA-1?
The current version of the protocol is using SHA-256. MTProto 1.0 used to rely on SHA-1 (see this FAQ for details).

In MTProto 2.0, SHA-1 is used only where the choice of hash function is irrelevant for security, e.g.:

- When generating new keys
- For computing 64-bit auth_key_id from auth_key
- For computing the 64-bit key_fingerprint in secret chat used for sanity checks (these are not the key visualizations – they use a different algorithm, see Hash Collisions for Diffie-Hellman keys)

### Q: Do you use IGE? IGE is broken!
Yes, we use IGE, but it is not broken in our implementation. The fact that we do not use IGE as MAC together with other properties of our system makes the known attacks on IGE irrelevant.

IGE, just as the ubiquitous CBC, is vulnerable to blockwise-adaptive CPA. But adaptive attacks are only a threat for as long as the same key can be used in several messages (not so in MTProto).

Adaptive attacks are even theoretically impossible in MTProto, because in order to be encrypted the message must be fully formed first, since the key is dependent on the message content. As for non-adaptive CPA, IGE is secure against them, as is CBC.

## Authentication
### Q: How is the server authenticated during DH key exchange?
The DH exchange is authenticated with the server's public RSA-key that is built into the client (the same RSA-key is also used for protection against MitM attacks).

### Q: How are clients authenticated?
Various secrets (nonce, server_nonce, new_nonce) exchanged during key generation guarantee that the DH-key can only be obtained by the instance that initiated the exchange.

Notice that new_nonce is transferred explicitly only once, inside an RSA-encrypted message from the client to the server.

### Q: How are Secret Chats authenticated?
Keys for end-to-end encrypted secret chats are generated by a new instance of DH key exchange, so they are known only to the parties involved and not to the server. To establish the identities of these parties and to ensure that no MitM is in place, it is recommended to compare identicons, generated from hashes of the DH secret chat keys (key visualizations).

### Q: How are Voice Calls authenticated?
Keys for end-to-end encrypted calls are generated using the Diffie-Hellman key exchange. Users who are on a call can ensure that there is no MitM by comparing key visualizations.

To make key verification practical in the context of a voice call, Telegram uses a three-message modification of the standard DH key exchange for calls:

- A->B : (generates a and) sends g_a_hash := hash(g^a)
- B->A : (stores g_a_hash, generates b and) sends g_b := g^b
- A->B : (computes key (g_b)a, then) sends g_a := ga
- B : checks hash(g_a) == g_a_hash, then computes key (g_a)^b

The idea is that Alice commits to a specific value of a (and of g_a), but does not reveal g_a to Bob (or Eve) until the very last step. Bob has to choose his value of b and g_b without knowing the true value of g_a. If Eve is performing a Man-in-the-Middle attack, she cannot change a depending on the value of g_b received from Bob and she also can't tune her value of b depending on g_a. As a result, Eve only gets one shot at injecting her parameters — and she must fire this shot with her eyes closed.

Thanks to this modification, it becomes possible to prevent eavesdropping (MitM attacks on DH) with a probability of more than 0.9999999999 by using just over 33 bits of entropy in the visualization. These bits are presented to the users in the form of four emoticons. We have selected a pool of 333 emoji that all look quite different from one another and can be easily described in simple words in any language.

You can read more about key verification for Telegram calls here.

### Q: Do you have Forward Secrecy?
Telegram's Secret chats support Perfect Forward Secrecy, you can read more about it here.

## Protection against known attacks
### Known-Plaintext Attacks
By definition, the known-plaintext attack (KPA) is an attack model for cryptanalysis where the attacker has samples of both the plaintext, and its encrypted version (ciphertext).

AES IGE that is used in MTProto is robust against KPA attacks (see this, if you wonder how one can securely use IGE). On top of that, the plaintext in MTProto always contains server_salt and session id.

### Chosen-Plaintext Attacks
By definition, a chosen-plaintext attack (CPA) is an attack model for cryptanalysis which presumes that the attacker has the capability to choose arbitrary plaintexts to be encrypted and obtain the corresponding ciphertexts.

MTProto uses AES in IGE mode (see this, if you wonder how one can securely use IGE) that is secure against non-adaptive CPAs. IGE is known to be not secure against blockwise-adaptive CPA, but MTProto fixes this in the following manner:

Each plaintext message to be encrypted always contains the following to be checked upon decryption:

- server salt (64-Bit)
- message sequence number
- time

On top of this, in order to replace the plaintext, you would also need to use the right AES key and iv, both dependent on the auth_key. This makes MTProto robust against a CPA.

### Chosen-Ciphertext Attacks
By definition, a chosen-ciphertext attack (CCA) is an attack model for cryptanalysis in which the cryptanalyst gathers information, at least in part, by choosing a ciphertext and obtaining its decryption under an unknown key. In the attack, an adversary has a chance to enter one or more known ciphertexts into the system and obtain the resulting plaintexts. From these pieces of information the adversary can attempt to recover the hidden secret key used for decryption.

Each time a message is decrypted in MTProto, a check is performed to see whether msg_key is equal to the SHA-256 of a fragment of the auth_key concatenated with the decrypted message (including 12…1024 bytes of random padding). The plaintext (decrypted data) also always contains message length, server salt and sequence number. This negates known CCAs.

### What about IND-CCA?
MTProto 2.0 satisfies the conditions for indistinguishability under chosen ciphertext attack (IND-CCA).

> Read more about IND-CCA in MTProto 1.0

### Replay attacks
Replay attacks are denied because each plaintext to be encrypted contains the server salt and the unique message id and sequence number.

This means that each message can only be sent once.

### Man-in-the-middle attacks
Telegram has two modes of communication — ordinary chats using client-server encryption and Secret Chats using end-to-end encryption.

Client-Server communication is protected from MiTM-attacks during DH key generation by means of a server RSA public key embedded into client software. After that, if both clients trust the server software, the Secret Chats between them are protected by the server from MiTM attacks.

The interface offers a way of comparing Secret Chat keys for users who do not trust the server. Visualizations of the key are presented in the form of identicons (example here). By comparing key visualizations users can make sure no MITM attack had taken place.

### Hash collisions for Diffie-Hellman Keys
Currently, the fingerprint uses 128-bits of SHA-1 concatenated with 160 bits from the SHA-256 of the key, yielding a total of 288 fingerprint bits, thus negating the possibility of hash-collision attacks.

Read more about fingerprints in earlier versions of Telegram

### Length extension attacks
By definition, length extension attacks are a type of attack when certain types of hashes are misused as message authentication codes, allowing for inclusion of extra information.

A message in MTProto consists of an msg_key, equal to the SHA-256 of a fragment of the auth_key concatenated with the plaintext (including 12…1024 bytes of random padding and some additional parameters), followed by the ciphertext. The attacker cannot append extra bytes to the end and recompute the SHA-256, since the SHA-256 is computed from the plaintext, not the ciphertext, and the attacker has no way to obtain the ciphertext corresponding to the extra plaintext bytes she may want to add.

Apart from that, changing the msg_key would also change the AES decryption key for the message in a way unpredictable for the attacker, so even the original prefix would decrypt to garbage — which would be immediately detected since the app performs a security check to ensure that the SHA-256 of the plaintext (combined with a fragment of the auth_key) matches the msg_key received.

## Encrypted CDNs
As of Telegram 4.2, we support encrypted CDNs for caching media from public channels with over 100.000 members. The CDN caching nodes are located in regions with significant Telegram traffic where we wouldn't want to place Telegram servers for various reasons.

> For technical details of the implementation, encryption and verification of data, see the CDN manual.

See this document for a Persian version of this FAQ.
بخش فارسی

### Q: Why did you decide to use CDNs?
We use our own distributed servers to speed up downloads in regions where freedom of speech is guaranteed — and even there we don't take this for granted. But when Telegram becomes immensely popular in other areas, we can only rely on CDNs which we treat rather like ISPs from the technical standpoint in that they only get encrypted data they can't decipher.

Thanks to this technology, the download speed for public photos and videos can become significantly higher in regions like Turkey, Indonesia, South America, India, Iran or Iraq without the slightest compromise in security.

### Q: Can the CDN decipher the files?
No. Each file that is to be sent to the CDN is encrypted with a unique key using AES-256-CTR encryption. The CDN can't access the data it stores because these keys are only accessible to the main MTProto server and to the authorized client.

### Q: Can the CDN substitute the data with their own version?
No. Data downloaded from CDN caching nodes is always verified by the receiving Telegram app by way of a hash: attackers won’t be able to replace any files with their own versions.

### Q: Can the CDN delete any files?
No. CDN nodes only cache encrypted copies of files, originals are stored on the Telegram servers. The user is notified about receiving the file by the Telegram server. If the CDN caching node doesn't give the file to the user, the user will receive the file from the Telegram server directly.

### Q: Can CDNs be used for censorship?
No. All original files are stored on the Telegram servers. The CDNs only get encrypted data — and they can't decipher it. They can't substitute any data. And in case of any problems with the CDN, the file will be simply delivered to the users directly from the Telegram servers. Users will always get their data, nobody can stop this.

### Q: Can I verify this?
Yes. Anyone can verify our CDN implementation by checking the source code of Telegram apps and inspecting traffic.

### Q: Does this affect private data?
No. The CDN caching nodes are not a part of the Telegram cloud. CDN caching nodes are used only for caching popular public media from massive channels. Private data never goes there.

### Q: Is this connected with government requests to move private data to their territory?
No. We haven't entered in any agreements with any government regarding the CDNs and the CDNs are not part of any deal. The only purpose of CDNs is to securely improve connectivity in high demand regions where Telegram can't place its servers.

### Q: Does this give some countries any influence over Telegram?
No. We have taken special precautions to make sure that no country gains any leverage over Telegram by way of the CDN caching nodes:

- The CDNs do not belong to Telegram – all the risks are on a third-party company that supplies us with CDN nodes around the world.
- We did not invest anything in these CDNs and will only be paying for traffic that is used to pass cached items from our main clusters and to the end users.

As the result, if any country decides to mess with the CDN in their region, they gain nothing except for reducing connectivity for their own citizens – and Telegram loses nothing of value.