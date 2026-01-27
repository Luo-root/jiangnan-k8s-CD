package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const KeyPath = "public/key.txt"

func WriteFile(filePath string, content string) error {
	// 1. 确保目录存在
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("【创建目录失败】: %w", err)
	}
	// 2. 再写文件
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("【写入文件失败】: %v", err)
	}
	return nil
}

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("【读取文件失败】: %v", err)
	}
	return string(content), nil
}
