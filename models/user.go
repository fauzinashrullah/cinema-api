package models

type User struct {
    ID			string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name		string `gorm:"not null" json:"name"`
    Email		string `gorm:"unique;not null" json:"email"`
    Password	string `gorm:"not null" json:"password"`
	CreatedAt	string `gorm:"autoCreateTime" json:"created_at"`
}
