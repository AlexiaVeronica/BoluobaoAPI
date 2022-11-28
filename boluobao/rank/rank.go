package rank

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
)

type Rank struct{}

func (r *Rank) GET_POPULARITY_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "view", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}

func (r *Rank) GET_BEST_SELLER_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "sale", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}

func (r *Rank) GET_NEW_NOVEL_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "newhit", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}

func (r *Rank) GET_MARK_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "mark", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}

func (r *Rank) GET_TICKET_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "ticket", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}

func (r *Rank) GET_BONUS_RANKS() BoluobaoStructs.Rank {
	var RankStruct BoluobaoStructs.Rank
	params := map[string]string{"page": "0", "size": "50", "rtype": "bonus", "ntype": "origin", "expand": "typeName,tags,ticket,latestchapter"}
	request.Get("ranks/week/novels").AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}
