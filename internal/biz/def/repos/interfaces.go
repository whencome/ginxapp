package repos

// CaptchaRepo 验证码相关接口
type CaptchaRepo interface {
    // Set sets the digits for the captcha id.
    Set(id string, value string) error
    // Get returns stored digits for the captcha id. Clear indicates
    // whether the captcha must be deleted from the store.
    Get(id string, clear bool) string
    // Verify captcha's answer directly
    Verify(id, answer string, clear bool) bool
    // StoreRequestId 报存请求与验证码ID的关系
    StoreRequestId(reqToken, id string) error
    // FetchRequestId 根据请求token获取验证码ID
    FetchRequestId(reqToken string) string
}

// CacheRepo 缓存数据接口
type CacheRepo interface {
    // Call 执行有缓存的调用
    Call(ret interface{}, cacheKey string, ttl int64, f func() (interface{}, error)) error
    // LockCall 执行带锁的调用，因为使用的是redis，因此相当于实现了一个简单的分布式锁
    LockCall(lockKey string, ttl int64, f func() error) error
    // Store 缓存数据
    Store(cacheKey string, ttl int64, data interface{}) error
    // Fetch 取出对应的缓存内容
    Fetch(ret interface{}, cacheKey string) error
    // Remove 删除指定缓存
    Remove(cacheKey string)
}
