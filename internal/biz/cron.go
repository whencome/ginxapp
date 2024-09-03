package biz

import (
    "github.com/whencome/goutil/log"
    "github.com/whencome/goutil/timeutil"
)

// CronUseCase 定时任务用例
type CronUseCase struct {
}

func NewCronUseCase() *CronUseCase {
    return &CronUseCase{}
}

// Echo 输出当前时间，用于定时任务测试
// 每10分钟执行一次
func (uc *CronUseCase) Echo(logger log.Logger) error {
    logger.Infof("now is %s ...", timeutil.Now().String())
    // 操作完成
    return nil
}
