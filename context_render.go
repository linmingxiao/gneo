package gneo

import (
	"net/http"
	"github.com/linmingxiao/gneo/render"
)

//数据返回基本样式
func NewRenderKV(status, msg string, code int) KV {
	return KV{
		"status": status,
		"code":   code,
		"msg":    msg,
	}
}

func (c *Context) FaiErr(err error) {
	if errX, ok := err.(*render.ErrorX); ok{
		c.Fai(errX.ErrCode, errX.ErrMsg, nil)
	} else {
		c.Fai(-1, err.Error(), nil)
	}
}

func (c *Context) FaiMsg(msg string) {
	c.Fai(-1, msg, nil)
}

func (c *Context) FaiKV(obj KV) {
	c.Fai(-1, "", obj)
}

func (c *Context) SucMsg(msg string) {
	c.Suc(0, msg, nil)
}

func (c *Context) SucKV(obj KV) {
	c.Suc(0, "", obj)
}

//按照错误码返回错误
func (c *Context) FaiCode(errCode int) {
	msg, ok:= render.MapErrorCode[errCode]
	if !ok{
		c.FaiCode(-1)
	} else {
		c.Fai(errCode, msg, nil)
	}
}

func (c *Context) Fai(code int, msg string, obj interface{}) {
	jsonData := NewRenderKV("fai", msg, code)
	if obj != nil {
		jsonData["data"] = obj
	}
	c.faiKV(jsonData)
}
func (c *Context) Suc(code int, msg string, obj interface{}) {
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
	if c.Sess != nil {
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
		jsonData["code"] = -1
	}
	if c.Sess != nil && c.Sess.TokIsNew {
		jsonData["tok"] = c.Sess.Token
	}
	c.JSON(http.StatusOK, jsonData)
}