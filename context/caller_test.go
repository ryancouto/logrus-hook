package hook

import (
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestCaller(t *testing.T) {
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.Debug("debug info")
	logrus.WithFields(logrus.Fields{
		"name": "john smith",
		"age":  23,
	}).Info("with fields debug info")

	logrus.Warn("warn")
	logrus.WithField("class", 5).Warn("with fields warn")
}

func TestBenchmarkCaller(t *testing.T) {
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()

	logrus.SetOutput(fd)

	for i := 0; i < 10; i++ {
		logrus.Debug("debug to file")
	}
}

func TestBenchmarkCallerWithField(t *testing.T) {
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	// logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()
	logrus.SetOutput(fd)

	for i := 0; i < 10; i++ {
		logrus.WithField("name", "john").Debug("debug to file with field")
	}
}

func TestBenchmarkCallerWithFields(t *testing.T) {
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()
	logrus.SetOutput(fd)

	for i := 0; i < 10; i++ {
		logrus.WithFields(logrus.Fields{
			"name":  "john smith",
			"age":   32,
			"class": 3,
		}).Info("info to file with field")
	}
}
