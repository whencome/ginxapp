package etc

import (
    "github.com/whencome/ginx"
    "github.com/whencome/goutil/log"
)

// AppConfig 定义应用配置
type AppConfig struct {
    Site  *ginx.ServerOptions `yaml:"site"`  // 站点配置
    CORS  *CORSConfig         `yaml:"cors"`  // 跨域配置
    Log   *log.Config         `yaml:"log"`   // log配置，使用logrus以及file-rotatelogs
    Redis *RedisConfig        `yaml:"redis"` // redis配置
    Jwt   *JWTConfig          `yaml:"jwt"`   // jwt配置
    Cron  *CronConfig         `yaml:"cron"`  // 定时任务配置
}

// RedisConfig redis配置
type RedisConfig struct {
    Addr     string `yaml:"addr"`
    Password string `yaml:"password"`
    Db       int    `yaml:"db"`
}

// CORSConfig 跨域配置
type CORSConfig struct {
    IsEnabled   bool     `yaml:"is_enabled"`
    IPWhiteList []string `yaml:"ip_whitelist"`
    IPBlackList []string `yaml:"ip_blacklist"`
}

// LogConfig 定义日志配置
type LogConfig struct {
    Level       string `yaml:"level"`        // 日志级别
    Path        string `yaml:"path"`         // 日志保存路径
    RotateHours int64  `yaml:"rotate_hours"` // 多少小时切割一次
    KeepDays    int64  `yaml:"keep_days"`    // 保存时间
}

// JWTConfig 定义jwt配置
type JWTConfig struct {
    SigningKey  string `yaml:"signing-key"`  // jwt签名
    ExpiresTime int64  `yaml:"expires-time"` // 过期时间
    BufferTime  int64  `yaml:"buffer-time"`  // 缓冲时间
    Issuer      string `yaml:"issuer"`       // 签发者
}

// CronConfig 定时任务配置
type CronConfig struct {
    IsEnabled bool   `yaml:"is_enabled"` // 是否启用定时任务
    LogDir    string `yaml:"log_dir"`    // 日志存储目录
}
