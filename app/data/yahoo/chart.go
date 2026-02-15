package yahoo

// type YahooChartRequest struct {
// 	Stock    string    `json:"stock"`
// 	Range    RangeData `json:"range"`
// 	Interval RangeData `json:"interval"`
// }

type RangeData string

const (
	ONE_MINUTE  RangeData = "1m"
	ONE_HOUR    RangeData = "1h"
	ONE_DAY     RangeData = "1d"
	FIVE_DAYS   RangeData = "5d"
	ONE_MONTH   RangeData = "1mo"
	THREE_MONTH RangeData = "3mo"
	SIX_MONTH   RangeData = "6mo"
	ONE_YEAR    RangeData = "1y"
	TWO_YEAR    RangeData = "2y"
	FIVE_YEAR   RangeData = "5y"
	TEN_YEAR    RangeData = "10y"
	YTD         RangeData = "ytd"
	MAX         RangeData = "max"
)

type YahooChartResponse struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []ChartResult `json:"result"`
	Error  interface{}   `json:"error"`
}

type ChartResult struct {
	Meta       Meta       `json:"meta"`
	Timestamp  []int64    `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}

type Meta struct {
	Currency             string        `json:"currency"`
	Symbol               string        `json:"symbol"`
	ExchangeName         string        `json:"exchangeName"`
	FullExchangeName     string        `json:"fullExchangeName"`
	InstrumentType       string        `json:"instrumentType"`
	FirstTradeDate       int64         `json:"firstTradeDate"`
	RegularMarketTime    int64         `json:"regularMarketTime"`
	HasPrePostMarketData bool          `json:"hasPrePostMarketData"`
	GMTOffset            int           `json:"gmtoffset"`
	Timezone             string        `json:"timezone"`
	ExchangeTimezoneName string        `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64       `json:"regularMarketPrice"`
	FiftyTwoWeekHigh     float64       `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64       `json:"fiftyTwoWeekLow"`
	RegularMarketDayHigh float64       `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64       `json:"regularMarketDayLow"`
	RegularMarketVolume  int64         `json:"regularMarketVolume"`
	LongName             string        `json:"longName"`
	ShortName            string        `json:"shortName"`
	ChartPreviousClose   float64       `json:"chartPreviousClose"`
	PriceHint            int           `json:"priceHint"`
	CurrentTradingPeriod TradingPeriod `json:"currentTradingPeriod"`
	DataGranularity      string        `json:"dataGranularity"`
	Range                string        `json:"range"`
	ValidRanges          []string      `json:"validRanges"`
}

type TradingPeriod struct {
	Pre     MarketSession `json:"pre"`
	Regular MarketSession `json:"regular"`
	Post    MarketSession `json:"post"`
}

type MarketSession struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	GMTOffset int    `json:"gmtoffset"`
}

type Indicators struct {
	Quote    []Quote    `json:"quote"`
	AdjClose []AdjClose `json:"adjclose"`
}

type Quote struct {
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Close  []float64 `json:"close"`
	High   []float64 `json:"high"`
	Volume []int64   `json:"volume"`
}

type AdjClose struct {
	AdjClose []float64 `json:"adjclose"`
}
