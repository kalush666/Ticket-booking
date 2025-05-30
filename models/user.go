package models

import "gorm.io/gorm"

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User struct {
	ID        uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string   `json:"email" gorm:"uniqueIndex;not null"`
	Role      UserRole `json:"role" gorm:"text;default:'attendee';not null"`
	Password  string   `json:"-" gorm:"not null"`
	CreatedAt string   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt string   `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) AfterCreate(db *gorm.DB) error{
	if u.ID == 1{
		db.Model(u).Update("role", Manager)
	}
	return nil
}