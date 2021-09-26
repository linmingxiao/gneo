package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/linmingxiao/gneo/logx"
)

// go-redis
type (
	ConnConfig struct {


		// sentinel
		SentinelAddr []string `json:",optional"`
		MasterName   string   `json:",optional"`
		SentinelPass string   `json:",optional"`
		SlaveOnly    bool     `json:",optional"`

		// single redis
		Addr string `json:",optional"`

		// common
		Pass     string `json:",optional"`
		DB       int    `json:",optional"`
		PoolSize int    `json:",optional"`
		MinIdle  int    `json:",optional"`
	}
	GoRedisX struct {
		Cli *redis.Client
		Ctx context.Context
	}
)

var GlobalRedisMap map[string] *GoRedisX = make(map[string] *GoRedisX)

func GetSingletonRedis(cf *ConnConfig) *GoRedisX {
	masterName := cf.MasterName
	if GlobalRedisMap[masterName] != nil{
		return GlobalRedisMap[masterName]
	}
	r := NewGoRedis(cf)
	GlobalRedisMap[masterName] = r
	return r
}

// 直接连接redis
// go-redis 底层自带连接池功能，不需要你再管理了。
// 看看 go-redis/redis.go 中的代码：
// 第330行 func (c *baseClient) process(ctx context.Context, cmd Cmder)
// 第288行 func (c *baseClient) withConn(
// 第292行 cn, err := c.getConn(ctx)
func NewGoRedis(cf *ConnConfig) *GoRedisX {
	rds := GoRedisX{Ctx: context.Background()}

	if cf.Addr != "" {
		rds.Cli = redis.NewClient(&redis.Options{
			Addr:         cf.Addr,
			Password:     cf.Pass,
			DB:           cf.DB,
			PoolSize:     cf.PoolSize,
			MinIdleConns: cf.MinIdle,
			OnConnect: func(ctx context.Context, cn *redis.Conn) error {
				logx.Infof("%s——%s connected.", cf.MasterName, cn.String())
				return nil
			},
		})
		logx.Infof("Redis %s —— %s created.", cf.MasterName, cf.Addr)
	} else if cf.SentinelAddr != nil {
		// 通过sentinel连接 redis
		rds.Cli = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:       cf.MasterName,
			SentinelAddrs:    cf.SentinelAddr,
			SentinelPassword: cf.SentinelPass,
			Password:         cf.Pass,
			SlaveOnly:        cf.SlaveOnly,
			DB:               cf.DB,
			PoolSize:         cf.PoolSize,
			MinIdleConns:     cf.MinIdle,
			OnConnect: func(ctx context.Context, cn *redis.Conn) error {
				logx.Info("%s —— %s connected.", cf.MasterName, cn.String())
				return nil
			},
		})
		logx.Infof("Redis %s created.", cf.MasterName)
	}

	return &rds
}

//// 通过sentinel连接 redis
//func NewGoRedisBySentinel(cf *ConnConfig) *GoRedisX {
//	rds := GoRedisX{Ctx: context.Background()}
//	rds.Cli = redis.NewFailoverClient(&redis.FailoverOptions{
//		MasterName:       cf.MasterName,
//		SentinelAddrs:    cf.SentinelAddr,
//		SentinelPassword: cf.SentinelPass,
//		Password:         cf.Pass,
//		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
//			logx.Info(fmt.Sprintf("%s connected.", cn.String()))
//			return nil
//		},
//	})
//	logx.Info(fmt.Sprintf("Redis %s created.", cf.Addr))
//	return &rds
//}
