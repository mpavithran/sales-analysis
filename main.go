package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/mpavithran/sales-analysis/config"
	"github.com/mpavithran/sales-analysis/controllers"
	"github.com/mpavithran/sales-analysis/models"
	"github.com/mpavithran/sales-analysis/repositories"
	"github.com/mpavithran/sales-analysis/routes"
	"github.com/mpavithran/sales-analysis/services"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Order{})

	analysisRepo := repositories.NewAnalysisRepository(config.DB)
	analysisService := services.NewAnalysisService(analysisRepo)
	analysisController := controllers.NewAnalysisController(analysisService)

	r := gin.Default()
	routes.AnalysisRoutes(r, analysisController)

	fmt.Println("Server started on :8080")
	r.Run(":8080")
}
