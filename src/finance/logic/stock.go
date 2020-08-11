package logic

import (
	"alfred/workflow/output"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Stock struct {
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	PrevClosePrice    float64   `json:"prev_close_price"`
	OpenPrice         float64   `json:"open_price"`
	CurrentPrice      float64   `json:"current_price"`
	CurrentGainRatio  float64   `json:"current_gain_ratio"`
	CurrentGainAmount float64   `json:"current_gain_amount"`
	HighPrice         float64   `json:"high_price"`
	LowPrice          float64   `json:"low_price"`
	VolNum            float64   `json:"vol_num"`
	VolAmount         float64   `json:"vol_amount"`
	Time              time.Time `json:"time"`
}

func (s *Stock) ToItem(chart string) (item *output.Item) {
	item = &output.Item{
		Valid:     true,
		Title:     fmt.Sprintf("%s  %.2f (%+.2f %+.2f%%)", s.Name, s.CurrentPrice, s.CurrentGainAmount, s.CurrentGainRatio),
		SubTitle:  fmt.Sprintf("昨收: %.2f, 今开: %.2f, 振幅: %.2f ~ %.2f (%.2f%%) [%s]", s.PrevClosePrice, s.OpenPrice, s.LowPrice, s.HighPrice, (s.HighPrice-s.LowPrice)*100/s.PrevClosePrice, s.Time.Format(DtLayout)),
		Icon:      &output.Icon{Path: output.IconDefault},
		Arguments: fmt.Sprintf(DetailStockUri, s.Code),
	}

	if s.CurrentGainAmount > 0 {
		item.Icon.Path = IconRise
	} else if s.CurrentGainAmount < 0 {
		item.Icon.Path = IconFall
	} else {
		item.Icon.Path = IconFlat
	}

	if chart == ChartDaily {
		item.QuicklookUrl = fmt.Sprintf(ChartDailyStockUri, s.Code, rand.Float64())
	} else {
		item.QuicklookUrl = fmt.Sprintf(ChartTodayStockUri, s.Code, rand.Float64())
	}

	return
}

func (s *Stock) GetValue() float64 {
	return s.CurrentPrice
}

func (s *Stock) GetGainRatio() float64 {
	return s.CurrentGainRatio
}

func parseStock(code, content string) (item IItem, err error) {
	stock := &Stock{Code: code}
	info := strings.Split(content, ",")

	if len(info) < 32 {
		return
	}

	if stock.Name, err = gbkToUtf8(info[0]); err != nil {
		return
	}

	if stock.OpenPrice, err = strconv.ParseFloat(info[1], 10); err != nil {
		return
	}

	if stock.PrevClosePrice, err = strconv.ParseFloat(info[2], 10); err != nil {
		return
	}

	if stock.CurrentPrice, err = strconv.ParseFloat(info[3], 10); err != nil {
		return
	}

	if stock.HighPrice, err = strconv.ParseFloat(info[4], 10); err != nil {
		return
	}

	if stock.LowPrice, err = strconv.ParseFloat(info[5], 10); err != nil {
		return
	}

	if stock.VolNum, err = strconv.ParseFloat(info[8], 10); err != nil {
		return
	}

	if stock.VolAmount, err = strconv.ParseFloat(info[9], 10); err != nil {
		return
	}

	if stock.Time, err = parseTime(info[30], info[31]); err != nil {
		return
	}

	stock.CurrentGainAmount = stock.CurrentPrice - stock.PrevClosePrice
	stock.CurrentGainRatio = stock.CurrentGainAmount * 100 / stock.PrevClosePrice

	return stock, nil
}
