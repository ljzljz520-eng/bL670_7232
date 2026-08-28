package importer

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"training-review/internal/service"
)

type CSVReader struct {
	reader  *csv.Reader
	headers map[string]int
}

func NewCSVReader(input io.Reader) (*CSVReader, error) {
	reader := csv.NewReader(input)
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("read csv header: %w", err)
	}
	positions := make(map[string]int, len(header))
	for index, value := range header {
		positions[strings.ToLower(strings.TrimSpace(value))] = index
	}
	return &CSVReader{reader: reader, headers: positions}, nil
}

func (r *CSVReader) value(row []string, key string) string {
	index, ok := r.headers[key]
	if !ok || index >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[index])
}

func (r *CSVReader) ReadRows() ([]service.ImportRow, []string, error) {
	rows := make([]service.ImportRow, 0)
	issues := make([]string, 0)
	line := 1
	for {
		line++
		record, err := r.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("read csv row %d: %w", line, err)
		}
		row := service.ImportRow{CourseID: r.value(record, "course_id"), Participant: r.value(record, "participant"), Email: r.value(record, "email"), Department: r.value(record, "department"), Tags: splitTags(r.value(record, "tags"))}
		if row.CourseID == "" || row.Participant == "" {
			issues = append(issues, fmt.Sprintf("row %d missing identity", line))
			continue
		}
		rows = append(rows, row)
	}
	return rows, issues, nil
}

func splitTags(value string) []string {
	if value == "" {
		return nil
	}
	pieces := strings.Split(value, "|")
	result := make([]string, 0, len(pieces))
	for _, piece := range pieces {
		if strings.TrimSpace(piece) != "" {
			result = append(result, strings.TrimSpace(piece))
		}
	}
	return result
}

func RequiredHeaders() []string {
	return []string{"course_id", "participant", "email", "department", "tags"}
}

func MissingHeaders(headers map[string]int) []string {
	missing := make([]string, 0)
	for _, name := range RequiredHeaders() {
		if _, ok := headers[name]; !ok {
			missing = append(missing, name)
		}
	}
	return missing
}
