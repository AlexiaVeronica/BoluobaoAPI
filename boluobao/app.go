package boluobao

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
)

type Boluobao struct {
	Book   *book.Books
	Search *search.Search
	Rank   *rank.Rank
}

var API = Boluobao{
	Book:   &book.Books{},
	Search: &search.Search{},
	Rank:   &rank.Rank{},
}
