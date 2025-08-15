package common

import "github.com/pkg/errors"

var (
	ErrorNotFound          = errors.New("未找到")
	ErrInvalidType         = errors.New("无效的类型")
	ErrInvalidVersion      = errors.New("无效的版本号")
	ErrNoUncommittedEvents = errors.New("没有未提交的事件")
)
