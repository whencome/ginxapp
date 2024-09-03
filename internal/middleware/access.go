package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/ginxapp/pkg/kits"
    "github.com/whencome/ginxapp/pkg/trace"
    "github.com/whencome/ginxapp/pkg/xerr"
    "github.com/whencome/goutil"
)

// WithTrace 添加追踪信息
func WithTrace(c *gin.Context) error {
    trace.WrapGin(c)
    c.Next()
    return nil
}

// CheckAccess 检查登录权限
func CheckAccess(c *gin.Context) error {
    token := c.Request.Header.Get("Access-Token")
    if token == "" {
        return xerr.New(10001, "用户尚未登录，请登录后再操作")
    }
    j := kits.NewJWT(etc.AppConf.Jwt.SigningKey)
    claims, err := j.ParseToken(token)
    if err != nil {
        if err == kits.TokenExpired {
            return xerr.New(10002, "授权已过期，请重新登录")
        }
        return xerr.Newf(10000, "parse jwt token fail: %v, token = %s", err, token)
    }
    // 用户登录成功，记录用户信息
    c.Set("claims", claims)
    c.Set("user_id", claims.UserId)
    c.Set("user_name", claims.UserName)
    claimsData := goutil.MVal(claims.Data)
    c.Set("partner_id", claimsData.GetInt64("partner_id"))
    c.Set("store_id", claimsData.GetInt64("store_id"))
    c.Set("avatar", claimsData.GetString("avatar"))
    c.Set("role", claimsData.GetString("role"))
    c.Next()
    return nil
}
