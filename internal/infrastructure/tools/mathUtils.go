package tools

import (
	"fmt"
	"strconv"
	"time"
)

func SubtractStringNumbers(a, b string, n float64) (string, error) {
	// 1. 将字符串转为 float64
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "", fmt.Errorf("转换 %s 失败: %v", a, err)
	}

	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "", fmt.Errorf("转换 %s 失败: %v", b, err)
	}

	// 2. 计算减法
	result := numA - numB*n

	// 3. 将结果转为字符串
	return fmt.Sprintf("%v", result), nil
}
func CompareStringsWithFloat(a, b string, n float64) bool {
	// 将字符串转换为 float64
	floatA, errA := strconv.ParseFloat(a, 64)
	floatB, errB := strconv.ParseFloat(b, 64)

	if errA != nil || errB != nil {
		return false
	}

	// 计算 b * 2
	bTimesTwo := floatB * n

	// 比较 a 和 b * 2
	return floatA > bTimesTwo
}
func AddStringsAsFloats(a, b string) string {
	// 1. 将第一个字符串转换成 float64
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}

	// 2. 将第二个字符串转换成 float64
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}

	// 3. 相加并返回结果

	sum := num1 + num2
	amount := fmt.Sprintf("%f", sum)

	return amount[0 : len(amount)-3]
}
func Generate6DigitOrderNo() string {
	// 获取当前时间的秒数(0-59)和纳秒的后4位
	now := time.Now()
	seconds := now.Second()           // 0-59
	nanos := now.Nanosecond() % 10000 // 取纳秒的后4位

	// 组合成6位数: 秒数(2位) + 纳秒后4位
	return fmt.Sprintf("%02d%04d", seconds, nanos/100)
}

func CompareNumberStrings(a, b string) (int, error) {
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number string: %s", a)
	}

	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number string: %s", b)
	}

	if numA < numB {
		return -1, nil
	} else if numA > numB {
		return 1, nil
	}
	return 0, nil
}
