package errno

import (
	"errors"
	"fmt"
)

const (
	// System Code
	SuccessCode    = 0
	ServiceErrCode = 10001
	ParamErrCode   = 10002
	Unauthorized   = 10005
	AuthorizeFail  = 10006

	LoginErrCode            = 11001
	UserNotExistErrCode     = 11002
	UserAlreadyExistErrCode = 11003
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(SuccessCode, "成功")
	ServiceErr = NewErrNo(ServiceErrCode, "服务器错误")
	ParamErr   = NewErrNo(ParamErrCode, "参数错误")

	UnauthorizedErr  = NewErrNo(Unauthorized, "无权限")
	AuthorizeFailErr = NewErrNo(AuthorizeFail, "授权失败")

	LoginPasswordErr    = NewErrNo(LoginErrCode, "密码错误")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "用户不存在")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "用户已存在")
	UserMobileExistErr  = NewErrNo(UserAlreadyExistErrCode, "手机号已存在")
	TimeFormatErr       = NewErrNo(ParamErrCode, "日期格式错误")
	RecordNotFound      = NewErrNo(ParamErrCode, "记录不存在")
	RecordAlreadyExist  = NewErrNo(ParamErrCode, "记录已存在")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
