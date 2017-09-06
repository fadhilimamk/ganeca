package news

import (
	"time"

	"github.com/robfig/cron"
)

// ItemData is in memory cache to store all crawled news item data
var ItemData []Item

// NewsData is in memory cache to store all crawled news data
var NewsData []News

// UpdatedAt is last crawled time
var UpdatedAt time.Time

func init() {
	Source = "https://www.itb.ac.id/news/index/category/home"

	c := cron.New()
	c.AddFunc("@every 6h0m0s", fecthItems)

}

// Init is method to get news data if empty
func Init() {
	if len(ItemData) == 0 {
		fecthItems()
		fetchNews()
	}
}
