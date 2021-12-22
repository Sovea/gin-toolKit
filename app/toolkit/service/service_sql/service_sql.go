package service_sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/toolKit")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
}
func GetDB() *sql.DB {
	return db
}
func GetJSON(sqlString string, data ...interface{}) (string, error) {
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer stmt.Close()
	rows, err := stmt.Query(data...)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JustDo(sqlString string, args ...interface{}) error {
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func GetMapData(sqlString string, args ...interface{}) ([]map[string]interface{}, error) {
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}

	return list, nil
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}
