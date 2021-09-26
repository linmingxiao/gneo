package rescue

import "github.com/linmingxiao/gneo/logx"

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logx.Error(p)
	}
}
