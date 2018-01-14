package main

import (
	"encoding/json"
	"github.com/rivo/tview"
	"io/ioutil"
	"net/http"
)

/*type Crypto struct {
	id                 string
	name               string
	symbol             string
	rank               int
	price_usd          float64
	h_volume_usd       float64
	market_cap_usd     float64
	available_supply   float64
	total_supply       float64
	max_supply         float64
	percent_change_1h  float64
	percent_change_24h float64
	percent_change_7d  float64
	last_updated       int
	price_inr          float64
	h_volume_usd       float64
	market_cap_usd     float64
}*/

func GetBTC() string {
	url := "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	btcPrice := result["price_inr"].(string)
	return btcPrice
}

func GetETH() string {
	url := "https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	ethPrice := result["price_inr"].(string)
	return ethPrice
}

func GetXRP() string {
	url := "https://api.coinmarketcap.com/v1/ticker/ripple/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	xrpPrice := result["price_inr"].(string)
	return xrpPrice
}

func GetBCH() string {
	url := "https://api.coinmarketcap.com/v1/ticker/bitcoin-cash/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	bchPrice := result["price_inr"].(string)
	return bchPrice
}

func GetLTC() string {
	url := "https://api.coinmarketcap.com/v1/ticker/litecoin/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	ltcPrice := result["price_inr"].(string)
	return ltcPrice
}

func GetADA() string {
	url := "https://api.coinmarketcap.com/v1/ticker/cardano/?convert=INR"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	// Accessing Value using empty interface
	var value interface{}
	err = json.Unmarshal(body, &value)
	price := value.([]interface{})
	result := price[0].(map[string]interface{})
	ripPrice := result["price_inr"].(string)
	return ripPrice
}

func main() {
	btc := GetBTC()
	eth := GetETH()
	xrp := GetXRP()
	bch := GetBCH()
	ada := GetADA()
	ltc := GetLTC()
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Bitcoin (BTC)", "\u20B9 " + btc, 'a', nil).
		AddItem("Ethereum (ETH)", "\u20B9 " + eth, 'b', nil).
		AddItem("Ripple (XRP)", "\u20B9 "+ xrp, 'c', nil).
		AddItem("Bitcoin Cash (BCH)", "\u20B9 " + bch, 'd', nil).
		AddItem("Cardano (ADA)", "\u20B9 " + ada, 'e', nil).
		AddItem("Litecoin (LTC)", "\u20B9 " + ltc, 'b', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
