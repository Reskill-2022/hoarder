package errors

import "testing"

func TestError(t *testing.T) {
	err := New("test", 500)

	if err.Error() != "test" {
		t.Fatal("error message does not match")
	}
	if CodeFrom(err, 0) != 500 {
		t.Fatal("error code does not match")
	}
	if !CodeIs(err, 500) {
		t.Fatal("expected error code to match 500")
	}

	err2 := From(err, "test2", 501).(Error)
	if err2.Cause() != err {
		t.Fatal("expected cause to match original error")
	}
}
