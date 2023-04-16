package errorx

const (
	SuccessCode           = 10001
	AlreadyExistsCode     = 10002
	NotFoundCode          = 10003
	ServerErrorCode       = 10004
	ParameterErrorCode    = 10005
	LoginErrorCode        = 10006
	UnregisteredErrorCode = 1007
)

var (
	ErrMsgMap = map[int]string{
		SuccessCode:           "success",
		AlreadyExistsCode:     "already exists",
		NotFoundCode:          "not found",
		ServerErrorCode:       "server error",
		ParameterErrorCode:    "parameter error",
		LoginErrorCode:        "incorrect username or password",
		UnregisteredErrorCode: "user not registered",
	}
)

func ErrMsg(code int) string {
	msg, ok := ErrMsgMap[code]
	if ok {
		return msg
	}
	return "UnknownError"
}
