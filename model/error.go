package model

import (
	"fmt"
)

var LimitError = GameError{Msg: "数量约束错误", Code: 50001}

type ErrorCode int

type GameError struct {
	Msg    string // 信息
	Code   int    // 错误码
	Detail string // 详细输出
}

func (e *GameError) Error() string {
	return fmt.Sprintf("[%d][%s]%s", e.Code, e.Msg, e.Detail)
}
