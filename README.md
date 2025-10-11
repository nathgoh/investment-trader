# investment-trader
My attempt at creating an agentic investment bot

# Alpaca Trading API Documentation

This API provides a comprehensive interface to Alpaca's trading platform, supporting both paper and live trading accounts.

## Base URL
```
http://localhost:8080/api/v1
```

## Endpoints

### Health Check
- **GET** `/health`
  - Returns API health status
  - Response: `{"status": "ok"}`

---

## Account Management

### Get Paper Account
- **GET** `/account/paper`
  - Retrieves paper trading account information
  - Response: Account object with balance, equity, buying power, etc.

### Get Live Account
- **GET** `/account/live`
  - Retrieves live trading account information
  - Response: Account object with balance, equity, buying power, etc.

---

## Market Data

### Get Stock Quote
- **GET** `/marketdata/quotes/:symbol`
  - Retrieves stock quotes for a symbol
  - **Path Parameters:**
    - `symbol` - Stock symbol (e.g., AAPL)
  - **Query Parameters:**
    - `limit` - Number of quotes to retrieve (default: 1)
    - `startDate` - Start date in M/D/YYYY format (default: today)
  - **Example:** `/marketdata/quotes/AAPL?limit=10&startDate=1/1/2024`

---

## Trading - Orders

### Place Order
- **POST** `/orders`
  - Places a new order
  - **Request Body:**
    ```json
    {
      "symbol": "AAPL",
      "qty": 10,
      "side": "buy",
      "type": "market",
      "time_in_force": "day",
      "limit_price": 150.00,
      "stop_price": 145.00,
      "is_paper": true
    }
    ```
  - **Fields:**
    - `symbol` (required) - Stock symbol
    - `qty` (required) - Quantity to trade
    - `side` (required) - "buy" or "sell"
    - `type` (required) - "market", "limit", "stop", or "stop_limit"
    - `time_in_force` (required) - "day", "gtc", "opg", "cls", "ioc", or "fok"
    - `limit_price` (optional) - Required for limit orders
    - `stop_price` (optional) - Required for stop orders
    - `is_paper` (optional) - Use paper trading account (default: false)

### Get Orders
- **GET** `/orders`
  - Retrieves orders with optional filters
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)
    - `status` - Filter by status (open, closed, all)
    - `limit` - Maximum number of orders to return
    - `direction` - Sort direction (asc/desc)
    - `nested` - Include nested orders (true/false)
    - `after` - Filter orders after this time (RFC3339 format)
    - `until` - Filter orders until this time (RFC3339 format)
  - **Example:** `/orders?is_paper=true&status=open&limit=50`

### Get Order by ID
- **GET** `/orders/:id`
  - Retrieves a specific order
  - **Path Parameters:**
    - `id` - Order ID
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)
    - `nested` - Include nested orders (true/false)

### Cancel Order
- **DELETE** `/orders/:id`
  - Cancels a specific order
  - **Path Parameters:**
    - `id` - Order ID
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)

### Cancel All Orders
- **DELETE** `/orders`
  - Cancels all open orders
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)

---

## Trading - Positions

### Get All Positions
- **GET** `/positions`
  - Retrieves all open positions
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)

### Get Position by Symbol
- **GET** `/positions/:symbol`
  - Retrieves position for a specific symbol
  - **Path Parameters:**
    - `symbol` - Stock symbol
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)

### Close Position
- **DELETE** `/positions/:symbol`
  - Closes a position for a symbol
  - **Path Parameters:**
    - `symbol` - Stock symbol
  - **Request Body:**
    ```json
    {
      "qty": 10,
      "percentage": 50.0,
      "is_paper": true
    }
    ```
  - **Fields:**
    - `qty` (optional) - Quantity to close
    - `percentage` (optional) - Percentage of position to close
    - `is_paper` (optional) - Use paper account (default: false)
  - **Note:** Specify either `qty` or `percentage`, not both

### Close All Positions
- **DELETE** `/positions`
  - Closes all open positions
  - **Query Parameters:**
    - `is_paper` - Use paper account (true/false)
    - `cancel_orders` - Cancel open orders (true/false)

---

## Assets

### Get All Assets
- **GET** `/assets`
  - Retrieves all tradable assets
  - **Query Parameters:**
    - `status` - Filter by status (active, inactive)
    - `asset_class` - Filter by asset class (us_equity, crypto)

### Get Asset by Symbol
- **GET** `/assets/:symbol`
  - Retrieves information about a specific asset
  - **Path Parameters:**
    - `symbol` - Asset symbol

---

## Market Information

### Get Market Clock
- **GET** `/clock`
  - Retrieves current market clock information
  - Response includes: is_open, next_open, next_close

### Get Market Calendar
- **GET** `/calendar`
  - Retrieves market calendar
  - **Query Parameters:**
    - `start` - Start date (YYYY-MM-DD format)
    - `end` - End date (YYYY-MM-DD format)
  - **Example:** `/calendar?start=2024-01-01&end=2024-12-31`

---

## Order Types

### Market Order
Places an order at the current market price.
```json
{
  "symbol": "AAPL",
  "qty": 10,
  "side": "buy",
  "type": "market",
  "time_in_force": "day",
  "is_paper": true
}
```

### Limit Order
Places an order at a specific price or better.
```json
{
  "symbol": "AAPL",
  "qty": 10,
  "side": "buy",
  "type": "limit",
  "time_in_force": "gtc",
  "limit_price": 150.00,
  "is_paper": true
}
```

### Stop Order
Places a market order when the stop price is reached.
```json
{
  "symbol": "AAPL",
  "qty": 10,
  "side": "sell",
  "type": "stop",
  "time_in_force": "gtc",
  "stop_price": 145.00,
  "is_paper": true
}
```

### Stop Limit Order
Places a limit order when the stop price is reached.
```json
{
  "symbol": "AAPL",
  "qty": 10,
  "side": "sell",
  "type": "stop_limit",
  "time_in_force": "gtc",
  "limit_price": 144.00,
  "stop_price": 145.00,
  "is_paper": true
}
```

---

## Time In Force Options

- **day** - Order valid for the current trading day
- **gtc** - Good 'til cancelled
- **opg** - Execute at market open
- **cls** - Execute at market close
- **ioc** - Immediate or cancel
- **fok** - Fill or kill

---

## Environment Variables

Required environment variables in `.env` file:

```env
# Paper Trading
ALPACA_PAPER_API_KEY=your_paper_api_key
ALPACA_PAPER_SECRET_KEY=your_paper_secret_key

# Live Trading
ALPACA_LIVE_API_KEY=your_live_api_key
ALPACA_LIVE_API_SECRET_KEY=your_live_secret_key
```

---

## Running the Server

```bash
cd alpaca
go run cmd/main.go
```

The server will start on `http://localhost:8080`

---

## Example Usage with cURL

### Place a Market Order
```bash
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "symbol": "AAPL",
    "qty": 1,
    "side": "buy",
    "type": "market",
    "time_in_force": "day",
    "is_paper": true
  }'
```

### Get All Positions
```bash
curl http://localhost:8080/api/v1/positions?is_paper=true
```

### Get Account Information
```bash
curl http://localhost:8080/api/v1/account/paper
```

### Cancel All Orders
```bash
curl -X DELETE http://localhost:8080/api/v1/orders?is_paper=true
```