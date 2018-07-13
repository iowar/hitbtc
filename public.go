package hitbtc

import (
	"encoding/json"
	"strings"
	"time"
)

type Currency struct {
	Id                 string  `json:"id"`
	FullName           string  `json:"fullName"`
	Crypto             bool    `json:"crypto"`
	PayinEnabled       bool    `json:"payinEnabled"`
	PayinPaymentId     bool    `json:"payinPaymentId"`
	PayinConfirmations uint    `json:"payinConfirmations"`
	PayoutEnabled      bool    `json:"payoutEnabled"`
	PayoutIsPaymentId  bool    `json:"payoutIsPaymentId"`
	TransferEnabled    bool    `json:"transferEnabled"`
	Delisted           bool    `json:"delisted"`
	PayoutFee          float64 `json:"payoutFee,string"`
}

func (h *HitBtc) GetCurrency(id string) (currency Currency, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/currency/"+strings.ToUpper(id), respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &currency)
	return
}

func (h *HitBtc) GetCurrencies() (currencies []Currency, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/currency/", respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &currencies)
	return
}

type Symbol struct {
	Id                   string  `json:"id"`
	BaseCurrency         string  `json:"baseCurrency"`
	QuoteCurrency        string  `json:"quoteCurrency"`
	QuantityIncrement    float64 `json:"quantityIncrement,string"`
	TickSize             float64 `json:"tickSize,string"`
	TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
	ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
	FeeCurrency          string  `json:"feeCurrency"`
}

func (h *HitBtc) GetSymbol(id string) (symbol Symbol, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/symbol/"+strings.ToUpper(id), respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &symbol)
	return
}

func (h *HitBtc) GetSymbols() (symbols []Symbol, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/symbol/", respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &symbols)
	return
}

type Ticker struct {
	Ask         float64   `json:"ask,string"`
	Bid         float64   `json:"bid,string"`
	Last        float64   `json:"last,string"`
	Open        float64   `json:"open,string"`
	Low         float64   `json:"low,string"`
	High        float64   `json:"high,string"`
	Volume      float64   `json:"volume,string"`
	VolumeQuote float64   `json:"volumeQuote,string"`
	Timestamp   time.Time `json:"timeStamp"`
	Symbol      string    `json:"symbol"`
}

func (t *Ticker) UnmarshalJSON(b []byte) error {
	var ticker map[string]string

	err := json.Unmarshal(b, &ticker)
	if err != nil {
		return err
	}

	t.Ask, err = strconv.ParseFloat(ticker["ask"], 64)
	if err != nil {
		t.Ask = 0
	}

	t.Bid, err = strconv.ParseFloat(ticker["bid"], 64)
	if err != nil {
		t.Bid = 0
	}

	t.Last, err = strconv.ParseFloat(ticker["last"], 64)
	if err != nil {
		t.Last = 0
	}

	t.Open, err = strconv.ParseFloat(ticker["open"], 64)
	if err != nil {
		t.Open = 0
	}

	t.Low, err = strconv.ParseFloat(ticker["low"], 64)
	if err != nil {
		t.Low = 0
	}

	t.High, err = strconv.ParseFloat(ticker["high"], 64)
	if err != nil {
		t.High = 0
	}

	t.Volume, err = strconv.ParseFloat(ticker["volume"], 64)
	if err != nil {
		t.Volume = 0
	}

	t.VolumeQuote, err = strconv.ParseFloat(ticker["volumeQuote"], 64)
	if err != nil {
		t.VolumeQuote = 0
	}

	layout := "2006-01-02T15:04:05.000Z"

	t.Timestamp, err = time.Parse(layout, ticker["timestamp"])
	if err != nil {
		t.Timestamp = time.Time{}
	}

	t.Symbol = ticker["symbol"]

	return nil
}

func (h *HitBtc) GetTicker(symbol string) (ticker Ticker, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/ticker/"+strings.ToUpper(symbol), respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &ticker)
	return
}

func (h *HitBtc) GetTickers() (tickers []Ticker, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/ticker/", respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &tickers)

	return
}

type MarketTrade struct {
	Id        int       `json:"id"`
	Price     float64   `json:"price,string"`
	Quantity  float64   `json:"quantity,string"`
	Side      string    `json:"side"`
	Timestamp time.Time "json:timestamp"
}

func (h *HitBtc) GetTrades(symbol string) (trades []MarketTrade, err error) {

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest("/api/2/public/trades/"+strings.ToUpper(symbol), respch, errch)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &trades)
	return
}

type Book struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size,string"`
}

type OrderBook struct {
	Ask []Book `json:"ask"`
	Bid []Book `json:"bid"`
}

func (h *HitBtc) GetOrderBook(symbol string, args ...int) (orderbooks OrderBook, err error) {

	var limit int

	respch := make(chan []byte)
	errch := make(chan error)

	if len(args) > 0 {
		limit = args[0]
	} else {
		limit = 100
	}

	go h.publicRequest(
		"/api/2/public/orderbook/"+strings.ToUpper(symbol)+spr("?limit=%d", limit),
		respch,
		errch,
	)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &orderbooks)
	return
}

type Candle struct {
	Timestamp   time.Time `json:"timestamp"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Volume      float64   `json:"volume,string"`
	VolumeQuote float64   `json:"volumeQuote,string"`
}

// period list
// M1 (one minute), M3, M5, M15, M30, H1, H4, D1, D7, 1M (one month).
func (h *HitBtc) GetCandles(symbol, period string, limit int) (candles []Candle, err error) {

	var pfl int

	for _, v := range []string{"M1", "M3", "M5", "M15", "M30", "H1", "H4", "D1", "D7", "1M"} {
		if period == v {
			pfl++
		}
	}

	if pfl == 0 {
		return nil, Error(PeriodError)
	}

	respch := make(chan []byte)
	errch := make(chan error)

	go h.publicRequest(
		"/api/2/public/candles/"+strings.ToUpper(symbol)+spr("?limit=%d&period=%s", limit, period),
		respch,
		errch,
	)

	response := <-respch
	err = <-errch

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &candles)
	return
}
