//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "github.com/whencome/ginxapp/internal/biz"
    "github.com/whencome/ginxapp/internal/data"
    "github.com/whencome/ginxapp/internal/etc"
    "github.com/whencome/ginxapp/internal/handler"
    "github.com/whencome/ginxapp/internal/server"
    "github.com/whencome/goutil/log"

    "github.com/google/wire"
    "github.com/whencome/ginx"
)

// wireApp init app
func wireApp(logger log.Logger, site *ginx.ServerOptions, redisConf *etc.RedisConfig) (*App, func(), error) {
    panic(wire.Build(data.ProviderSet, biz.ProviderSet, handler.ProviderSet, server.ProviderSet, NewApp))
}
