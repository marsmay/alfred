package logic

import "regexp"

const ApiUri = "http://vip.stock.finance.sina.com.cn/forex/api/openapi.php/ForexService.getBankForex?_=%f"
const IconPath = "icon/%s.png"

const (
	OptBuyCash      = "cbc"
	OptBuyExchange  = "cbe"
	OptSellCash     = "csc"
	OptSellExchange = "cse"
)

var OptDesc = map[string]string{
	OptBuyCash:      "购钞",
	OptBuyExchange:  "购汇",
	OptSellCash:     "结钞",
	OptSellExchange: "结汇",
}

var BankDesc = map[string]string{
	"icbc":     "工商银行",
	"boc":      "中国银行",
	"abchina":  "农业银行",
	"bankcomm": "交通银行",
	"ccb":      "建设银行",
	"cmbchina": "招商银行",
	"spdb":     "浦发银行",
	"ecitic":   "中信银行",
	"cebbank":  "光大银行",
	"cib":      "兴业银行",
}

var CurrencyRegex = regexp.MustCompile(`[A-Z]{3}`)
