package model

import "gorm.io/gorm"

type DBConn struct {
	Writer *gorm.DB
	Reader *gorm.DB
}
