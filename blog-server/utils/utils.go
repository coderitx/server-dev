package utils

import (
	"go.uber.org/zap"
	"os"
)

// IsExists 文件夹是否存在
func IsExists(dir string) error {
	_, err := os.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			zap.S().Errorf("创建文件夹: %v 失败，error: %v", dir, err.Error())
			return err
		}
	}
	return nil
}
