package common

type HttpError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

