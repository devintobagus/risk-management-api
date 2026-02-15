package services

import (
	"fmt"
	"rm/app/clients"
	"rm/app/data/rm"
	"strings"

	"github.com/goravel/framework/facades"
)

type NewsService struct {
	appscriptClient  clients.AppScriptClient
	googleNewsClient clients.GoogleNewsClient
	aiService        *AIService
}

func NewNewsService(ai *AIService) *NewsService {
	return &NewsService{
		appscriptClient:  *clients.NewAppScriptClient(),
		googleNewsClient: *clients.NewGoogleNewsClient(),
		aiService:        ai,
	}
}

func (s *NewsService) buildNewsQuery(n *rm.NewsKeyword) string {
	var parts []string

	for _, t := range n.Tickers {
		parts = append(parts, fmt.Sprintf(`"%s"`, t))
	}

	for _, c := range n.Companies {
		parts = append(parts, fmt.Sprintf(`"%s"`, c))
	}

	for _, i := range n.Industry {
		parts = append(parts, i)
	}

	for _, m := range n.Macro {
		parts = append(parts, fmt.Sprintf(`"%s"`, m))
	}

	for _, p := range n.People {
		parts = append(parts, fmt.Sprintf(`"%s"`, p))
	}

	for _, e := range n.Events {
		parts = append(parts, fmt.Sprintf(`"%s"`, e))
	}

	return strings.Join(parts, " OR ")
}
func (s *NewsService) StockNews(
	stock string,
) (*[]rm.NewsResponse, error) {
	keywordString, err := facades.Cache().RememberForever(
		fmt.Sprintf("services.news.stock-news.%s", stock),
		func() (any, error) {
			stockList, err := s.appscriptClient.GetStockList(stock)
			if err != nil {
				return nil, err
			}

			if len(stockList.Data) == 0 {
				return nil, fmt.Errorf("stock not found")
			}

			keywords, err := s.aiService.NewsKeyword(stock)
			if err != nil {
				return nil, err
			}

			keywordString := s.buildNewsQuery(keywords)
			return keywordString, nil
		})

	if err != nil {
		return nil, err
	}

	keyword := keywordString.(string)

	rss, err := s.googleNewsClient.GetNewsRss(keyword)
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
