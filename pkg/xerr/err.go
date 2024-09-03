package xerr

import (
    "fmt"
    "net/http"
    "runtime/debug"
)

// XErr 自定义错误信息
type XErr struct {
    Code    int
    Message string
    Stack   string
}

// Error 实现error接口
func (e *XErr) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// String 实现fmt.Stringer接口
func (e *XErr) String() string {
    return e.Error()
}

// WithStack 增加记录调用栈信息
func (e *XErr) WithStack(stack string) *XErr {
    e.Stack = stack
    return e
}

// newErr 创建一个新的Err对象
func newErr(code int, msg string) *XErr {
    return &XErr{
        Code:    code,
        Message: msg,
        Stack:   string(debug.Stack()),
    }
}

// Wrap 对任意错误信息进行包装
func Wrap(err error) *XErr {
    if xe, ok := err.(*XErr); ok {
        if xe.Stack == "" {
            xe.Stack = string(debug.Stack())
        }
        return xe
    }
    return &XErr{
        Code:    http.StatusBadRequest,
        Message: err.Error(),
        Stack:   string(debug.Stack()),
    }
}

// From 根据error创建错误
func From(code int, err error) *XErr {
    return newErr(code, err.Error())
}

// New 创建自定义错误信息
func New(code int, msg string) *XErr {
    return newErr(code, msg)
}

// Newf 创建新的自带格式化的错误信息
func Newf(code int, format string, v ...interface{}) *XErr {
    return newErr(code, fmt.Sprintf(format, v...))
}

// BadRequest 错误的请求
func BadRequest(format string, args ...interface{}) *XErr {
    return Newf(http.StatusBadRequest, format, args...)
}

func Unauthorized(format string, args ...interface{}) *XErr {
    return Newf(http.StatusUnauthorized, format, args...)
}

func PaymentRequired(format string, args ...interface{}) *XErr {
    return Newf(http.StatusPaymentRequired, format, args...)
}

func Forbidden(format string, args ...interface{}) *XErr {
    return Newf(http.StatusForbidden, format, args...)
}

func NotFound(format string, args ...interface{}) *XErr {
    return Newf(http.StatusNotFound, format, args...)
}

func MethodNotAllowed(format string, args ...interface{}) *XErr {
    return Newf(http.StatusMethodNotAllowed, format, args...)
}

func NotAcceptable(format string, args ...interface{}) *XErr {
    return Newf(http.StatusNotAcceptable, format, args...)
}

func RequestTimeout(format string, args ...interface{}) *XErr {
    return Newf(http.StatusRequestTimeout, format, args...)
}

func Locked(format string, args ...interface{}) *XErr {
    return Newf(http.StatusLocked, format, args...)
}

func InternalServerError(format string, args ...interface{}) *XErr {
    return Newf(http.StatusInternalServerError, format, args...)
}

func BadGateway(format string, args ...interface{}) *XErr {
    return Newf(http.StatusBadGateway, format, args...)
}

func ServiceUnavailable(format string, args ...interface{}) *XErr {
    return Newf(http.StatusServiceUnavailable, format, args...)
}

func GatewayTimeout(format string, args ...interface{}) *XErr {
    return Newf(http.StatusGatewayTimeout, format, args...)
}
