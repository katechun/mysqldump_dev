package tools

import (
	mysql "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smallnest/rpcx/log"
	"strconv"
)

func init() {
	dao, err := mysql.Open(DbInfo.DbType, DbInfo.DbUser+":"+DbInfo.DbPass+"@tcp("+DbInfo.DbHost+":"+strconv.Itoa(DbInfo.DbPort)+")/"+DbInfo.DbName+"?charset="+DbInfo.DbCharset)
	if err != nil {
		log.Fatal(err)
	}
}

func Conn() (*mysql.DB, error) {
	dao, err := mysql.Open(DbInfo.DbType, DbInfo.DbUser+":"+DbInfo.DbPass+"@tcp("+DbInfo.DbHost+":"+strconv.Itoa(DbInfo.DbPort)+")/"+DbInfo.DbName+"?charset="+DbInfo.DbCharset)
	if err != nil {
		log.Fatal(err)
	}
	return dao, nil
}

func Prepare(Conn, sql string) (*mysql.Stmt, error) {
	return Conn.Prepare(sql)
}

func Query(stmt mysql.Stmt, values []interface{}) (*mysql.Rows, error) {
	rows, err := stmt.Query(values...)
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}
