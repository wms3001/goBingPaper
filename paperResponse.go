package goBingPaper

type Paper struct {
	Startdate     string      `json:"startdate"`
	Fullstartdate string      `json:"fullstartdate"`
	Enddate       string      `json:"enddate"`
	Url           string      `json:"url"`
	Urlbase       string      `json:"urlbase"`
	Copyright     string      `json:"copyright"`
	Copyrightlink string      `json:"copyrightlink"`
	Title         string      `json:"title"`
	Quiz          string      `json:"quiz"`
	Wp            bool        `json:"wp"`
	Hsh           string      `json:"hsh"`
	Drk           int         `json:"drk"`
	Top           int         `json:"top"`
	Bot           int         `json:"bot"`
	Hs            interface{} `json:"hs"`
}

type PaperResponse struct {
	LastUpdate string  `json:"LastUpdate"`
	Total      string  `json:"Total"`
	Language   string  `json:"Language"`
	Message    string  `json:"message"`
	Status     bool    `json:"status"`
	Success    bool    `json:"success"`
	Data       []Paper `json:"data"`
}
