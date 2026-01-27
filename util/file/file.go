package file

import (
	"fmt"
	"os"
)

const KeyPath = "public/key.txt"

func WriteFile(filePath string, content string) error {
	// 0644
	err := os.WriteFile(filePath, []byte(content), 0644)
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
