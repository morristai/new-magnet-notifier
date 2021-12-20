package client

import (
	"github.com/mozillazg/go-httpheader"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Options struct {
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

func NewHeader() http.Header {
	opt := Options{
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		AcceptEncoding:          "br",
		AcceptLanguage:          "en-US",
		CacheControl:            "max-age=0",
		Connection:              "keep-alive",
		Cookie:                  "tcc; c_cookie=3buqk19y2s; gaDts48g=q8h5pp9t; aby=2; use_alt_cdn=1; skt=GZV8Homxwp; skt=GZV8Homxwp; gaDts48g=q8h5pp9t; ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; expla=1; ppu_sub_9ef78edf998c4df1e1636c9a474d9f47=3; ppu_delay_9ef78edf998c4df1e1636c9a474d9f47=1",
		Dnt:                     "1",
		Host:                    "rarbg.to",
		Referer:                 "https://rarbg.to/torrents.php",
		SecChUa:                 "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Microsoft Edge\";v=\"96\"",
		SecChUaMobile:           "?0",
		SecChUaPlatform:         "macOS",
		SecFetchDest:            "document",
		SecFetchMode:            "navigate",
		SecFetchSite:            "same-origin",
		SecFetchUser:            "?1",
		UserAgent:               "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.62",
		UpgradeInsecureRequests: "1",
	}

	h, err := httpheader.Header(opt)
	if err != nil {
		log.Fatal(err)
	}
	return h
}
