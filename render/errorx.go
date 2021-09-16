package render

import (
	_ "reflect"
	"github.com/linmingxiao/gneo/logx"
)

var MapErrorCode map[int]string = map[int]string{
	0: "请求成功",

	-100: "鼎信汇金公募基金系统维护升级中，暂时无法访问，为此给您带来的不便深表歉意。",
	-1:   "未知错误",

	120:  "系统错误请联系客服",

	110:  "您还没有登录，请登录系统。",
	111:  "登录失效，请您退出后重新登录系统",
	112:  "当前账户还未注册，请先注册",
	113:  "账户名或密码错误",

	422:  "请求缺少参数或参数错误",
	429:  "未根据请求参数获得相应的返回结果",
	425:  "系统繁忙，请稍候再试。",
	501:  "服务器维护升级中，请稍后再试",

	//**************************公募基金相关错误
	1500: "您未开通基金账号",
	1501: "您已经开通基金账号",

	1600: "您还没有进行风险测评，请先进行风险测评",
	1601: "您的风险评估已经过期，请您重新评估。",
	1602: "根据反洗钱相关法规要求，自2020年5月8日起，我司对身份信息资料不完善的客户，将限制或中止办理相关业务。为了不影响您的基金交易，请您及时上传本人有效身份证件照片，完善身份信息资料。您所上传的证件照片仅用于投资者身份验证，我司会保证您的隐私安全。",
	1603: "您的账户还未进行实名认证，请上传本人有效身份证件照片，完善身份信息资料。您所上传的证件照片仅用于投资者身份验证，我司会保证您的隐私安全。",
	1604: "根据《证券期货投资者适当性管理办法》规定，请您重新进行风险评估",

	1700: "您还未完善反洗钱相关信息，请先进行完善",

	3000: "未找到此定投协议",
	3001: "此银行卡未持有此基金份额",
	3002: "请先终止定投协议",

	4001: "原交易密码不能为空",
	4002: "新交易密码不能为空",
	4003: "两次密码不一致",

	5001: "请选择银行",
	5002: "请填写银行账号",
	5003: "请选择银行分行信息",
	5004: "未知操作",
	5005: "未查询到银行卡",
	5006: "对不起,您未开通换卡的功能",
	5007: "对不起,您未开通签约的功能",
	5008: "尊敬的客户：您填写的银行信息与申请更换的银行信息不一致，请重新填写。",
	5009: "您输入的银行卡与绑定的银行卡不一致，请重新填写。",
	5010: "您未签约这张银行卡,请重新输入银行卡！",
	5011: "您已经签约汇付天下或者通联快捷相同银行的银行卡,不能再签约。",
	5012: "操作超时，请您重新进行操作",
	5013: "请输入银行预留手机号码",
	5014: "您还未添加银行卡",

	6001: "身份证号或邮政编码格式不正确,请您重新输入",
	6002: "该身份证号码已开户，不能重复开户",
	6003: "身份证号码冲突，请联系客服确认",
	6004: "两次输入的交易密码不一致",

	7001: "可用基金份额少于转换份额",

	8001: "此申请不存在，请联系客服确认",

	9001: "姓名、身份证或者验证码输入错误",
	9002: "验证码已过期,请重新发送",
	9003: "重置交易密码,请先进行银行验证",
	9004: "重置交易密码,请先进行手机验证",

	10000: "交易密码错误,请重新输入",
	10001: "交易密码锁定,请重置交易密码",
}

type ErrorX struct {
	ErrCode int
	ErrMsg string
}

func (err *ErrorX) Error() string {
	if msg, ok:= MapErrorCode[err.ErrCode]; err.ErrCode != -1 && ok{
		return msg
	} else if err.ErrMsg != ""{
		return err.ErrMsg
	} else {
		return "未知错误"
	}
}

func NewErrorX(pms ...interface{}) *ErrorX{
	if len(pms) == 2{
		return &ErrorX{
			ErrCode: pms[0].(int),
			ErrMsg: pms[1].(string),
		}
	} else if len(pms) == 1{
		pms0 := pms[0]
		switch pms0.(type) {
		case int:
			errCode := pms0.(int)
			if errMsg, ok := MapErrorCode[errCode]; ok{
				return &ErrorX{
					ErrCode: errCode,
					ErrMsg: errMsg,
				}
			} else {
				return &ErrorX{
					ErrCode: errCode,
					ErrMsg: "未知错误",
				}
			}
		case string:
			errMsg := pms0.(string)
			return &ErrorX{
				ErrCode: -1,
				ErrMsg: errMsg,
			}
		default:
			logx.Error("NewErrorX参数错误")
			panic("NewErrorX参数错误")
		}
		//if reflect.TypeOf(pms[0]).Kind() == reflect.Int{
		//	errCode := pms[0].(int)
		//	if errMsg, ok := MapErrorCode[errCode]; ok{
		//		return &ErrorX{
		//			ErrCode: errCode,
		//			ErrMsg: errMsg,
		//		}
		//	} else {
		//		return &ErrorX{
		//			ErrCode: errCode,
		//			ErrMsg: "未知错误",
		//		}
		//	}
		//} else if reflect.TypeOf(pms[0]).Kind() == reflect.String{
		//	errMsg := pms[0].(string)
		//	return &ErrorX{
		//		ErrCode: -1,
		//		ErrMsg: errMsg,
		//	}
		//} else {
		//	logx.Error("NewErrorX参数错误")
		//	panic("NewErrorX参数错误")
		//}
	} else {
		logx.Error("NewErrorX参数错误")
		panic("NewErrorX参数错误")
	}
}