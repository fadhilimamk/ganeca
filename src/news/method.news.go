package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fadhilimamk/ganeca/src/global"
)

func fecthItems() {
	fmt.Println("Fetching news items ...")
	start := time.Now()

	param := map[string][]string{
		"tahun": []string{"2017"},
	}

	resp, err := http.PostForm(Source, param)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".p-t-b-10").Each(func(index int, item *goquery.Selection) {
		meta := item.ChildrenFiltered(".col-md-10")

		title := meta.
			ChildrenFiltered("a").
			ChildrenFiltered("strong").
			Text()
		date := global.GetCurrentTime()
		description := meta.ChildrenFiltered("p").Text()
		image, _ := item.ChildrenFiltered(".col-md-2").ChildrenFiltered("a").ChildrenFiltered("img").Attr("src")

		// fixing leading and trailing whitespace
		description = strings.TrimSpace(description)
		description = strings.Replace(description, "\n", " ", -1)

		newsItem := NewItem(title, date, description, image)
		ItemData = append(ItemData, newsItem)

	})

	fmt.Println("Finish fetching news items. Got ", len(ItemData), " items on ", (time.Now().Sub(start).Nanoseconds()), "ns")
}
