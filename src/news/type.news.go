package news

// Source is variable to store source url
var Source string

// News object represent detail of the news
type News struct {
	Title   string
	Author  string
	Date    int64
	Content string
	Image   string
	Images  []string
}

// Item object represent small portion of the news.
// Can be used for display data on list
type Item struct {
	ID          int64
	Title       string
	Date        int64
	Description string
	Image       string
	URL         string
}

// NewNews is method to create News object (constructor)
func NewNews(title string, author string, date int64, content string, image string, images []string) News {
	return News{
		Title:   title,
		Author:  author,
		Date:    date,
		Content: content,
		Image:   image,
		Images:  images,
	}
}

// NewItem is method to create Item object (constructor)
func NewItem(id int64, title string, date int64, description string, image string, url string) Item {
	return Item{
		ID:          id,
		Title:       title,
		Date:        date,
		Description: description,
		Image:       image,
		URL:         url,
	}
}
