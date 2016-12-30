package gemini

import (
	"encoding/json"
	"strconv"
)

// Active Orders
func (g *GeminiAPI) ActiveOrders() ([]Order, error) {

	url := g.url + ACTIVE_ORDERS_URL
	req := requestParams{
		"request": ACTIVE_ORDERS_URL,
		"nonce":   getNonce(),
	}

	var orders []Order

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &orders)

	return orders, nil
}

// Past Trades
func (g *GeminiAPI) PastTrades(symbol string, limitTrades int, timestamp int64) ([]Trade, error) {

	url := g.url + PAST_TRADES_URL

	req := requestParams{
		"request":      PAST_TRADES_URL,
		"nonce":        getNonce(),
		"symbol":       symbol,
		"limit_trades": limitTrades,
		"timestamp":    timestamp,
	}

	var trades []Trade

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &trades)

	return trades, nil
}

// New Order
func (g *GeminiAPI) NewOrder(symbol, clientOrderId string, amount, price float64, side string, options []string) (Order, error) {

	url := g.url + NEW_ORDER_URL
	req := requestParams{
		"request":         NEW_ORDER_URL,
		"nonce":           getNonce(),
		"client_order_id": clientOrderId,
		"symbol":          symbol,
		"amount":          strconv.FormatFloat(amount, 'f', -1, 64),
		"price":           strconv.FormatFloat(price, 'f', -1, 64),
		"side":            side,
		"type":            "exchange limit",
	}

	if options != nil {
		req["options"] = options
	}

	var order Order

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// Order Status
func (g *GeminiAPI) OrderStatus(orderId OrderId) (Order, error) {

	url := g.url + ORDER_STATUS_URL
	req := requestParams{
		"request":  ORDER_STATUS_URL,
		"nonce":    getNonce(),
		"order_id": orderId,
	}

	var order Order

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// Cancel All
func (g *GeminiAPI) CancelAll() (CancelResponse, error) {
	url := g.url + CANCEL_ALL_URL
	req := requestParams{
		"request": CANCEL_ALL_URL,
		"nonce":   getNonce(),
	}

	var res CancelResponse

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Cancel Session
func (g *GeminiAPI) CancelSession() (GenericResponse, error) {
	url := g.url + CANCEL_SESSION_URL
	req := requestParams{
		"request": CANCEL_SESSION_URL,
		"nonce":   getNonce(),
	}

	var res GenericResponse

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Cancel Order
func (g *GeminiAPI) CancelOrder(orderId OrderId) (Order, error) {

	url := g.url + CANCEL_ORDER_URL
	req := requestParams{
		"request":  CANCEL_ORDER_URL,
		"nonce":    getNonce(),
		"order_id": orderId,
	}

	var order Order

	body, err := g.request("POST", url, req, nil)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// Heartbeat
func (g *GeminiAPI) Heartbeat() (GenericResponse, error) {

	url := g.url + HEARTBEAT_URL
	req := requestParams{
		"request": HEARTBEAT_URL,
		"nonce":   getNonce(),
	}

	var res GenericResponse

	body, err := g.request("POST", url, req, nil)

	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}