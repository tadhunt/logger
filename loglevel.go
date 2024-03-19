package logger

//go:generate go run golang.org/x/tools/cmd/stringer -type=LogLevel

/*
 * The go generate line using go run (and tools.go file) solve a couple of problems WRT
 * dependency mgmt.  Details here: https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/
 */
import (
	"fmt"
)

type LogLevel int

const (
	LogLevel_PACKET LogLevel = iota
	LogLevel_DEBUG
	LogLevel_INFO
	LogLevel_WARN
	LogLevel_ERROR
	LogLevel_FATAL
)

func NewLogLevelFromString(s string) (LogLevel, error) {
	for level := LogLevel_PACKET; level <= LogLevel_FATAL; level++ {
		if s == level.String() {
			return level, nil
		}
	}

	return LogLevel_PACKET, fmt.Errorf("bad loglevel: '%s'", s)
}
