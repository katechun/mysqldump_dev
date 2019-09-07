package export

import (
	mysql "database/sql"
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

type Dump struct {
	DbNames []string
	Tables  []string
	Rows    []string
}

type selectrows interface {
	Config()
	Connect()
	Select()
}

func (p Dump) SelectRows() []interface{} {

	DbInfo := DbInfo{
		DbType:    "mysql",
		DbHost:    "127.0.0.1",
		DbUser:    "root",
		DbPass:    "123456",
		DbPort:    3306,
		DbName:    "test",
		DbCharset: "utf8",
	}

	dao, err := mysql.Open(DbInfo.DbType, DbInfo.DbUser+":"+DbInfo.DbPass+"@tcp("+DbInfo.DbHost+":"+strconv.Itoa(DbInfo.DbPort)+")/"+DbInfo.DbName+"?charset="+DbInfo.DbCharset)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := dao.Prepare(sql)

}
