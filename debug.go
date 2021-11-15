package gneo

import (
	"github.com/linmingxiao/gneo/logx"
	"runtime"
	"strconv"
	"strings"
)

const ginSupportMinGoVer = 13

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(gin.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return ginMode == modeDebugCode
}

// DebugPrintRouteFunc indicates debug log output format.
var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {

	if logx.IsDebug() {
		nuHandlers := len(handlers)
		handlerName := nameOfFunction(handlers.Last())
		if DebugPrintRouteFunc == nil {
			logx.Debugf("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
		} else {
			DebugPrintRouteFunc(httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
}

func debugPrint(format string, values ...interface{}) {
	if logx.IsDebug() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		logx.Debugf(format, values...)
	}
}

func getMinVer(v string) (uint64, error) {
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	if first == last {
		return strconv.ParseUint(v[first+1:], 10, 64)
	}
	return strconv.ParseUint(v[first+1:last], 10, 64)
}

func debugPrintWARNINGDefault() {
	if v, e := getMinVer(runtime.Version()); e == nil && v <= ginSupportMinGoVer {
		logx.Warnf(`Now GNEO requires Go 1.13+.`)
	}
	logx.Warnf(`Creating an Engine instance with the Logger and Recovery middleware already attached.`)
}

func debugPrintWARNINGNew() {
	logx.Warnf(`Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GNEO_MODE=release
 - using code:	gneo.SetMode(gneo.ReleaseMode)
`)
}

func debugPrintError(err error) {
	if err != nil && logx.IsDebug() {
		logx.Debugf("[ERROR] %v\n", err)
	}
}
