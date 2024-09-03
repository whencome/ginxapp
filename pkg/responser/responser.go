package responser

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/whencome/ginxapp/pkg/xerr"
)

type ApiResponseMessage struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type ApiResponser struct{}

func (r *ApiResponser) buildMsg(c *gin.Context, code int, v interface{}) ApiResponseMessage {
    msg := ApiResponseMessage{
        Code:    code,
        Message: "",
        Data:    nil,
    }
    // SUCCESS
    if code == http.StatusOK {
        msg.Message = "success"
        msg.Data = v
        msg.Code = 0
        return msg
    }
    // FAIL
    e, ok := v.(error)
    if ok {
        if xe, ok := e.(*xerr.XErr); ok {
            msg.Code = xe.Code
            msg.Message = xe.String()
        } else {
            msg.Message = e.Error()
        }
    } else {
        msg.Message = fmt.Sprintf("%s", v)
    }
    if msg.Code == 0 {
        msg.Code = 400
    }
    return msg
}

func (r *ApiResponser) Response(c *gin.Context, code int, v interface{}) {
    msg := r.buildMsg(c, code, v)
    c.JSON(http.StatusOK, msg)
    c.Abort()
}

func (r *ApiResponser) Success(c *gin.Context, v interface{}) {
    r.Response(c, http.StatusOK, v)
}

func (r *ApiResponser) Fail(c *gin.Context, v interface{}) {
    r.Response(c, http.StatusBadRequest, v)
}
