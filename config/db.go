package db

import (
	"database/sql"
	"fmt"

	//Mysql ...
	_ "github.com/go-sql-driver/mysql"
)

// Con ...
// var Con, err = sql.Open("mysql", "auraj25xu2pirozw:ddiuegczyopso39b@tcp(jw0ch9vofhcajqg7.cbetxkdyhwsb.us-east-1.rds.amazonaws.com)/kzi50q30ap322cl4")

var Con, err = sql.Open("mysql", "jv:libra2010@/db_clinica")

//TestarConn ...
func TestarConn() {
	if err != nil {
		panic(err.Error())
	}

	if Con.Ping() != nil {
		panic(err.Error())
	}

	fmt.Println("Banco: OK")
}
