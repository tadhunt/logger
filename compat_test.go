package logger

import (
	"fmt"
	"testing"
)

func TestErrFmt(t *testing.T) {
	log := NewCompatLogWriter(LogLevel_INFO)

	err := log.ErrFmt("test")
	actual := fmt.Sprintf("%v", err)
	expected := "compat_test.go:11 TestErrFmt(): test"

	if actual != expected {
		t.Fatalf("actual {%s} expected {%s}", actual, expected)
	}
}
