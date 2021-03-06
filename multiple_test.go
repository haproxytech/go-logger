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
	"time"
)

func Test_gologgers_Print(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Print("test")
	time.Sleep(time.Millisecond)

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
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Trace("test")
	time.Sleep(time.Millisecond)

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
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Debug("test")
	time.Sleep(time.Millisecond)

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
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Info("test")
	time.Sleep(time.Millisecond)

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
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Warning("test")
	time.Sleep(time.Millisecond)

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
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Error("test")
	time.Sleep(time.Millisecond)

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

func Test_gologgers_Fatal(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	flog, ok := log.(*gologgers)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatal("test")
	time.Sleep(time.Millisecond)

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
	if !exitCalled {
		t.FailNow()
	}
}

func Test_gologgers_Panic(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	log.Panic("test")
}

func Test_gologgers_Printf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Printf("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Print" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Print" {
		t.FailNow()
	}
}

func Test_gologgers_Tracef(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Tracef("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Trace" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Trace" {
		t.FailNow()
	}
}

func Test_gologgers_Debugf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Debugf("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Debug" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Debug" {
		t.FailNow()
	}
}

func Test_gologgers_Infof(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Infof("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Info" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Info" {
		t.FailNow()
	}
}

func Test_gologgers_Warningf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Warningf("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Warning" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Warning" {
		t.FailNow()
	}
}

func Test_gologgers_Errorf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	log.Errorf("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Error" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Error" {
		t.FailNow()
	}
}

func Test_gologgers_Fatalf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)
	flog, ok := log.(*gologgers)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatalf("test %s", "str")
	time.Sleep(time.Millisecond)

	if len(mocked1.args) != 1 {
		t.FailNow()
	}
	if mocked1.args[0] != "test str" {
		t.FailNow()
	}
	if mocked1.caller != "Print" {
		t.FailNow()
	}
	if len(mocked2.args) != 1 {
		t.FailNow()
	}
	if mocked2.args[0] != "test str" {
		t.FailNow()
	}
	if mocked2.caller != "Print" {
		t.FailNow()
	}
	if !exitCalled {
		t.FailNow()
	}
}

func Test_gologgers_Panicf(t *testing.T) {
	t.Parallel()

	mocked1 := &loggerMock{} //nolint:exhaustivestruct
	mocked2 := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked1, mocked2)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	log.Panicf("test %s", "str")
}
