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
func main() {
	btc := GetBTC()
	eth := GetETH()
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Bitcoin (BTC)", btc, 'a', nil).
		AddItem("Ethereum (ETH)", eth, 'b', nil).
		AddItem("Ripple (XRP)", "To Be Implemented", 'c', nil).
		AddItem("Bitcoin Cash (BCH)", "To Be Implemented", 'd', nil).
		AddItem("Carnado (ADA)", "To Be Implemented", 'e', nil).
		AddItem("Litecoin (LTC)", "To Be Implemented", 'b', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
