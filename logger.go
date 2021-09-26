package gneo

import (
	"fmt"
	"github.com/linmingxiao/gneo/logx"
	"net/http"
	"time"
)

type LogFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time
	StatusCode   int
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	ErrorMessage string
	BodySize     int
	Keys         map[string]interface{}
}

var routeLogFormatter = func(param LogFormatterParams) string {
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

func Logger() HandlerFunc{
	return func(c *Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		param := LogFormatterParams{
			Request: c.Request,
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
		logx.Log(routeLogFormatter(param))
	}
}



























