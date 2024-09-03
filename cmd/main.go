package main

import (
    "flag"

    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/ginxapp/internal/server"
    "github.com/whencome/ginxapp/pkg/responser"
    "github.com/whencome/goutil/cachex"
    "github.com/whencome/goutil/fileutil"
    "github.com/whencome/goutil/log"

    "github.com/whencome/ginx"
)

var app *App

// App 定义全局应用
type App struct {
    httpSvr *ginx.HTTPServer
    cronSvr *server.CronServer
    clean   func()
}

// NewApp create global
func NewApp(httpSvr *ginx.HTTPServer, cronSvr *server.CronServer) (*App, error) {
    app := new(App)
    app.httpSvr = httpSvr
    app.cronSvr = cronSvr
    return app, nil
}

// Start 启动应用
func (app *App) Start() {
    // 启动cron server
    if app.cronSvr != nil && etc.AppConf.Cron.IsEnabled {
        go func() {
            app.cronSvr.Start()
            log.Infof("cron server started...")
        }()
    }
    // 启动http服务
    if app.httpSvr == nil {
        log.Panicf("http server has not been initialized")
    }
    if _, err := app.httpSvr.Start(); err != nil {
        log.Panicf("start http server fail: %v", err)
    }
    log.Infof("http server stated on port :%d in %s mode", etc.AppConf.Site.Port, etc.AppConf.Site.Mode)
}

// Stop 停止应用
func (app *App) Stop() {
    if app.cronSvr != nil {
        app.cronSvr.Stop()
        app.cronSvr = nil
    }
    if app.httpSvr != nil {
        app.httpSvr.Stop()
        app.httpSvr = nil
    }
    if app.clean != nil {
        app.clean()
    }
}

// loadConfig 加载配置
func loadConfig() {
    // 获取参数配置
    cfgFile := flag.String("f", "etc/app.yaml", "app config file")
    flag.Parse()
    // 检查配置文件是否存在
    if !fileutil.Exists(*cfgFile) {
        log.Panicf("配置文件%s不存在", *cfgFile)
        return
    }
    // 加载配置
    err := etc.Load(*cfgFile)
    if err != nil {
        log.Panicf("加载配置文件%s失败：%v", *cfgFile, err)
        return
    }
}

// runApp 运行服务
func runApp() {
    // 初始化日志组件
    logger := log.New(etc.AppConf.Log)
    // 初始化cachex
    cachex.SetCacheKeyPrefix("device")
    // 注册api响应对象
    ginx.UseApiResponser(new(responser.ApiResponser))
    // 注册ginx全局日志对象
    ginx.UseLogger(logger)

    // 初始化并启动app
    var err error
    var cleanup func()
    app, cleanup, err = wireApp(logger, etc.AppConf.Site, etc.AppConf.Redis)
    if err != nil {
        log.Panicf("wire app fail: %v", err)
    }
    app.clean = cleanup
    app.Start()
}

// main app entry
func main() {
    // 加载配置
    loadConfig()
    // 运行服务
    runApp()
    // 等待程序退出
    ginx.Wait(func() {
        app.Stop()
    })
}
