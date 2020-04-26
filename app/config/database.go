package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/go-gorp/gorp"
)

func NewDBConnection() *gorp.DbMap {
	return getMysqlConn()
}

//func cloudSqlConn(user, pass, instance, env string) (*sql.DB, error) {
//	cfg := mysql.Cfg(instance, user, pass)
//	cfg.DBName = "sn_project_management"
//	conn, err := mysql.DialCfg(cfg)
//	return conn, err
//}

func getMysqlConn() *gorp.DbMap {
	var dsn string

	if env := os.Getenv("PROJECT_ENV"); env == "production" {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		instance := os.Getenv("INSTANCE_CONNECTION_NAME")
		//conn, err := cloudSqlConn(user, pass, instance, env)
		dsn = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s",
			user,
			pass,
			instance,
			"sn_project_management?loc=Local&parseTime=true")
	} else {
		// go-sql-driverのmysqlの場合、datatime型などをselectする場合、パラメータ:parseTime=trueをつけてOpenしないとselect出来ない。
		//dsn := "root:12345678@tcp(127.0.0.1:3306)/sn_project_management?parseTime=true"
		//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/sn_project_management?loc=Local&parseTime=true", user, pass, host, port)
		// docker
		dsn = "root@tcp(db)/sn_project_management?loc=Local&parseTime=true"
	}
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: conn, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(model.Project{}, "projects")
	dbmap.AddTableWithName(model.Task{}, "tasks")

	return dbmap
}
