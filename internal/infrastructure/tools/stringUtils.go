package tools

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

// IsEmpty 检查字符串是否为空
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty 检查字符串是否非空
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Substring 截取字符串
func Substring(s string, start, end int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	if start > end {
		return ""
	}
	return string(runes[start:end])
}

// ContainsAny 检查字符串是否包含任意一个给定的子串
func ContainsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// ContainsAll 检查字符串是否包含所有给定的子串
func ContainsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// EqualsIgnoreCase 忽略大小写比较字符串
func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// RemoveAll 移除字符串中所有指定的子串
func RemoveAll(s, remove string) string {
	return strings.ReplaceAll(s, remove, "")
}

// RemoveAny 移除字符串中任意一个指定的子串
func RemoveAny(s string, removes ...string) string {
	for _, remove := range removes {
		s = strings.ReplaceAll(s, remove, "")
	}
	return s
}

// IsNumeric 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlpha 检查字符串是否只包含字母
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric 检查字符串是否只包含字母和数字
func IsAlphaNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsBlank 检查字符串是否只包含空白字符
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotBlank 检查字符串是否包含非空白字符
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// DefaultIfEmpty 如果字符串为空则返回默认值
func DefaultIfEmpty(s, defaultValue string) string {
	if IsEmpty(s) {
		return defaultValue
	}
	return s
}

// Truncate 截断字符串并在末尾添加后缀（如果需要）
func Truncate(s string, maxLength int, suffix string) string {
	if maxLength <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= maxLength {
		return s
	}
	if len(runes) <= len(suffix) {
		return string(runes[:maxLength])
	}
	return string(runes[:maxLength-len(suffix)]) + suffix
}

// Join 连接字符串切片
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Split 分割字符串
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Capitalize 首字母大写
func Capitalize(s string) string {
	if IsEmpty(s) {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Uncapitalize 首字母小写
func Uncapitalize(s string) string {
	if IsEmpty(s) {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// UpperCase 转换为大写
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// LowerCase 转换为小写
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// CountMatches 统计子串出现的次数
func CountMatches(s, sub string) int {
	return strings.Count(s, sub)
}

// DeleteWhitespace 删除所有空白字符
func DeleteWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// RandomString 生成随机字符串
func RandomString(length int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[num.Int64()]
	}
	return string(result)
}

// IsEmail 检查字符串是否是有效的电子邮件地址
func IsEmail(s string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(s)
}

// IsURL 检查字符串是否是有效的URL
func IsURL(s string) bool {
	pattern := `^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`
	return regexp.MustCompile(pattern).MatchString(s)
}

// LeftPad 左侧填充字符串
func LeftPad(s string, length int, padStr string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(padStr, length-len(s))
	return padding + s
}

// RightPad 右侧填充字符串
func RightPad(s string, length int, padStr string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(padStr, length-len(s))
	return s + padding
}

// Strip 去除字符串两端的指定字符
func Strip(s string, stripChars string) string {
	return strings.Trim(s, stripChars)
}

// StripStart 去除字符串开头的指定字符
func StripStart(s string, stripChars string) string {
	return strings.TrimLeft(s, stripChars)
}

// StripEnd 去除字符串末尾的指定字符
func StripEnd(s string, stripChars string) string {
	return strings.TrimRight(s, stripChars)
}

// Abbreviate 缩写字符串
func Abbreviate(s string, maxWidth int, abbrevMarker string) string {
	if maxWidth <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= maxWidth {
		return s
	}
	if maxWidth <= len(abbrevMarker) {
		return string(runes[:maxWidth])
	}
	return string(runes[:maxWidth-len(abbrevMarker)]) + abbrevMarker
}

// SwapCase 交换字符串的大小写
func SwapCase(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return unicode.ToUpper(r)
	}, s)
}

// Wrap 用指定字符包裹字符串
func Wrap(s string, wrapWith string) string {
	return wrapWith + s + wrapWith
}

// Unwrap 去除包裹字符串的指定字符
func Unwrap(s string, wrapWith string) string {
	if strings.HasPrefix(s, wrapWith) && strings.HasSuffix(s, wrapWith) {
		return s[len(wrapWith) : len(s)-len(wrapWith)]
	}
	return s
}
