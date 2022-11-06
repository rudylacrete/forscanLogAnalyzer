package models

var ForscanFields = []string{
	"RPM",
	"VSS",
	"TURBO",
}

type ForscanLogs struct {
	Fields []string    `json:"fields"`
	Values [][]float64 `json:"values"`
}
