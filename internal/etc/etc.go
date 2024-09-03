package etc

import (
    "os"
    "sync"

    "gopkg.in/yaml.v2"
)

// AppConf 全局站点配置
var AppConf *AppConfig = nil

// 用于全局配置更新时
var locker sync.Mutex

// Load 加载配置
// f - 配置文件地址
func Load(f string) error {
    conf := new(AppConfig)
    yamlFile, err := os.ReadFile(f)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(yamlFile, conf)
    if err != nil {
        return err
    }
    locker.Lock()
    AppConf = conf
    locker.Unlock()
    return nil
}
