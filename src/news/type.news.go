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
	title       string
	date        int64
	description string
	image       string
	url         string
}

// NewNews is method to create News object (constructor)
func NewNews(title string, author string, date int64, content string, image string, images []string) News {
	return News{
		title:   title,
		author:  author,
		date:    date,
		content: content,
		image:   image,
		images:  images,
	}
}

// NewItem is method to create Item object (constructor)
func NewItem(title string, date int64, description string, image string, url string) Item {
	return Item{
		title:       title,
		date:        date,
		description: description,
		image:       image,
		url:         url,
	}
}

// BEGIN ---------------------- Method for Item object ------------------------

// GetTitle return item title
func (i Item) GetTitle() string {
	return i.title
}

// SetTitle change item title
func (i Item) SetTitle(title string) {
	i.title = title
}

// GetDate return item date
func (i Item) GetDate() int64 {
	return i.date
}

// SetDate change item date
func (i Item) SetDate(date int64) {
	i.date = date
}

// GetDescription return item description
func (i Item) GetDescription() string {
	return i.description
}

// SetDescription change item description
func (i Item) SetDescription(description string) {
	i.description = description
}

// GetImage return item image
func (i Item) GetImage() string {
	return i.image
}

// SetImage change item image
func (i Item) SetImage(image string) {
	i.image = image
}

// GetURL return item image
func (i Item) GetURL() string {
	return i.url
}

// SetURL change item image
func (i Item) SetURL(url string) {
	i.url = url
}
