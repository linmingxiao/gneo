package jwtx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/linmingxiao/gneo/logx"
	"github.com/linmingxiao/gneo/skill/lang"
	"regexp"
	"strings"
)

// tok=t:NFRRcE81WDFQSEZJQUptZkpJ.v9EN6bWz8KU6sKRrcEId1OKUKqYx0hed2zSpCQImvc
func FetchSid(tok string) (string, string) {
	start := strings.Index(tok, TokenPrefix)
	dot := strings.Index(tok, ".")
	// 格式明显不对，直接返回空
	if start != 0 || dot <= 0 {
		// return "", "", errors.New("Can't parse sid. ")
		return "", ""
	}
	sid := tok[2:dot]
	if len(sid) <= 18 {
		// return "", "", errors.New("Sid length error. ")
		return "", ""
	}
	sHmac := tok[(dot + 1):]
	return sid, sHmac
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++==
// 闪电侠SID：为24位的字符串
func GenToken(secret string) (string, string) {
	sid := genSid(24)
	logx.Infof("New sid: %s", sid)
	tok := TokenPrefix + genSign(sid, secret)
	logx.Infof("New token: %s", tok)
	return sid, tok
}

// 按照指定长度length, 自动生成随机的Sid字符串，
func genSid(length int) string {
	src := lang.GetRandomBytes(length)
	sid := base64.StdEncoding.EncodeToString(src)
	sid = cleanString(sid)

	if length > len(sid) {
		length = len(sid)
	}
	return sid[:length]
}

func genSign(val, secret string) string {
	signSHA256 := genSignSHA256([]byte(val), []byte(secret))
	return val + "." + cleanString(signSHA256)
}

func genSignSHA256(data, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)

	// toBase64
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func cleanString(src string) string {
	regExp := regexp.MustCompile("[+=]*")
	//regExp := regexp.MustCompile("[+=/]*")
	return regExp.ReplaceAllString(src, "")
}
