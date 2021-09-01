package gneo

import (
	"github.com/linmingxiao/gneo/logx"
	"github.com/linmingxiao/gneo/skill/stat"
)

// 全局的初始化日志
func InitLogger(cfg *logx.LogConfig) {
	logx.MustSetup(*cfg)

	// log初始化完毕，接下来解析
	if cfg.NeedCpuMem {
		stat.StartCpuMemCollect()
	}
}
