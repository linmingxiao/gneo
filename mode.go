package gneo

import (
	"io"
	"os"

	"github.com/linmingxiao/gneo/binding"
)

// EnvGinMode indicates environment name for gin mode.
const EnvGinMode = "GIN_MODE"

const (
	//开发调试模式
	DebugMode = "debug"
	//测试模式
	TestMode = "test"
	//预生产模式
	ReleaseMode = "release"
	//生产模式
	ProductMode = "product"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++
// 当前运行处于啥模式：
const (
	modeDebugCode = iota
	modeTestCode
	modeReleaseCode
	modeProductCode
)

// 日志文件的目标系统
const (
	LogTypeConsole    = "console"
	LogTypeELK        = "elk"
	LogTypePrometheus = "prometheus"
)



// DefaultWriter is the default io.Writer used by Gin for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		gin.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by Gin to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var ginMode = modeDebugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		ginMode = modeDebugCode
	case ReleaseMode:
		ginMode = modeReleaseCode
	case TestMode:
		ginMode = modeTestCode
	case ProductMode:
		ginMode = modeProductCode
	default:
		panic("gin mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

// DisableBindValidation closes the default validator.
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode returns currently gin mode.
func Mode() string {
	return modeName
}
