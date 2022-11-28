package rank

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
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
func (r *Rank) GET_SFACG_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"size": "50", "rtype": "view", "ntype": "origin", "expand": "typeName,tags,sysTags,ticket,latestchapter"}
	if r.RtypeIndex >= 6 {
		fmt.Println("RtypeIndex must be less than 6")
		r.RtypeIndex = 0
	}
	params["page"] = strconv.Itoa(r.Page)
	params["rtype"] = []string{"view", "sale", "newhit", "mark", "ticket", "bonus"}[r.RtypeIndex]
	request.Get(r.rank_api()).AddAll(params).NewRequests().Unmarshal(&RankStruct).WriteResultString()
	return RankStruct
}
