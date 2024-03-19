package logger

// Modified from https://wycd.net/posts/2014-07-02-logging-function-names-in-go.html

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
)

var Log = NewLogWriter()

type LogWriter interface{
	Printf(format string, args ...any)
	Fatalf(format string, args ...any)
	SetPrefix(prefix string)
	Errorf(f string, args ...any) error
	PrettyPrint(arg interface{})
}

type logWriter struct{
}

func NewLogWriter() LogWriter {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	return logWriter{}
}

func (lw logWriter) Printf(format string, args ...any) {
	s := Format(format, args...)
	log.Printf("%s", s)
}

func (lw logWriter) Fatalf(format string, args ...any) {
	s := Format(format, args...)
	log.Fatalf("%s", s)
}

func (lw logWriter) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

func (lw logWriter) Errorf(f string, args ...any) error {
	return errorf(2, f, args...)
}

func (lw logWriter) PrettyPrint(arg interface{}) {
	out, err := json.MarshalIndent(arg, "", " ")
	if err != nil {
		lw.Printf("[can't jsonify]: %v", arg)
		return
	}

	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		text := scanner.Text()
		s := Format("%s", text)
		log.Printf("%s", s)
	}
}
