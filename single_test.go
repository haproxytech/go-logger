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
	"testing"
)

func Test_gologger_Print(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Print("test")

	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Print" {
		t.FailNow()
	}
}

func Test_gologger_Trace(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Trace("test")
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_gologger_Debug(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Debug("test")
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Debug" {
		t.FailNow()
	}
}

func Test_gologger_Info(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Info("test")

	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Info" {
		t.FailNow()
	}
}

func Test_gologger_Warning(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Warning("test")

	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Warning" {
		t.FailNow()
	}
}

func Test_gologger_Error(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Error("test")

	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Error" {
		t.FailNow()
	}
}

func Test_gologger_Panic(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Panic("test")

	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "test" {
		t.FailNow()
	}
	if mocked.caller != "Panic" {
		t.FailNow()
	}
}

func Test_gologger_Printf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Printf("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Printf" {
		t.FailNow()
	}
}

func Test_gologger_Tracef(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Tracef("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_gologger_Debugf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Debugf("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Debugf" {
		t.FailNow()
	}
}

func Test_gologger_Infof(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Infof("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Infof" {
		t.FailNow()
	}
}

func Test_gologger_Warningf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Warningf("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Warningf" {
		t.FailNow()
	}
}

func Test_gologger_Errorf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Errorf("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Errorf" {
		t.FailNow()
	}
}

func Test_gologger_Panicf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{}
	log := New(mocked)

	log.Panicf("test %s", "str")

	if mocked.format != "test %s" {
		t.FailNow()
	}
	if len(mocked.args) != 1 {
		t.FailNow()
	}
	if mocked.args[0] != "str" {
		t.FailNow()
	}
	if mocked.caller != "Panicf" {
		t.FailNow()
	}
}
