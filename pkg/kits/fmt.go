package kits

import (
    "fmt"
    "math"
    "strings"

    "github.com/whencome/goutil/timeutil"
)

// FormatDuration 格式化时长信息
func FormatDuration(d float64) string {
    if d <= 0 {
        return "00:00"
    }
    m := math.Floor(d / 60)
    s := d - m
    return fmt.Sprintf("%d:%.2f", int64(m), s)
}

// MoneyY2C 将元转换成分(注意精度问题)
func MoneyY2C(y float64) int64 {
    return int64(math.Round(y * 100))
}

// MoneyC2Y 将分转换成元(注意精度问题)
func MoneyC2Y(c int64) float64 {
    return float64(c) / 100
}

// FormatMoney 将分转换成元(注意精度问题), 并转换成字符串
func FormatMoney(c int64) string {
    return fmt.Sprintf("%.2f", float64(c)/100)
}

// FormatShortMoney 将分转换成元,保留最少的小数位数
func FormatShortMoney(c int64) string {
    m := FormatMoney(c)
    m = strings.TrimRight(m, ".0")
    return m
}

// FormatDateTime 将unix时间戳格式化日期时间格式
func FormatDateTime(t int64) string {
    return timeutil.DateTimeFromUnixTime(t)
}

// FormatDate 将unix时间戳格式化日期格式
func FormatDate(t int64) string {
    return timeutil.DateFromUnixTime(t)
}
