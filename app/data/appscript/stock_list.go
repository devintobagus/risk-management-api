package appscript

type StockListResponse struct {
	Success bool            `json:"success"`
	Data    []StockListData `json:"data"`
}

type StockListData struct {
	ID           int64  `json:"id"`
	Symbol       string `json:"symbol"`
	CompanyName  string `json:"company_name"`
	RegisterDate string `json:"register_date"`
	TotalStocks  string `json:"total_stocks"`
	Board        string `json:"board"`
}
