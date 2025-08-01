package models

import "time"

type Schedule struct {
    ID			string		`gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    FilmID		string		`gorm:"not null"`
    TheaterID	string		`gorm:"not null"`
    ShowTime	time.Time	`gorm:"not null"`

	Film		Film		`gorm:"foreignKey:FilmID"`
    Theater		Theater		`gorm:"foreignKey:TheaterID"`
}