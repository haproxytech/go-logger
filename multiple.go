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

type gologgers struct {
	loggers []Logger
}

func (l *gologgers) Print(args ...interface{}) {
	for _, log := range l.loggers {
		log.Print(args...)
	}
}

func (l *gologgers) Trace(args ...interface{}) {
	for _, log := range l.loggers {
		log.Trace(args...)
	}
}

func (l *gologgers) Debug(args ...interface{}) {
	for _, log := range l.loggers {
		log.Debug(args...)
	}
}

func (l *gologgers) Info(args ...interface{}) {
	for _, log := range l.loggers {
		log.Info(args...)
	}
}

func (l *gologgers) Warning(args ...interface{}) {
	for _, log := range l.loggers {
		log.Warning(args...)
	}
}

func (l *gologgers) Error(args ...interface{}) {
	for _, log := range l.loggers {
		log.Error(args...)
	}
}

func (l *gologgers) Panic(args ...interface{}) {
	for _, log := range l.loggers {
		log.Panic(args...)
	}
}

func (l *gologgers) Printf(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Printf(format, args...)
	}
}

func (l *gologgers) Tracef(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Tracef(format, args...)
	}
}

func (l *gologgers) Debugf(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Debugf(format, args...)
	}
}

func (l *gologgers) Infof(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Infof(format, args...)
	}
}

func (l *gologgers) Warningf(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Warningf(format, args...)
	}
}

func (l *gologgers) Errorf(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Errorf(format, args...)
	}
}

func (l *gologgers) Panicf(format string, args ...interface{}) {
	for _, log := range l.loggers {
		log.Panicf(format, args...)
	}
}
