package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Game struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Score struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	GameID    uint      `json:"game_id"`
	Game      Game      `gorm:"foreignKey:GameID" json:"game"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
