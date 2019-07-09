package lib

const (
	Success       string = "0000" // 成功
	SystemError   string = "5000" // 系统错误
	MissParam     string = "5001" // 参数缺失
	InvalidParam  string = "5002" // 参数错误
	ApiNotFound   string = "5003" // 接口未找到
	SignError     string = "5004" // 签名错误
	RequestExpire string = "5005" // 请求已失效
	SystemBusy    string = "5006" // 系统忙

	MustLogin    string = "12306" //未登录
	SessionEmpty string = "10086" //微信未授权

	// 错误代码从4开始
	UploadFailed    string = "4001" // 上传失败
	DeleteFailed    string = "4002" // 删除失败
	NotFound        string = "4004" //不存在
	DataExist       string = "4005" //已存在
	BusinessError   string = "4006" //业务错误
	SessionKeyError string = "6001" //获取session_key失败

	GenerateCodeError       string = "7001" //邀请码生成失败
	GenerateAuthError       string = "7002" //权限不足
	CodeError               string = "7003" //邀请码错误
	PasswordOrUsernameError string = "7004" //用户名或密码错误
	MainManagerCanBeOnlyOne string = "7005" //用户名或密码错误
	CanNotLoginByUsername   string = "7006" //请联系主管理员扫码登录

	CannotModifyPassword string = "8001" //只有主管理员才能修改密码
	CannotDelUser        string = "8002"
	UserNotExist         string = "8003"
	NotSuperManager      string = "8004" //
	CannotModifyInfo     string = "8005" //只有主管理员才能修改
	ForbiddenDelUser     string = "8006"
	DecryptPhoneError    string = "8007" //手机解密失败

	ActAssistanceEnough         string = "9001"
)

// 获取错误代码对应的提示信息
func Msg(code string) string {
	msgMap := map[string]string{
		Success:       "success",
		SystemError:   "系统错误",
		MissParam:     "参数缺失",
		InvalidParam:  "参数错误",
		ApiNotFound:   "接口未找到",
		SignError:     "签名错误",
		RequestExpire: "请求已失效",
		SystemBusy:    "系统忙",

		MustLogin:       "未登录",
		SessionEmpty:    "微信未授权",
		UploadFailed:    "上传失败",
		DeleteFailed:    "删除失败",
		NotFound:        "不存在",
		DataExist:       "已存在",
		SessionKeyError: "获取session_key失败",

		GenerateCodeError:       "邀请码生成失败",
		GenerateAuthError:       "没权限生成邀请码",
		CodeError:               "邀请码错误",
		PasswordOrUsernameError: "用户名或密码错误",
		MainManagerCanBeOnlyOne: "主管理员已用另一个微信账号登录",
		CanNotLoginByUsername:   "请联系主管理员扫码登录",

		CannotModifyPassword: "只有主管理员才能修改密码",
		CannotDelUser:        "主管理员才有删除权限",
		UserNotExist:         "用户不存在",
		NotSuperManager:      "你没权限进行此操作",
		CannotModifyInfo:     "只有主管理员才能修改",
		ForbiddenDelUser:     "不能删除主管理员",
		DecryptPhoneError:    "手机解密失败",

		ActAssistanceEnough:  "您今日助力次数已达上限",
	}

	msg, exist := msgMap[code]

	if exist == false { // 不存在默认返回空字符串
		msg = ""
	}

	return msg
}
