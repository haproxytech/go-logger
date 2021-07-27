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

func Test_gologgers_Print(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Print("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Print" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Print" {
		t.FailNow()
	}
}

func Test_gologgers_Trace(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Trace("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Trace" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Trace" {
		t.FailNow()
	}
}

func Test_gologgers_Debug(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Debug("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Debug" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Debug" {
		t.FailNow()
	}
}

func Test_gologgers_Info(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Info("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Info" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Info" {
		t.FailNow()
	}
}

func Test_gologgers_Warning(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Warning("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Warning" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Warning" {
		t.FailNow()
	}
}

func Test_gologgers_Error(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Error("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Error" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Error" {
		t.FailNow()
	}
}

func Test_gologgers_Panic(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Panic("test")
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test" {
		t.FailNow()
	}
	if mocked1.caller != "Panic" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test" {
		t.FailNow()
	}
	if mocked2.caller != "Panic" {
		t.FailNow()
	}
}

func Test_gologgers_Printf(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Printf("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Printf" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Printf" {
		t.FailNow()
	}
}

func Test_gologgers_Tracef(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Tracef("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Tracef" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_gologgers_Debugf(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Debugf("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Debugf" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Debugf" {
		t.FailNow()
	}
}

func Test_gologgers_Infof(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Infof("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Infof" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Infof" {
		t.FailNow()
	}
}

func Test_gologgers_Warningf(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Warningf("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Warningf" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Warningf" {
		t.FailNow()
	}
}

func Test_gologgers_Errorf(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Errorf("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Errorf" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Errorf" {
		t.FailNow()
	}
}

func Test_gologgers_Panicf(t *testing.T) {
	mocked1 := &loggerMock{}
	mocked2 := &loggerMock{}
	log := New(mocked1, mocked2)
	log.Panicf("test %s", "str")
	if mocked1.format != "test %s" {
		t.FailNow()
	}
	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "str" {
		t.FailNow()
	}
	if mocked1.caller != "Panicf" {
		t.FailNow()
	}
	if mocked2.format != "test %s" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "str" {
		t.FailNow()
	}
	if mocked2.caller != "Panicf" {
		t.FailNow()
	}
}
