package logic

import "alfred/workflow/output"

type IItem interface {
	ToItem(chart string) (item *output.Item)
	GetValue() float64
	GetGainRatio() float64
}

type SortItems []IItem

func (s SortItems) Len() int      { return len(s) }
func (s SortItems) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type SortItemsByValue struct{ SortItems }

func (s SortItemsByValue) Less(i, j int) bool {
	return s.SortItems[i].GetValue() < s.SortItems[j].GetValue()
}

type SortItemsByRatio struct{ SortItems }

func (s SortItemsByRatio) Less(i, j int) bool {
	return s.SortItems[i].GetGainRatio() < s.SortItems[j].GetGainRatio()
}
