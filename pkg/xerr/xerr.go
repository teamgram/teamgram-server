// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xerr

/*
import "github.com/zeromicro/go-zero/core/logx"

type XErr struct {
	err error
	msg string
}

func New(err error) error {
	return wrap(err, "")
}

func NewMsg(err error, msg string) error {
	return wrap(err, msg)
}

func NewWithErrorLog(err error, logFormat string, logParams ...any) error {
	logError(err, logFormat, logParams...)
	return wrap(err, "")
}

func NewMsgWithErrorLog(err error, msg string, logFormat string, logParams ...any) error {
	logError(err, logFormat, logParams...)
	return wrap(err, msg)
}

func NewWithInfoLog(err error, logFormat string, logParams ...any) error {
	logInfo(err, logFormat, logParams...)
	return wrap(err, "")
}

func NewMsgWithInfoLog(err error, msg string, logFormat string, logParams ...any) error {
	logInfo(err, logFormat, logParams...)
	return wrap(err, msg)
}

func (e *XErr) Error() string {
	return e.msg
}

func (e *XErr) Unwrap() error {
	return e.err
}

func wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	if msg == "" {
		msg = err.Error()
	}
	return &XErr{
		err: err,
		msg: msg,
	}
}

func logError(err error, format string, params ...any) {
	if err == nil || format == "" {
		return
	}
	logx.Errorf(format, params...)
}

func logInfo(err error, format string, params ...any) {
	if err == nil || format == "" {
		return
	}
	logx.Infof(format, params...)
}
*/
