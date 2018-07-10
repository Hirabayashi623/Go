package main

import (
	"fmt"
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
	rows, err := db.Query("select * from tbl_sample")
	
	if err != nil {
		panic(err.Error())
	}
	
	// カラム名の一覧を取得する
	columns, err := rows.Columns()
	
	if err != nil {
		panic(err.Error())
	}
	
	// カラム数分のデータの入れ物を準備してあげる
	values   := make([]sql.RawBytes, len(columns))
	// 
	scanArgs := make([]interface{}, len(values))
	fmt.Println(scanArgs)
	
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	// 取得したレコード分処理
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		
		var value string
		
		for i, col := range values {
			if col == nil {
				value = "NULL"
			}else{
				value = string(col)
			}
			
			fmt.Printf("%s : %s\n", columns[i], value)
		}
	}
}