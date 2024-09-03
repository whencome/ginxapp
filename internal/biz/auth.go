package biz

import (
    "github.com/gin-gonic/gin"
    "github.com/whencome/ginx"
    "github.com/whencome/ginxapp/internal/biz/def"
    "github.com/whencome/ginxapp/internal/biz/def/repos"
    "github.com/whencome/ginxapp/internal/biz/def/requests"
    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/ginxapp/pkg/kits"
)

type AuthUseCase struct {
    cacheRepo repos.CacheRepo
}

func NewAuthUseCase(cacheRepo repos.CacheRepo) *AuthUseCase {
    return &AuthUseCase{
        cacheRepo: cacheRepo,
    }
}

// RefreshToken 刷新token
func (uc *AuthUseCase) RefreshToken(c *gin.Context, req *requests.RefreshTokenRequest) (ginx.Response, error) {
    j := kits.NewJWT(etc.AppConf.Jwt.SigningKey)
    claims, err := j.ParseToken(req.RefreshToken)
    if err != nil {
        return nil, err
    }
    userId := claims.UserId
    loginUserInfo := &def.LoginUserInfo{
        Uid:      userId,
        Name:     "",
        Mobile:   "",
        OpenId:   "",
        Nickname: "",
        Avatar:   "",
    }
    return uc.makeLogin(c, loginUserInfo)
}

// makeLogin 生成登录信息
func (uc *AuthUseCase) makeLogin(c *gin.Context, user *def.LoginUserInfo) (*requests.LoginResponse, error) {
    // 构建登录信息
    j := kits.NewJWT(etc.AppConf.Jwt.SigningKey)
    claims := j.CreateClaims(kits.BaseClaims{
        UserId:   user.Uid,
        UserName: user.Nickname,
        Data: map[string]interface{}{
            "openid": user.OpenId,
            "avatar": user.Avatar,
        },
    }, etc.AppConf.Jwt.Issuer, etc.AppConf.Jwt.BufferTime, etc.AppConf.Jwt.ExpiresTime)
    accessToken, refreshToken, err := j.CreateToken(claims)
    if err != nil {
        return nil, err
    }
    resp := &requests.LoginResponse{
        Uid:          user.Uid,
        Mobile:       user.Mobile,
        Nickname:     user.Nickname,
        OpenId:       user.OpenId,
        Avatar:       user.Avatar,
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }
    // 返回登录结果
    return resp, nil
}

// Login 登录
func (uc *AuthUseCase) Login(c *gin.Context, req *requests.LoginRequest) (ginx.Response, error) {
    // 需自行获取登录用户信息
    loginUserInfo := &def.LoginUserInfo{
        Uid:      0,
        Name:     "",
        Mobile:   "",
        OpenId:   "",
        Nickname: "",
        Avatar:   "",
    }
    return uc.makeLogin(c, loginUserInfo)
}
