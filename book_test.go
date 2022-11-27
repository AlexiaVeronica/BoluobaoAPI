package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println(boluobao.GET_BOOK_INFORMATION("551946"))
}
func TestCATALOGUE(t *testing.T) {
	fmt.Println(boluobao.GET_CATALOGUE("551946"))
}
