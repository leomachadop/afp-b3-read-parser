package model

type StockQuote struct {
	DataPregao string  `json:"data-pregao"`
	CODBDI     string  `json:"codbdi"`
	CODNEG     string  `json:"codneg"`
	TPMERC     string  `json:"tpmerc"`
	NOMRES     string  `json:"nomres"`
	ESPECI     string  `json:"especi"`
	PRAZOT     string  `json:"prazot"`
	MODREF     string  `json:"modref"`
	PREABE     float64 `json:"preabe"`
	PREMAX     float64 `json:"premax"`
	PREMIN     float64 `json:"premin"`
	PREMED     float64 `json:"premed"`
	PREULT     float64 `json:"preult"`
	PREOFC     float64 `json:"preofc"`
	PREOFV     float64 `json:"preofv"`
	TOTNEG     int     `json:"totneg"`
	QUATOT     int     `json:"quatot"`
	VOLTOT     int     `json:"voltot"`
	PREEXE     float64 `json:"preexe"`
	INDOPC     string  `json:"indopc"`
	DATVEN     string  `json:"datven"`
	FATCOT     string  `json:"fatcot"`
	PTOEXE     string  `json:"ptoexe"`
	CODISI     string  `json:"codisi"`
	DISMES     string  `json:"dismes"`
}
