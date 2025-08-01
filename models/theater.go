package models

type Theater struct {
    ID		string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name	string `gorm:"not null"`
    City	string `gorm:"not null"`
}