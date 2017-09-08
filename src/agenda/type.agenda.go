package agenda

// Source is variable to store source url
var Source string

// Agenda : data stracture for agenda
type Agenda struct {
	ID    int64  `json:"id"`
	Date  int64  `json:"date"`
	Title string `json:"title"`
	Place string `json:"place"`
	Info  string `json:"info"`
}

// NewAgenda : constructor for Agenda object
func NewAgenda(id, date int64, title, place, info string) Agenda {
	return Agenda{
		ID:    id,
		Date:  date,
		Title: title,
		Place: place,
		Info:  info,
	}
}
