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

func Test_loggerLimiter_Trace(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Trace("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_TraceEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelDebug)

	log.Trace("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Debug(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Debug("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_DebugEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelPrint)

	log.Debug("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Print(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Print("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_PrintEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelInfo)

	log.Print("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Info(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Info("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_InfoEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelWarning)

	log.Info("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Warning(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Warning("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_WarningEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelError)

	log.Warning("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Error(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Error("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_ErrorEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelFatal)

	log.Error("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Fatal(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)
	flog, ok := log.(loggerLimiter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatal("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
	if !exitCalled {
		t.FailNow()
	}
}

func Test_loggerLimiter_FatalEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelPanic)
	flog, ok := log.(loggerLimiter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatal("str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
	if exitCalled {
		t.FailNow()
	}
}

func Test_loggerLimiter_Panic(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Panic("str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Tracef(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Tracef("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_TracefEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelDebug)

	log.Tracef("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Debugf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Debugf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_DebugfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelPrint)

	log.Debugf("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Printf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Printf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_PrintfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelInfo)

	log.Printf("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Infof(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Infof("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_InfofEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelWarning)

	log.Infof("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Warningf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Warningf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_WarningfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelError)

	log.Warningf("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Errorf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Errorf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

func Test_loggerLimiter_ErrorfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelFatal)

	log.Errorf("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
}

func Test_loggerLimiter_Fatalf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)
	flog, ok := log.(loggerLimiter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatalf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
	if !exitCalled {
		t.FailNow()
	}
}

func Test_loggerLimiter_FatalfEmpty(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)

	log = Limit(log, LevelPanic)
	flog, ok := log.(loggerLimiter)
	if !ok {
		t.FailNow()
	}
	exitCalled := false
	flog.exitFunc = func(code int) {
		exitCalled = true
	}
	log = flog

	log.Fatalf("%s", "str")

	if len(mocked.args) > 0 {
		t.FailNow()
	}
	if exitCalled {
		t.FailNow()
	}
}

func Test_loggerLimiter_Panicf(t *testing.T) {
	t.Parallel()

	mocked := &loggerMock{} //nolint:exhaustivestruct
	log := New(mocked)
	log = Limit(log, LevelTrace)

	log.Panicf("%s", "str")

	if len(mocked.args) < 1 {
		t.FailNow()
	}
}

/*
	LevelTrace   LogLevel = 1
	LevelDebug   LogLevel = 2
	LevelPrint   LogLevel = 3
	LevelInfo    LogLevel = 4
	LevelWarning LogLevel = 5
	LevelError   LogLevel = 6
	LevelFatal   LogLevel = 7
	LevelPanic   LogLevel = 8
*/
