package test_util

import (
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
)

func LoadEnv() {
	// テスト実行場所からの相対パスが基本となるが、実行者の都合によって変わるため、
	// 本ファイルからの相対パスを常に取得するようにすることで対処する。
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot get current file info")
	}

	currentDir := filepath.Dir(filename)

	filePath := filepath.Join(currentDir, "../../../envs", ".env.test")

	err := godotenv.Load(filePath)
	if err != nil {
		panic("Error loading .env file")
	}
}
