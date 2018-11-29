libphonenumber
==============

golang port of Google's libphonenumber

[![forthebadge](http://forthebadge.com/images/badges/no-ragrets.svg)](http://forthebadge.com)

[![Build Status](https://travis-ci.org/ttacon/libphonenumber.svg?branch=master)](https://travis-ci.org/ttacon/libphonenumber)
[![GoDoc](https://godoc.org/github.com/ttacon/libphonenumber?status.png)](https://godoc.org/github.com/ttacon/libphonenumber)

Status
======

This library is fully stable and is used in production by several companies.

Examples
========

Super simple to use.

### To get a phone number

```go
num, err := libphonenumber.Parse("6502530000", "US")
```

### To format a number

```go
// num is a *libphonenumber.PhoneNumber
formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
```
