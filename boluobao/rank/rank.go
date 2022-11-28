package rank

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
)

type Rank struct {
	RtypeIndex int //view, sale, newhit, mark, ticket, bonus
	Month      bool
	All        bool
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
	params := map[string]string{"page": "0", "size": "50", "rtype": "view", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	params["rtype"] = []string{"view", "sale", "newhit", "mark", "ticket", "bonus"}[r.RtypeIndex]
	request.Get(r.rank_api()).AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}
