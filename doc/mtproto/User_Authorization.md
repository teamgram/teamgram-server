# User Authorization
> Forwarded from [User Authorization](https://core.telegram.org/api/auth)

Authorization is associated with a client’s encryption key identifier: auth_key_id. No additional parameters need to be passed into methods following authorization.

Authorization requires that a text message containing an authorization code first be sent to the user’s phone. This may be done using the auth.sendCode method. If the message takes too long (30 seconds) to arrive at the phone, the auth.sendCall method may be invoked and passed a code identifier and a phone number, in which case the user will receive a phone call, and a synthesized voice will give him or her the code previously sent in the text message.

The auth.sendCode method will also return the phone_registered field, which indicates whether or not the user with this phone number is registered in the system. If phone_registered == boolTrue, authorization requires that auth.signIn be invoked. Otherwise, basic information must be requested from the user and the new user registration method (auth.signUp) must be invoked.

When phone_registered === boolFalse, the auth.signIn method can be used to pre-validate the code entered from the text message. The code was entered correctly if the method returns Error 400 PHONE_NUMBER_UNOCCUPIED.

As a result of authorization, the client key, auth_key_id, becomes associated with the user, and each subsequent API call with this key will be executed with that user’s identity. The authorization method itself returns the relevant user and the association expiration time, expires. It is best to immediately store the User ID locally in a binding with the key. At some point in time (>= expires), the association will expire, and the authorization process will have to be repeated.

Only a small portion of the API methods are available to unauthorized users:

- auth.sendCode
- auth.sendCall
- auth.checkPhone
- auth.signUp
- auth.signIn
- auth.importAuthorization
- help.getConfig
- help.getNearestDc

Other methods will result in an error: 401 UNAUTHORIZED.
