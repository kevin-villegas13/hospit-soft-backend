package entity

import (
	"hospit-soft-backend/internal/domain/enums"
	"hospit-soft-backend/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string     `gorm:"type:string;primaryKey;"`
	FirstName string     `gorm:"type:varchar(50)"`
	LastName  string     `gorm:"type:varchar(50)"`
	Phone     string     `gorm:"type:varchar(20)"`
	Password  string     `gorm:"type:varchar(255)"`
	Email     string     `gorm:"unique"`
	Role      enums.Role `gorm:"type:varchar(20)"`
	Active    bool       `gorm:"default:true"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = utils.GenerateUUID()
	return
}
