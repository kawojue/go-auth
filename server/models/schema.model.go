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
	TOTP         TOTP     `gorm:"foreignKey:UserID"`
	Profile      Profiles `gorm:"foreignKey:UserID"`
}

type Profiles struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string
	LastName  string
	Age       int
	Avatar    Avatars   `gorm:"foreignKey:UserID"`
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}

type TOTP struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Otp       string
	OtpExpiry string
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}

type Avatars struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Url    string
	Path   string
	Type   string
	UserId uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}
