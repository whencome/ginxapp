package requests

// LoginResponse 登录响应结果
type LoginResponse struct {
    Uid          int64  `json:"uid"`
    Mobile       string `json:"mobile"`
    Nickname     string `json:"nickname"`
    OpenId       string `json:"open_id"`
    Avatar       string `json:"avatar"`
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}
