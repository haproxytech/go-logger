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

type gologger struct {
	logger Logger
}

func (l *gologger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *gologger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *gologger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *gologger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *gologger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *gologger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *gologger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *gologger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *gologger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *gologger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *gologger) Infof(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *gologger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

func (l *gologger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *gologger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
