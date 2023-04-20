// Code generated by protoc-gen-go-error-generator. DO NOT EDIT.
// Version v1.0.0
package errors

import (
	errors "github.com/classtorch/go-error-generator-examples/internal/errors"
)

var (
	Msg = map[int32]*errors.Error{
		0: &errors.Error{Code: 0, Msg: "成功"},
		1: &errors.Error{Code: 1, Msg: "登录失效，请重新登录"},
		2: &errors.Error{Code: 2, Msg: "账号不存在"},
		3: &errors.Error{Code: 3, Msg: "订单不存在"},
	}

	Msg_English = map[int32]*errors.Error{
		0: &errors.Error{Code: 0, Msg: "success"},
		1: &errors.Error{Code: 1, Msg: "token failure，please login"},
		2: &errors.Error{Code: 2, Msg: "account not exist"},
		3: &errors.Error{Code: 3, Msg: "order not exist"},
	}
)