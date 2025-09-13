package routes

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nathgoh/investment-trader/alpaca/internal/utils"
)

func Handler(ctx context.Context) *gin.Engine {

	// Router setup with CORS middleware
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8501",
	}
	router.Use(cors.New(config))

	router.GET(utils.API_URL_PATH+"/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
