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

// Convert converts all `higher` level messages to a loglevel ones.
func Convert(log Logger, logLevel LogLevel) Logger { //nolint:ireturn
	conv := loggerConverter{ //nolint:exhaustivestruct
		logger:   log,
		logLevel: logLevel,
		exitFunc: os.Exit,
	}
	switch logLevel {
	case LevelTrace:
		conv.thresholdFunc = conv.logger.Trace
		conv.thresholdFuncf = conv.logger.Tracef
	case LevelDebug:
		conv.thresholdFunc = conv.logger.Debug
		conv.thresholdFuncf = conv.logger.Debugf
	case LevelPrint:
		conv.thresholdFunc = conv.logger.Print
		conv.thresholdFuncf = conv.logger.Printf
	case LevelInfo:
		conv.thresholdFunc = conv.logger.Info
		conv.thresholdFuncf = conv.logger.Infof
	case LevelWarning:
		conv.thresholdFunc = conv.logger.Warning
		conv.thresholdFuncf = conv.logger.Warningf
	case LevelError:
		conv.thresholdFunc = conv.logger.Error
		conv.thresholdFuncf = conv.logger.Errorf
	case LevelFatal:
		conv.thresholdFunc = conv.Fatal
		conv.thresholdFuncf = conv.Fatalf
	case LevelPanic:
		conv.thresholdFunc = conv.logger.Panic
		conv.thresholdFuncf = conv.logger.Panicf
	}

	return conv
}

type loggerConverter struct {
	logger         Logger
	logLevel       LogLevel
	exitFunc       func(code int)
	thresholdFunc  func(args ...interface{})
	thresholdFuncf func(format string, args ...interface{})
}

func (l loggerConverter) Print(args ...interface{}) {
	if l.logLevel < LevelPrint {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Print(args...)
}

func (l loggerConverter) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l loggerConverter) Debug(args ...interface{}) {
	if l.logLevel < LevelDebug {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Debug(args...)
}

func (l loggerConverter) Info(args ...interface{}) {
	if l.logLevel < LevelInfo {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Info(args...)
}

func (l loggerConverter) Warning(args ...interface{}) {
	if l.logLevel < LevelWarning {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Warning(args...)
}

func (l loggerConverter) Error(args ...interface{}) {
	if l.logLevel < LevelError {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Error(args...)
}

func (l loggerConverter) Fatal(args ...interface{}) {
	if l.logLevel < LevelFatal {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Print(args...)
	l.exitFunc(1)
}

func (l loggerConverter) Panic(args ...interface{}) {
	if l.logLevel < LevelPanic {
		l.thresholdFunc(args...)

		return
	}
	l.logger.Panic(args...)
}

func (l loggerConverter) Printf(format string, args ...interface{}) {
	if l.logLevel < LevelPrint {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Printf(format, args...)
}

func (l loggerConverter) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l loggerConverter) Debugf(format string, args ...interface{}) {
	if l.logLevel < LevelDebug {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Debugf(format, args...)
}

func (l loggerConverter) Infof(format string, args ...interface{}) {
	if l.logLevel < LevelInfo {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Infof(format, args...)
}

func (l loggerConverter) Warningf(format string, args ...interface{}) {
	if l.logLevel < LevelWarning {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Warningf(format, args...)
}

func (l loggerConverter) Errorf(format string, args ...interface{}) {
	if l.logLevel < LevelError {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Errorf(format, args...)
}

func (l loggerConverter) Fatalf(format string, args ...interface{}) {
	if l.logLevel < LevelFatal {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Printf(format, args...)
	l.exitFunc(1)
}

func (l loggerConverter) Panicf(format string, args ...interface{}) {
	if l.logLevel < LevelPanic {
		l.thresholdFuncf(format, args...)

		return
	}
	l.logger.Panicf(format, args...)
}
