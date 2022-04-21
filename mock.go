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

type loggerMock struct {
	caller string
	format string
	args   []string
}

func argsToStr(args []interface{}) (result []string) {
	for _, arg := range args {
		a := arg.(string) //nolint:forcetypeassert
		result = append(result, a)
	}

	return result
}

func (m *loggerMock) save(caller string, args []interface{}) {
	m.caller = caller
	m.format = ""
	m.args = argsToStr(args)
}

func (m *loggerMock) saveWithFormat(caller string, format string, args []interface{}) {
	m.caller = caller
	m.format = format
	m.args = argsToStr(args)
}

func (m *loggerMock) Print(args ...interface{}) {
	m.save("Print", args)
}

func (m *loggerMock) Trace(args ...interface{}) {
	m.save("Trace", args)
}

func (m *loggerMock) Debug(args ...interface{}) {
	m.save("Debug", args)
}

func (m *loggerMock) Info(args ...interface{}) {
	m.save("Info", args)
}

func (m *loggerMock) Warning(args ...interface{}) {
	m.save("Warning", args)
}

func (m *loggerMock) Error(args ...interface{}) {
	m.save("Error", args)
}

func (m *loggerMock) Fatal(args ...interface{}) {
}

func (m *loggerMock) Panic(args ...interface{}) {
	m.save("Panic", args)
}

func (m *loggerMock) Printf(format string, args ...interface{}) {
	m.saveWithFormat("Printf", format, args)
}

func (m *loggerMock) Tracef(format string, args ...interface{}) {
	m.saveWithFormat("Tracef", format, args)
}

func (m *loggerMock) Debugf(format string, args ...interface{}) {
	m.saveWithFormat("Debugf", format, args)
}

func (m *loggerMock) Infof(format string, args ...interface{}) {
	m.saveWithFormat("Infof", format, args)
}

func (m *loggerMock) Warningf(format string, args ...interface{}) {
	m.saveWithFormat("Warningf", format, args)
}

func (m *loggerMock) Errorf(format string, args ...interface{}) {
	m.saveWithFormat("Errorf", format, args)
}

func (m *loggerMock) Fatalf(format string, args ...interface{}) {
}

func (m *loggerMock) Panicf(format string, args ...interface{}) {
	m.saveWithFormat("Panicf", format, args)
}
