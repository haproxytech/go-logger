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
	"bytes"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()

	buf := new(bytes.Buffer)
	got := New(buf, "", 0)

	// check that we don't panic
	got.Debug("0")
	got.Debugf("1")
	got.Error("2")
	got.Errorf("3")
	got.Info("4")
	got.Infof("5")
	got.Print("6")
	got.Printf("7")
	got.Trace("8")
	got.Tracef("9")
	got.Warning("a")
	got.Warningf("b")

	// check for log output
	result := buf.String()
	if result != "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\na\nb\n" {
		t.Errorf("want 0\n1\n2\n3\n4\n5\n6\n7\n8\n9\na\nb\n got %s", result)
	}
}

func TestNewPanic1(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	got := New(os.Stdout, "", log.LstdFlags)

	got.Panic("aaaa")
}

func TestNewPanic2(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	got := New(os.Stdout, "", log.LstdFlags)

	got.Panicf("aaaa")
}

func TestDefaultConform(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()

	got := Default()

	// check that we don't panic
	got.Debug("0")
	got.Debugf("1")
	got.Error("2")
	got.Errorf("3")
	got.Info("4")
	got.Infof("5")
	got.Print("6")
	got.Printf("7")
	got.Trace("8")
	got.Tracef("9")
	got.Warning("1")
	got.Warningf("b")
}

func TestDefaultPanic1(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	got := Default()

	got.Panic("aaaa")
}

func TestDefaultPanic2(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	got := Default()

	got.Panicf("aaaa")
}
