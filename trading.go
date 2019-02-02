package hitbtc

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

type Balance struct {
	Currency  string  `json:"currency"`
	Available float64 `json:"available,string"`
	Reserved  float64 `json:"reserved,string"`
}

func (h *HitBtc) GetBalances() (balances []Balance, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("get", "/api/2/trading/balance", nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &balances)
	return
}

type Order struct {
	ClientOrderId string    `json:"clientOrderId"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	TimeInForce   string    `json:"timeInForce"`
	Quantity      float64   `json:"quantity,string"`
	Price         float64   `json:"price,string"`
	CumQuantity   float64   `json:"cumQuantity,string"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	StopPrice     float64   `json:"stopPrice,string"`
	ExpireTime    time.Time `json:"expireTime"`
}

func (h *HitBtc) GetOrder(clientorderid string) (order Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("get", "/api/2/order/"+clientorderid, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &order)
	return
}

func (h *HitBtc) GetOrders(args ...string) (orders []Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var action = "/api/2/order"

	if len(args) > 0 {
		action = spr("%s/?symbol=%s", action, args[0])
	}

	go h.tradeRequest("get", action, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &orders)
	return
}

func (h *HitBtc) GetOrderStatus(clientorderid string) (order Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var parameters = make(map[string]string)
	parameters["clientOrderId"] = clientorderid

	go h.tradeRequest("get", "/api/2/history/order", parameters, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	var orders []Order
	err = json.Unmarshal(response, &orders)
	if err != nil {
		return
	}

	if len(orders) > 0 {
		order = orders[0]
		return
	}

	err = errors.New("Order Not Found!")
	return
}

func (h *HitBtc) Buy(symbol string, price, quantity float64) (order Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var parameters = make(map[string]string)

	parameters["side"] = "buy"
	parameters["symbol"] = symbol
	parameters["price"] = strconv.FormatFloat(float64(price), 'f', 8, 64)
	parameters["quantity"] = strconv.FormatFloat(float64(quantity), 'f', 8, 64)

	go h.tradeRequest("post", "/api/2/order/", parameters, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &order)
	return
}

func (h *HitBtc) Sell(symbol string, price, quantity float64) (order Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var parameters = make(map[string]string)

	parameters["side"] = "sell"
	parameters["symbol"] = symbol
	parameters["price"] = strconv.FormatFloat(float64(price), 'f', 8, 64)
	parameters["quantity"] = strconv.FormatFloat(float64(quantity), 'f', 8, 64)

	go h.tradeRequest("post", "/api/2/order/", parameters, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &order)
	return
}

func (h *HitBtc) CancelOrder(clientorderid string) (order Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("delete", "/api/2/order/"+clientorderid, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &order)
	return
}

func (h *HitBtc) CancelOrders(args ...string) (orders []Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var action = "/api/2/order"

	if len(args) > 0 {
		action = spr("%s/?symbol=%s", action, args[0])
	}

	go h.tradeRequest("delete", action, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &orders)
	return
}

type Fee struct {
	TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
	ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
}

func (h *HitBtc) GetFee(symbol string) (fee Fee, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("get", "/api/2/trading/fee/"+symbol, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &fee)
	return
}

func (h *HitBtc) GetOrderHistory(symbol string, limit int) (orders []Order, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var parameters = make(map[string]string)

	if symbol != "all" {
		parameters["symbol"] = strings.ToUpper(symbol)
	}

	parameters["limit"] = strconv.Itoa(limit)

	go h.tradeRequest("get", "/api/2/history/order", parameters, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &orders)
	return
}

type Trade struct {
	Id            uint64    `json:"id"`
	OrderId       uint64    `json:"orderId"`
	ClientOrderId string    `json:"clientOrderId"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	Quantity      float64   `json:"quantity,string"`
	Price         float64   `json:"price,string"`
	Fee           float64   `json:"fee,string"`
	Timestamp     time.Time `json:"timestamp"`
}

func (h *HitBtc) GetTradeHistory(symbol string, limit int) (trades []Trade, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	var parameters = make(map[string]string)

	if symbol != "all" {
		parameters["symbol"] = strings.ToUpper(symbol)
	}

	parameters["limit"] = strconv.Itoa(limit)

	go h.tradeRequest("get", "/api/2/history/trades", parameters, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &trades)
	return
}

func (h *HitBtc) GetTradesByOrder(orderid uint64) (trades []Trade, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	action := spr("/api/2/history/order/%d/trades", orderid)

	go h.tradeRequest("get", action, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &trades)
	return
}

func (h *HitBtc) GetAccountBalances() (balances []Balance, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("get", "/api/2/account/balance", nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &balances)
	return
}

type Deposit struct {
	Address   string `json:"address"`
	PaymentId string `json:"paymentId"`
}

func (h *HitBtc) GetDepositAddress(symbol string) (deposit Deposit, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("get", "/api/2/account/crypto/address/"+symbol, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &deposit)
	return
}

func (h *HitBtc) NewDepositAddress(symbol string) (deposit Deposit, err error) {
	respch := make(chan []byte)
	errch := make(chan error)

	go h.tradeRequest("post", "/api/2/account/crypto/address/"+symbol, nil, respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &deposit)
	return
}
