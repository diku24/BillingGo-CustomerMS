package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

//var logrus = logrus.New()

type FileHook struct {
	File *os.File
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func (hook *FileHook) Levels() []logrus.Level {
	//return []logrus.Level{logrus.WarnLevel}
	return logrus.AllLevels
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	msg := fmt.Sprintf("%s: %s\n", entry.Time.Format("2023-01-02 15:09:09"), entry.Message)
	_, err := hook.File.WriteString(msg)

	return err
}

func Init() {

	logrus.SetReportCaller(true)

	//logrus.SetFormatter(&logrus.TextFormatter{})
	formatter := logrus.TextFormatter{
		TimestampFormat:        "2023-05-06 15:04:05",
		DisableColors:          false,
		FullTimestamp:          false,
		DisableLevelTruncation: true,
		QuoteEmptyFields:       true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
		// Customizing delimiters
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFile:  "file",
			//logrus.FieldKeyFunc:  "caller",
		},
	}
	//Log as JSON Instead of default ASCII formatter
	logrus.SetFormatter(&formatter)

	// Open a file for appending logs
	file, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	hook := &FileHook{File: file}
	logrus.AddHook(hook)

	//Output tos stdout instead of default stderr
	//can be any io.writter
	multi := io.MultiWriter(file, os.Stdout)
	logrus.SetOutput(multi)
}
