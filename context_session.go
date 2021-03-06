package gneo

import (
	"github.com/linmingxiao/gneo/connx/redis"
	"github.com/linmingxiao/gneo/jwtx"
	"github.com/linmingxiao/gneo/logx"
)




type CtxSessionFunc interface {
	InitSession() error //初始化Session
}

func (c *Context)InitSession(ctxSessConf *jwtx.CtxSessionConfig) error {
	var token string = ""
	var tokIsNew bool = false
	var sid string = ""
	if tokQ, ok := c.GetQuery("tok"); ok{
		logx.Debug("Get query token.")
		token = tokQ
	} else if tokP, ok := c.GetPostForm("tok"); ok{
		logx.Debug("Get form token.")
		token = tokP
	} else if c.Sess != nil && len(c.Sess.Token) > 10{
		logx.Debug("Get session token.")
		c.Sess.TokIsNew = false;
		token = c.Sess.Token;
	} else {
		logx.Debug("Request has no token, need to new one.")
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
		logx.Debug(">>>>>>>>>>>>>>>>>>>>>>>>>>>New Session>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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
		logx.Debugf("Session has one token: %s", token)
	}
	if c.Sess.Redis == nil {
		logx.Debug("Init session redis...")
		c.Sess.Redis = redis.GetSingletonRedis(&c.Sess.RedisConnCnf)
	}
	if tokIsNew{
		c.Sess.Values = KV{
			"sid": sid,
		}
		logx.Debug("Create a new session and save.")
		_, err := c.Sess.SaveToRedis()
		if err != nil {
			return err
		}
	} else {
		logx.Debug("Load session from redis.")
		_, err := c.Sess.LoadFromRedis(sid)
		if err != nil{
			return err
		}
	}
	c.Sess.IsReady = true
	return nil
}