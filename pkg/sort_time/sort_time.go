package sort_time

import "sort"

type MapSorter []SortItem

type SortItem struct {
	Key string      `json:"key"`
	Val interface{} `json:"val"`
}

func (ms MapSorter) Len() int {
	return len(ms)
}
func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Key < ms[j].Key // 按键排序
	// return ms[i].Value <ms[j].Value //按值排序
}
func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func MapSort(m map[string]string) []map[string]string {
	ms := make(MapSorter, 0, len(m))

	for k, v := range m {
		ms = append(ms, SortItem{k, v})
	}
	sort.Sort(ms)
	var result []map[string]string
	for _, p := range ms {
		result = append(result, map[string]string{p.Key: p.Val.(string)})
	}
	return result
}
