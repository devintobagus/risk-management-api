package config

import "github.com/goravel/framework/facades"

func init() {
	config := facades.Config()
	config.Add("client", map[string]any{
		"gemini": map[string]any{
			"api_key": config.Env("GEMINI_API_KEY"),
		},
		"yahoo": map[string]any{
			"host": config.Env("YAHOO_API_BASE_URL"),
		},
		"appscript": map[string]any{
			"stock_list_url": config.Env("APPSCRIPT_STOCK_LIST_BASE_URL"),
		},
		"mistral": map[string]any{
			"api_key": config.Env(("MISTRAL_API_KEY")),
		},
		"google_news": map[string]any{
			"host": config.Env("GOOGLE_NEWS_BASE_URL"),
		},
	})
}
