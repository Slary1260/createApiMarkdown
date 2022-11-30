/*
 * @Author: tj
 * @Date: 2022-10-10 09:28:09
 * @LastEditors: tj
 * @LastEditTime: 2022-11-30 13:55:51
 * @FilePath: \createApiMarkdown\logger\logrus.go
 */
package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	out  io.Writer
	once sync.Once
)

// DefaultLogrusLogger DefaultLogrusLogger
func DefaultLogrusLogger() {
	NewLogrusLogger("", DefaultLogFilePath(), logrus.InfoLevel)
}

func GetOutWriter() io.Writer {
	return out
}

func cresteWriter(dir, fileName string) {
	once.Do(func() {
		lumberjackLogrotate := &lumberjack.Logger{
			Filename:   filepath.Join(dir, fileName),
			MaxSize:    20, // Max megabytes before log is rotated
			MaxBackups: 10, // Max number of old log files to keep
			MaxAge:     30, // Max number of days to retain log files
			Compress:   true,
		}

		// logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.RFC3339})
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
		logrus.SetReportCaller(true)

		// 编译为dll文件 不能使用os.Stdout
		out = io.MultiWriter(os.Stdout, lumberjackLogrotate)
	})
}

// NewLogrusLogger NewLogrusLogger
func NewLogrusLogger(dir, fileName string, level logrus.Level) {
	cresteWriter(dir, fileName)
	logrus.SetOutput(out)
	logrus.SetLevel(level)
}

// DefaultLogFilePath DefaultLogFilePath
func DefaultLogFilePath() string {
	var logFilePath string
	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))

	switch runtime.GOOS {
	case "android":
		// TODO android logFilePath
		// logFilePath = "/storage/emulated/0/Android/data/com.gdh.project/files/project.log"

	case "windows":
		// TODO windows logFilePath
		// logFilePath = filepath.Join(os.Getenv("AppData"), "project/log", "project.log")
		logFilePath = filepath.Join("./log", fileName)

	case "darwin":
		// logFilePath = "~/Library/Application Support/project/project.log"
		logFilePath = filepath.Join("./log", fileName)

	default:
		// xdgCfg := os.Getenv("XDG_CONFIG_HOME")
		// if xdgCfg != "" {
		// 	logFilePath = filepath.Join(xdgCfg, "project", "project.log")
		// } else {
		// 	logFilePath = filepath.Join("~/.config/project/", "project.log")
		// }
		logFilePath = filepath.Join("./log", fileName)
	}
	return logFilePath
}
