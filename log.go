package logging

import (
	"log"
	"log/syslog"
	"os"
)

/**
 * Package name: logging
 * Project name: Logging
 * Created by: Praveen Premaratne
 * Created on: 22/03/2020 20:14
 */

var LogFilePath string
var Tag string

type logging struct {
	level   string
	message string
}

type logger interface {
	write() error
}

// Error defines the log implementation for error logging
func Error(msg string) {
	var l logging
	l.level = "error"
	l.message = msg
	if err := l.write(); err != nil {
		panic(err)
	}
	os.Exit(1)
}

// Warning defines log implementation for warning logging
func Warning(msg string) {
	var l logging
	l.level = "warning"
	l.message = msg
	if err := l.write(); err != nil {
		panic(err)
	}
}

// Info defines log implementation for info logging
func Info(msg string) {
	var l logging
	l.level = "info"
	l.message = msg
	if err := l.write(); err != nil {
		panic(err)
	}
}

// Debug defines log implementation for debug logging
func Debug(msg string) {
	var l logging
	l.level = "debug"
	l.message = msg
	if err := l.write(); err != nil {
		panic(err)
	}
}

// write creates log file if not exist and write logs to that file.
// If the file creation fails, it will attempt to utilise syslog as a log destination
// and if that fails, it will return the error back to caller for handling.
func (l *logging) write() error {
	logFile, fErr := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if fErr != nil {
		logFile, err := syslog.New(syslog.LOG_NOTICE, Tag)
		if err != nil {
			return err
		} else {
			defer logFile.Close()

			log.SetOutput(logFile)

			log.Printf("%s: %s\n", l.level, l.message)
		}
	} else {
		defer logFile.Close()

		log.SetOutput(logFile)

		log.Printf("%s: %s\n", l.level, l.message)
	}
	return fErr
}
