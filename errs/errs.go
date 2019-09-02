/**
自定义错误码
规则，错误码均为6位数字
前三位用于区分 业务类型（新业务错误码在现有的顺序上累加）
后三位用于区分 错误类型（新错误码在现有的顺序上累加哦）
*/
package errs

/**
系统错误码
*/
var (
	Success      = &Error{100000, "Success"}
	ErrorArgs    = &Error{100001, "参数错误"}
	ErrorService = &Error{100002, "网络开小差了，一会儿再来试试吧~"}
	ErrorNoData  = 100003
)

// Error ...
type Error struct {
	Code int
	Msg  string
}

// Error 设置错误信息
func (err Error) Error() string {
	return err.Msg
}

// NewError 返回新的错误，根据http状态码，msg
func New(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}
