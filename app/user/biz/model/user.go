package model

// 定义用户模型
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique_index"`
	PasswordHashed string `gorm:"type:varchar(255) not null`
}

func (User) TableName() string {
	return "user"
}
