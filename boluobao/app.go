package boluobao

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/bookshelf"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/recommend"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/tag"
)

type Boluobao struct {
	Book      *book.Books
	Search    *search.Search
	Rank      *rank.Rank
	BookShelf *bookshelf.BookShelf
	Tags      *tag.Tags
	User      *account.User
	Recommend *recommend.Recommend
}

var API = Boluobao{
	Book:      &book.Books{},
	Search:    &search.Search{},
	Rank:      &rank.Rank{},
	BookShelf: &bookshelf.BookShelf{},
	Tags:      &tag.Tags{},
	User:      &account.User{},
	Recommend: &recommend.Recommend{},
}
