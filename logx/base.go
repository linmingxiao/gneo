// Copyright 2020 GoFast Author(http://chende.ren). All rights reserved.
// Use of this source code is governed by a MIT license
package logx

import (
	"errors"
	"io"
	"os"
	"sync"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	Reset   = "\033[0m"
)

var (
	//DefWriter      io.Writer = os.Stdout
	DefErrorWriter io.Writer = os.Stderr
)

const (
	InfoLevel   = iota // InfoLevel logs everything
	ErrorLevel         // ErrorLevel includes errors, slows, stacks
	SevereLevel        // SevereLevel only log severe messages
)

const (
	timeFormat     = "2006-01-02T15:04:05.000Z07"
	timeFormatMini = "01-02 15:04:05"

	accessFilename = "access.log"
	errorFilename  = "error.log"
	severeFilename = "severe.log"
	slowFilename   = "slow.log"
	statFilename   = "stat.log"

	consoleMode = "console"
	volumeMode  = "volume"

	levelAlert  = "alert"
	levelInfo   = "info"
	levelError  = "error"
	levelSevere = "severe"
	levelFatal  = "fatal"
	levelSlow   = "slow"
	levelStat   = "stat"

	callerInnerDepth = 5
	flags            = 0x0
)

var (
	ErrLogPathNotSet        = errors.New("log path must be set")
	ErrLogNotInitialized    = errors.New("log not initialized")
	ErrLogServiceNameNotSet = errors.New("log service name must be set")

	writeConsole bool
	logLevel     uint32
	// 6个不同等级的日志输出
	infoLog   WriterCloser
	errorLog  WriterCloser
	severeLog WriterCloser
	slowLog   WriterCloser
	statLog   WriterCloser
	//stackLog  io.Writer

	once        sync.Once
	initialized uint32
	options     logOptions
)

type (
	logEntry struct {
		Timestamp string `json:"@timestamp"`
		Level     string `json:"lv"`
		Duration  string `json:"duration,omitempty"`
		Content   string `json:"ct"`
	}

	logOptions struct {
		gzipEnabled          bool
		logStackArchiveMills int
		keepDays             int
	}

	LogOption func(options *logOptions)

	Logger interface {
		Error(...interface{})
		Errorf(string, ...interface{})
		Info(...interface{})
		Infof(string, ...interface{})
		Slow(...interface{})
		Slowf(string, ...interface{})
		WithDuration(time.Duration) Logger
	}
)
