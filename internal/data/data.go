package data

import (
    "context"

    "github.com/gomodule/redigo/redis"
    "github.com/google/wire"
    "github.com/whencome/ginxapp/internal/biz/def/repos"
    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/goutil/log"
)

var ProviderSet = wire.NewSet(
    wire.Bind(new(repos.CaptchaRepo), new(*captchaRepo)),
    wire.Bind(new(repos.CacheRepo), new(*cacheRepo)),
    NewData,
    NewCaptchaRepo,
    NewCacheRepo,
)

// GetRedisPool 获取redis连接池
func GetRedisPool(redisConf *etc.RedisConfig) *redis.Pool {
    return &redis.Pool{
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", redisConf.Addr)
            if err != nil {
                return nil, err
            }
            if redisConf.Password != "" {
                if _, err := c.Do("AUTH", redisConf.Password); err != nil {
                    c.Close()
                    return nil, err
                }
            }
            if _, err := c.Do("SELECT", redisConf.Db); err != nil {
                c.Close()
                return nil, err
            }
            return c, nil
        },
    }
}

// Data 定义数据层资源封装
// 注意： api只连接redis，不连接数据库
type Data struct {
    redisConf *etc.RedisConfig
    RedisPool *redis.Pool
}

// NewData .
func NewData(
    logger log.Logger,
    redisConf *etc.RedisConfig,
) (*Data, func(), error) {
    // 构造返回值
    data := &Data{
        redisConf: redisConf,
    }

    // 连接redis
    data.RedisPool = GetRedisPool(redisConf)

    // 注册清除函数
    cleanUp := func() {
        if data.RedisPool != nil {
            data.RedisPool.Close()
            data.RedisPool = nil
        }
    }

    // 返回结果
    return data, cleanUp, nil
}

// Redis 获取redis连接
func (d *Data) Redis() redis.Conn {
    return d.RedisPool.Get()
}

func (d *Data) RedisCtx(ctx context.Context) (redis.Conn, error) {
    return d.RedisPool.GetContext(ctx)
}
