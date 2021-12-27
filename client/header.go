package client

import (
	c "github.com/morristai/rarbg-notifier/common"
	"github.com/mozillazg/go-httpheader"
	"log"
	"net/http"
)

func NewHeader(cookie string) http.Header {
	opt := c.Headers{
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		AcceptEncoding:          "br",
		AcceptLanguage:          "en-US",
		CacheControl:            "max-age=0",
		Connection:              "keep-alive",
		Cookie:                  cookie,
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
		log.Fatalln(err)
	}
	return h
}

func NewImdbHeader(imdbCookie string) http.Header {
	h := NewHeader("")
	h.Set("Cookie", imdbCookie)
	return h
}
