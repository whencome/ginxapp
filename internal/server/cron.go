package server

import (
    "fmt"
    "path/filepath"
    "sync"

    cron "github.com/robfig/cron/v3"
    "github.com/whencome/ginxapp/internal/biz"
    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/goutil/log"
)

// cronTask 定义一个定时任务方法
type cronTask func(l log.Logger) error

type CronServer struct {
    cfg     *etc.CronConfig // 配置
    c       *cron.Cron
    cronUc  *biz.CronUseCase
    loggers map[string]log.Logger
    mu      sync.Mutex
}

func NewCronServer(cronUc *biz.CronUseCase) *CronServer {
    // create server
    c := &CronServer{
        cfg:     etc.AppConf.Cron,
        c:       cron.New(),
        cronUc:  cronUc,
        loggers: make(map[string]log.Logger),
        mu:      sync.Mutex{},
    }
    // init tasks
    c.init()
    return c
}

func (s *CronServer) Start() {
    s.c.Start()
}

func (s *CronServer) Stop() {
    s.c.Stop()
}

// init 服务初始化
func (s *CronServer) init() {
    // 测试：
    s.c.AddFunc("*/10 * * * *", s.newTask("timer_echo", s.cronUc.Echo))
}

// getTaskLogger 获取任务的日志对象
func (s *CronServer) getTaskLogger(taskName string) log.Logger {
    // 检查日志对象是否已经存在
    if l, ok := s.loggers[taskName]; ok {
        return l
    }
    // 重新创建日志对象
    fileName := fmt.Sprintf("cron_%s.log", taskName)
    cfg := &log.Config{
        Level:        "info",
        Output:       "file",
        Path:         filepath.Join(s.cfg.LogDir, fileName),
        RotationTime: "240h",
        MaxKeepTime:  "720h",
    }
    logger := log.Instance(cfg)
    s.mu.Lock()
    s.loggers[taskName] = logger
    s.mu.Unlock()
    return logger
}

// newTask 创建一个新任务
// taskName 任务名称，将作为日志文件的名称，因此尽量不要使用特殊符号以及空格等
func (s *CronServer) newTask(taskName string, task cronTask) func() {
    return func() {
        logger := s.getTaskLogger(taskName)
        logger.Infof("start cron %s ...", taskName)
        defer func() {
            if err := recover(); err != nil {
                logger.Errorf("cron %s recovered from err: %v", taskName, err)
            }
        }()
        // 执行任务
        err := task(logger)
        if err != nil {
            logger.Errorf("cron exit with error: %v", err)
            return
        }
        logger.Info("task execute success")
    }
}
