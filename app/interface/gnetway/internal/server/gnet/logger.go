// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package gnet

import (
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/zeromicro/go-zero/core/logx"
)

type Logger struct {
	logger logx.Logger
}

func NewLogger() logging.Logger {
	return &Logger{
		logger: logx.WithCallerSkip(2),
	}
}

// Debugf logs messages at DEBUG level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Infof logs messages at INFO level.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warnf logs messages at WARN level.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Errorf logs messages at ERROR level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatalf logs messages at FATAL level.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
