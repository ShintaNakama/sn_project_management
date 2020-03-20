package main

import (
	"log"
	"net/http"

  _ "github.com/go-sql-driver/mysql"

	"github.com/ShintaNakama/sn_project_management/app"
)

const dsn = "root@tcp(db)/sn_project_management"

func main() {
	//conn, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	return nil, err
	//}
	//db := DB{conn: conn}

	//conn.SetConnMaxLifetime(10 * time.Second)
	//conn.SetMaxOpenConns(10)
	//conn.SetMaxIdleConns(10)

	//if err := conn.Ping(); err != nil {
	//	return nil, err
	//}
	if err := http.ListenAndServe(":8888", app.NewMux()); err != nil {
		log.Println(err)
	}
}

