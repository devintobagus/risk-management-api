package google

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Source  Source `xml:"source"`
}

type Source struct {
	Name string `xml:",chardata"`
	URL  string `xml:"url,attr"`
}
