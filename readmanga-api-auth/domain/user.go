package domain

import "time"

type User struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Nickname      string         `json:"nickname"`
	Avatar        string         `json:"avatar"`
	CreatedAt     time.Time      `json:"created_at"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:To"`
}

type Notification struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	To      int    `json:"to"`
	From    int    `json:"from"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Link    string `json:"link"`
	View    bool   `json:"view"`
}
