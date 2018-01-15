package main

import (
	"encoding/json"
	"github.com/rivo/tview"
	"io/ioutil"
	"net/http"
)

func GetCurr() (string, string, string, string, string, string)  {
	url := "https://api.coinmarketcap.com/v1/ticker/?limit=6&convert=INR"
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

	result0 := price[0].(map[string]interface{})
	btcPrice := result0["price_inr"].(string)

  result1 := price[1].(map[string]interface{})
  ethPrice := result1["price_inr"].(string)

  result2 := price[2].(map[string]interface{})
  xrpPrice := result2["price_inr"].(string)

  result3 := price[3].(map[string]interface{})
  bchPrice := result3["price_inr"].(string)

  result4 := price[4].(map[string]interface{})
  adaPrice := result4["price_inr"].(string)

  result5 := price[5].(map[string]interface{})
  ltcPrice := result5["price_inr"].(string)

	return btcPrice, ethPrice, xrpPrice, bchPrice, adaPrice, ltcPrice
}

func main() {
	btc, eth, xrp, bch, ada, ltc := GetCurr()
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
