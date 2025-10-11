package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/gin-gonic/gin"
	"github.com/nathgoh/investment-trader/alpaca/internal/trading"
	"github.com/shopspring/decimal"
)

// PlaceOrderRequest represents the request body for placing an order
type PlaceOrderRequest struct {
	Symbol      string  `json:"symbol" binding:"required"`
	Qty         float64 `json:"qty" binding:"required"`
	Side        string  `json:"side" binding:"required"` // "buy" or "sell"
	Type        string  `json:"type" binding:"required"` // "market", "limit", "stop", "stop_limit"
	TimeInForce string  `json:"time_in_force" binding:"required"` // "day", "gtc", "opg", "cls", "ioc", "fok"
	LimitPrice  *float64 `json:"limit_price,omitempty"`
	StopPrice   *float64 `json:"stop_price,omitempty"`
	IsPaper     bool    `json:"is_paper"`
}

// PlaceOrder handles placing a new order
func PlaceOrder(c *gin.Context) {
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert qty to decimal
	qty := decimal.NewFromFloat(req.Qty)

	// Convert side
	var side alpaca.Side
	switch strings.ToLower(req.Side) {
	case "buy":
		side = alpaca.Buy
	case "sell":
		side = alpaca.Sell
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid side, must be 'buy' or 'sell'"})
		return
	}

	// Convert order type
	var orderType alpaca.OrderType
	switch strings.ToLower(req.Type) {
	case "market":
		orderType = alpaca.Market
	case "limit":
		orderType = alpaca.Limit
	case "stop":
		orderType = alpaca.Stop
	case "stop_limit":
		orderType = alpaca.StopLimit
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order type"})
		return
	}

	// Convert time in force
	var timeInForce alpaca.TimeInForce
	switch strings.ToLower(req.TimeInForce) {
	case "day":
		timeInForce = alpaca.Day
	case "gtc":
		timeInForce = alpaca.GTC
	case "opg":
		timeInForce = alpaca.OPG
	case "cls":
		timeInForce = alpaca.CLS
	case "ioc":
		timeInForce = alpaca.IOC
	case "fok":
		timeInForce = alpaca.FOK
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid time_in_force"})
		return
	}

	// Convert prices to decimal if provided
	var limitPrice, stopPrice *decimal.Decimal
	if req.LimitPrice != nil {
		lp := decimal.NewFromFloat(*req.LimitPrice)
		limitPrice = &lp
	}
	if req.StopPrice != nil {
		sp := decimal.NewFromFloat(*req.StopPrice)
		stopPrice = &sp
	}

	order, err := trading.PlaceOrder(req.IsPaper, req.Symbol, qty, side, orderType, timeInForce, limitPrice, stopPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrders retrieves orders with optional filters
func GetOrders(c *gin.Context) {
	isPaper := c.Query("is_paper") == "true"
	
	// Parse query parameters
	var status, direction, symbols *string
	var limit *int
	var after, until *time.Time
	var nested *bool

	if s := c.Query("status"); s != "" {
		status = &s
	}
	if d := c.Query("direction"); d != "" {
		direction = &d
	}
	if sym := c.Query("symbols"); sym != "" {
		symbols = &sym
	}
	if l := c.Query("limit"); l != "" {
		if limitInt, err := strconv.Atoi(l); err == nil {
			limit = &limitInt
		}
	}
	if n := c.Query("nested"); n != "" {
		if nestedBool, err := strconv.ParseBool(n); err == nil {
			nested = &nestedBool
		}
	}
	if a := c.Query("after"); a != "" {
		if afterTime, err := time.Parse(time.RFC3339, a); err == nil {
			after = &afterTime
		}
	}
	if u := c.Query("until"); u != "" {
		if untilTime, err := time.Parse(time.RFC3339, u); err == nil {
			until = &untilTime
		}
	}

	orders, err := trading.GetOrders(isPaper, status, limit, after, until, direction, nested, symbols)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder retrieves a single order by ID
func GetOrder(c *gin.Context) {
	orderID := c.Param("id")
	isPaper := c.Query("is_paper") == "true"
	nested := c.Query("nested") == "true"

	order, err := trading.GetOrder(isPaper, orderID, nested)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CancelOrder cancels an order by ID
func CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	isPaper := c.Query("is_paper") == "true"

	err := trading.CancelOrder(isPaper, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})
}

// CancelAllOrders cancels all open orders
func CancelAllOrders(c *gin.Context) {
	isPaper := c.Query("is_paper") == "true"

	err := trading.CancelAllOrders(isPaper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all orders cancelled successfully"})
}

// GetPositions retrieves all positions
func GetPositions(c *gin.Context) {
	isPaper := c.Query("is_paper") == "true"

	positions, err := trading.GetPositions(isPaper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, positions)
}

// GetPosition retrieves a single position by symbol
func GetPosition(c *gin.Context) {
	symbol := c.Param("symbol")
	isPaper := c.Query("is_paper") == "true"

	position, err := trading.GetPosition(isPaper, symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, position)
}

// ClosePositionRequest represents the request body for closing a position
type ClosePositionRequest struct {
	Qty        *float64 `json:"qty,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
	IsPaper    bool     `json:"is_paper"`
}

// ClosePosition closes a position for a symbol
func ClosePosition(c *gin.Context) {
	symbol := c.Param("symbol")
	
	var req ClosePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var qty, percentage *decimal.Decimal
	if req.Qty != nil {
		q := decimal.NewFromFloat(*req.Qty)
		qty = &q
	}
	if req.Percentage != nil {
		p := decimal.NewFromFloat(*req.Percentage)
		percentage = &p
	}

	order, err := trading.ClosePosition(req.IsPaper, symbol, qty, percentage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CloseAllPositions closes all positions
func CloseAllPositions(c *gin.Context) {
	isPaper := c.Query("is_paper") == "true"
	cancelOrders := c.Query("cancel_orders") == "true"

	responses, err := trading.CloseAllPositions(isPaper, cancelOrders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// GetAssets retrieves all assets
func GetAssets(c *gin.Context) {
	var status, assetClass *string

	if s := c.Query("status"); s != "" {
		status = &s
	}
	if ac := c.Query("asset_class"); ac != "" {
		assetClass = &ac
	}

	assets, err := trading.GetAssets(status, assetClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assets)
}

// GetAsset retrieves a single asset by symbol
func GetAsset(c *gin.Context) {
	symbol := c.Param("symbol")

	asset, err := trading.GetAsset(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}

// GetClock retrieves the market clock
func GetClock(c *gin.Context) {
	clock, err := trading.GetClock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clock)
}

// GetCalendar retrieves the market calendar
func GetCalendar(c *gin.Context) {
	var start, end *time.Time

	if s := c.Query("start"); s != "" {
		if startTime, err := time.Parse("2006-01-02", s); err == nil {
			start = &startTime
		}
	}
	if e := c.Query("end"); e != "" {
		if endTime, err := time.Parse("2006-01-02", e); err == nil {
			end = &endTime
		}
	}

	calendar, err := trading.GetCalendar(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, calendar)
}

// Legacy handlers for account endpoints (using net/http)
func GetPaperAccountGin(c *gin.Context) {
	client := trading.GetClient(true)
	account, err := client.GetAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func GetLiveAccountGin(c *gin.Context) {
	client := trading.GetClient(false)
	account, err := client.GetAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

// GetStockQuoteGin is a Gin wrapper for the existing GetStockQuote handler
func GetStockQuoteGin(c *gin.Context) {
	// Create a ResponseRecorder to capture the response
	w := &ginResponseWriter{c: c}
	GetStockQuote(w, c.Request)
}

// ginResponseWriter wraps gin.Context to implement http.ResponseWriter
type ginResponseWriter struct {
	c *gin.Context
	statusCode int
	written bool
}

func (w *ginResponseWriter) Header() http.Header {
	return w.c.Writer.Header()
}

func (w *ginResponseWriter) Write(data []byte) (int, error) {
	if !w.written {
		w.written = true
		if w.statusCode == 0 {
			w.statusCode = http.StatusOK
		}
		w.c.Status(w.statusCode)
	}
	
	// Parse JSON and re-encode through Gin for consistency
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err == nil {
		w.c.JSON(w.statusCode, jsonData)
		return len(data), nil
	}
	
	return w.c.Writer.Write(data)
}

func (w *ginResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
