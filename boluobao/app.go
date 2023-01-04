package boluobao

import "github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"

type Boluobao struct {
	Book *book.Books
}

var API = Boluobao{
	Book: &book.Books{},
}
