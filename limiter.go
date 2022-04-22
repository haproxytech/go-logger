// Copyright 2021 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"os"
)

type LogLevel uint8

const (
	LevelTrace   LogLevel = 1
	LevelDebug   LogLevel = 2
	LevelPrint   LogLevel = 3
	LevelInfo    LogLevel = 4
	LevelWarning LogLevel = 5
	LevelError   LogLevel = 6
	LevelFatal   LogLevel = 7
	LevelPanic   LogLevel = 8
)

// limits printing only above threashold.
func Limit(log Logger, logLevel LogLevel) Logger { //nolint:ireturn
	return loggerLimiter{
		logger:   log,
		logLevel: logLevel,
		exitFunc: os.Exit,
	}
}

type loggerLimiter struct {
	logger   Logger
	logLevel LogLevel
	exitFunc func(code int)
}

func (l loggerLimiter) Print(args ...interface{}) {
	if l.logLevel > LevelPrint {
		return
	}
	l.logger.Print(args...)
}

func (l loggerLimiter) Trace(args ...interface{}) {
	if l.logLevel > LevelTrace {
		return
	}
	l.logger.Trace(args...)
}

func (l loggerLimiter) Debug(args ...interface{}) {
	if l.logLevel > LevelDebug {
		return
	}
	l.logger.Debug(args...)
}

func (l loggerLimiter) Info(args ...interface{}) {
	if l.logLevel > LevelInfo {
		return
	}
	l.logger.Info(args...)
}

func (l loggerLimiter) Warning(args ...interface{}) {
	if l.logLevel > LevelWarning {
		return
	}
	l.logger.Warning(args...)
}

func (l loggerLimiter) Error(args ...interface{}) {
	if l.logLevel > LevelError {
		return
	}
	l.logger.Error(args...)
}

func (l loggerLimiter) Fatal(args ...interface{}) {
	if l.logLevel > LevelFatal {
		return
	}
	l.logger.Print(args...)
	l.exitFunc(1)
}

func (l loggerLimiter) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l loggerLimiter) Printf(format string, args ...interface{}) {
	if l.logLevel > LevelPrint {
		return
	}
	l.logger.Printf(format, args...)
}

func (l loggerLimiter) Tracef(format string, args ...interface{}) {
	if l.logLevel > LevelTrace {
		return
	}
	l.logger.Tracef(format, args...)
}

func (l loggerLimiter) Debugf(format string, args ...interface{}) {
	if l.logLevel > LevelDebug {
		return
	}
	l.logger.Debugf(format, args...)
}

func (l loggerLimiter) Infof(format string, args ...interface{}) {
	if l.logLevel > LevelInfo {
		return
	}
	l.logger.Infof(format, args...)
}

func (l loggerLimiter) Warningf(format string, args ...interface{}) {
	if l.logLevel > LevelWarning {
		return
	}
	l.logger.Warningf(format, args...)
}

func (l loggerLimiter) Errorf(format string, args ...interface{}) {
	if l.logLevel > LevelError {
		return
	}
	l.logger.Errorf(format, args...)
}

func (l loggerLimiter) Fatalf(format string, args ...interface{}) {
	if l.logLevel > LevelFatal {
		return
	}
	l.logger.Printf(format, args...)
	l.exitFunc(1)
}

func (l loggerLimiter) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
