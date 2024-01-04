package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email        string    `gorm:"uniqueIndex"`
	Username     string    `gorm:"uniqueIndex"`
	Password     string
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	RefreshToken string
	TOTP         TOTP `gorm:"foreignKey:UserID"`
	Info         Info `gorm:"foreignKey:UserID"`
}

type Info struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generator_v4()"`
	FirstName string
	LastName  string
	Age       int
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}

type TOTP struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Otp       string
	OtpExpiry string
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}
