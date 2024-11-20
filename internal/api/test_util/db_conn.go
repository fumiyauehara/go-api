package test_util

import (
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func MysqlTestDBConn() *gorm.DB {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PW"),
		os.Getenv("MYSQL_HOST"),
		util.ConvertStringToInt(os.Getenv("MYSQL_PORT")),
		os.Getenv("MYSQL_DB"))

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Sprintf("failed to connect database: %s\n", err))
	} else {
		return db
	}
}
