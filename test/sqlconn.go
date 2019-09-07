package main

import (
	mysql "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type DbInfo struct {
	DbType    string
	DbHost    string
	DbUser    string
	DbPass    string
	DbPort    int
	DbName    string
	DbCharset string
}

func main() {
	DbInfo := DbInfo{
		DbType:    "mysql",
		DbHost:    "127.0.0.1",
		DbUser:    "root",
		DbPass:    "123456",
		DbPort:    3306,
		DbName:    "test",
		DbCharset: "utf8",
	}

	var names []interface{}
	names = append(names, "test")
	sql := "select CONSTRAINT_NAME,TABLE_NAME,COLUMN_NAME,REFERENCED_TABLE_SCHEMA,REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME from information_schema.`KEY_COLUMN_USAGE` where REFERENCED_TABLE_SCHEMA = ?"
	dao, err := mysql.Open(DbInfo.DbType, DbInfo.DbUser+":"+DbInfo.DbPass+"@tcp("+DbInfo.DbHost+":"+strconv.Itoa(DbInfo.DbPort)+")/"+DbInfo.DbName+"?charset="+DbInfo.DbCharset)

	defer dao.Close()

	if err != nil {
		fmt.Println("Mysql connect error!")
		return
	}

	stmt, err := dao.Prepare(sql)
	if err != nil {
		fmt.Println("Mysql prepare error!")
		return
	}

	rows, err := stmt.Query(names...)

	if err != nil {
		fmt.Println("Mysql Query error!")
		return
	}

	columns, err := rows.Columns()
	vs := make([]mysql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range vs {
		scans[i] = &vs[i]
	}

	var result []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]interface{})
		for i, col := range vs {
			if col != nil {
				each[columns[i]] = FilterHolder(string(col))
			} else {
				each[columns[i]] = nil
			}
		}

		result = append(result, each)

	}
	fmt.Println(result)
}

func FilterHolder(content string) string {
	newContent := ""
	for _, value := range content {
		if value != 65533 {
			newContent += string(value)
		}
	}
	return newContent
}
