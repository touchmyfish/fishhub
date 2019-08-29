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

	if err := beego.LoadAppConfig("ini", *confPath); err != nil {
		panic(fmt.Sprintf("Beego load app config failed, err: %s", err))
	}

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
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

	if err := orm.RegisterDataBase("default", "postgres", sqlConn); err != nil {
		panic(err)
	}

	if err := orm.RunSyncdb("default", true, true); err != nil {
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
