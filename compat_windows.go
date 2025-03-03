//go:build windows

package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type compatLogWriter struct {
	id     int
	logger *log.Logger
	prefix string
	level  LogLevel

	wl      sync.Mutex
	wbuf    *bytes.Buffer
	scanner *bufio.Scanner
	scount  int
}

func NewCompatLogWriter(level LogLevel) CompatLogWriter {
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(buf)

	log := &compatLogWriter{
		id:     Registry.NewId(),
		logger: log.New(os.Stderr, "", log.Ltime | log.Lmicroseconds),
		prefix: "",
		level:  level,

		wl:      sync.Mutex{},
		wbuf:    buf,
		scanner: scanner,
		scount:  0,
	}

	Registry.Add(log)

	return log
}

func (lw *compatLogWriter) Syslog(enabled bool) error {
	if enabled {
		return fmt.Errorf("windows does not support syslog")
	}
	return nil
}

func (lw *compatLogWriter) SetLevel(level LogLevel) {
	oldLevel := lw.level
	lw.level = level

	if lw.level <= LogLevel_INFO {
		s := Format("LogLevel %s → %s", oldLevel, level)
		lw.logger.Printf("[INFO] %s", s)
	}
}

func (lw *compatLogWriter) Level() LogLevel {
	return lw.level
}

func (lw *compatLogWriter) SetPrefix(prefix string) {
	lw.prefix = prefix
	lw.logger.SetPrefix(prefix)
}

func (lw *compatLogWriter) Packetf(format string, args ...any) {
	if lw.level > LogLevel_PACKET {
		return
	}

	s := "[PACKET] " + Format(format, args...)

	lw.logger.Printf("%s", s)
}

func (lw *compatLogWriter) Debugf(format string, args ...any) {
	if lw.level > LogLevel_DEBUG {
		return
	}

	s := "[DEBUG] " + Format(format, args...)

	lw.logger.Printf("%s", s)
}

func (lw *compatLogWriter) Infof(format string, args ...any) {
	if lw.level > LogLevel_INFO {
		return
	}

	s := "[INFO] " + Format(format, args...)

	lw.logger.Printf("%s", s)
}

func (lw *compatLogWriter) Warnf(format string, args ...any) {
	if lw.level > LogLevel_WARN {
		return
	}

	s := "[WARN] " + Format(format, args...)

	lw.logger.Printf("%s", s)
}

func (lw *compatLogWriter) Errorf(f string, args ...any) {
	if lw.level > LogLevel_ERROR {
		return
	}

	s := "[ERROR] " + Format(f, args...)

	lw.logger.Printf("%s", s)
}

func (lw *compatLogWriter) Fatalf(format string, args ...any) {
	s := "[FATAL] " + Format(format, args...)

	lw.logger.Fatalf("%s", s)
}

func (lw *compatLogWriter) ErrFmt(format string, args ...any) error {
	err := errorf(2, lw.prefix + format, args...)
	return err
}

func (lw *compatLogWriter) Prefix() string {
	return lw.prefix
}

func (lw *compatLogWriter) Write(data []byte) (int, error) {
	lw.wl.Lock()
	defer lw.wl.Unlock()

	lw.wbuf.Write(data)
	lw.scan("[LOG]")

	return len(data), nil
}

// must be called with lw.wl locked
func (lw *compatLogWriter) scan(prefix string) {
	for {
		data := lw.wbuf.Bytes()
		i := bytes.IndexByte(data, '\n')
		if i < 0 {
			return
		}
		text := data[:i]
		lw.wbuf.Next(i + 1)

		s := prefix + " " + string(text)
		lw.logger.Printf("%s", s)
		lw.scount++
	}
}

func (lw *compatLogWriter) Flush() {
	lw.wl.Lock()
	defer lw.wl.Unlock()

	data := lw.wbuf.Bytes()
	if len(data) == 0 {
		return
	}

	if data[len(data)-1] != '\n' {
		lw.wbuf.Write([]byte{'\n'})
	}

	lw.scan("[LOG]")
}

func (lw *compatLogWriter) Json(label string, data any, indent bool) {
	var raw []byte
	var err error

	if indent {
		raw, err = json.MarshalIndent(data, "", "  ")
	} else {
		raw, err = json.Marshal(data)
	}

	if err != nil {
		s := Format("[JSON-ERR] marshal: %v", err)
		lw.logger.Printf("%s", s)
		return
	}

	lines := strings.Split(string(raw), "\n")
	lines[0] = fmt.Sprintf("%s: %s", label, lines[0])

	for _, line := range lines {
		s := Format("[JSON] %s", line)
		lw.logger.Printf("%s", s)
	}
}

func (lw *compatLogWriter) SetId(id int) {
	lw.id = id
}

func (lw *compatLogWriter) Id() int {
	return lw.id
}
