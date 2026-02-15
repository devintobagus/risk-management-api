package rm

type NewsKeyword struct {
	Companies []string `json:"companies"`
	Tickers   []string `json:"tickers"`
	People    []string `json:"people"`
	Macro     []string `json:"macro"`
	Events    []string `json:"events"`
	Industry  []string `json:"industry"`
}

type NewsKeywordRequest struct {
	Query string `json:"query" form:"query" validate:"required"`
}
