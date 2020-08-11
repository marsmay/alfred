package logic

import (
	"alfred/workflow/config"
	"alfred/workflow/output"
	"alfred/workflow/util"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type Config struct {
	Sort  string   `json:"sort"`
	Chart string   `json:"chart"`
	Codes []string `json:"codes"`
}

func parseTime(ymd, his string) (t time.Time, err error) {
	loc, err := time.LoadLocation(TimeLocal)

	if err == nil {
		t, err = time.ParseInLocation(TimeLayout, ymd+" "+his, loc)
	}

	return
}

func parseData(content string) (items []IItem, err error) {
	pieces := DataRegex.FindAllStringSubmatch(content, -1)
	items = make([]IItem, 0, len(pieces))

	for _, piece := range pieces {
		if len(piece) < 3 {
			continue
		}

		var parseFunc func(string, string) (IItem, error)

		if strings.HasPrefix(piece[1], FuPrefix) {
			parseFunc = parseFund
		} else {
			parseFunc = parseStock
		}

		item, e := parseFunc(piece[1], piece[2])

		if e != nil {
			err = e
			return
		}

		if item != nil {
			items = append(items, item)
		}
	}

	return
}

func getData(codes []string) (items []IItem, err error) {
	for k, v := range codes {
		if strings.HasPrefix(v, FuPrefix) && !strings.HasPrefix(v, FuCodePrefix) {
			codes[k] = FuCodePrefix + strings.TrimPrefix(v, FuPrefix)
		}
	}

	uri := fmt.Sprintf(ApiUri, rand.Float64(), strings.Join(codes, ","))
	body, err := util.HttpRequest(uri, nil, nil)

	if err != nil {
		return
	}

	return parseData(string(body))
}

func getDataResult(codes []string, sortBy, chart string) (res *output.Result) {
	if len(codes) == 0 {
		return output.NewNotice("No results", nil)
	}

	res = &output.Result{Items: make([]*output.Item, 0, len(codes))}
	items, err := getData(codes)

	if err != nil {
		return output.NewNotice("Query data failed", err)
	}

	if len(items) == 0 {
		return output.NewNotice("No results", nil)
	}

	switch sortBy {
	case SortValueAsc:
		sort.Sort(SortItemsByValue{items})
	case SortValueDesc:
		sort.Sort(sort.Reverse(SortItemsByValue{items}))
	case SortRatioAsc:
		sort.Sort(SortItemsByRatio{items})
	case SortRatioDesc:
		sort.Sort(sort.Reverse(SortItemsByRatio{items}))
	}

	for _, item := range items {
		res.Items = append(res.Items, item.ToItem(chart))
	}

	return
}

func modifyList(codes []string, opt, code string) (newCodes []string) {
	newCodes = make([]string, 0, len(codes)+1)

	if opt == OptAdd {
		newCodes = append(newCodes, code)
	}

	for _, v := range codes {
		if v != code {
			newCodes = append(newCodes, v)
		}
	}

	return
}

func parseCode(str string) (valid bool, code string) {
	items := CodeRegex.FindStringSubmatch(str)

	if len(items) != 3 {
		return
	}

	valid, code = true, str
	prefix, num := items[1], items[2]

	if prefix == FuPrefix {
		code = FuCodePrefix + num
	}

	return
}

func Excute() (result *output.Result) {
	conf := &Config{}

	if err := config.Load(AppName, conf); err != nil {
		return output.NewNotice("Load settings failed", err)
	}

	if len(os.Args) <= 1 || os.Args[1] == "" {
		return getDataResult(conf.Codes, conf.Sort, conf.Chart)
	}

	args := strings.Split(strings.ToLower(os.Args[1]), " ")

	switch args[0] {
	case OptClean:
		conf.Codes = []string{}
		return config.SaveResult(AppName, conf)
	case OptAdd, OptDel:
		if len(args) != 2 {
			return output.NewNotice("Invalid parameters", nil)
		}

		valid, code := parseCode(args[1])

		if !valid {
			return output.NewNotice("Invalid parameters", nil)
		}

		conf.Codes = modifyList(conf.Codes, args[0], code)
		return config.SaveResult(AppName, conf)
	case OptSort:
		if len(args) != 2 || !SortAll[args[1]] {
			return output.NewNotice("Invalid parameters", nil)
		}

		conf.Sort = args[1]
		return config.SaveResult(AppName, conf)
	case OptChart:
		if len(args) != 2 || !ChartAll[args[1]] {
			return output.NewNotice("Invalid parameters", nil)
		}

		conf.Chart = args[1]
		return config.SaveResult(AppName, conf)
	default:
		valid, code := parseCode(args[0])

		if !valid {
			return output.NewNotice("Invalid parameters", nil)
		}

		return getDataResult([]string{code}, SortNone, conf.Chart)
	}
}
