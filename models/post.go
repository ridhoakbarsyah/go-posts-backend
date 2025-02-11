package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(200)"`
	Content     string    `gorm:"type:text"`
	Category    string    `gorm:"type:varchar(100)"`
	CreatedDate time.Time `gorm:"type:timestamp"`
	UpdatedDate time.Time `gorm:"type:timestamp"`
	Status      string    `gorm:"type:varchar(100)"`
}
