package rank

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

type Rank struct {
	TypeName string
	Month    bool
	All      bool
	Page     int
	Size     int
}

var TypeName = struct {
	ViewRank   string
	BestRank   string
	NewRank    string
	TicketRank string
	MarkRank   string
	BonusRank  string
}{
	ViewRank:   "view",
	BestRank:   "sale",
	NewRank:    "newhit",
	TicketRank: "ticket",
	MarkRank:   "mark",
	BonusRank:  "bonus",
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
		r.Size = 50
	}
	if r.TypeName == "sale" && r.Size > 40 {
		r.Size = 40 // 畅销榜最大值为40
		r.Page = 0  // 畅销榜只有一页
	}
	if r.TypeName == "newhit" || r.TypeName == "bonus" {
		r.Month = true // 新作榜和打赏榜只有月榜
	}

	params["page"] = strconv.Itoa(r.Page)
	params["rtype"] = r.TypeName
	params["size"] = strconv.Itoa(r.Size)
	request.Get(r.rank_api()).AddAll(params).NewRequests().Unmarshal(&RankStruct)
	return RankStruct
}
