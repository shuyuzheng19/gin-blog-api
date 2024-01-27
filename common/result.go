package common

// R 全局返回
// @Description 全局返回
type R struct {
	Code    int         `json:"code"`           //状态码
	Message string      `json:"message"`        //返回的消息
	Data    interface{} `json:"data,omitempty"` //返回的数据
}

// 成功状态码
const okStatusCode = 200

// 错误状态码
const errorStatusCode = 500

// 默认返回状态码
const defaultErrorMessage = "服务器无法处理请求"

type ErrorCode int

const (
	BadRequestCode      ErrorCode = 400          //参数验证错误
	Unauthorized        ErrorCode = 401          //无效身份
	Forbidden           ErrorCode = 403          //没权限
	ServerError         ErrorCode = 500          //服务器错误
	FailCode            ErrorCode = iota + 10001 //失败
	FieldValidationFail                          //字段验证失败
	RegisteredCode                               //注册失败
	SendEmailFailCode                            //发送邮件失败
	EmailCodeValidate                            //验证邮箱验证码失败
	CreateTokenFail                              //生成token失败
	ParseTokenFail                               //解析token失败
	TokenExpireFail                              //Token失效
	LoginFail                                    //登录失败
	NoLogin                                      //还未登录
	AddFileFail                                  //上传文件失败
	NoFile                                       //请求体里文件为空
	SaveBlogFail                                 //保存博客失败
	NotFoundBlog                                 //找不到博客
	UpdateBlogFail                               //修改博客失败
	GetUpdateFail                                //获取修改博客失败
	SaveFail                                     //添加失败
	UpdateFail                                   //修改失败
	NotFoundDir                                  //找不到文件就
	ClearFileFail                                //清空文件失败
	ConvertFail                                  //转换失败
	NotJson                                      //不是json格式
	PathEmptyFail                                //路径为空
	OpenFileFail                                 //打开文件失败
)

var messageMap = map[ErrorCode]string{
	FailCode:            "处理失败",
	BadRequestCode:      "参数验证失败",
	RegisteredCode:      "注册失败",
	SendEmailFailCode:   "发送邮件失败",
	EmailCodeValidate:   "邮件验证码不正确",
	CreateTokenFail:     "创建Token失败",
	ParseTokenFail:      "解析Token失败",
	LoginFail:           "账号或密码错误",
	NoLogin:             "你还未登录",
	Forbidden:           "您没有权限处理",
	AddFileFail:         "添加文件失败",
	NoFile:              "没有文件,请选择文件",
	SaveBlogFail:        "添加博客失败",
	NotFoundBlog:        "该博客不存在",
	UpdateBlogFail:      "修改博客失败",
	GetUpdateFail:       "获取修改博客失败",
	SaveFail:            "添加失败",
	UpdateFail:          "修改失败",
	NotFoundDir:         "找不到该文件夹",
	ClearFileFail:       "清空文件失败",
	ConvertFail:         "转换失败",
	NotJson:             "这不是个json格式的字符串",
	TokenExpireFail:     "Token已失效",
	Unauthorized:        "无效身份信息",
	PathEmptyFail:       "请填入路径",
	OpenFileFail:        "处理文件失败",
	FieldValidationFail: "字段验证未通过",
	ServerError:         "服务器无法正常处理请求",
}

// Success 返回成功 data 返回的数据
func Success(data interface{}) R {
	return R{Code: okStatusCode, Message: "success", Data: data}
}

// Error 服务器错误
func Error() R {
	return R{Code: errorStatusCode, Message: "error"}
}

// Fail 返回失败 code 失败状态码 message 失败消息
func Fail(code ErrorCode, message string) R {
	return R{Code: int(code), Message: message}
}

// AutoFail 自动返回对应的错误 code 对应上方的状态码常量
func AutoFail(code ErrorCode) R {
	return R{Code: int(code), Message: messageMap[code]}
}
