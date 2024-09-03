package biz

import (
    "encoding/base64"
    "fmt"
    "image/color"
    "strings"

    "github.com/whencome/ginxapp/internal/biz/def/repos"
    "github.com/whencome/ginxapp/pkg/encrypt"

    "github.com/gin-gonic/gin"
    "github.com/mojocn/base64Captcha"
)

// CaptchaUseCase 验证码用例
type CaptchaUseCase struct {
    repo    repos.CaptchaRepo
    captcha base64Captcha.Captcha
}

func NewCaptchaUseCase(repo repos.CaptchaRepo) *CaptchaUseCase {
    // 配置验证码的参数
    driverString := base64Captcha.DriverString{
        Height:          38,
        Width:           109,
        NoiseCount:      0,
        ShowLineOptions: 2 | 4,
        Length:          6,
        Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
        BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
        Fonts:           []string{"wqy-microhei.ttc"},
    }
    // ConvertFonts 按名称加载字体
    driver := driverString.ConvertFonts()
    captcha := base64Captcha.NewCaptcha(driver, repo)
    // 创建CaptchaUseCase
    return &CaptchaUseCase{
        repo:    repo,
        captcha: *captcha,
    }
}

// GetToken 获取验证码标识
func (uc *CaptchaUseCase) GetToken(c *gin.Context, t string) string {
    token := fmt.Sprintf("%s_%s", t, c.ClientIP())
    f := encrypt.Md5Short(token)
    return f
}

// Make 生成验证码，返回图片字节切片
func (uc *CaptchaUseCase) Make(token string) ([]byte, error) {
    id, imgBase64, _, err := uc.captcha.Generate()
    if err != nil {
        return nil, err
    }
    // 缓存请求与id的关系
    err = uc.repo.StoreRequestId(token, id)
    if err != nil {
        return nil, err
    }
    commaPos := strings.Index(imgBase64, ",")
    if commaPos > 0 {
        imgBase64 = imgBase64[commaPos+1:]
    }
    return base64.StdEncoding.DecodeString(imgBase64)
}

// Verify 验证验证码
func (uc *CaptchaUseCase) Verify(token, code string) bool {
    captchaId := uc.repo.FetchRequestId(token)
    return uc.captcha.Verify(captchaId, code, true)
}

// VerifyAndKeep 验证并保留验证码
func (uc *CaptchaUseCase) VerifyAndKeep(token, code string) bool {
    captchaId := uc.repo.FetchRequestId(token)
    return uc.captcha.Verify(captchaId, code, false)
}
