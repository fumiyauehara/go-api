package api

import (
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBConn(writer model.DBConfig, reader model.DBConfig) model.DBConn {
	common := func(conf model.DBConfig) *gorm.DB {
		dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
			conf.DBUser,
			conf.DBPw,
			conf.DBHost,
			conf.DBPort,
			conf.DBName)

		var driver gorm.Dialector
		switch conf.Type {
		case "mysql":
			driver = mysql.Open(dsn)
		case "postgres":
			driver = postgres.Open(dsn)
		default:
			panic("unknown db type")
		}

		db, err := gorm.Open(driver, &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		return db
	}

	return model.DBConn{
		Writer: common(writer),
		Reader: common(reader),
	}
}
