package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/whencome/ginx"
    "github.com/whencome/ginxapp/internal/biz"
    "github.com/whencome/ginxapp/internal/biz/def/requests"
    "github.com/whencome/ginxapp/pkg/xerr"
)

type AuthHandler struct {
    accountUc *biz.AuthUseCase
    captchaUc *biz.CaptchaUseCase
}

func NewAuthHandler(accountUc *biz.AuthUseCase, captchaUc *biz.CaptchaUseCase) *AuthHandler {
    return &AuthHandler{
        accountUc: accountUc,
        captchaUc: captchaUc,
    }
}

// RegisterRoute 注册路由
func (h *AuthHandler) RegisterRoute(g *gin.RouterGroup) {
    // 登录，支持密码和验证码登录
    g.POST("/login", ginx.NewApiHandler(requests.LoginRequest{}, h.Login))
    // 刷新token
    g.POST("/refresh_token", ginx.NewApiHandler(requests.RefreshTokenRequest{}, h.RefreshToken))
}

// Login 登录，支持密码登录和短信验证码登录
func (h *AuthHandler) Login(c *gin.Context, r ginx.Request) (ginx.Response, error) {
    req, ok := r.(*requests.LoginRequest)
    if !ok {
        return nil, xerr.ErrInvalidRequest
    }
    // 验证图形验证码
    token := h.captchaUc.GetToken(c, req.Token)
    isOk := h.captchaUc.Verify(token, req.ImageCode)
    if !isOk {
        return nil, xerr.BadRequest("图形验证码不正确")
    }
    // 登录
    return h.accountUc.Login(c, req)
}

// RefreshToken 刷新token
func (h *AuthHandler) RefreshToken(c *gin.Context, r ginx.Request) (ginx.Response, error) {
    req, ok := r.(*requests.RefreshTokenRequest)
    if !ok {
        return nil, xerr.ErrInvalidRequest
    }
    return h.accountUc.RefreshToken(c, req)
}
