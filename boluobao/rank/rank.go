package rank

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

type Rank struct {
	RtypeIndex int //view, sale, newhit, mark, ticket, bonus
	Month      bool
	All        bool
	Page       int
}

func (r *Rank) rank_api() string {
	if r.Month {
		return "ranks/month/novels"
	} else if r.All {
		return "ranks/all/novels"
	} else {
		return "ranks/week/novels"
	}
}
func (r *Rank) GET_SFACG_RANKS() Template.Rank {
	var RankStruct Template.Rank
	params := map[string]string{"size": "50", "ntype": "origin", "expand": "typeName,tags,sysTags,ticket,latestchapter"}
	if r.RtypeIndex >= 6 {
		fmt.Println("RtypeIndex must be less than 6")
		r.RtypeIndex = 0
	}
	params["page"] = strconv.Itoa(r.Page)
	params["rtype"] = map[int]string{0: "view", 1: "sale", 2: "newhit", 3: "mark", 4: "ticket", 5: "bonus"}[r.RtypeIndex]

	request.Get(r.rank_api()).AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}
