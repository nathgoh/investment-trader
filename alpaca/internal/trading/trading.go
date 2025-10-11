package trading

import (
	"log"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

var (
	paperClient *alpaca.Client
	liveClient  *alpaca.Client
)

// Initialize clients
func init() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Initialize paper trading client
	paperAPIKey := os.Getenv("ALPACA_PAPER_API_KEY")
	paperAPISecret := os.Getenv("ALPACA_PAPER_SECRET_KEY")
	if paperAPIKey != "" && paperAPISecret != "" {
		paperClient = alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    paperAPIKey,
			APISecret: paperAPISecret,
			BaseURL:   "https://paper-api.alpaca.markets",
		})
	}

	// Initialize live trading client
	liveAPIKey := os.Getenv("ALPACA_LIVE_API_KEY")
	liveAPISecret := os.Getenv("ALPACA_LIVE_API_SECRET_KEY")
	if liveAPIKey != "" && liveAPISecret != "" {
		liveClient = alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    liveAPIKey,
			APISecret: liveAPISecret,
			BaseURL:   "https://api.alpaca.markets",
		})
	}
}

// GetClient returns the appropriate client based on the isPaper flag
func GetClient(isPaper bool) *alpaca.Client {
	if isPaper {
		return paperClient
	}
	return liveClient
}

// PlaceOrder places a new order
func PlaceOrder(isPaper bool, symbol string, qty decimal.Decimal, side alpaca.Side, orderType alpaca.OrderType, timeInForce alpaca.TimeInForce, limitPrice, stopPrice *decimal.Decimal) (*alpaca.Order, error) {
	client := GetClient(isPaper)
	
	req := alpaca.PlaceOrderRequest{
		Symbol:      symbol,
		Qty:         &qty,
		Side:        side,
		Type:        orderType,
		TimeInForce: timeInForce,
	}

	if limitPrice != nil {
		req.LimitPrice = limitPrice
	}
	if stopPrice != nil {
		req.StopPrice = stopPrice
	}

	order, err := client.PlaceOrder(req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrders retrieves orders with optional filters
func GetOrders(isPaper bool, status *string, limit *int, after, until *time.Time, direction *string, nested *bool, symbols *string) ([]alpaca.Order, error) {
	client := GetClient(isPaper)
	
	req := alpaca.GetOrdersRequest{}
	
	if status != nil {
		req.Status = *status
	}
	if limit != nil {
		req.Limit = *limit
	}
	if after != nil {
		req.After = *after
	}
	if until != nil {
		req.Until = *until
	}
	if direction != nil {
		req.Direction = *direction
	}
	if nested != nil {
		req.Nested = *nested
	}
	// Note: Symbols filter not supported in this version

	orders, err := client.GetOrders(req)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrder retrieves a single order by ID
func GetOrder(isPaper bool, orderID string, nested bool) (*alpaca.Order, error) {
	client := GetClient(isPaper)
	
	order, err := client.GetOrder(orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CancelOrder cancels an order by ID
func CancelOrder(isPaper bool, orderID string) error {
	client := GetClient(isPaper)
	
	err := client.CancelOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}

// CancelAllOrders cancels all open orders
func CancelAllOrders(isPaper bool) error {
	client := GetClient(isPaper)
	
	err := client.CancelAllOrders()
	if err != nil {
		return err
	}

	return nil
}

// GetPositions retrieves all positions
func GetPositions(isPaper bool) ([]alpaca.Position, error) {
	client := GetClient(isPaper)
	
	positions, err := client.GetPositions()
	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetPosition retrieves a single position by symbol
func GetPosition(isPaper bool, symbol string) (*alpaca.Position, error) {
	client := GetClient(isPaper)
	
	position, err := client.GetPosition(symbol)
	if err != nil {
		return nil, err
	}

	return position, nil
}

// ClosePosition closes a position for a symbol
func ClosePosition(isPaper bool, symbol string, qty *decimal.Decimal, percentage *decimal.Decimal) (*alpaca.Order, error) {
	client := GetClient(isPaper)
	
	req := alpaca.ClosePositionRequest{}
	if qty != nil {
		req.Qty = *qty
	}
	if percentage != nil {
		req.Percentage = *percentage
	}

	order, err := client.ClosePosition(symbol, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CloseAllPositions closes all positions
func CloseAllPositions(isPaper bool, cancelOrders bool) ([]alpaca.Order, error) {
	client := GetClient(isPaper)
	
	req := alpaca.CloseAllPositionsRequest{
		CancelOrders: cancelOrders,
	}
	
	responses, err := client.CloseAllPositions(req)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

// GetAssets retrieves all assets
func GetAssets(status, assetClass *string) ([]alpaca.Asset, error) {
	// Use paper client for asset queries (same for both)
	client := paperClient
	if client == nil {
		client = liveClient
	}
	
	req := alpaca.GetAssetsRequest{}
	if status != nil {
		req.Status = *status
	}
	if assetClass != nil {
		req.AssetClass = *assetClass
	}

	assets, err := client.GetAssets(req)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

// GetAsset retrieves a single asset by symbol
func GetAsset(symbol string) (*alpaca.Asset, error) {
	// Use paper client for asset queries (same for both)
	client := paperClient
	if client == nil {
		client = liveClient
	}
	
	asset, err := client.GetAsset(symbol)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

// GetClock retrieves the market clock
func GetClock() (*alpaca.Clock, error) {
	// Use paper client for clock queries (same for both)
	client := paperClient
	if client == nil {
		client = liveClient
	}
	
	clock, err := client.GetClock()
	if err != nil {
		return nil, err
	}

	return clock, nil
}

// GetCalendar retrieves the market calendar
func GetCalendar(start, end *time.Time) ([]alpaca.CalendarDay, error) {
	// Use paper client for calendar queries (same for both)
	client := paperClient
	if client == nil {
		client = liveClient
	}
	
	req := alpaca.GetCalendarRequest{}
	if start != nil {
		req.Start = *start
	}
	if end != nil {
		req.End = *end
	}

	calendar, err := client.GetCalendar(req)
	if err != nil {
		return nil, err
	}

	return calendar, nil
}
