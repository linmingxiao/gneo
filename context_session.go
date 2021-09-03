package gneo

import (
	"fmt"
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
		logx.Info("Get query token.")
		token = tokQ
	} else if tokP, ok := c.GetPostForm("tok"); ok{
		logx.Info("Get form token.")
		token = tokP
	} else if c.Sess != nil && len(c.Sess.Token) > 10{
		logx.Info("Get session token.")
		c.Sess.TokIsNew = false;
		token = c.Sess.Token;
	} else {
		logx.Info("Request has no token, need to new one.")
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
		logx.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>>New Session>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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
	} else {
		fmt.Print(c.Sess)
		logx.Info(fmt.Printf("Session has one token: %s", token))
	}
	if c.Sess.Redis == nil {
		logx.Info("Init session redis...")
		c.Sess.Redis = redis.GetSingletonRedis(&c.Sess.RedisConnCnf)
	}
	if tokIsNew{
		c.Sess.Values = KV{
			"Sid": sid,
		}
		logx.Info("Create a new session and save.")
		_, err := c.Sess.SaveToRedis()
		if err != nil {
			return err
		}
	} else {
		logx.Info("Load session from redis.")
		_, err := c.Sess.LoadFromRedis(sid)
		if err != nil{
			return err
		}
	}
	c.Sess.IsReady = true
	return nil
}