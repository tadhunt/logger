package logger

import (
	"bytes"
	"bufio"
	"encoding/json"
	"testing"
)


type TestLogWriter struct {
	t *testing.T
}

func NewTestLogWriter(t *testing.T) LogWriter {
	return &TestLogWriter{
		t: t,
	}
}

func (lw TestLogWriter) Printf(format string, args ...any) {
	s := Format(format, args...)
	lw.t.Log(s)
}

func (lw TestLogWriter) Fatalf(format string, args ...any) {
	s := Format(format, args...)
	lw.t.Fatalf("%s", s)
}

func (lw TestLogWriter) SetPrefix(prefix string) {
}

func (lw TestLogWriter) Errorf(f string, args ...any) error {
	return errorf(2, f, args...)
}

func (lw TestLogWriter) PrettyPrint(arg interface{}) {
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
		lw.t.Log(s)
	}
}

type testCompatLogWriter struct {
	id     int
	t      *testing.T
	level  LogLevel
	prefix string
}

func NewTestCompatLogWriter(t *testing.T) CompatLogWriter {
	return &testCompatLogWriter{
		id:     0,
		t:      t,
		level:  LogLevel_PACKET,
		prefix: "",
	}
}

func (lw *testCompatLogWriter) Syslog(enabled bool) error {
	return nil
}

func (lw *testCompatLogWriter) SetLevel(level LogLevel) {
	lw.level = level
}

func (lw *testCompatLogWriter) Level() LogLevel {
	return lw.level
}

func (lw *testCompatLogWriter) SetPrefix(prefix string) {
	lw.prefix = prefix
}

func (lw *testCompatLogWriter) Packetf(format string, args ...any) {
	if lw.level > LogLevel_PACKET {
		return
	}

	s := Format(format, args...)
	lw.t.Log("[PACKET] " + s)
}

func (lw *testCompatLogWriter) Debugf(format string, args ...any) {
	if lw.level > LogLevel_DEBUG {
		return
	}

	s := Format(format, args...)
	lw.t.Log("[DEBUG] " + s)
}

func (lw *testCompatLogWriter) Infof(format string, args ...any) {
	if lw.level > LogLevel_INFO {
		return
	}

	s := Format(format, args...)
	lw.t.Log("[INFO] " + s)
}

func (lw *testCompatLogWriter) Warnf(format string, args ...any) {
	if lw.level > LogLevel_WARN {
		return
	}

	s := Format(format, args...)
	lw.t.Log("[WARN] " + s)
}

func (lw *testCompatLogWriter) Errorf(f string, args ...any) {
	if lw.level > LogLevel_ERROR {
		return
	}

	s := Format(f, args...)
	lw.t.Log("[ERROR] " + s)
}

func (lw *testCompatLogWriter) Fatalf(format string, args ...any) {
	s := Format(format, args...)
	lw.t.Fatalf("[FATAL] %s", s)
}

func (lw *testCompatLogWriter) ErrFmt(format string, args ...any) error {
	return errorf(2, format, args...)
}

func (lw *testCompatLogWriter) Prefix() string {
	return lw.prefix
}

func (lw *testCompatLogWriter) Write(data []byte) (int, error) {
	panic("unimplemented")
}

func (lw *testCompatLogWriter) Json(label string, data any, indent bool) {
	panic("unimplemented")
}

func (lw *testCompatLogWriter) Flush() {
}

func (lw *testCompatLogWriter) SetId(id int) {
	lw.id = id
}

func (lw *testCompatLogWriter) Id() int {
	return lw.id
}
