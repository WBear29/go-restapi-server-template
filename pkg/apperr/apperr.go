package apperr

import "net/http"

const (
	ERROR_INVALID_PATM_PARAMS AppErrorCode = 1000 + iota
	ERROR_INVALID_QUERY_PARAMS
	ERROR_INVALID_REQUEST_BODY
	ERROR_INTERNAL_SERVER_ERROR
)

type AppErrorCode int

type Err interface {
	Status() int
	Code() int
	Message() string
	Log() string
}

var errorStatusMap = map[AppErrorCode]int{
	ERROR_INVALID_PATM_PARAMS:   http.StatusBadRequest,
	ERROR_INVALID_QUERY_PARAMS:  http.StatusBadRequest,
	ERROR_INVALID_REQUEST_BODY:  http.StatusBadRequest,
	ERROR_INTERNAL_SERVER_ERROR: http.StatusInternalServerError,
}

var errorMessageMap = map[AppErrorCode]string{
	ERROR_INVALID_PATM_PARAMS:   "Invalid path params",
	ERROR_INVALID_QUERY_PARAMS:  "Invalid query params",
	ERROR_INVALID_REQUEST_BODY:  "Invalid request body",
	ERROR_INTERNAL_SERVER_ERROR: "Internal server error",
}

type appErr struct {
	code AppErrorCode
	log  string
}

func New(code AppErrorCode, log string) Err {
	return &appErr{code, log}
}

func (e appErr) Status() int {
	return errorStatusMap[e.code]
}

func (e appErr) Code() int {
	return int(e.code)
}

func (e appErr) Message() string {
	return errorMessageMap[e.code]
}

func (e appErr) Log() string {
	return e.log
}
