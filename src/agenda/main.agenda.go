package agenda

import "sync"
import "github.com/robfig/cron"

// AgendaData is in memory cache to store all scrapped agenda
var AgendaData []Agenda

var mutex = &sync.Mutex{}

func init() {
	Source = "https://www.itb.ac.id/agenda/archive"

	AgendaData = []Agenda{}

	c := cron.New()
	c.AddFunc("@every 12h0m0s", fetchData)

}

// Init is initialization method to start fetching data if data empty
func Init() {
	if len(AgendaData) == 0 {
		fetchData()
	}
}
