package gneo

import "github.com/qinchende/gofast/logx"

type APPConfig struct {
	LogConfig              logx.LogConfig
	Name                   string `cnf:",NA,def=GoFastSite"`
	Addr                   string `cnf:",def=0.0.0.0:8099"`
	RunMode                string `cnf:",def=debug,enum=debug|test|product"` // 当前模式[debug|test|product]
	SecureJsonPrefix       string `cnf:",NA,def=while(1);"`
	MaxMultipartMemory     int64  `cnf:",def=33554432"` // 最大上传文件的大小，默认32MB
	SecondsBeforeShutdown  int64  `cnf:",def=1000"`     // 退出server之前等待的毫秒，等待清理释放资源
	RedirectTrailingSlash  bool   `cnf:",def=false"`    // 探测url后面加减'/'之后是否能匹配路由（这个时代默认不需要了）
	HandleMethodNotAllowed bool   `cnf:",def=false"`
	DisableDefNotAllowed   bool   `cnf:",def=false"`
	DisableDefNoRoute      bool   `cnf:",def=false"`
	ForwardedByClientIP    bool   `cnf:",def=true"`
	RemoveExtraSlash       bool   `cnf:",def=false"`                       // 规范请求的URL
	UseRawPath             bool   `cnf:",def=false"`                       // 默认取原始的Path，不需要自动转义
	UnescapePathValues     bool   `cnf:",def=true"`                        // 默认把URL中的参数值做转义
	PrintRouteTrees        bool   `cnf:",def=false"`                       // 是否打印出当前路由数
	FitReqTimeout          int64  `cnf:",def=3000"`                        // 每次请求的超时时间（单位：毫秒）
	FitMaxReqContentLen    int64  `cnf:",def=33554432"`                    // 最大请求字节数
	FitMaxReqCount         int32  `cnf:",def=1000000,range=[0:100000000]"` // 最大请求处理数
	FitJwtSecret           string `cnf:",NA"`                              // JWT认证的秘钥
	FitLogType             string `cnf:",def=json,enum=json|sdx"`
	modeType               int8   `cnf:",NA"` // 内部记录状态
}

