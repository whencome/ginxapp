package data

import (
    "fmt"

    "github.com/whencome/goutil/cachex"
)

type captchaRepo struct {
    data *Data
}

func NewCaptchaRepo(d *Data) *captchaRepo {
    return &captchaRepo{
        data: d,
    }
}

// Set sets the digits for the captcha id.
func (repo *captchaRepo) Set(id string, value string) error {
    rds := repo.data.Redis()
    defer rds.Close()
    cacheKey := fmt.Sprintf("captcha:%s", id)
    return cachex.Store(rds, cacheKey, 600, value)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (repo *captchaRepo) Get(id string, clear bool) string {
    rds := repo.data.Redis()
    defer rds.Close()
    cacheKey := fmt.Sprintf("captcha:%s", id)
    var value string
    err := cachex.Fetch(&value, rds, cacheKey)
    if err != nil {
        return ""
    }
    return value
}

// Verify captcha's answer directly
func (repo *captchaRepo) Verify(id, answer string, clear bool) bool {
    // 在验证之后立即删除缓存
    rds := repo.data.Redis()
    defer rds.Close()
    cacheKey := fmt.Sprintf("captcha:%s", id)
    defer func() {
        cachex.Remove(rds, cacheKey)
    }()
    // 比较值
    value := repo.Get(id, clear)
    if value == answer {
        return true
    }
    return false
}

// StoreRequestId 报存请求与验证码ID的关系
func (repo *captchaRepo) StoreRequestId(reqToken, id string) error {
    rds := repo.data.Redis()
    defer rds.Close()
    cacheKey := fmt.Sprintf("captcha:request:%s", reqToken)
    return cachex.Store(rds, cacheKey, 600, id)
}

// FetchRequestId 根据请求token获取验证码ID
func (repo *captchaRepo) FetchRequestId(reqToken string) string {
    rds := repo.data.Redis()
    defer rds.Close()
    cacheKey := fmt.Sprintf("captcha:request:%s", reqToken)
    var id string
    err := cachex.Fetch(&id, rds, cacheKey)
    if err != nil {
        return ""
    }
    return id
}
