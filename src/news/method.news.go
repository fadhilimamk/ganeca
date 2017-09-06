package news

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fadhilimamk/ganeca/src/global"
	"github.com/fadhilimamk/ganeca/src/log"
)

func fecthItems() {
	log.Info("Fetching news items ...")
	start := time.Now()
	ItemData = []Item{}

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
		dateString := meta.ChildrenFiltered("span").Text()
		description := meta.ChildrenFiltered("p").Text()
		image, _ := item.ChildrenFiltered(".col-md-2").ChildrenFiltered("a").ChildrenFiltered("img").Attr("src")
		url, _ := meta.ChildrenFiltered("a").Attr("href")

		dateTime := global.DateStringToTime(dateString)
		date := dateTime.Unix()

		// fixing leading and trailing whitespace
		description = strings.TrimSpace(description)
		description = strings.Replace(description, "\n", " ", -1)

		newsItem := NewItem(title, date, description, image, url)
		ItemData = append(ItemData, newsItem)

	})

	log.Info("Finish fetching news items. Got ", len(ItemData), " items on ", (time.Now().Sub(start).Nanoseconds()), "ns")
	UpdatedAt = time.Now()
}

func fetchNews() {
	log.Info("Fetching news item details ...")
	NewsData = []News{}

	for _, item := range ItemData {
		doc, err := goquery.NewDocument(item.GetURL())
		if err != nil {
			log.Fatal(err)
		}

		rawData := doc.Find(".view")

		title := item.GetTitle()
		image := item.GetImage()
		date := item.GetDate()
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

		NewsData = append(NewsData, NewNews(title, author, date, content, image, images))
	}

	for i, item := range NewsData {
		fmt.Printf("#%d\t%s\n", i, item.title)
		fmt.Printf("\t%s\n", item.author)
		fmt.Printf("\t%d\n", item.date)
		fmt.Printf("\t%s\n", item.content)
		fmt.Printf("\t%s\n", item.image)
		fmt.Println(item.images)
		fmt.Println()
	}

}
