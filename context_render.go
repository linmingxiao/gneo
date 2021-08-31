package gneo

import (
	"net/http"
)

//数据返回基本样式
func NewRenderKV(status, msg string, code int32) KV {
	return KV{
		"status": status,
		"code":   code,
		"msg":    msg,
	}
}


func (c *Context) FaiErr(err error) {
	c.Fai(0, err.Error(), nil)
}

func (c *Context) FaiMsg(msg string) {
	c.Fai(0, msg, nil)
}

func (c *Context) FaiKV(obj KV) {
	c.Fai(0, "", obj)
}

func (c *Context) SucMsg(msg string) {
	c.Suc(0, msg, nil)
}

func (c *Context) SucKV(obj KV) {
	c.Suc(0, "", obj)
}



func (c *Context) Fai(code int32, msg string, obj interface{}) {
	jsonData := NewRenderKV("fai", msg, code)
	if obj != nil {
		jsonData["data"] = obj
	}
	c.faiKV(jsonData)
}
func (c *Context) Suc(code int32, msg string, obj interface{}) {
	jsonData := NewRenderKV("suc", msg, code)
	if obj != nil {
		jsonData["data"] = obj
	}
	c.sucKV(jsonData)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (c *Context) sucKV(jsonData KV) {
	if jsonData == nil {
		jsonData = make(KV)
	}
	jsonData["status"] = "suc"
	if jsonData["msg"] == nil {
		jsonData["msg"] = ""
	}
	if jsonData["code"] == nil {
		jsonData["code"] = 0
	}
	if c.Sess != nil && c.Sess.TokIsNew {
		jsonData["tok"] = c.Sess.Token
	}
	c.JSON(http.StatusOK, jsonData)
}

func (c *Context) faiKV(jsonData KV) {
	if jsonData == nil {
		jsonData = make(KV)
	}
	jsonData["status"] = "fai"
	if jsonData["msg"] == nil {
		jsonData["msg"] = ""
	}
	if jsonData["code"] == nil {
		jsonData["code"] = 0
	}
	if c.Sess != nil && c.Sess.TokIsNew {
		jsonData["tok"] = c.Sess.Token
	}
	c.JSON(http.StatusOK, jsonData)
}