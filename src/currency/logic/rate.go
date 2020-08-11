package logic

import (
	"alfred/workflow/output"
	"fmt"
)

type Rate struct {
	Bank  string  `json:"bank"`
	Price float64 `json:"price"`
}

func (r *Rate) String() string {
	return fmt.Sprintf("%+v", *r)
}

type Rates []*Rate

func (s Rates) Len() int           { return len(s) }
func (s Rates) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Rates) Less(i, j int) bool { return s[i].Price < s[j].Price }

type ExchangeRate struct {
	Currency     string `json:"currency"`
	BuyCash      Rates  `json:"buy_cash"`
	BuyExchange  Rates  `json:"buy_exchange"`
	SellCash     Rates  `json:"sell_cash"`
	SellExchange Rates  `json:"sell_exchange"`
}

func (er *ExchangeRate) String() string {
	return fmt.Sprintf("%+v", *er)
}

func (er *ExchangeRate) GetResult(opt string, amount float64) (result *output.Result) {
	result = output.NewNotice("No results", nil)

	var rates Rates

	switch opt {
	case OptBuyCash:
		rates = er.SellCash
	case OptBuyExchange:
		rates = er.SellExchange
	case OptSellCash:
		rates = er.BuyCash
	case OptSellExchange:
		rates = er.BuyExchange
	}

	if len(rates) == 0 {
		return
	}

	result.Items = make([]*output.Item, 0, len(rates))

	for _, rate := range rates {
		value := amount * rate.Price / 100

		result.Items = append(result.Items, &output.Item{
			Valid:     true,
			Title:     fmt.Sprintf("%.2f CNY", value),
			SubTitle:  fmt.Sprintf("%s : CNY = 1 : %.6f", er.Currency, rate.Price/100),
			Icon:      output.NewIcon(fmt.Sprintf(IconPath, rate.Bank)),
			Arguments: fmt.Sprintf("%s%s %.2f %s = %.2f CNY", BankDesc[rate.Bank], OptDesc[opt], amount, er.Currency, value),
		})
	}

	return
}
