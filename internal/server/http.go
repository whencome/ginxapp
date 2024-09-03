package server

import (
    "github.com/gin-gonic/gin"
    "github.com/whencome/ginx"
    "github.com/whencome/ginxapp/internal/handler"
    "github.com/whencome/ginxapp/internal/middleware"
    "github.com/whencome/goutil/log"
)

// NewHTTPServer 创建一个新的http服务器
func NewHTTPServer(
    siteConf *ginx.ServerOptions,
    logger log.Logger,
    publicHandler *handler.PublicHandler,
    authHandler *handler.AuthHandler,
// handler列表
) *ginx.HTTPServer {
    // 初始化并启动http server
    httpSvr := ginx.NewServer(siteConf)
    httpSvr.PostInit(func(r *gin.Engine) error {
        r.Use(ginx.NewHandler(middleware.WithTrace))
        g1 := r.Group("")
        aBucket := ginx.NewBucket(
            g1,
            authHandler,
            publicHandler,
            // any handlers that are public add here
        )
        // 以下handler需要进行权限验证
        g2 := r.Group("")
        g2.Use(
            ginx.NewHandler(middleware.CheckAccess),
        )
        bBucket := ginx.NewBucket(
            g2,
            // any handlers that should your login first add here
        )
        aBucket.AddBucket(bBucket)
        aBucket.Register()
        // 返回结果
        return nil
    })
    // 返回结果
    return httpSvr
}
