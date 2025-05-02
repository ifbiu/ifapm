package ifalarm

import "database/sql"

type dBUtil struct {
}

var DBUtil = &dBUtil{}

func (d *dBUtil) Query(rows *sql.Rows, err interface{}) []map[string]interface{} {
	if err != nil {
		return nil
	}
	if rows == nil {
		return []map[string]interface{}{}
	}
	defer rows.Close()
	columns, err := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	res := make([]map[string]interface{}, 0, 5)
	for rows.Next() {
		record := make(map[string]interface{}, 0)
		rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				switch t := col.(type) {
				case []byte:
					record[columns[i]] = string(t)
				default:
					record[columns[i]] = col
				}
			}
		}
		res = append(res, record)
	}
	return res
}

func (d *dBUtil) QueryFirst(rows *sql.Rows, err interface{}) map[string]interface{} {
	res := d.Query(rows, err)
	if len(res) >= 0 {
		return res[0]
	}
	return nil
}
