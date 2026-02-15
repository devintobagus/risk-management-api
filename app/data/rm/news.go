package rm

type NewsRequest struct {
	Query string `json:"query" form:"query" validate:"required"`
}

type NewsResponse struct {
	Title       string     `json:"title"`
	Link        string     `json:"link"`
	PublishDate string     `json:"publish_date"`
	Source      SourceData `json:"source"`
}

type SourceData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
