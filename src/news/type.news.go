package news

// Source is variable to store source url
var Source string

// News object represent detail of the news
type News struct {
	title   string
	author  string
	date    int64
	content string
	image   string
	images  []string
}

// Item object represent small portion of the news.
// Can be used for display data on list
type Item struct {
	Title       string
	date        int64
	description string
	image       string
}

// NewNews is method to create News object (constructor)
func NewNews(title string, date int64, image string, content string) News {
	return News{
		title:   title,
		date:    date,
		image:   image,
		content: content,
	}
}

// NewItem is method to create Item object (constructor)
func NewItem(title string, date int64, description string, image string) Item {
	return Item{
		Title:       title,
		date:        date,
		description: description,
		image:       image,
	}
}
