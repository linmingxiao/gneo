package jwtx

var (
	TokenPrefix   = "t:"
	SessKeyPrefix = "tls:"
)

type BaseSessionFunc interface {
	Get(string) interface{}
	Set(string, interface{})
	Del(string)
	Save()
	Expire(int32)
}


type Session struct {
	Sid string //Session ID
	Token string
	Values map[string] interface{} //Session 内容体
	Saved bool //判断当前 Session 是否已保存
}

func (ss *Session)Get(key string) interface{} {
	if ss.Values == nil{
		return nil
	}
	return ss.Values[key]
}

func (ss *Session)Set(key string, val interface{}){
	ss.Values[key] = val
	ss.Saved = false
}

func (ss *Session)SetKV(kv map[string]interface{}){
	for k, v:= range kv{
		ss.Values[k] = v
	}
	ss.Saved = false
}







