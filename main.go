package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"github.com/ShintaNakama/sn_project_management/app/interactor"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/middleware"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/router"
)

// go-sql-driverのmysqlの場合、datatime型などをselectする場合、パラメータ:parseTime=trueをつけてOpenしないとselect出来ない。
const dsn = "root@tcp(db)/sn_project_management?parseTime=true"

func main() {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	conn.SetConnMaxLifetime(10 * time.Second)
	//conn.LogMode(true)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	// echoを使わず標準パッケージでhttp.Handleを定義しようとしたが、controllerの各メソッドを実装したinterfaceを適用させることが出来なかった。。
	// interfaceだと、戻り値が w http.ResponseWriter, r *http.Request でもinterfaceを通して呼ぶと戻り値が(w http.ResponseWriter, r *http.Request)にはならない？？
	e := echo.New()

	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()

	i := interactor.NewInteractor(conn)
	c := i.NewAppController()

	router.NewRouter(e, c)
	middleware.NewMiddleware(e)
	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
