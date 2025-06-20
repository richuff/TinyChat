package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	res, err := json.Marshal(h)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code:  code,
		Data:  data,
		Total: total,
	}
	res, err := json.Marshal(h)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func RespOkList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, 0, data, total)
}
