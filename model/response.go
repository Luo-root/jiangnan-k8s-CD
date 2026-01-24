package model

type Response struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
	Data any        `json:"data"`
}
