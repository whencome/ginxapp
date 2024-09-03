package data

import (
    "github.com/whencome/goutil/cachex"
)

// cacheRepo 实现缓存得数据接口
type cacheRepo struct {
    data *Data
}

func NewCacheRepo(d *Data) *cacheRepo {
    return &cacheRepo{
        data: d,
    }
}

// Call 执行有缓存的调用
func (repo *cacheRepo) Call(ret interface{}, cacheKey string, ttl int64, f func() (interface{}, error)) error {
    return cachex.PCall(ret, repo.data, cacheKey, ttl, f)
}

// LockCall 执行带锁的调用，因为使用的是redis，因此相当于实现了一个简单的分布式锁
func (repo *cacheRepo) LockCall(lockKey string, ttl int64, f func() error) error {
    return cachex.PLockCall(nil, repo.data, lockKey, ttl, true, func() (interface{}, error) {
        return nil, f()
    })
}

// Store 缓存数据
func (repo *cacheRepo) Store(cacheKey string, ttl int64, data interface{}) error {
    return cachex.PStore(repo.data, cacheKey, ttl, data)
}

// Fetch 取出对应的缓存内容
func (repo *cacheRepo) Fetch(ret interface{}, cacheKey string) error {
    return cachex.PFetch(ret, repo.data, cacheKey)
}

// Remove 删除指定缓存
func (repo *cacheRepo) Remove(cacheKey string) {
    cachex.PRemove(repo.data, cacheKey)
}
