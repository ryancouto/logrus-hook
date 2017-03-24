package hook

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
)

type CallerHook struct {
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	switch len(entry.Data) {
	case 0:
		// used without logrus.WithFields
		entry.Data["context"] = hook.caller(7)
	default:
		// used with logrus.WithFields
		entry.Data["context"] = hook.caller(5)
	}

	return nil
}

func (hook *CallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *CallerHook) caller(skip int) string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	}
	// Unable to determine caller
	return "???"
}
