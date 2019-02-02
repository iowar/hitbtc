# HitBTC API in GO

## Install
```
$ go get -u github.com/iowar/hitbtc
```

## Public Api

No need authentication to use it.

~~~go
package main

import (
        "fmt"

        hbtc "github.com/iowar/hitbtc"
)

func main() {
        hitbtc, err := hbtc.NewClient("", "")
        if err != nil {
                panic(err)
        }

        r, _ := hitbtc.GetCurrency("ipl")
        //r, _ := hitbtc.GetCurrencies()
        //r, _ := hitbtc.GetSymbol("iplbtc")
        //r, _ := hitbtc.GetSymbols()
        //r, _ := hitbtc.GetTicker("iplbtc")
        //r, _ := hitbtc.GetTickers()
        //r, _ := hitbtc.GetTrades("iplbtc")
        //r, _ := hitbtc.GetOrderBook("iplbtc", 10)
        //r, _ := hitbtc.GetCandles("iplbtc", "M3", 10)

        fmt.Println(r)
}
~~~

## Trading Api
Api keys must be set.

~~~go
package main

import (
        "fmt"

        hbtc "github.com/iowar/hitbtc"
)

const (
        api_key    = ""
        api_secret = ""
)

func main() {
        hitbtc, err := hbtc.NewClient(api_key, api_secret)
        if err != nil {
                panic(err)
        }

        r, err := hitbtc.GetBalances()
        //r, err := hitbtc.GetOrder("<valid_orderid>")
        //r, err := hitbtc.GetOrders("iplbtc")
        //r, err := hitbtc.GetOrderStatus("clientOrderId")
        //r, err := hitbtc.Buy("iplbtc", 0.00000001, 10000)
        //r, err := hitbtc.Sell("iplbtc", 1, 1000)
        //r, err := hitbtc.CancelOrder("<valid_orderid>")
        //r, err := hitbtc.CancelOrders("iplbtc")
        //r, err := hitbtc.GetFee("iplbtc")
        //r, err := hitbtc.GetOrderHistory("iplbtc", 10)
        //r, err := hitbtc.GetTradeHistory("iplbtc", 10)
        //r, err := hitbtc.GetTradesByOrder(orderid_type_uint64)
        //r, err := hitbtc.GetAccountBalances()
        //r, err := hitbtc.GetDepositAddress("BTC")
        //r, err := hitbtc.NewDepositAddress("BTC")

        if err != nil {
                panic(err)
        }

        fmt.Println(r)
}

~~~

### Some Notes
* Orders may be delayed by the hitbtc when used buy and sell options. 
* Throttle set as 9 req/s.
* There is no ws support.

License
----
[MIT](https://github.com/iowar/hitbtc/blob/master/LICENSE)

