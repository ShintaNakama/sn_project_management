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

func getMysqlConn() *gorp.DbMap {
	// go-sql-driverのmysqlの場合、datatime型などをselectする場合、パラメータ:parseTime=trueをつけてOpenしないとselect出来ない。
	//dsn := "root@tcp(db)/sn_project_management?loc=Local&parseTime=true"
	//dsn := "root:12345678@tcp(127.0.0.1:3306)/sn_project_management?parseTime=true"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/sn_project_management?loc=Local&parseTime=true", user, pass, host, port)
	//dsn := "root@tcp(db)/sn_project_management?loc=Local&parseTime=true"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//conn.SetConnMaxLifetime(10 * time.Second)
	////conn.LogMode(true)
	//conn.SetMaxOpenConns(10)
	//conn.SetMaxIdleConns(10)

	//if err := conn.Ping(); err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if err := conn.Close(); err != nil {
	//		e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
	//	}
	//}()

	dbmap := &gorp.DbMap{Db: conn, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(model.Project{}, "projects")
	dbmap.AddTableWithName(model.Task{}, "tasks")
	//project := dbmap.AddTableWithName(model.Project{}, "projects")
	//project.ColMap("ID").Rename("id")
	//project.ColMap("Name").Rename("name")
	//project.ColMap("Description").Rename("description")
	//project.ColMap("CreatedAt").Rename("created_at")
	//project.ColMap("UpdatedAt").Rename("updated_at")
	//project.ColMap("Completed").Rename("completed")

	return dbmap
}
