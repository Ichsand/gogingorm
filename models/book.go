package models

import "time"

type Book struct {
	ID        string `gorm:"primary_key" json:"id"`
	Title     string `gorm:"type:varchar(255);NOT NULL" json:"title" binding:"required"`
	Author    string `gorm:"type:varchar(255); NOT NULL" json:"author"`
	Publisher string `gorm:"type:varchar(255)" json:"publisher"`
	Year      int    `gorm:"type:int(11); NOT NULL" json:"year"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Books []Book
