package routes

import (
	"github.com/fauzinashrullah/cinema-api/controllers"
	"github.com/fauzinashrullah/cinema-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)

    schedule := r.Group("/api/schedules")
    schedule.Use(middleware.RequireAuth())
    {
        schedule.GET("", controllers.GetSchedules)
        schedule.POST("", controllers.CreateSchedule)
        schedule.PUT(":id", controllers.UpdateSchedule)
        schedule.DELETE(":id", controllers.DeleteSchedule)
    }
}