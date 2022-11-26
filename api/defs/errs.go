package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

// 初始化两个ErrorResponse和Err两个对象
var (
	// ErrorRequestBodyParseFailed : Request 传入的消息体解析失败或者无法解析
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{
		Error: "Request body is not correct", ErrorCode: "001",
	}}
	// ErrorNotAuthUser : 用户不存在
	ErrorNotAuthUser = ErrorResponse{HttpSC: 401, Error: Err{
		Error: "User authentication failed.", ErrorCode: "002",
	}}
)
