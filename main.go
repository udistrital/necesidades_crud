package main

import (
	_ "github.com/udistrital/necesidades_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	if beego.BConfig.RunMode == "dev" {
		// Database alias.
		// name := "default"

		// // Drop table and re-create.
		// force := false

		// // Print log.
		// verbose := true

		// Error.
		// err := orm.RunSyncdb(name, force, verbose)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

}

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerror.CustomErrorController{})
	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/necesidades_crud/necesidades_crud.log"}`)

	apistatus.Init()
	beego.Run()
}
