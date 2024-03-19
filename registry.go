package logger

import(
	"fmt"
	"strconv"
	"sync"
)

type CompatRegistry struct {
	sync.Mutex
	nextId  int
	loggers map[int]CompatLogWriter
}

var Registry = &CompatRegistry{
	nextId: 1,
	loggers: make(map[int]CompatLogWriter),
}


func (r *CompatRegistry) NewId() int {
	r.Lock()
	defer r.Unlock()

	id := r.nextId
	r.nextId++

	return id
}

func (r *CompatRegistry) Add(log CompatLogWriter) {
	r.Lock()
	defer r.Unlock()

	id := log.Id()

	_, found := r.loggers[id]
	if found {
		log.Warnf("duplicate id %d", id)
	}

	r.loggers[id] = log
}

func (r *CompatRegistry) List() []CompatLogWriter {
	r.Lock()
	defer r.Unlock()

	results := make([]CompatLogWriter, len(r.loggers))

	i := 0
	for _, log := range r.loggers {
		results[i] = log
		i++
	}

	return results
}

func (r *CompatRegistry) LookupById(id int) CompatLogWriter {
	r.Lock()
	defer r.Unlock()

	return r.loggers[id]
}

func (r *CompatRegistry) Command(args []string) []string {
	if args == nil || len(args) < 1 {
		return []string{"bad args"}
	}

	cmd := args[0]
	args = args[1:]

	switch cmd {
	case "list":
		loggers := r.List()
		results := make([]string, len(loggers))
		for i, log := range loggers {
			results[i] = fmt.Sprintf("id {%d} prefix {%s} level {%s}", log.Id(), log.Prefix(), log.Level())
		}
		return results
	case "level":
		if len(args) != 2 {
			return []string{"usage: level {id} {loglevel}"}
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return []string{fmt.Sprintf("id '%s': syntax error: %v", args[0], err)}
		}

		level, err := NewLogLevelFromString(args[1])
		if err != nil {
			return []string{fmt.Sprintf("level: syntax error: %v", err)}
		}

		log := Registry.LookupById(id)
		if log == nil {
			return []string{fmt.Sprintf("id %d: not found", id)}
		}

		log.SetLevel(level)

		return []string{"ok"}
	}

	return []string{"unsupported command"}
}
