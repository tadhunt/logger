package logger

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

func Format(format string, args ...any) string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	prefix := "%s:%d %s: "

	a := make([]any, 0)
	a = append(a, filepath.Base(file), line, fnName)
	a = append(a, args...)

	return fmt.Sprintf(prefix+format, a...)
}

func FuncInfo(f interface{}) string {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer())

	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	file, line := fn.FileLine(fn.Entry())
	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
}

func Caller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "<unknown>"
	}

	pcs := []uintptr{pc}

	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()

	return fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
}

// TODO(tadhunt): Messy: almost identical to Format() except preserves error semantics (%w)
func errorf(depth int, format string, args ...any) error {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	prefix := "%s:%d %s: "

	a := make([]any, 0)
	a = append(a, filepath.Base(file), line, fnName)
	a = append(a, args...)

	return fmt.Errorf(prefix+format, a...)
}

func Errorf(format string, args ...any) error {
	return errorf(2, format, args...)
}
