package common

type Headers struct {
	Accept                  string `header:"Accept"`
	AcceptEncoding          string `header:"Accept-Encoding"`
	AcceptLanguage          string `header:"Accept-Language"`
	CacheControl            string `header:"Cache-Control"`
	Connection              string `header:"Connection"`
	Cookie                  string `header:"Cookie"`
	Dnt                     string `header:"DNT"`
	Host                    string `header:"HOST"`
	Referer                 string `header:"Referer"`
	SecChUa                 string `header:"sec-ch-ua"`
	SecChUaMobile           string `header:"sec-ch-ua-mobile"`
	SecChUaPlatform         string `header:"sec-ch-ua-platform"`
	SecFetchDest            string `header:"Sec-Fetch-Dest"`
	SecFetchMode            string `header:"Sec-Fetch-Mode"`
	SecFetchSite            string `header:"Sec-Fetch-Site"`
	SecFetchUser            string `header:"Sec-Fetch-User"`
	UserAgent               string `header:"User-Agent"`
	UpgradeInsecureRequests string `header:"Upgrade-Insecure-Requests"`
}

type VideoInfo struct {
	Title      string `json:"title"`
	Year       uint16 `json:"year"`
	Poster     string `json:"poster"`
	Size       string `json:"size"`
	Resolution string `json:"resolution,omitempty"`
	Source     string `json:"source,omitempty"`
	Formats    string `json:"formats,omitempty"`
	Audio      string `json:"audio,omitempty"`
	Encoding   string `json:"encoding,omitempty"`
	Language   string `json:"language,omitempty"`
	Imdb       string `json:"imdb,omitempty"`
	Others     string `json:"others,omitempty"`
}
