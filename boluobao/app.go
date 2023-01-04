package boluobao

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/bookshelf"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/booktag"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/ranking"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/recommend"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/task"
)

type Boluobao struct {
	Book      *book.Books
	Search    *search.Search
	Rank      *ranking.Rank
	BookShelf *bookshelf.BookShelf
	Tags      *booktag.Tags
	User      *account.User
	Recommend *recommend.Recommend
	Task      *task.Task
}

var API = Boluobao{
	Book:      &book.Books{},
	Search:    &search.Search{},
	Rank:      &ranking.Rank{},
	BookShelf: &bookshelf.BookShelf{},
	Tags:      &booktag.Tags{},
	User:      &account.User{},
	Recommend: &recommend.Recommend{},
	Task:      &task.Task{},
}
