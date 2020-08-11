package logic

import "regexp"

const AppName = "cn.ac.may.finance"
const ApiUri = "https://hq.sinajs.cn/?_=%f&list=%s"
const TimeLocal = "PRC"
const TimeLayout = "2006-01-02 15:04:05"
const DtLayout = "01/02 15:04"
const FuPrefix = "fu"
const FuCodePrefix = "fu_"

const (
	ChartTodayFuUri    = "https://image.sinajs.cn/newchart/v5/fundpre/min/%s.gif?%f"
	ChartTodayStockUri = "https://image.sinajs.cn/newchart/min/n/%s.gif?%f"
	ChartDailyFuUri    = "https://image.sinajs.cn/newchart/v5/fund/nav/b/%s.gif?%f"
	ChartDailyStockUri = "http://image.sinajs.cn/newchart/daily/n/%s.gif?%f"

	DetailFuUri    = "http://finance.sina.com.cn/fund/quotes/%s/bc.shtml"
	DetailStockUri = "https://finance.sina.com.cn/realstock/company/%s/nc.shtml"
)

const (
	IconRise = "icon/rise.png"
	IconFall = "icon/fall.png"
	IconFlat = "icon/flat.png"
)

const (
	OptAdd   = "add"
	OptDel   = "del"
	OptClean = "clean"
	OptSort  = "sort"
	OptChart = "chart"
)

const (
	SortNone      = "nn"
	SortValueAsc  = "va"
	SortValueDesc = "vd"
	SortRatioAsc  = "ra"
	SortRatioDesc = "rd"
)

const (
	ChartToday = "today"
	ChartDaily = "daily"
)

var DataRegex = regexp.MustCompile(`var hq_str_(\w+)="(.*?)";`)
var CodeRegex = regexp.MustCompile(`^(fu|sh|sz)(\d{6})$`)

var SortAll = map[string]bool{
	SortNone:      true,
	SortValueAsc:  true,
	SortValueDesc: true,
	SortRatioAsc:  true,
	SortRatioDesc: true,
}

var ChartAll = map[string]bool{
	ChartToday: true,
	ChartDaily: true,
}
