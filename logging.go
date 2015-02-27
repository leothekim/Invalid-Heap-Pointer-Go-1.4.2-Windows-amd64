package main

import (
	"github.com/tolsen/slogger/v2"
	"github.com/tolsen/slogger/v2/async_appender"
	"github.com/tolsen/slogger/v2/rolling_file_appender"

	"fmt"
	"os"
)

func LogErrorHandler(err error) {
	fmt.Fprintln(os.Stderr, "Logging error occurred:", err.Error())
}

func NewAsyncRollingFileAppender(logPath string, maxFileSize int64, maxRotatedLogs int, rotateIfExists bool, headerGenerator func() []string) (*async_appender.AsyncAppender, error) {
	var stringWriterCallback func(*os.File) slogger.StringWriter

	rollingAppender, err := rolling_file_appender.NewWithStringWriter(
		logPath,
		maxFileSize,
		maxRotatedLogs,
		rotateIfExists,
		headerGenerator,
		stringWriterCallback,
	)

	if err != nil {
		return nil, err
	}

	async_appender := async_appender.New(
		rollingAppender,
		4096,
		LogErrorHandler,
	)

	return async_appender, nil
}
