package controllers

import (
	"net/http"
	"time"

	"github.com/fauzinashrullah/cinema-api/config"
	"github.com/fauzinashrullah/cinema-api/models"

	"github.com/gin-gonic/gin"
)

func GetSchedules(c *gin.Context) {
    var schedules []models.Schedule
    config.DB.Preload("Film").Preload("Theater").Find(&schedules)
    c.JSON(http.StatusOK, schedules)
}

type ScheduleInput struct {
    FilmID    string    `json:"film_id"`
    TheaterID string    `json:"theater_id"`
    ShowTime  time.Time `json:"show_time"`
}

func CreateSchedule(c *gin.Context) {
    var input ScheduleInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

	var film models.Film
    if err := config.DB.First(&film, "id = ?", input.FilmID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
        return
    }

    var theater models.Theater
    if err := config.DB.First(&theater, "id = ?", input.TheaterID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Theater not found"})
        return
    }
	
    schedule := models.Schedule{
        FilmID:    input.FilmID,
        TheaterID: input.TheaterID,
        ShowTime:  input.ShowTime,
		Film:      film,
        Theater:   theater,
    }
    config.DB.Create(&schedule)

	config.DB.Preload("Film").Preload("Theater").First(&schedule, "id = ?", schedule.ID)
    c.JSON(http.StatusCreated, schedule)
}

func UpdateSchedule(c *gin.Context) {
    id := c.Param("id")
    var schedule models.Schedule
    if err := config.DB.First(&schedule, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
        return
    }

    var input ScheduleInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    schedule.FilmID = input.FilmID
    schedule.TheaterID = input.TheaterID
    schedule.ShowTime = input.ShowTime
    config.DB.Save(&schedule)

	config.DB.Preload("Film").Preload("Theater").First(&schedule, "id = ?", id)
    c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.Schedule{}, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}