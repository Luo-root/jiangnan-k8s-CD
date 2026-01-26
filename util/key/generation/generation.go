package generation

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		// rand.Int(rand.Reader, big.NewInt(n)) - 生成 [0, n) 的随机数]
		// rand.Reader - 随机数生成器 从操作系统获取密码学安全的随机数据
		// big.NewInt(n) - 设置上限
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("【生成密钥失败】: %v", err)
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}
