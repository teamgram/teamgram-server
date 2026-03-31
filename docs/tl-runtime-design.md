# TL Runtime Design

## Goal

Define a stable server-side TL runtime model that supports:

- multi-layer encoding
- abstract result types
- concrete constructor decoding
- business-friendly handler signatures

This document focuses on the runtime core only. It does not define code generator internals.

## Core Principles

1. `Decode` is constructor-id-driven.
2. `Encode` is layer-driven.
3. Methods are concrete request constructors.
4. Method results are abstract TL types.
5. Business logic returns concrete constructors wrapped as abstract result types.

## Decode Model

Incoming TL payloads already carry a constructor id.

Because of that, decode does not need to know the peer layer. Runtime decode only needs:

- `clazz_id -> concrete constructor type`
- `clazz_id -> decode function`

The decode path should always resolve bytes into a concrete constructor object first.

Example:

- `0x5e002502 -> TLAuthSentCode`
- `0x2390fe44 -> TLAuthSentCodeSuccess`
- `0xe0955a3c -> TLAuthSentCodePaymentRequired`

For abstract result families such as `auth.SentCode`, the runtime may then wrap the decoded concrete constructor into an abstract wrapper like `AuthSentCode`.

## Encode Model

Encode is different from decode.

When the server returns a TL object to a client, the runtime must choose the correct constructor id for the client's negotiated layer.

Encode therefore needs:

- `predicate + layer -> clazz_id`
- optionally `type + layer -> allowed constructors`

The business layer should not choose constructor ids directly. It should choose semantic result constructors, and the runtime should encode them using the target layer.

## Methods, Types, and Constructors

The runtime must keep these concepts distinct:

- `method`
- `type`
- `constructor`

Example:

- method: `auth.sendCode`
- result type: `auth.SentCode`
- constructors:
  - `auth.sentCode`
  - `auth.sentCodeSuccess`
  - `auth.sentCodePaymentRequired`

The method declares the abstract result type. The business layer chooses one concrete constructor from that family.

## Recommended Handler Shape

For a method like:

```tl
auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
```

the server-side handler should be modeled as:

```go
func (c *AuthorizationCore) AuthSendCode(in *tg.TLAuthSendCode) (*tg.AuthSentCode, error)
```

This is the preferred shape because:

- request is a concrete method constructor
- response is an abstract result type
- the business layer can return different result constructors without changing the method contract

## Abstract Result Wrappers

Each abstract TL result type should have a wrapper.

Example:

```go
type AuthSentCode struct {
    Clazz AuthSentCodeClazz `json:"_clazz"`
}
```

Each concrete constructor in the family should implement the corresponding abstract interface:

```go
type AuthSentCodeClazz interface {
    iface.TLObject
    AuthSentCodeClazzName() string
}
```

This allows method signatures to stay stable while business logic can still return different constructor variants.

## Constructor-to-Type Promotion

Every concrete constructor that belongs to an abstract type should expose a helper that promotes it into the abstract wrapper.

Example:

```go
func (m *TLAuthSentCode) ToAuthSentCode() *AuthSentCode {
    if m == nil {
        return nil
    }
    return &AuthSentCode{Clazz: m}
}
```

Equivalent helpers should exist for all constructors in the same family:

- `TLAuthSentCode -> ToAuthSentCode()`
- `TLAuthSentCodeSuccess -> ToAuthSentCode()`
- `TLAuthSentCodePaymentRequired -> ToAuthSentCode()`

This keeps business code simple and explicit.

## Business Layer Rules

The business layer should:

- construct concrete result constructors
- convert them into abstract wrappers
- avoid dealing with `clazz_id`
- avoid dealing with peer `layer`

Recommended style:

```go
resp := &tg.TLAuthSentCodeSuccess{
    Authorization: auth,
}
return resp.ToAuthSentCode(), nil
```

or:

```go
resp := &tg.TLAuthSentCode{
    Type:          codeType,
    PhoneCodeHash: hash,
    NextType:      nextType,
    Timeout:       &timeout,
}
return resp.ToAuthSentCode(), nil
```

The business layer should prefer concrete Go types and wrapper helpers over predicate strings.

## Predicate Usage

Predicates remain important, but mainly for runtime metadata, logging, and diagnostics.

They are appropriate for:

- runtime registration
- layer-based constructor lookup
- logging
- metrics
- tracing
- debugging

They should not be the primary mechanism for business branching when concrete Go types already exist.

Business code should prefer:

- `ToXxx()` helpers
- `Match(...)`
- direct concrete types

instead of raw predicate string comparisons.

## Method Skeleton Responsibilities

Generated method skeletons should do the following:

1. decode the concrete request object
2. call the business handler
3. validate the abstract result wrapper
4. encode the concrete wrapped constructor using the client layer

At minimum, method skeletons should validate:

- response is not `nil`
- response wrapper contains a non-`nil` `Clazz`

Example validation:

```go
if resp == nil {
    return nil, fmt.Errorf("auth.sendCode: nil result")
}
if resp.Clazz == nil {
    return nil, fmt.Errorf("auth.sendCode: nil auth.SentCode clazz")
}
```

This validation should happen before encode.

## Runtime Summary

The runtime model can be summarized as:

- request: concrete method constructor
- response contract: abstract result type
- business result: concrete constructor
- decode: constructor-id-driven
- encode: layer-driven

This is the recommended foundation for server-side TL support.
