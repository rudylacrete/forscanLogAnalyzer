package models

var ForscanFields = []string{
	"rpm",
	"vss",
	"turbo",
	"fuelpw",
}

type ForscanLogs struct {
	Fields []string    `json:"fields"`
	Values [][]float64 `json:"values"`
}
