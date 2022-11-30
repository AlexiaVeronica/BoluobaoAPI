package tag

import "strconv"

func TagsParams(TagID string, page int) map[string]string {
	return map[string]string{
		"systagids":      TagID,
		"isfree":         "both",
		"size":           "50",
		"charcountbegin": "0",
		"updatedays":     "-1",
		"expand":         "typeName,sysTags,discount,discountExpireDate",
		"sort":           "latest",
		"page":           strconv.Itoa(page),
		"charcountend":   "0",
		"isfinish":       "both",
	}
}
