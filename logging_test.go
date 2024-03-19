package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestErrWrap(t *testing.T) {
	origErr := context.Canceled

	err1 := Errorf("foo: %w", origErr)
	err2 := Log.Errorf("bar: %w", origErr)

	if !errors.Is(err1, context.Canceled) {
		t.Fatalf("err1 failed: %v", err1)
	}

	if !errors.Is(err2, context.Canceled) {
		t.Fatalf("err2 failed: %v", err1)
	}

	fmt.Printf("err1: %v\n", err1)
	fmt.Printf("err2: %v\n", err2)

	s := fmt.Sprintf("%v", err1)
	fields := strings.Split(s, ":")

	if fields[0] != "logging_test.go" {
		t.Fatalf("err1: expected file logging_test, got %s", fields[0])
	}

	s = fmt.Sprintf("%v", err2)
	fields = strings.Split(s, ":")

	if fields[0] != "logging_test.go" {
		t.Fatalf("err2: expected file logging_test, got %s", fields[0])
	}
}
