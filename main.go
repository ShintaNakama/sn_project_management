package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"github.com/ShintaNakama/sn_project_management/app/config"
	"github.com/ShintaNakama/sn_project_management/app/interactor"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/middleware"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/router"
)

func main() {
	conn := config.NewDBConnection()
	// echoを使わず標準パッケージでhttp.Handleを定義しようとしたが、controllerの各メソッドを実装したinterfaceを適用させることが出来なかった。。
	// interfaceだと、戻り値が w http.ResponseWriter, r *http.Request でもinterfaceを通して呼ぶと戻り値が(w http.ResponseWriter, r *http.Request)にはならない？？
	e := echo.New()

	i := interactor.NewInteractor(conn)
	c := i.NewAppController()

	router.NewRouter(e, c)
	middleware.NewMiddleware(e)
	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
