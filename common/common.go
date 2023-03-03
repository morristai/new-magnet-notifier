package common

import (
	"time"
)

type Config struct {
	Rarbg struct {
		Headers Headers
		Movie   struct {
			Url1080p string
			Url2160p string
		}
		TvShow struct {
			Url1080p string
		}
	}
	Imdb struct {
		Cookie string
	}
	Cache struct {
		Json1080p string
		Json2160p string
	}
	Discord struct {
		Token    string
		Channels []string
		Members  []string
	}
	Log struct {
		Path    string
		Level   string
		IsDebug bool
	}
}

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

type LeaderboardCache struct { // Should I use pointer inside a struct?
	VideoList map[string]*VideoInfo
	Newest9   map[string]*VideoInfo
	Notified  map[string]time.Time // send discord notification time
	Time      time.Time
}

type VideoInfo struct {
	Url            string         `json:"url"`
	Title          string         `json:"title"`
	Year           int            `json:"year"`
	Poster         string         `json:"poster"`
	Size           string         `json:"size"`
	Genre          string         `json:"genre,omitempty"`
	Resolution     string         `json:"resolution,omitempty"`
	Source         string         `json:"source,omitempty"`
	Format         string         `json:"formats,omitempty"`
	Audio          string         `json:"audio,omitempty"`
	Encoding       string         `json:"encoding,omitempty"`
	Language       string         `json:"language,omitempty"`
	ImdbUrl        string         `json:"imdb,omitempty"`
	ProlificReview ProlificReview `json:"prolific_review,omitempty"`
	Rating         string         `json:"rating,omitempty"`
	Others         string         `json:"others,omitempty"`
	// HelpfulReview HelpfulReview `json:"helpfulReview,omitempty"`
}

type ProlificReview struct {
	Mean float32 `json:"mean,omitempty"`
	Std  float32 `json:"std,omitempty"`
	// Comments
}

type HelpfulReview struct {
	Mean float32 `json:"mean,omitempty"`
	Std  float32 `json:"std,omitempty"`
	// Comments
}
