package requests

import (
    "github.com/whencome/ginxapp/pkg/validator"
    "github.com/whencome/ginxapp/pkg/xerr"
)

// PageRequest 分页请求
type PageRequest struct {
    Page     int `form:"page,default=1" label:"分页页数" binding:"omitempty,gt=0"`
    PageSize int `form:"page_size,default=10" label:"分页页数" binding:"omitempty,gt=0"`
}

func (r *PageRequest) GetPage() int {
    return r.Page
}

func (r *PageRequest) GetPageSize() int {
    return r.PageSize
}

func (r *PageRequest) Offset() int {
    if r.Page < 1 {
        return 0
    }
    return (r.Page - 1) * r.PageSize
}

// AdjustPagination 自动校正分页请求参数
func (r *PageRequest) AdjustPagination() {
    if r.Page < 1 {
        r.Page = 1
    }
    if r.PageSize < 1 || r.PageSize > 200 {
        r.PageSize = 10
    }
}

type TokenRequest struct {
    Token string `form:"t" label:"令牌"`
}

// LoginRequest 账户登录请求，适用于密码登录和短信验证码登录
type LoginRequest struct {
    Mobile    string `form:"mobile" label:"手机号" binding:"required"`
    Token     string `form:"t" label:"令牌" binding:"required"`           // 用于图形验证码唯一性识别
    ImageCode string `form:"img_code" label:"图形验证码" binding:"required"` // 图形验证码，部分场景需要
    Password  string `form:"password" label:"密码"`
}

func (r *LoginRequest) Validate() error {
    if !validator.IsMobile(r.Mobile) {
        return xerr.BadRequest("手机号格式不正确")
    }
    return nil
}

// RefreshTokenRequest 刷新token
type RefreshTokenRequest struct {
    RefreshToken string `form:"refresh_token" binding:"required"`
}
