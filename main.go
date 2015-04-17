package sql2json

import (
	"database/sql"
	"encoding/json"
	"strings"
)

func Run(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)

	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for i, _ := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return "", err
		}

		entry := make(map[string]interface{}, 0)
		for i, key := range columns {
			var v interface{}
			rawValue := values[i]

			s, ok := rawValue.([]byte)
			if ok {
				v = strings.TrimSpace(string(s))
			} else {
				v = rawValue
			}

			entry[key] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), err
}
