package kits

import (
    "database/sql"
    "errors"
    "regexp"
    "strings"

    "github.com/whencome/ginxapp/pkg/xerr"
    "gorm.io/gorm"
)

// IsNoRowsErr 判断是否是查询结果为空的错误
func IsNoRowsErr(err error) bool {
    if errors.Is(err, sql.ErrNoRows) || errors.Is(err, gorm.ErrRecordNotFound) {
        return true
    }
    return false
}

// ErrIgnoreNoRows 检查错误信息，过滤掉查询为空的错误，然后返回错误信息
func ErrIgnoreNoRows(err error) error {
    if err == nil || IsNoRowsErr(err) {
        return nil
    }
    return xerr.Wrap(err)
}

// ErrCheckNoRows 检查错误信息，并根据参数判断是否过滤掉查询为空的错误，然后返回错误信息
func ErrCheckNoRows(err error, ignoreNoRowsErr bool) error {
    if err == nil || ignoreNoRowsErr && IsNoRowsErr(err) {
        return nil
    }
    return err
}

// GetPaymentMethodName 获取支付方式名称
func GetPaymentMethodName(payMethod string) string {
    payMethod = strings.TrimSpace(payMethod)
    if payMethod == "" {
        return "-"
    }
    if m, e := regexp.MatchString(`^(.*)?(Wei(X|x)in|WxProvider).*$`, payMethod); e == nil && m {
        return "微信"
    }
    if m, e := regexp.MatchString(`^Alipay.*$`, payMethod); e == nil && m {
        return "支付宝"
    }
    if payMethod == "Balance" {
        return "余额"
    }
    return "-"
}

// MaskName 对用户姓名打码，此方法同一将第一个字修改为“*”
func MaskName(name string) string {
    runes := []rune(name)
    if len(runes) > 0 {
        runes[0] = '*'
        return string(runes)
    }
    return name
}

// MaskAddress 对地址进行打码处理
func MaskAddress(address string) string {
    runes := []rune(address)
    count := len(runes)
    if count < 10 {
        return address
    }
    mask := make([]rune, 0)
    mask = append(mask, runes[0:6]...)
    mask = append(mask, []rune("****")...)
    mask = append(mask, runes[count-4:count]...)
    return string(mask)
}

// MaskMobileNo 对手机号进行打码处理
func MaskMobileNo(mobile string) string {
    runes := []rune(mobile)
    count := len(runes)
    if count < 10 {
        return mobile
    }
    mask := make([]rune, 0)
    mask = append(mask, runes[0:3]...)
    mask = append(mask, []rune("****")...)
    mask = append(mask, runes[count-4:count]...)
    return string(mask)
}
