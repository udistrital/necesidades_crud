package crudoperationstestcases

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
)

func TestMain(m *testing.M) {
	pgUser := os.Getenv("NECESIDADES_CRUD_DB_USER")
	pgPass := os.Getenv("NECESIDADES_CRUD_DB_PASSWORD")
	pgUrls := os.Getenv("NECESIDADES_CRUD_DB_HOST")
	pgDb := os.Getenv("NECESIDADES_CRUD_DB_NAME")
	pgSchema := os.Getenv("NECESIDADES_CRUD_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

	exitVal := m.Run()

	os.Exit(exitVal)
}
