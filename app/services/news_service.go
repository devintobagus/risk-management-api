package services

import (
	"rm/app/clients"
	"rm/app/data/rm"
)

type NewsService struct {
	googleNewsClient clients.GoogleNewsClient
}

func NewNewsService() *NewsService {
	return &NewsService{
		googleNewsClient: *clients.NewGoogleNewsClient(),
	}
}

func (s *NewsService) News(
	query string,
) (*[]rm.NewsResponse, error) {
	rss, err := s.googleNewsClient.GetNewsRss(query)
	if err != nil {
		return nil, err
	}
	var news []rm.NewsResponse
	for _, rssItem := range rss.Channel.Items {
		news = append(news, rm.NewsResponse{
			Title:       rssItem.Title,
			Link:        rssItem.Link,
			PublishDate: rssItem.PubDate,
			Source: rm.SourceData{
				Name: rssItem.Source.Name,
				Url:  rssItem.Source.URL,
			},
		})
	}

	return &news, nil
}
