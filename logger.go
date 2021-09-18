package gneo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mattn/go-isatty"
)

type LoggerConfig struct {
	Formatter LogFormatter
	Output io.Writer
	SkipPaths []string
}

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time
	StatusCode   int
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	ErrorMessage string
	isTerm       bool
	BodySize     int
	Keys         map[string]interface{}
}

var defaultLogFormatter = func(param LogFormatterParams) string {
	if param.Latency > time.Minute{
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[ %s - %#v ] %s | %s | %v \n%s",
		param.Method, param.Path,

		param.ClientIP,

		param.TimeStamp.Format("2006/01/02-15:04:05"),

		param.Latency,

		param.ErrorMessage,
	)
}

func ErrorLogger() HandlerFunc{
	return ErrorLoggerT(ErrorTypeAny)
}

func ErrorLoggerT(typ ErrorType) HandlerFunc{
	return func(c *Context){
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0{
			c.JSON(-1, errors)
		}
	}
}

func Logger() HandlerFunc{
	return LoggerWithConfig(LoggerConfig{})
}

func LoggerWithFormatter(f LogFormatter) HandlerFunc{
	return LoggerWithConfig(LoggerConfig{
		Formatter: f,
	})
}

func LoggerWithWriter(out io.Writer, notLogged ...string) HandlerFunc{
	return LoggerWithConfig(LoggerConfig{
		Output: out,
		SkipPaths: notLogged,
	})
}

func LoggerWithConfig(conf LoggerConfig) HandlerFunc{
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = DefaultWriter
	}

	notlogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request: c.Request,
				isTerm:  isTerm,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			fmt.Fprint(out, formatter(param))
		}
	}
}



























