package main

import (
	"encoding/json"
	"github.com/classtorch/go-error-generator-examples/internal/errors"
	"github.com/classtorch/go-error-generator-examples/internal/golang/account"
	"net/http"
	"net/url"
	"strconv"
)

const (
	Language_Chinese = 0 //中文
	Language_Enligh  = 1 //英文
)

func main() {
	http.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {
		languageStr := r.Header.Get("language")
		language, _ := strconv.Atoi(languageStr)
		values, _ := url.ParseQuery(r.URL.RawQuery)
		phone := values["phone"][0]
		err := checkPhone(phone)
		response(w, language, err)
	})
	http.ListenAndServe(":8080", nil)
}

func response(w http.ResponseWriter, language int, err error) {
	e := account.SUCCESS
	var ok bool
	if err != nil {
		if e, ok = err.(*errors.Error); !ok {
			e = account.UnKnown
		}
	}
	if language == Language_Chinese {
		if e, ok = account.Msg[e.Code]; !ok {
			e = account.UnKnown
		}
	} else if language == Language_Enligh {
		if e, ok = account.Msg_English[e.Code]; !ok {
			e = account.UnKnown
		}
	}
	res := make(map[string]interface{})
	res["code"] = e.Code
	res["msg"] = e.Msg
	json, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Write(json)
}

func checkPhone(phoneNo string) error {
	if len(phoneNo) != 11 {
		return account.InValid_Phone
	}
	return nil
}
