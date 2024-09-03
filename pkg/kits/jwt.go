package kits

import (
    "errors"
    "time"

    jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
    SigningKey []byte
}

var (
    TokenExpired     = errors.New("Token is expired")
    TokenNotValidYet = errors.New("Token not active yet")
    TokenMalformed   = errors.New("That's not even a token")
    TokenInvalid     = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
    BaseClaims
    BufferTime int64
    jwt.RegisteredClaims
}

type BaseClaims struct {
    UserId   int64                  // 用户ID
    UserName string                 // 用户名/昵称
    Data     map[string]interface{} // 用于保存其它需要的数据
}

func NewJWT(jwtKey string) *JWT {
    return &JWT{
        []byte(jwtKey),
    }
}

func (j *JWT) CreateClaims(baseClaims BaseClaims, issuer string, bufferTime, expiresTime int64) CustomClaims {
    claims := CustomClaims{
        BaseClaims: baseClaims,
        BufferTime: bufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
        RegisteredClaims: jwt.RegisteredClaims{
            NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),                      // 签名生效时间
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expiresTime))), // 过期时间 7天  配置文件
            Issuer:    issuer,                                                                       // 签名的发行者
        },
    }
    return claims
}
func (j *CustomClaims) Setissuer(issuer string) {
    j.Issuer = issuer
}

func (j *CustomClaims) SetBufferTime(bufferTime int64) {
    j.BufferTime = bufferTime
}

func (j *CustomClaims) SetExpiresAt(expiresTime int64) {
    j.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expiresTime)))
}

// CreateToken 创建并返回access token以及refresh token
func (j *JWT) CreateToken(claims CustomClaims) (string, string, error) {
    // 创建access token
    accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
    if err != nil {
        return "", "", err
    }
    // 创建refresh token， refresh token不需要保存任何用户信息
    rClaims := claims
    rClaims.BaseClaims = BaseClaims{
        UserId: claims.UserId,
    }
    rClaims.ExpiresAt.Add(time.Hour * 240) // 给10天宽限期
    refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims).SignedString(j.SigningKey)
    if err != nil {
        return "", "", err
    }
    return accessToken, refreshToken, nil
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, string, error) {
    // 解析token
    _, err := j.ParseToken(oldToken)
    if err != nil {
        return "", "", err
    }
    return j.CreateToken(claims)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
    dstClaims := new(CustomClaims)
    _, err := jwt.ParseWithClaims(tokenString, dstClaims, func(token *jwt.Token) (i interface{}, e error) {
        return j.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return dstClaims, TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return dstClaims, TokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return dstClaims, TokenNotValidYet
            } else {
                return dstClaims, TokenInvalid
            }
        }
    }
    return dstClaims, nil
}

// RefreshToken 无感刷新token
func (j *JWT) RefreshToken(accessToken, refreshToken string) (string, string, error) {
    _, err := j.ParseToken(refreshToken)
    if err != nil {
        return "", "", err
    }
    claims, err := j.ParseToken(accessToken)
    if err == nil || err == TokenExpired {
        return j.CreateToken(*claims)
    }
    return "", "", err
}
