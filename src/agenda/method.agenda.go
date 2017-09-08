package agenda

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fadhilimamk/ganeca/src/global"
	"github.com/fadhilimamk/ganeca/src/log"
)

// AgendaDataResult is temporary slice for agenda while on scrapping
var AgendaDataResult []Agenda

func fetchData() {
	log.Info("Fetching events and egendas ...")
	start := time.Now()
	AgendaDataResult = []Agenda{}

	param := map[string][]string{
		"tahun": []string{"2017"},
	}

	resp, err := http.PostForm(Source, param)
	if err != nil {
		log.Fatal(err.Error())
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err.Error())
	}

	agendaRaw := doc.Find("tbody")

	var i int64
	agendaRaw.Find("tr").Each(func(index int, item *goquery.Selection) {

		// Extracting date
		dayString := item.Find(".dayofmonth").Text()
		day, _ := strconv.Atoi(dayString)
		monthYearMeta := item.Find(".shortdate").Text()
		monthYearData := strings.Split(monthYearMeta, ",")
		shortMonthString := strings.TrimSpace(monthYearData[0])
		month := global.ShortMonthToMonth(shortMonthString)
		yearString := strings.TrimSpace(monthYearData[1])
		year, _ := strconv.Atoi(yearString)
		date := global.MakeSimpleDate(year, month, day)

		// Extracting title
		title := item.Find(".agenda-events").Text()
		title = strings.TrimSpace(title)

		// Extracting url
		url, _ := item.Find("a").Attr("href")

		// Create Agenda Object
		agendaItem := NewAgenda(i, date.Unix(), title, "", "")

		// Appending and completion
		if strings.TrimSpace(title) != "" {
			AgendaDataResult = append(AgendaDataResult, agendaItem)
			go fetchDetail(i, url)
			i++
		}

	})

	if len(AgendaDataResult) != 0 {
		AgendaData = AgendaDataResult
	} else {
		log.Info("Got empty data! Using older version instead")
	}

	log.Info("Finish fetching agenda items. Got ", len(AgendaData), " items for ", (time.Now().Sub(start).Nanoseconds()), "ns")

}

func fetchDetail(idx int64, url string) {
	doc, _ := goquery.NewDocument(url)

	detail := doc.Find("dl")

	detail.Find("dt").Each(func(index int, item *goquery.Selection) {

		switch strings.ToLower(item.Text()) {
		case "where":
			AgendaDataResult[idx].Place = item.Next().Text()
			break
		case "description":
			AgendaDataResult[idx].Info = item.Next().Text()
			break
		}

	})

}
