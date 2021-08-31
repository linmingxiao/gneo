package gneo

import (
	"github.com/linmingxiao/gneo/connx/redis"
	"github.com/linmingxiao/gneo/jwtx"
)

var (
	GSess *jwtx.CtxSession
)


type CtxSessionFunc interface {
	InitSession() error //初始化Session
	InitToken() string  //初始化 Token
}

func (c *Context)InitSession(ctxSessConf *jwtx.CtxSessionConfig) {
	if GSess != nil {
		return
	}
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
				Token: c.Sess.Token,
				Saved: false,
			},
			CtxSessionConfig: *ctxSessConf,
			TokIsNew: tokIsNew,
		}
	}
	if c.Sess.Redis == nil {
		c.Sess.Redis = redis.NewGoRedis(&c.Sess.RedisConnCnf)
	}
	if tokIsNew{
		_, err := c.Sess.SaveToRedis()
		if err != nil {
			panic(err)
		}
	} else {
		_, err := c.Sess.LoadFromRedis(sid)
		if err != nil{
			panic(err)
		}
	}
	GSess = c.Sess
	c.Sess.IsReady = true
}