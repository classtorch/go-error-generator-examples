package main

import (
	rawErrors "errors"
	"fmt"
	"github.com/classtorch/go-error-generator-examples/internal/errors"
	"github.com/classtorch/go-error-generator-examples/internal/golang/account"
	errors2 "github.com/classtorch/go-error-generator-examples/internal/golang/errors"
	"strconv"
)

const (
	AppId_Chinese = 0
	AppId_Enlish  = 1
)

func main() {
	err := checkAccountExist("10000000001")
	if err == nil {
		return
	}
	code, msg := convertError(AppId_Chinese, err)
	fmt.Printf("chinese,code:%d, msg:%s\n", code, msg)
	code, msg = convertError(AppId_Enlish, err)
	fmt.Printf("english, code:%d, msg:%s", code, msg)
}

func checkAccountExist(phoneNo string) error {
	if phoneNo != "10000000000" {
		return account.ACCOUNT_NOT_EXISTS
	}
	return nil
}

func convertError(appId int, err error) (int, string) {
	codeStr := ""
	msg := ""
	var e *errors.Error
	if rawErrors.As(err, &e) {
		codeStr = strconv.Itoa(int(e.Code))
		msg = e.Msg
		if appId == AppId_Enlish {
			mErr, ok := errors2.Msg_English[e.Code]
			if ok {
				msg = mErr.Msg
			}
		} else {
			mErr, ok := errors2.Msg[e.Code]
			if ok {
				msg = mErr.Msg
			}
		}
	} else {
		codeStr = "-1"
		msg = err.Error()
	}
	retCode, _ := strconv.Atoi(codeStr)
	return retCode, msg
}
