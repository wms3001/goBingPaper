package goBingPaper

import (
	"encoding/json"
	"github.com/wms3001/http"
)

type BingPaper struct {
}

var url = "https://raw.onmicrosoft.cn/Bing-Wallpaper-Action/main/data/zh-CN_all.json"

func (bingPaper *BingPaper) GetPaper() PaperResponse {
	paperResponse := PaperResponse{}
	m := map[string]string{
		"Content-Type": "application/json",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	json.Unmarshal([]byte(resp.Data), &paperResponse)
	return paperResponse
}
