package errcode

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ErrCode int

const (
	Success ErrCode = 0
	ErrNone ErrCode = Success

	errBase         ErrCode = 10000
	ErrExecFailed   ErrCode = errBase + 1
	ErrInvalidParam ErrCode = errBase + 2
	ErrInvalidToken ErrCode = errBase + 3
	ErrServerError  ErrCode = errBase + 4
	ErrReportError  ErrCode = errBase + 5
	ErrPassWdError  ErrCode = errBase + 6

	errAuthBase           ErrCode = 10100
	ErrAuthFailed         ErrCode = errAuthBase + 1
	ErrAuthInvalidParam   ErrCode = errAuthBase + 2
	ErrAuthPasswdNotMatch ErrCode = errAuthBase + 3
	ErrAuthNameNotExist   ErrCode = errAuthBase + 4
	ErrAuthInvalidStatus  ErrCode = errAuthBase + 5
	ErrAuthDB             ErrCode = errAuthBase + 6
	ErrUnauthorized               = errAuthBase + 7
	ErrPasswdStrength             = errAuthBase + 8
	ErrPasswdLength               = errAuthBase + 9
	ErrUserNameExist              = errAuthBase + 10
	ErrPasswdDec                  = errAuthBase + 11
	ErrLoginLocked                = errAuthBase + 12
	ErrSign                       = errAuthBase + 13

	errSimulatorBase         ErrCode = 10200
	ErrSimulatorCommon       ErrCode = errSimulatorBase + 1
	ErrSimulatorComputingErr ErrCode = errSimulatorBase + 2
	ErrSimulatorCancelErr    ErrCode = errSimulatorBase + 3

	// 第三方报告
	errTReportBase            ErrCode = 10300
	ErrTReportCommon          ErrCode = errTReportBase + 1
	ErrTReportUploadErr       ErrCode = errTReportBase + 2
	ErrTReportUploadLimitErr  ErrCode = errTReportBase + 3
	ErrTReportUploadFormatErr ErrCode = errTReportBase + 4

	// 企业相关
	errCompanyBase          ErrCode = 10400
	ErrCompanyNotMaintained         = errCompanyBase + 1
	ErrCompanyNameRequired          = errCompanyBase + 2
	ErrProjectNameRequired          = errCompanyBase + 3
	ErrDateRequired                 = errCompanyBase + 4
	ErrEntNoFound                   = errCompanyBase + 5
	ErrEntBizGet                    = errCompanyBase + 6
)

var errCodeMap = map[ErrCode]string{
	ErrNone:         "Success",
	ErrExecFailed:   "操作失败",
	ErrInvalidParam: "参数错误",
	ErrInvalidToken: "认证失败",
	ErrServerError:  "服务器异常",
	ErrReportError:  "不能生成报告",
	ErrPassWdError:  "密码长度不够12位或者密码强度不够",

	// Auth
	ErrAuthFailed:         "请先登录",
	ErrAuthInvalidParam:   "用户名或者密码为空",
	ErrAuthPasswdNotMatch: "用户名或密码不正确",
	ErrAuthNameNotExist:   "用户名不存在",
	ErrAuthInvalidStatus:  "用户状态异常",
	ErrAuthDB:             "数据库错误",
	ErrUnauthorized:       "没有权限",
	ErrPasswdStrength:     "密码必须同时包含数字、大小写字母、特殊字符",
	ErrPasswdLength:       "密码长度必须在8位到32位之间",
	ErrUserNameExist:      "用户名已存在",
	ErrPasswdDec:          "登陆失败",
	ErrLoginLocked:        "登录失败次数过多，账号已被锁定30分钟",
	ErrSign:               "gis接口签名失败",

	ErrSimulatorCommon:       "仿真失败",
	ErrSimulatorComputingErr: "任务计算中",
	ErrSimulatorCancelErr:    "非计算中，不能操作取消",

	ErrTReportCommon:          "报告文件生成失败",
	ErrTReportUploadErr:       "文件上传失败",
	ErrTReportUploadLimitErr:  "文件大小超过限制了",
	ErrTReportUploadFormatErr: "文件格式不正确",

	ErrCompanyNotMaintained: "企业名称未在系统中维护",
	ErrCompanyNameRequired:  "企业名称不能为空",
	ErrProjectNameRequired:  "项目名称不能为空",
	ErrDateRequired:         "日期不能为空",
	ErrEntNoFound:           "企业不存在",
	ErrEntBizGet:            "获取企业信息失败",
}

type CodeError struct {
	stackErr error
	Code     ErrCode
	Msg      string
}

type CodeErrorResponse struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
}

func NewCodeError(code ErrCode, err error) error {
	msg := err.Error()
	if msg == "" {
		msg = errCodeMap[code]
	} else {
		msg = cleanRPCError(msg)
	}
	return &CodeError{Code: code, Msg: msg, stackErr: errors.WithStack(err)}
}

func NewDefaultError(err error) error {
	return NewCodeError(ErrExecFailed, err)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// HandleError 自定义错误
func HandleError(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
		logx.Errorf("%+v", e.stackErr)
		return http.StatusOK, e.Data()
	default:
		logx.Errorf("%+v", err)
		return http.StatusOK, &CodeErrorResponse{
			Code: ErrExecFailed,
			Msg:  errCodeMap[ErrExecFailed],
		}
	}
}

// OKResponse : Success Response for HTTP API
type OKResponse struct {
	Code ErrCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(data interface{}) *OKResponse {
	return &OKResponse{
		Code: ErrNone,
		Msg:  errCodeMap[ErrNone],
		Data: data,
	}
}

func ValidateErr(err error) *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: ErrInvalidParam,
		Msg:  err.Error(),
	}
}

func cleanRPCError(errMsg string) string {

	// 定义要去除的前缀
	const prefix = "rpc error: code = Unknown desc ="

	// 检查错误信息是否包含前缀
	if strings.HasPrefix(errMsg, prefix) {
		// 去除前缀并返回干净的错误信息
		return fmt.Sprintf("%s", strings.TrimPrefix(errMsg, prefix))
	}

	// 如果没有前缀，直接返回原始错误
	return errMsg
}
