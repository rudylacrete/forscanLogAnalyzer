package lib

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/rudylacrete/forscanLogAnalyzer/models"
)

type headerMapping map[string]int

func fprintf(logger io.Writer, format string, args ...interface{}) {
	if logger != nil {
		fmt.Fprintf(logger, format, args...)
	}
}

func parseHeader(header []string) (hm headerMapping, fields []string) {
	hm = make(headerMapping)
	for i, val := range header {
		for _, label := range models.ForscanFields {
			if strings.Contains(strings.ToLower(val), strings.ToLower(label)) {
				if _, ok := hm[label]; !ok {
					fields = append(fields, strings.ToLower(label))
				}
				hm[label] = i
			}
		}
	}
	return hm, fields
}

func ParseFile(file string, logger io.Writer) (logs *models.ForscanLogs, err error) {
	fprintf(logger, "Opening file %s", file)
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}
	parser := csv.NewReader(bytes.NewReader(f))
	parser.Comma = ';'
	headerParsed := false
	var hm headerMapping
	var fields []string
	logs = &models.ForscanLogs{}
	for {
		record, err := parser.Read()
		// no more lines to parse
		if err == io.EOF {
			break
		}
		if err != nil {
			fprintf(logger, "An error occured during csv parsing: %s", err)
			continue
		}
		if !headerParsed {
			hm, fields = parseHeader(record)
			fmt.Printf("Got headers: %v | fields = %v", hm, fields)
			logs.Fields = fields
			headerParsed = true
			continue
		}
		entry := make([]float64, 0, len(hm))
		for f, i := range hm {
			if record[i] == "-" {
				entry = append(entry, 0)
				continue
			}
			v, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				fprintf(logger, "An error occured while parsing field %s: %s", f, err)
				continue
			}
			entry = append(entry, v)
		}
		logs.Values = append(logs.Values, entry)
	}
	return logs, nil
}
