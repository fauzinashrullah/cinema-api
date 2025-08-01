package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fauzinashrullah/cinema-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=cinema port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	DB = db
	fmt.Println("Database connected.")

	db.AutoMigrate(
        &models.User{},
        &models.Film{},
        &models.Theater{},
        &models.Schedule{},
    )

	// Dummy user (password: password123, hashed)
    hashed := "$2a$12$KFtyzpTed5pkV01arCx.auoHN1VT6lwASx2y8d/KlgiJrtdY.Fw9y"
    db.FirstOrCreate(&models.User{}, models.User{
        Email:    "admin@example.com",
        Name:     "Admin",
        Password: hashed,
    })

    // Dummy film
    var film models.Film
    db.FirstOrCreate(&film, models.Film{
        Title:       "Avengers",
        Duration:    150,
        Description: "Superhero film",
    })

    // Dummy theater
    var theater models.Theater
    db.FirstOrCreate(&theater, models.Theater{
        Name: "XXI Tunjungan",
        City: "Surabaya",
    })

    // Dummy schedule
    db.FirstOrCreate(&models.Schedule{}, models.Schedule{
        FilmID:    film.ID,
        TheaterID: theater.ID,
        ShowTime:  time.Now().Add(24 * time.Hour),
    })
}
