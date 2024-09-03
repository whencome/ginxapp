package trace

import (
    "context"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/whencome/goutil"
)

const (
    // TraceIdKey 全局追踪ID key
    TraceIdKey string = "x-trace-id"
)

// newTraceId create a new trace id
func newTraceId() string {
    return uuid.New().String()
}

// Context get a context with trace_id
func Context() context.Context {
    return context.WithValue(context.Background(), TraceIdKey, newTraceId())
}

// TraceId get trace id from context
func TraceId(ctx context.Context) string {
    if ctx == nil {
        return ""
    }
    v := ctx.Value(TraceIdKey)
    if v == nil {
        return ""
    }
    return goutil.String(v)
}

// MustTraceId get trace id from context, if no trace id were found, create a new one
func MustTraceId(ctx context.Context) string {
    traceId := TraceId(ctx)
    if traceId != "" {
        return traceId
    }
    return newTraceId()
}

// Wrap check whether the ctx has trace_id or not, if not, set one
func Wrap(ctx context.Context) context.Context {
    traceId := TraceId(ctx)
    if traceId != "" {
        return ctx
    }
    return context.WithValue(ctx, TraceIdKey, newTraceId())
}

// WrapGin set a trace id to gin context
func WrapGin(c *gin.Context) context.Context {
    traceId := TraceId(c.Request.Context())
    if traceId == "" {
        // 检查头信息
        traceId = c.GetHeader(TraceIdKey)
    }
    // 没有头信息，创建一个信息
    if traceId == "" {
        traceId = newTraceId()
    }
    ctx := context.WithValue(c.Request.Context(), TraceIdKey, newTraceId())
    c.Request = c.Request.WithContext(ctx)
    return ctx
}

// GinTraceId get trace id from gin context
func GinTraceId(c *gin.Context) string {
    return MustTraceId(c.Request.Context())
}
