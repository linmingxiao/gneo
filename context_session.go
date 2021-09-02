package gneo

import (
	"github.com/linmingxiao/gneo/connx/redis"
	"github.com/linmingxiao/gneo/jwtx"
	"github.com/linmingxiao/gneo/logx"
)




type CtxSessionFunc interface {
	InitSession() error //初始化Session
	InitToken() string  //初始化 Token
}

func (c *Context)InitSession(ctxSessConf *jwtx.CtxSessionConfig) error {
	var token string = ""
	var tokIsNew bool = false
	var sid string = ""
	if tokQ, ok := c.GetQuery("tok"); ok{
		token = tokQ
	} else if tokP, ok := c.GetPostForm("tok"); ok{
		token = tokP
	} else {
		sid, token = jwtx.GenToken(ctxSessConf.Secret)
		tokIsNew = true
	}
	if len(token) < 10 {
		panic("初始化用户Session失败, token无效")
	}
	if !tokIsNew{
		sid, _ = jwtx.FetchSid(token)
	}
	if c.Sess == nil {
		c.Sess = &jwtx.CtxSession{
			Session: jwtx.Session{
				Sid:   sid,
				Token: token,
				Saved: false,
				Values: make(KV),
			},
			CtxSessionConfig: *ctxSessConf,
			TokIsNew: tokIsNew,
		}
	}
	if c.Sess.Redis == nil {
		logx.DebugPrint("First init session redis...")
		c.Sess.Redis = redis.GetSingletonRedis(&c.Sess.RedisConnCnf)
	}
	if tokIsNew{
		c.Sess.Values = KV{
			"Sid": sid,
		}
		logx.DebugPrint("Create a new session and save.")
		_, err := c.Sess.SaveToRedis()
		if err != nil {
			return err
		}
	} else {
		logx.DebugPrint("Load session from redis.")
		_, err := c.Sess.LoadFromRedis(sid)
		if err != nil{
			return err
		}
	}
	c.Sess.IsReady = true
	return nil
}