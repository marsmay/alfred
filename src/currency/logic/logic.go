package logic

import (
	"alfred/workflow/output"
	"os"
	"strconv"
	"strings"
)

func Excute() (result *output.Result) {
	if len(os.Args) <= 1 || os.Args[1] == "" {
		return output.NewNotice("Invalid parameters", nil)
	}

	args := strings.Split(os.Args[1], " ")

	if len(args) != 3 {
		return output.NewNotice("Invalid parameters", nil)
	}

	opt, currency := strings.ToLower(args[0]), strings.ToUpper(args[2])

	if OptDesc[opt] == "" {
		return output.NewNotice("Invalid parameters", nil)
	}

	if !CurrencyRegex.Match([]byte(currency)) {
		return output.NewNotice("Invalid currency", nil)
	}

	amount, err := strconv.ParseFloat(args[1], 10)

	if err != nil {
		return output.NewNotice("Invalid amount", err)
	}

	if amount <= 0 {
		return output.NewNotice("Invalid amount", nil)
	}

	rates, err := getRates()

	if err != nil {
		return output.NewNotice("Query data failed", err)
	}

	currencyRates := rates[currency]

	if currencyRates == nil {
		return output.NewNotice("No results", nil)
	}

	return currencyRates.GetResult(opt, amount)
}
