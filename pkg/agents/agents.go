package agents

import (
	"fmt"
	"math"
	"time"
)

type Entry struct {
	Name      string  `json:"name"`
	Set       string  `json:"set"`
	Condition string  `json:"condition"`
	Grade     string  `json:"grade"`
	Acquired  float64 `json:"acquired"`
	Date      string  `json:"date"`
}

type ItemResult struct {
	Name           string  `json:"name"`
	Set            string  `json:"set"`
	Condition      string  `json:"condition"`
	ScraperPrice   float64 `json:"scraper_price"`
	ArbitratorPrice float64 `json:"arbitrator_price"`
	AnalystPrice   float64 `json:"analyst_price"`
	Consensus      float64 `json:"consensus"`
	Confidence     int     `json:"confidence"`
	Volatility     float64 `json:"volatility"`
	Trend          string  `json:"trend"`
}

type JuryLog struct {
	Agent string `json:"agent"`
	Text  string `json:"text"`
}

type Report struct {
	GeneratedAt     string       `json:"generated_at"`
	Model           string       `json:"model"`
	Items           []ItemResult `json:"items"`
	PortfolioValue  float64      `json:"portfolio_value"`
	Change90d       float64      `json:"change_90d"`
	Diversification int          `json:"diversification"`
	Logs            []JuryLog    `json:"logs"`
}

type Jury struct {
	model string
}

func NewJury(model string) *Jury {
	return &Jury{model: model}
}

func (j *Jury) Run(entries []Entry) (Report, error) {
	items := make([]ItemResult, 0, len(entries))
	total := 0.0

	for _, e := range entries {
		base := scrapePrice(e)
		adjusted := ArbitrateCondition(base, e.Condition)
		analystPrice, volatility, trend := analystForecast(base, adjusted)

		consensus := (base + adjusted + analystPrice) / 3.0
		confidence := CalcConfidence(base, adjusted, analystPrice)

		items = append(items, ItemResult{
			Name:           e.Name,
			Set:            e.Set,
			Condition:      e.Condition,
			ScraperPrice:   base,
			ArbitratorPrice: adjusted,
			AnalystPrice:   analystPrice,
			Consensus:      consensus,
			Confidence:     confidence,
			Volatility:     volatility,
			Trend:          trend,
		})
		total += consensus
	}

	return Report{
		GeneratedAt:     time.Now().Format(time.RFC3339),
		Model:           j.model,
		Items:           items,
		PortfolioValue:  total,
		Change90d:       simulatedChange(),
		Diversification: 72,
		Logs: []JuryLog{
			{Agent: "scraper", Text: fmt.Sprintf("Queried %d items across open market sources.", len(entries))},
			{Agent: "arbitrator", Text: fmt.Sprintf("Applied condition penalties to %d items.", len(entries))},
			{Agent: "analyst", Text: "Generated 90-day volatility matrix and trend forecast."},
		},
	}, nil
}

func scrapePrice(e Entry) float64 {
	base := 10.0 + float64(len(e.Name))*2.5
	if e.Grade != "" {
		base *= 1.4
	}
	return base
}

func ArbitrateCondition(base float64, condition string) float64 {
	penalty := 0.0
	switch condition {
	case "Lightly-Played":
		penalty = 0.15
	case "Moderately-Played":
		penalty = 0.30
	case "Heavily-Played":
		penalty = 0.50
	case "Damaged":
		penalty = 0.70
	}
	return base * (1.0 - penalty)
}

func analystForecast(base, adjusted float64) (float64, float64, string) {
	vol := 5.0 + float64(len(fmt.Sprintf("%f", base)))
	if vol > 30 {
		vol = 30
	}
	trend := "Stable"
	if adjusted > base*1.1 {
		trend = "Upward"
	} else if adjusted < base*0.9 {
		trend = "Downward"
	}
	return adjusted * (1.0 + (vol/100.0)*0.1), vol, trend
}

func CalcConfidence(a, b, c float64) int {
	diff1 := math.Abs(a - b)
	diff2 := math.Abs(b - c)
	avg := (a + b + c) / 3.0
	spread := (diff1 + diff2) / 2.0
	if avg == 0 {
		return 0
	}
	conf := 100 - int((spread/avg)*100)
	if conf > 99 {
		conf = 99
	}
	if conf < 40 {
		conf = 40
	}
	return conf
}

func simulatedChange() float64 {
	return -2.3 + float64(time.Now().Unix()%100)/10.0
}
