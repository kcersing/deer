package errno

import (
	"errors"
)

type BaseResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *BaseResp {
	return &BaseResp{Code: err.ErrCode, Message: err.ErrMsg}
}
