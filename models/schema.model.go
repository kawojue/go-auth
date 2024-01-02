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
	RefreshToken []byte    `gorm:"type:bytea"`
}
