package tools

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
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

// IsValidEthereumAddress 检查字符串是否为有效的以太坊地址
func IsValidEthereumAddress(address string) bool {
	// 1. 基本格式检查
	if !isValidEthereumAddressFormat(address) {
		return false
	}

	// 2. 如果是带校验和的地址，验证校验和
	if hasChecksum(address) {
		return verifyChecksum(address)
	}

	return true
}

// isValidEthereumAddressFormat 检查基本地址格式
func isValidEthereumAddressFormat(address string) bool {
	// 去除可能的0x前缀
	addr := strings.ToLower(address)
	if strings.HasPrefix(addr, "0x") {
		addr = addr[2:]
	}

	// 长度检查（40个十六进制字符）
	if len(addr) != 40 {
		return false
	}

	// 正则表达式检查是否为有效的十六进制字符串
	matched, _ := regexp.MatchString("^[0-9a-f]{40}$", addr)
	return matched
}

// hasChecksum 检查地址是否包含校验和（大小写混合）
func hasChecksum(address string) bool {
	// 如果地址全小写或全大写，则没有校验和
	return address != strings.ToLower(address) && address != strings.ToUpper(address)
}

// verifyChecksum 验证以太坊地址的校验和
func verifyChecksum(address string) bool {
	// 去除0x前缀
	addr := address[2:]

	// 计算地址的Keccak256哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(strings.ToLower(addr)))
	hashBytes := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	// 验证每个字符的大小写
	for i := 0; i < len(addr); i++ {
		// 哈希的对应位
		hashPos := i
		if i >= len(hashHex) {
			hashPos = i % len(hashHex)
		}
		hashChar := hashHex[hashPos]

		// 当前地址字符
		addrChar := addr[i]

		// 根据哈希值决定字符是否应该大写
		if (hashChar >= '8' && hashChar <= 'f') && addrChar >= 'a' && addrChar <= 'f' {
			// 哈希位大于等于8，字符应为大写但实际是小写
			return false
		} else if (hashChar >= '0' && hashChar <= '7') && addrChar >= 'A' && addrChar <= 'F' {
			// 哈希位小于8，字符应为小写但实际是大写
			return false
		}
	}

	return true
}
