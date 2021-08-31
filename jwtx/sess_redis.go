package jwtx

import (
	"github.com/linmingxiao/gneo/connx/redis"
	"github.com/linmingxiao/gneo/internal/bytesconv"
	"github.com/linmingxiao/gneo/internal/json"
	"time"
)

type CtxSessionConfig struct {
	RedisConnCnf redis.ConnConfig `json:",optional"`                                    // 用 Redis 做持久化
	CheckTokenIP bool             `json:",optional,default=true"`                       // 看是否检查 token ip 地址
	AuthField    string           `json:",optional,default=user_id"`                    // 标记当前登录用户字段是 user_id
	Secret       string           `json:",optional"`                                    // token秘钥
	TTL          int32            `json:",optional,default=14400,range=[0:2000000000]"` // session有效期 默认 3600*4 秒
	TTLNew       int32            `json:",optional,default=180,range=[0:2000000000]"`   // 首次产生的session有效期 默认 60*3 秒
}


type CtxRedisSessionFunc interface {
	LoadFromRedis(string) (bool, error) //从 Redis 中 Load Session
	SaveToRedis() (string, error) //将 Session 保存到 Redis
	SetRedisExpire() (bool, error) //重新设定 Redis Session过期时间
	DestroyRedis() (int64, error)//销毁 Redis Session
}

type CtxSession struct {
	Session
	CtxSessionConfig
	Redis    *redis.GoRedisX
	IsReady  bool
	TokIsNew bool
}

func (ctxSess *CtxSession)LoadFromRedis(sid interface{})(bool, error) {
	sessKey := SessKeyPrefix + ctxSess.Sid
	if sid != nil {
		sessKey = SessKeyPrefix + sid.(string)
	}
	str, err := ctxSess.Redis.Get(sessKey)
	if str == "" || err != nil {
		str = `{}`
	}
	err = json.Unmarshal(bytesconv.StringToBytes(str), ctxSess.Values)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ctxSess *CtxSession)SaveToRedis()(string, error)  {
	str, _ := json.Marshal(ctxSess.Values)
	ttl := ctxSess.TTL
	if ctxSess.TokIsNew && ctxSess.Values[ctxSess.AuthField] == nil {
		ttl = ctxSess.TTLNew
	}
	return ctxSess.Redis.Set(SessKeyPrefix+ctxSess.Sid, str, time.Duration(ttl)*time.Second)
}

func (ctxSess *CtxSession)SetRedisExpire(ttl int32)(bool, error){
	if ttl < 0{
		if ctxSess.TokIsNew {
			ttl = ctxSess.TTLNew
		} else {
			ttl = ctxSess.TTL
		}
	}
	return ctxSess.Redis.Expire(SessKeyPrefix + ctxSess.Sid, time.Duration(ttl)*time.Second)
}

func (ctxSess *CtxSession)DestroyRedis() (int64, error) {
	return ctxSess.Redis.Del(SessKeyPrefix + ctxSess.Sid)
}
