package news

// Source is variable to store source url
var Source string

// News object represent detail of the news
type News struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Date    int64    `json:"date"`
	Content string   `json:"content"`
	Image   string   `json:"image"`
	Images  []string `json:"images"`
}

// Item object represent small portion of the news.
// Can be used for display data on list
type Item struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Date        int64  `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
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
