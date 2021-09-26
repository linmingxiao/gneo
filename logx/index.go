package logx

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var logger *Logger
var loggerOnce sync.Once = sync.Once{}


type Logger struct {
	ioWriter io.Writer
	mode string //log mode
	level int8
}

// 全局的初始化日志
func InitLogger(cfg *LogConfig) {

	if cfg == nil{
		cfg = &LogConfig{

		}
	}

	loggerOnce.Do(func() {
		var (
			ioWriter io.Writer
			fileName string = cfg.Path + "/" + cfg.FilePrefix + ".log"
		)
		if cfg.Mode == modeConsole{
			ioWriter = io.MultiWriter(os.Stdout)
		} else if cfg.Mode == modeFile{
			f, _ := os.Create(fileName)
			ioWriter = io.MultiWriter(f)
		} else if cfg.Mode == modeVolume{
			f, _ := os.Create(fileName)
			ioWriter = io.MultiWriter(os.Stdout, f)
		}
		logger = &Logger{
			ioWriter: ioWriter,
			mode: cfg.Mode,
			level: cfg.Level,
		}
	})
}

//********************************************************
//对外的方法

func Trace(values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Trace]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}

func Debug(values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Debug]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}

func Info(values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Info]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}

func Log(values ...interface{})  {
	fmt.Fprintln(logger.ioWriter, values...)
}

func Warn(values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Warn]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}

func Error(values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Error]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}

func Fatal(format string, values ...interface{})  {
	if logger.level >= trace{
		params := []interface{}{"[Fatal]"}
		params = append(params, values...)
		fmt.Fprintln(logger.ioWriter, params...)
	}
}


func Tracef(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Trace] "+format, values...)
	}
}

func Debugf(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Debug] "+format, values...)
	}
}

func Infof(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Info] "+format, values...)
	}
}

func Warnf(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Warn] "+format, values...)
	}
}

func Errorf(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Error] "+format, values...)
	}
}

func Fatalf(format string, values ...interface{})  {
	if logger.level >= trace{
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(logger.ioWriter, "[Fatal] "+format, values...)
	}
}
