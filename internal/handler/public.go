package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/whencome/ginx"
    "github.com/whencome/ginxapp/internal/biz"
    "github.com/whencome/ginxapp/internal/biz/def/requests"
    "github.com/whencome/ginxapp/pkg/xerr"
)

// PublicHandler 公共服务，用于提供不需要用户登录的一些接口
type PublicHandler struct {
    captchaUc *biz.CaptchaUseCase
}

func NewPublicHandler(captchaUc *biz.CaptchaUseCase) *PublicHandler {
    return &PublicHandler{
        captchaUc: captchaUc,
    }
}

// RegisterRoute 注册路由
func (h *PublicHandler) RegisterRoute(g *gin.RouterGroup) {
    // 图形验证码
    g.GET("/captcha/image", ginx.NewApiHandler(requests.TokenRequest{}, h.ImgCode))
}

// ImgCode 生成验证码图片
func (h *PublicHandler) ImgCode(c *gin.Context, r ginx.Request) (ginx.Response, error) {
    req, ok := r.(*requests.TokenRequest)
    if !ok {
        return nil, xerr.ErrInvalidRequest
    }
    // 生成验证码
    token := h.captchaUc.GetToken(c, req.Token)
    imgData, err := h.captchaUc.Make(token)
    if err != nil {
        return nil, err
    }
    // 将验证码图片写入响应体
    c.Data(200, "image/png", imgData)
    // 返回nil不会处理结果
    return nil, nil
}
