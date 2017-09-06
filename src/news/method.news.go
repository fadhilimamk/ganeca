package news

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fadhilimamk/ganeca/src/global"
	"github.com/fadhilimamk/ganeca/src/log"
)

func fecthItems() {

	// Fetching All Item is atomic process, to avoid collision with fetchNews
	mutex.Lock()

	log.Info("Fetching news items ...")
	start := time.Now()
	ItemDataResult := []Item{}

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

	var i int64
	doc.Find(".p-t-b-10").Each(func(index int, item *goquery.Selection) {
		meta := item.ChildrenFiltered(".col-md-10")

		title := meta.
			ChildrenFiltered("a").
			ChildrenFiltered("strong").
			Text()

		dateString := meta.ChildrenFiltered("span").Text()
		description := meta.ChildrenFiltered("p").Text()
		image, _ := item.ChildrenFiltered(".col-md-2").ChildrenFiltered("a").ChildrenFiltered("img").Attr("src")
		url, _ := meta.ChildrenFiltered("a").Attr("href")

		dateTime := global.DateStringToTime(dateString)
		date := dateTime.Unix()

		// fixing leading and trailing whitespace
		description = strings.TrimSpace(description)
		description = strings.Replace(description, "\n", " ", -1)

		newsItem := NewItem(i, title, date, description, image, url)
		if strings.TrimSpace(title) != "" {
			ItemDataResult = append(ItemDataResult, newsItem)
			i++
		}

	})

	if len(ItemDataResult) > 0 {
		ItemData = ItemDataResult
	} else {
		log.Info("Got empty news item data! Using older version instead")
	}

	log.Info("Finish fetching news items. Got ", len(ItemData), " items on ", (time.Now().Sub(start).Nanoseconds()), "ns")
	UpdatedAt = time.Now()

	mutex.Unlock()

}

func fetchNews() {

	mutex.Lock()

	log.Info("Fetching news item details ...")

	for _, item := range ItemData {
		doc, err := goquery.NewDocument(item.URL)
		if err != nil {
			log.Fatal(err)
		}

		rawData := doc.Find(".view")

		title := item.Title
		image := item.Image
		date := item.Date
		rawAuthor := strings.TrimSpace(strings.Replace(rawData.Find(".date2").Text(), "\n", "", -1))
		rawAuthor = global.RemoveDuplicateSpaceInString(rawAuthor)
		authorRaw := strings.Split(rawAuthor, "-")
		author := authorRaw[0]

		// getting content
		var buffer bytes.Buffer
		rawData.Find("p").Each(func(index int, item *goquery.Selection) {
			rawContent := strings.TrimSpace(item.Text())
			buffer.WriteString(rawContent)
			buffer.WriteString("\n")
			buffer.WriteString("\n")
		})
		content := buffer.String()

		// list all images
		images := []string{}
		rawData.Find("img").Each(func(index int, item *goquery.Selection) {
			img, _ := item.Attr("src")
			if img != image {
				images = append(images, img)
			}
		})

		NewsData[item.ID] = NewNews(title, author, date, content, image, images)
	}

	mutex.Unlock()

}

// GetNewsDetail return news with specific id
func GetNewsDetail(id int64) News {
	return NewsData[id]
}
