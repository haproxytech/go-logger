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

package standard

import (
	"io"
	"log"
	"os"

	loger "github.com/haproxytech/go-logger"
)

// New returns default log from Go 'log' package, but that logger is compatible with go-logger interface.
func New(out io.Writer, prefix string, flag int) loger.Logger { //nolint:ireturn
	return loggerStandard{
		logger:   log.New(out, prefix, flag),
		exitFunc: os.Exit,
	}
}

// Default returns default log from Go 'log' package, but that logger is compatible with go-logger interface.
func Default() loger.Logger { //nolint:ireturn
	return loggerStandard{
		logger:   log.Default(),
		exitFunc: os.Exit,
	}
}

type loggerStandard struct {
	logger   *log.Logger
	exitFunc func(code int)
}

func (l loggerStandard) Print(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Trace(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Info(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Warning(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Error(args ...interface{}) {
	l.logger.Println(args...)
}

func (l loggerStandard) Fatal(args ...interface{}) {
	l.logger.Print(args...)
	l.exitFunc(1)
}

func (l loggerStandard) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l loggerStandard) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Tracef(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Infof(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Warningf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Errorf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l loggerStandard) Fatalf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
	l.exitFunc(1)
}

func (l loggerStandard) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
