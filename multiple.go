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
	"fmt"
)

type gologgers struct {
	loggers  []Logger
	exitFunc func(code int)
}

func (l *gologgers) Print(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Print(args...)
	}
}

func (l *gologgers) Trace(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Trace(args...)
	}
}

func (l *gologgers) Debug(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Debug(args...)
	}
}

func (l *gologgers) Info(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Info(args...)
	}
}

func (l *gologgers) Warning(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Warning(args...)
	}
}

func (l *gologgers) Error(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Error(args...)
	}
}

func (l *gologgers) Fatal(args ...interface{}) {
	for _, log := range l.loggers {
		log := log
		go log.Print(args...)
	}
	l.exitFunc(1)
}

func (l *gologgers) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	for _, log := range l.loggers {
		log := log
		go log.Print(s)
	}
	panic(s)
}

func (l *gologgers) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Print(s)
	}
}

func (l *gologgers) Tracef(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Trace(s)
	}
}

func (l *gologgers) Debugf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Debug(s)
	}
}

func (l *gologgers) Infof(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Info(s)
	}
}

func (l *gologgers) Warningf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Warning(s)
	}
}

func (l *gologgers) Errorf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Error(s)
	}
}

func (l *gologgers) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Print(s)
	}
	l.exitFunc(1)
}

func (l *gologgers) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for _, log := range l.loggers {
		log := log
		go log.Print(s)
	}
	panic(s)
}
