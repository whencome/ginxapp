// Package validator 定义常用数据的验证方法
package validator

import (
    "regexp"
    "strconv"
    "strings"
    "time"
)

var (
    // 手机号的正则表达式
    mobilePattern = `^1\d{10}$`
    // 日期验证正则表达式
    datePattern = `^\d{4}\-\d{1,2}\-\d{1,2}`
    // http地址正则表达式
    httpUrlPattern = `^(http|https):\/\/[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
)

// IsMobile 判断给定内容是否是手机号
func IsMobile(s string) bool {
    if s == "" {
        return false
    }
    match, _ := regexp.MatchString(mobilePattern, s)
    return match
}

// IsHttpUrl 判断给定的值是否时有效的http地址
func IsHttpUrl(s string) bool {
    if s == "" {
        return false
    }
    match, _ := regexp.MatchString(httpUrlPattern, s)
    return match
}

// IsIdCardNo 判断给定的内容是否是有效的身份证号
func IsIdCardNo(id string) bool {
    // 首先检查长度是否为18位
    if len(id) != 18 {
        return false
    }
    // 检查前17位是否全为数字
    for _, digit := range id[:17] {
        if !strings.ContainsRune("0123456789", rune(digit)) {
            return false
        }
    }
    // 计算最后一位校验码
    weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
    sum := 0
    for i, digit := range id[:17] {
        parsedDigit, _ := strconv.Atoi(string(digit))
        sum += parsedDigit * weights[i]
    }

    mods := sum % 11
    checkCodes := "10X98765432"
    // 检查最后一位校验码是否正确
    lastChar := id[17]
    if lastChar == 'X' || lastChar == 'x' {
        return checkCodes[mods] == 'X'
    } else {
        lastParsed, _ := strconv.Atoi(string(lastChar))
        return lastParsed == int(checkCodes[mods]-'0')
    }
}

// IsDate 判断给定内容是否是有效的日期格式
func IsDate(s string) bool {
    if s == "" {
        return false
    }
    if match, err := regexp.MatchString(datePattern, s); err != nil || !match {
        return false
    }
    _, err := time.Parse("2006-01-02", s)
    if err != nil {
        return false
    }
    return true
}
