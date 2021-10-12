package main

import (
	_ "github.com/udistrital/necesidades_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	"github.com/udistrital/auditoria"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+
		":"+beego.AppConfig.String("PGpass")+
		"@"+beego.AppConfig.String("PGurls")+
		"/"+beego.AppConfig.String("PGdb")+
		"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")
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

	//Prueba de auditoria
	auditoria.InitMiddleware()
	apistatus.Init()
	beego.Run()
}
