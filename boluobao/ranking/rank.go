package ranking

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"strconv"
)

type Rank struct {
	TypeName string
	Month    bool
	All      bool
	Page     int
	Size     int
}

type TypeName struct {
	ViewRank   string
	BestRank   string
	NewRank    string
	TicketRank string
	MarkRank   string
	BonusRank  string
}

func (r *Rank) TypeNameInit() TypeName {
	return TypeName{
		ViewRank:   "view",
		BestRank:   "sale",
		NewRank:    "newhit",
		TicketRank: "ticket",
		MarkRank:   "mark",
		BonusRank:  "bonus",
	}
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
	params := map[string]string{"ntype": "origin", "expand": "typeName,tags,sysTags,ticket,latestchapter"}
	if r.Size == 0 {
		r.Size = 50 //默认50
	}
	if r.TypeName == "sale" {
		if r.Size > 40 {
			r.Size = 40 // 畅销榜最大值为40
		}
		if r.Page > 0 {
			r.Page = 0 // 畅销榜只有一页
		}
		if r.All {
			r.All = false // 畅销榜只有周榜, 月榜
		}
	}
	if r.TypeName == "newhit" || r.TypeName == "bonus" {
		r.Month = true // 新作榜和打赏榜只有月榜
	}

	params["page"] = strconv.Itoa(r.Page)
	params["rtype"] = r.TypeName
	params["size"] = strconv.Itoa(r.Size)
	api.Get(r.rank_api()).AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}
