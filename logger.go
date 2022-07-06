package logger

// Modified from https://wycd.net/posts/2014-07-02-logging-function-names-in-go.html

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"reflect"
)

type LogWriter struct{}

func (f LogWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(3)
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

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)
	return len(p), nil
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

func New() *log.Logger {
	return log.New(LogWriter{}, "", 0)
}
