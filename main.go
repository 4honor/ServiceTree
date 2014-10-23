package main

import (
    "fmt"
	_ "ServiceTree/docs"
	_ "ServiceTree/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    
    db_type := beego.AppConfig.String("database::db_type")
    db_user := beego.AppConfig.String("database::username")
    db_password := beego.AppConfig.String("database::password")
    db_host := beego.AppConfig.String("database::host")
    db_port := beego.AppConfig.String("database::port")
    db_name := beego.AppConfig.String("database::dbname")

    connect_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",db_user, db_password, db_host, db_port, db_name)
    beego.SetLogger("file", `{"filename":"./logs/tree.log"}`)
	orm.RegisterDataBase("default", db_type, connect_str)
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

