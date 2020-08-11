package logic

import (
	"alfred/workflow/output"
	"alfred/workflow/util"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Fund struct {
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	NetValue          float64   `json:"net_value"`
	EstimateNetValue  float64   `json:"estimate_net_value"`
	EstimateGainRatio float64   `json:"estimate_gain_ratio"`
	AccuNetValue      float64   `json:"accu_net_value"`
	Time              time.Time `json:"time"`
}

func (f *Fund) ToItem(chart string) (item *output.Item) {
	code := strings.TrimPrefix(f.Code, FuCodePrefix)
	item = &output.Item{
		Valid:     true,
		Title:     fmt.Sprintf("%s  %.4f (%+.2f%%)", f.Name, f.EstimateNetValue, f.EstimateGainRatio),
		SubTitle:  fmt.Sprintf("净值: %.4f, 累计净值: %.4f (%s)", f.NetValue, f.AccuNetValue, f.Time.Format(DtLayout)),
		Icon:      output.NewIcon(output.IconDefault),
		Arguments: fmt.Sprintf(DetailFuUri, code),
	}

	if f.EstimateGainRatio > 0 {
		item.Icon.Path = IconRise
	} else if f.EstimateGainRatio < 0 {
		item.Icon.Path = IconFall
	} else {
		item.Icon.Path = IconFlat
	}

	if chart == ChartDaily {
		item.QuicklookUrl = fmt.Sprintf(ChartDailyFuUri, code, rand.Float64())
	} else {
		item.QuicklookUrl = fmt.Sprintf(ChartTodayFuUri, code, rand.Float64())
	}

	return
}

func (f *Fund) GetValue() float64 {
	return f.EstimateNetValue
}

func (f *Fund) GetGainRatio() float64 {
	return f.EstimateGainRatio
}

func parseFund(code, content string) (item IItem, err error) {
	fund := &Fund{Code: code}
	info := strings.Split(content, ",")

	if len(info) < 8 {
		return
	}

	if fund.Name, err = util.GbkToUtf8(info[0]); err != nil {
		return
	}

	if fund.EstimateNetValue, err = strconv.ParseFloat(info[2], 10); err != nil {
		return
	}

	if fund.NetValue, err = strconv.ParseFloat(info[3], 10); err != nil {
		return
	}

	if fund.AccuNetValue, err = strconv.ParseFloat(info[4], 10); err != nil {
		return
	}

	if fund.EstimateGainRatio, err = strconv.ParseFloat(info[6], 10); err != nil {
		return
	}

	if fund.Time, err = parseTime(info[7], info[1]); err != nil {
		return
	}

	return fund, nil
}
