package main

import (
	"flag"
	"fmt"

	_ "fishhub/fish/models"
	_ "fishhub/fish/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	confPath := flag.String("c", "conf/app.conf", "Config.")

	flag.Parse()

	beego.LoadAppConfig("ini", *confPath)

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}

	dbName := beego.AppConfig.String("postgres::dbname")
	dbUser := beego.AppConfig.String("postgres::dbuser")
	dbHost := beego.AppConfig.String("postgres::dbhost")
	dbPassWord := beego.AppConfig.String("postgres::dbpassword")
	dbPort := beego.AppConfig.String("postgres::dbport")
	sslMode := beego.AppConfig.String("postgres::sslmode")

	sqlConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassWord, dbHost, dbPort, dbName, sslMode)

	logs.Info("PostgreSQL data source is %s", sqlConn)

	err = orm.RegisterDataBase("default", "postgres", sqlConn)
	if err != nil {
		panic(err)
	}
	err = orm.RunSyncdb("default", true, true)
	if err != nil {
		panic(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
