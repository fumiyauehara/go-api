package model

type Tenant struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(255);not null"`
	Address string `gorm:"type:varchar(255)"`
	Tel     string `gorm:"type:varchar(20)"`
}
