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

import "testing"

func Test_loggerConverter_Trace(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Trace("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_Debug(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Debug("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_DebugEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Debug("str")

	if mocked.caller != "Debug" {
		t.FailNow()
	}
}

func Test_loggerConverter_Print(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Print("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_PrintEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Print("str")

	if mocked.caller != "Print" {
		t.FailNow()
	}
}

func Test_loggerConverter_Info(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Info("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_InfoEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Info("str")

	if mocked.caller != "Info" {
		t.FailNow()
	}
}

func Test_loggerConverter_Warning(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Warning("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_WarningEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Warning("str")

	if mocked.caller != "Warning" {
		t.FailNow()
	}
}

func Test_loggerConverter_Error(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Error("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_ErrorEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Error("str")

	if mocked.caller != "Error" {
		t.FailNow()
	}
}

func Test_loggerConverter_Fatal(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Fatal("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_FatalEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)
	flog, ok := log.(loggerConverter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatal("str")

	if mocked.caller != "Print" {
		t.FailNow()
	}
	if !exitCalled {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Panic("str")

	if mocked.caller != "Trace" {
		t.FailNow()
	}
}

func Test_loggerConverter_PanicOK(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelPanic)

	log.Panic("str")

	if mocked.caller != "Panic" {
		t.FailNow()
	}
}

func Test_loggerConverter_Tracef(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Tracef("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_Debugf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Debugf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_DebugfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Debugf("%s", "str")

	if mocked.caller != "Debugf" {
		t.FailNow()
	}
}

func Test_loggerConverter_Printf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Printf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_PrintfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Printf("%s", "str")

	if mocked.caller != "Printf" {
		t.FailNow()
	}
}

func Test_loggerConverter_Infof(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Infof("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_InfofEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Infof("%s", "str")

	if mocked.caller != "Infof" {
		t.FailNow()
	}
}

func Test_loggerConverter_Warningf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Warningf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_WarningfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Warningf("%s", "str")

	if mocked.caller != "Warningf" {
		t.FailNow()
	}
}

func Test_loggerConverter_Errorf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Errorf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_ErrorfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)

	log.Errorf("%s", "str")

	if mocked.caller != "Errorf" {
		t.FailNow()
	}
}

func Test_loggerConverter_Fatalf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)
	flog, ok := log.(loggerConverter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatalf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
	if exitCalled {
		t.FailNow()
	}
}

func Test_loggerConverter_FatalfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Convert(log, LevelPanic)
	flog, ok := log.(loggerConverter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatalf("%s", "str")

	if mocked.caller != "Printf" {
		t.FailNow()
	}
	if !exitCalled {
		t.FailNow()
	}
}

func Test_loggerConverter_Panicf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelTrace)

	log.Panicf("%s", "str")

	if mocked.caller != "Tracef" {
		t.FailNow()
	}
}

func Test_loggerConverter_PanicfOK(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelPanic)

	log.Panicf("%s", "str")

	if mocked.caller != "Panicf" {
		t.FailNow()
	}
}

// test all levels

func Test_loggerConverter_Panic1(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelDebug)

	log.Panic("%s", "str")

	if mocked.caller != "Debug" {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic2(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelPrint)

	log.Panic("%s", "str")

	if mocked.caller != "Print" {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic3(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelInfo)

	log.Panic("%s", "str")

	if mocked.caller != "Info" {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic4(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelWarning)

	log.Panic("%s", "str")

	if mocked.caller != "Warning" {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic5(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelError)

	log.Panic("%s", "str")

	if mocked.caller != "Error" {
		t.FailNow()
	}
}

func Test_loggerConverter_Panic6(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Convert(log, LevelFatal)
	flog, ok := log.(loggerConverter)
	if !ok {
		t.FailNow()
	}
	flog.exitFunc = func(code int) {}
	flog.thresholdFunc = flog.Fatal // othervise is wrong func and exitFunc is not overidden

	log = flog

	log.Panic("%s", "str")

	if mocked.caller != "Print" {
		t.FailNow()
	}
}
