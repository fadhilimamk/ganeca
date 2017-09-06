// Package news responsible to fetching data and put it on exported variable,
// so another package can access it easily. This package also provide helper
// function to ease data processing.
package news

import (
	"sync"
	"time"

	"github.com/robfig/cron"
)

// ItemData is in memory cache to store all crawled news item data
var ItemData []Item

// NewsData is in memory cache to store all crawled news data
var NewsData map[int64]News

// UpdatedAt is last crawled time
var UpdatedAt time.Time

var mutex = &sync.Mutex{}

func init() {
	Source = "https://www.itb.ac.id/news/index/category/home"

	ItemData = []Item{}
	NewsData = map[int64]News{}

	c := cron.New()
	c.AddFunc("@every 5h0m0s", fecthItems)
	c.AddFunc("@every 2h0m0s", fetchNews)

}

// Init is method to get news data if empty
func Init() {
	if len(ItemData) == 0 {
		fecthItems()
		fetchNews()
	}
}
