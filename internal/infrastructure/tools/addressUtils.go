package tools

import (
	"regexp"
	"strings"
)

// IsValidAddress 验证波场地址是否有效
func IsValidAddress(address string) bool {
	address = strings.TrimSpace(address)

	// 检查基础地址格式 (T开头，34位)
	if len(address) == 34 && strings.HasPrefix(address, "T") {
		return isValidBase58Address(address)
	}

	// 检查十六进制地址格式 (41开头，42位)
	if len(address) == 42 && strings.HasPrefix(address, "41") {
		return isValidHexAddress(address)
	}

	return false
}

// 验证Base58格式地址
func isValidBase58Address(address string) bool {
	// Base58字符集
	base58Alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	// 检查每个字符是否在Base58字符集中
	for _, c := range address {
		if !strings.ContainsRune(base58Alphabet, c) {
			return false
		}
	}
	return true
}

// 验证十六进制格式地址
func isValidHexAddress(address string) bool {
	matched, _ := regexp.MatchString(`^41[0-9a-fA-F]{40}$`, address)
	return matched
}
