// Copyright 2020 GoFast Author(http://chende.ren). All rights reserved.
// Use of this source code is governed by a MIT license
package logx

var currConfig *LogConfig

type LogConfig struct {
	ServiceName string `json:",optional"`
	Mode        string `json:",default=console,options=console|file|volume"`
	Level       int8   `json:",default=0"`    // 记录日志的级别
	Path        string `json:",default=logs"` // 日志文件路径定义到文件夹
	FilePrefix  string `json:",default=gneo"` // 日志文件名统一前缀
	Compress    bool   `json:",optional"`     // 是否压缩
	KeepDays    int    `json:",optional"`
	NeedCpuMem  bool   `json:",default=true"`
	style       int8   `inner:",default=0"` // 日志模板样式
}

const (
	modeConsole string = "console" //只输出控制台
	modeFile string = "file" //只输出文件
	modeVolume string = "volume" //同时输出控制台和文件
)

//日志级别
const (
	//*********info级别日志*******
	m_trace int8 = iota  //详细日志 最全类型 0
	m_debug //调试模式日志 1
	m_info //信息级别日志 通常都为有效信息 2

	//**********错误级别日志**********
	m_warn //警告级别日志 需注意 3
	m_error //错误级别日志 需处理 4
	m_fatal //灾难级错误 会导致系统down掉 5
)

// 日志样式类型
const (
	styleJson int8 = iota
	styleJsonMini
	styleSdx
	styleSdxMini
)
