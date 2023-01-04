package boluobao

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
)

type Boluobao struct {
	Book   *book.Books
	Search *search.Search
}

var API = Boluobao{
	Book:   &book.Books{},
	Search: &search.Search{},
}
