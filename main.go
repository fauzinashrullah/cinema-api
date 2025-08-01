package main

import (
	"github.com/fauzinashrullah/cinema-api/config"
	"github.com/fauzinashrullah/cinema-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

    r := gin.Default()
    routes.RegisterRoutes(r)

    r.Run(":8080")
}
