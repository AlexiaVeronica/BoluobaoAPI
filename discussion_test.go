package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/discussion"
	"testing"
)

func TestDiscussion(t *testing.T) {
	discussion.Feeds(0, "hot")
}
