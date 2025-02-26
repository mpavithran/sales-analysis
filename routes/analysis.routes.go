package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mpavithran/sales-analysis/controllers"
)

func AnalysisRoutes(router *gin.Engine, analysisController *controllers.AnalysisController) {
	analysisGroup := router.Group("/analysis")
	{
		analysisGroup.POST("/upload-csv", analysisController.UploadCSV)
		analysisGroup.GET("/revenue", analysisController.GetRevenue)
		analysisGroup.GET("/top-products", analysisController.TopProducts)
	}
}
