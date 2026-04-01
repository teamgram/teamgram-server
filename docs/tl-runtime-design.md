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

### Single-Constructor Types

If a TL type has only one constructor in the entire schema, it should not be modeled as an abstract wrapper.

In that case, the runtime and business layer should use the concrete constructor type directly.

Example:

- if `FooResult` has exactly one constructor in the whole schema
- then a method returning `FooResult` should directly return `*TLFooResult`

Recommended rule:

- single-constructor type: generate only the concrete type
- multi-constructor type: generate wrapper + abstract interface

Do not generate these for single-constructor types:

- wrapper type
- `XXXClazz` interface
- `ToXxxType()` promotion helper
- `Match(...)`

This keeps the runtime model smaller and avoids paying abstraction cost where no polymorphism exists.

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
- direct concrete types
- ordinary Go `type switch` for large unions

instead of raw predicate string comparisons.

## Accessing Abstract Types

Abstract wrappers should remain thin. Their primary purpose is:

- method result typing
- constructor family grouping
- layer-aware encoding through the concrete constructor

The primary access pattern for business code should be `ToXxx()` helpers.

Example:

```go
if v, ok := resp.ToAuthSentCodeSuccess(); ok {
    // handle success variant
    return
}

if v, ok := resp.ToAuthSentCode(); ok {
    // handle normal sentCode variant
    return
}
```

This style is preferred because it is:

- idiomatic Go
- explicit
- type-safe
- easy to read and refactor

More examples:

```go
if v, ok := savedRingtone.ToAccountSavedRingtone(); ok {
    // concrete constructor: account.savedRingtone
    _ = v
    return
}

if v, ok := savedRingtone.ToAccountSavedRingtoneConverted(); ok {
    // concrete constructor: account.savedRingtoneConverted
    _ = v.Document
    return
}
```

```go
if v, ok := sentCode.ToAuthSentCode(); ok {
    _ = v.PhoneCodeHash
    return
}

if v, ok := sentCode.ToAuthSentCodeSuccess(); ok {
    _ = v.Authorization
    return
}

if v, ok := sentCode.ToAuthSentCodePaymentRequired(); ok {
    _ = v.StoreProduct
    return
}
```

### Match and Switch

The runtime should not generate `Match(f ...interface{})`.

Reasons:

- it relies on variadic `interface{}` callbacks
- it is less idiomatic than direct typed access
- it is harder to read and refactor
- it scales poorly for large constructor families

The runtime also does not need to generate a custom `Switch(...)` API by default.

For ordinary business logic, `ToXxx()` and plain Go `type switch` are sufficient and more readable.

### Small vs Large Constructor Families

Abstract types should be treated differently depending on constructor count.

For small constructor families such as `auth.SentCode`:

- keep the abstract wrapper
- generate `ToXxx()` helpers

For large constructor families such as `Update`:

- keep the abstract wrapper
- keep `ToXxx()` helpers
- expose `ClazzName()` / predicate metadata
- prefer ordinary Go `type switch` in business code
- do not rely on a giant generated `Match(...)` as the main API

Example:

```go
switch v := upd.Clazz.(type) {
case *tg.TLUpdateNewMessage:
    // ...
case *tg.TLUpdateDeleteMessages:
    // ...
}
```

This keeps the generated API smaller and makes business logic easier to follow.

## Concrete Field Naming

Concrete constructor structs should keep the original TL field semantics in their Go field names.

Example:

```tl
idVal id:long = IdVal;
idVals id:Vector<long> = IdVal;
seqIdVal id:long = IdVal;
```

Recommended generated concrete fields:

```go
type TLIdVal struct {
    Id int64
}

type TLIdVals struct {
    Id []int64
}

type TLSeqIdVal struct {
    Id int64
}
```

The generator may still need family-level disambiguation metadata when an abstract type contains same-named fields with different wire types. That internal disambiguation should not leak into the concrete constructor field names.

Recommended rule:

- concrete constructor fields keep the original TL semantic name
- abstract family merge metadata may use internal type-based disambiguation

Do not expose family-level suffixes such as `Id_INT64` or `Id_VECTORINT64` on concrete constructor structs unless there is an actual collision inside the same concrete struct.

More examples:

```go
switch v := user.Clazz.(type) {
case *tg.TLUser:
    _ = v.FirstName
case *tg.TLUserEmpty:
    // deleted or unavailable user
}
```

```go
switch v := media.Clazz.(type) {
case *tg.TLMessageMediaPhoto:
    _ = v.Photo
case *tg.TLMessageMediaDocument:
    _ = v.Document
case *tg.TLMessageMediaUnsupported:
    // ignore unsupported media
}
```

## Layer Field Evolution

When the same predicate evolves across layers but keeps the same semantic field, the generated business-facing struct should expose a single stable field whenever possible.

This is especially important for server-side responses, where business code should be able to assign once and let the runtime encode correctly for different client layers.

Example:

```tl
layer 220-221:
urlAuthResultAccepted#8f8c0e4e url:string = UrlAuthResult;

layer 222-223:
urlAuthResultAccepted#623a8fa0 flags:# url:flags.0?string = UrlAuthResult;
```

Recommended business-facing generated shape:

