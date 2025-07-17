package tools

import (
	"fmt"
	"strconv"
	"time"
)

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
