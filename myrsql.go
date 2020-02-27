package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func check(err error) { //因为要多次检查错误，所以干脆自己建立一个函数。
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	db, err := sql.Open("mysql", "desoon:K6vS4UcWwGgx@133@tcp(115.29.12.12:3306)/desoon")
	check(err)

	//query
	type info struct {
		order_sn      string    `db:"order_id"`
		userid    string `db:"userid"`
		//age     int    `db:"age"`
		//sex     string `db:"sex"`
		//salary  int    `db:"salary"`
		//work    string `db:"work"`
		//inparty string `db:"inparty"`
	}
	rows, err := db.Query("SELECT order_sn,userid FROM myr_order limit 0,10")
	check(err)
	for rows.Next() {
		var s info
		err = rows.Scan(&s.order_sn,&s.userid)
		check(err)
		fmt.Println(s)
	}
	rows.Close()
}