```go
type TLUrlAuthResultAccepted struct {
    ClazzID    uint32  `json:"_id"`
    ClazzName2 string  `json:"_name"`
    Url        *string `json:"url"`
}
```

Recommended encoding behavior:

- for layers where `url` is required, `Url` must be non-nil
- for layers where `url` is optional, `Url == nil` means the flag is not set
- decode should normalize both wire shapes back into the same semantic field

The key rule is:

- business-facing structs should prefer one semantic field over multiple per-layer conflict fields
- layer-specific shape differences should be handled by encode/decode projection

This keeps business code simple:

```go
resp := &tg.TLUrlAuthResultAccepted{
    Url: &url,
}
```

The same value can then be serialized for different client layers without forcing business code to set:

- `Url_STRING`
- `Url_FLAGSTRING`

### When a Single Semantic Field Is Appropriate

Use a single semantic field when the layer change is representational rather than semantic, such as:

- required `string` becoming optional `flags.?string`
- required object becoming optional `flags.?object`
- old and new layers representing the same identifier with different optionality

### When a Single Semantic Field Is Not Enough

If the same field name changes in a way that materially changes its domain, such as:

- `int` to `string`
- `vector<int>` to `vector<object>`
- scalar to union/object

then the generator should not force multiple conflicting concrete fields into one business-facing struct by default.

In those cases, prefer one of:

- a stable semantic model plus layer-aware projection
- an explicitly versioned wire representation when no stable semantic field exists

### Practical Rule

For generated server-facing TL structs:

- first try to expose one semantic field that business code can assign once
- only fall back to suffixed conflict fields when the same concrete struct truly must carry incompatible representations that cannot be normalized safely

This rule matters more than the exact suffix format because the main goal is not naming cleanliness by itself. The main goal is making one assignment work across supported client layers.

## Registry Model

The runtime registry should be split by responsibility instead of forcing all lookups through one map.

### Decode Registry

Decode is constructor-id-driven, so it should have a dedicated registry:

- `clazz_id -> concrete constructor factory`
- `clazz_id -> predicate`
- `clazz_id -> abstract parent type`

This registry does not need layer as an input because the incoming payload already contains the constructor id.

Example:

- `0x5e002502 -> auth.sentCode -> auth.SentCode`
- `0x2390fe44 -> auth.sentCodeSuccess -> auth.SentCode`
- `0xe0955a3c -> auth.sentCodePaymentRequired -> auth.SentCode`

This registry is used to:

- create concrete result/request objects from bytes
- decode wrapped abstract result families
- support logging and diagnostics after decode

### Encode Registry

Encode is layer-driven, so it needs a different registry:

- `predicate + layer -> clazz_id`
- `predicate + layer -> encode version availability`
- optionally `type + layer -> allowed constructor predicates`

This registry is used when the server already has a semantic object and needs to serialize it for a specific client layer.

Example:

- `auth.sendCode + layer -> method clazz_id`
- `auth.sentCode + layer -> constructor clazz_id`
- `auth.sentCodeSuccess + layer -> constructor clazz_id`

### Type Registry

Abstract result types should have an explicit registry view:

- `type name -> constructor predicate set`
- optionally `type name + layer -> constructor predicate set`

Example:

- `auth.SentCode -> {auth.sentCode, auth.sentCodeSuccess, auth.sentCodePaymentRequired}`

This registry is useful for:

- method/result validation
- future compatibility checks
- introspection and debugging

### Method Registry

Methods should also be explicit runtime metadata, not only generated code.

A method registry should minimally expose:

- `method predicate`
- request constructor predicate
- abstract result type name

Example:

- `auth.sendCode -> request: auth.sendCode, result: auth.SentCode`

This lets the runtime reason about what a handler is allowed to return without relying only on handwritten conventions.

### Practical Summary

The runtime should not rely on a single registry.

It should have at least these logical views:

- decode view: `clazz_id -> constructor`
- encode view: `predicate + layer -> clazz_id`
- type view: `type -> constructors`
- method view: `method -> result type`

These views may share the same underlying metadata store, but they should remain distinct in design and API.

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

## Large Layer Spans

When the server supports a large number of schema layers, the runtime must remain lightweight.

In that environment, the runtime should not try to dynamically model every field-level evolution across every supported layer.

Instead, the design should follow this rule:

- generator heavy
- runtime light

### What the Generator Should Do

The generator should absorb most schema complexity at generation time:

- load all supported layer schemas
- compute predicate evolution across layers
- compute constructor id changes across layers
- materialize versioned encode/decode branches
- generate abstract wrappers only for multi-constructor types
- generate method skeletons with correct abstract result types

### What the Runtime Should Do

The runtime should stay focused on a small set of operations:

- decode by constructor id
- encode by predicate and target layer
- wrap and unwrap abstract result types
- validate method results
- expose metadata for logging and diagnostics

The runtime should not become a general-purpose field projection engine.

### Why This Matters

If the runtime tries to dynamically reconcile field-level differences for every layer:

- complexity will grow with every new schema layer
- business logic will be polluted by compatibility rules
- the protocol layer will become harder to reason about
- maintenance cost will become unacceptable

For wide layer support, the correct architecture is:

- schema complexity handled at generation time
- runtime complexity minimized to dispatch and lookup

This is the only scalable model for long-running Telegram-compatible server support.
