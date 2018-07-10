package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//fmt.Println("Hello world!")
	
	db, err := sql.Open("mysql", "root:mysql@/go")
	
	if err != nil {
		panic(err.Error())
	}
	
	// 後でクローズするようにスケジューリング
	defer db.Close()
	
	// SQL実行
	_, err = db.Exec("insert into tbl_sample values(99, 'ru-')")
	
	if err != nil {
		panic(err.Error())
	}
}