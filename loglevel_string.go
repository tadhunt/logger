// Code generated by "stringer -type=LogLevel"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LogLevel_PACKET-0]
	_ = x[LogLevel_DEBUG-1]
	_ = x[LogLevel_INFO-2]
	_ = x[LogLevel_WARN-3]
	_ = x[LogLevel_ERROR-4]
	_ = x[LogLevel_FATAL-5]
}

const _LogLevel_name = "LogLevel_PACKETLogLevel_DEBUGLogLevel_INFOLogLevel_WARNLogLevel_ERRORLogLevel_FATAL"

var _LogLevel_index = [...]uint8{0, 15, 29, 42, 55, 69, 83}

func (i LogLevel) String() string {
	if i < 0 || i >= LogLevel(len(_LogLevel_index)-1) {
		return "LogLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LogLevel_name[_LogLevel_index[i]:_LogLevel_index[i+1]]
}
