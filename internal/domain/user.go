package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	UserRoleUnspecified Role = "UNSPECIFIED"
	UserRoleUser        Role = "USER"
	UserRoleAdmin       Role = "ADMIN"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email     string    `gorm:"unique;not null"`
	PassHash  []byte    `gorm:"not null"`
	Role      Role      `gorm:"not null"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdateAt  time.Time `gorm:"updated_at"`
}
