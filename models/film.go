package models

type Film struct {
    ID			string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Title		string `gorm:"not null"`
    Duration	int    `gorm:"not null"`
    Description	string
}
