package logger

import(
	"testing"
)

func TestWrite(t *testing.T) {
	log := NewCompatLogWriter(LogLevel_DEBUG)

	lines := []string{
		"this is line 1\n",
		"this is line 2\n",
		"this is line 3\n",
	}

	for i, line := range lines {
		data := []byte(line)
		n, err := log.Write(data)
		if n != len(line) {
			t.Fatalf("line %d bad n %d expected %d", i, n, len(line))
		}
		if err != nil {
			t.Fatalf("line %d: unexpected err %v", i, err)
		}
	}

	l := log.(*compatLogWriter)

	if l.scount != len(lines) {
		t.Fatalf("got %d lines expected %d", l.scount, len(lines))
	}
}

func TestWritePartial(t *testing.T) {
	log := NewCompatLogWriter(LogLevel_DEBUG)

	lines := []string{
		"this is line 1",
		"\n",
		"this is line 2",
		"\n",
		"this is line 3",
	}

	for i, line := range lines {
		data := []byte(line)
		n, err := log.Write(data)
		if n != len(line) {
			t.Fatalf("line %d bad n %d expected %d", i, n, len(line))
		}
		if err != nil {
			t.Fatalf("line %d: unexpected err %v", i, err)
		}
	}

	l := log.(*compatLogWriter)

	if l.scount != 2 {
		t.Fatalf("got %d lines expected %d", l.scount, 2)
	}

	l.Flush()

	if l.scount != 3 {
		t.Fatalf("got %d lines expected %d", l.scount, 3)
	}
}
