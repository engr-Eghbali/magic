package myplugins

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}

func Open(dbtype string, dbuser string, dbpass string, dbname string) (db *sql.DB, err error) {
	db, err = sql.Open(dbtype, dbuser+":"+dbpass+"@/"+dbname+"?charset=utf8")
	return db, err
}
func Close(db *sql.DB) {
	db.Close()
}

/////open connection

func DBexec(stmt *sql.Stmt, args ...interface{}) (res sql.Result, err error) {
	res, err = stmt.Exec(args)
	return res, err
}

func Insert(db *sql.DB, table string, columns string, values string) (stmt *sql.Stmt, err error) {

	stmt, err = db.Prepare("INSERT INTO `" + table + "`(" + columns + ") VALUES (" + values + ")")
	checkErr(err)
	return stmt, err

}

/////////////insert
