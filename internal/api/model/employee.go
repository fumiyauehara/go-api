package model

type RequestEmployee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Employee struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;column:email"` // カラム名を明示的に指定
	TenantID int    `gorm:"column:tenant_id"`                        // 外部キーのカラム名
	Tenant   Tenant `gorm:"foreignKey:TenantID"`                     // テナントモデルへの関連
}

type ViewEmployee struct {
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	TenantID int    `gorm:"column:tenant_id"`
}
