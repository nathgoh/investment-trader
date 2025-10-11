package routes

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nathgoh/investment-trader/alpaca/api/handlers"
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

	// Health check
	router.GET(utils.API_URL_PATH+"/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Account endpoints
	router.GET(utils.API_URL_PATH+"/account/paper", handlers.GetPaperAccountGin)
	router.GET(utils.API_URL_PATH+"/account/live", handlers.GetLiveAccountGin)

	// Market data endpoints
	router.GET(utils.API_URL_PATH+"/marketdata/quotes/:symbol", handlers.GetStockQuoteGin)

	// Trading - Order endpoints
	router.POST(utils.API_URL_PATH+"/orders", handlers.PlaceOrder)
	router.GET(utils.API_URL_PATH+"/orders", handlers.GetOrders)
	router.GET(utils.API_URL_PATH+"/orders/:id", handlers.GetOrder)
	router.DELETE(utils.API_URL_PATH+"/orders/:id", handlers.CancelOrder)
	router.DELETE(utils.API_URL_PATH+"/orders", handlers.CancelAllOrders)

	// Trading - Position endpoints
	router.GET(utils.API_URL_PATH+"/positions", handlers.GetPositions)
	router.GET(utils.API_URL_PATH+"/positions/:symbol", handlers.GetPosition)
	router.DELETE(utils.API_URL_PATH+"/positions/:symbol", handlers.ClosePosition)
	router.DELETE(utils.API_URL_PATH+"/positions", handlers.CloseAllPositions)

	// Asset endpoints
	router.GET(utils.API_URL_PATH+"/assets", handlers.GetAssets)
	router.GET(utils.API_URL_PATH+"/assets/:symbol", handlers.GetAsset)

	// Market info endpoints
	router.GET(utils.API_URL_PATH+"/clock", handlers.GetClock)
	router.GET(utils.API_URL_PATH+"/calendar", handlers.GetCalendar)

	return router
}
