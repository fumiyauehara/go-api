package middleware_test

import (
	"github.com/fumiyauehara/go-api/internal/api/test_util"
	"gorm.io/gorm"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

func setup() {
	test_util.LoadEnv()
	db = test_util.MysqlTestDBConn()
}

func teardown() {
}
