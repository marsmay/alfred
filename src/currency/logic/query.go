package logic

import (
	"alfred/workflow/util"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Resp struct {
	Result *RespResult `json:"result"`
}

type RespResult struct {
	Status *RespStatus `json:"status"`
	Data   *RespRates  `json:"data"`
}

type RespStatus struct {
	Code int `json:"code"`
}

type RespRates struct {
	Rates map[string][]*RespBank `json:"bank"`
}

type RespBank struct {
	Bank              string `json:"bank"`
	BuyCashPrice      string `json:"xc_buy_price"`
	BuyExchangePrice  string `json:"xh_buy_price"`
	SellCashPrice     string `json:"xc_sell_price"`
	SellExchangePrice string `json:"xh_sell_price"`
}

func parseData(content []byte) (rates map[string]*ExchangeRate, err error) {
	resp := &Resp{}

	if err = json.Unmarshal(content, resp); err != nil {
		return
	}

	if resp.Result == nil || resp.Result.Data == nil || resp.Result.Data.Rates == nil {
		err = fmt.Errorf("invalid response")
		return
	}

	rates = make(map[string]*ExchangeRate, len(resp.Result.Data.Rates))

	for currency, banks := range resp.Result.Data.Rates {
		if rates[currency] == nil {
			rates[currency] = &ExchangeRate{
				Currency:     currency,
				BuyCash:      make(Rates, 0, 8),
				BuyExchange:  make(Rates, 0, 8),
				SellCash:     make(Rates, 0, 8),
				SellExchange: make(Rates, 0, 8),
			}
		}

		for _, bank := range banks {
			if v, e := strconv.ParseFloat(bank.BuyCashPrice, 10); e == nil && v > 0 {
				rates[currency].BuyCash = append(rates[currency].BuyCash, &Rate{bank.Bank, v})
			}

			if v, e := strconv.ParseFloat(bank.BuyExchangePrice, 10); e == nil && v > 0 {
				rates[currency].BuyExchange = append(rates[currency].BuyExchange, &Rate{bank.Bank, v})
			}

			if v, e := strconv.ParseFloat(bank.SellCashPrice, 10); e == nil && v > 0 {
				rates[currency].SellCash = append(rates[currency].SellCash, &Rate{bank.Bank, v})
			}

			if v, e := strconv.ParseFloat(bank.SellExchangePrice, 10); e == nil && v > 0 {
				rates[currency].SellExchange = append(rates[currency].SellExchange, &Rate{bank.Bank, v})
			}
		}
	}

	for _, rate := range rates {
		sort.Sort(sort.Reverse(rate.BuyCash))
		sort.Sort(sort.Reverse(rate.BuyExchange))
		sort.Sort(rate.SellCash)
		sort.Sort(rate.SellExchange)
	}

	return
}

func getRates() (rates map[string]*ExchangeRate, err error) {
	uri := fmt.Sprintf(ApiUri, rand.Float64())
	body, err := util.HttpRequest(uri, nil, nil)

	if err != nil {
		return
	}

	return parseData(body)
}
