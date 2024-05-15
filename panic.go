package go_testing

import (
	"testing"
)

func AssertPanic(t *testing.T, panicError string, f func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("%s should have panicked", t.Name())
		}
		if r != panicError {
			ErrorMessage("Expected", panicError)

			if str, ok := r.(string); ok {
				SuccessMessage("Received", str)
			} else {
				t.Fatalf("%v - is not a string", r)
			}

			t.Fatalf("panic error message mismatch")
		}
	}()
	f()
}
