package goBingPaper

import (
	"log"
	"testing"
)

func TestBingPaper_GetPaper(t *testing.T) {
	bingPaper := BingPaper{}
	res := bingPaper.GetPaper()
	log.Println(res)
}
