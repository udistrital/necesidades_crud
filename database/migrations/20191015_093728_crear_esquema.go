package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquema_20191015_093728 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquema_20191015_093728{}
	m.Created = "20191015_093728"

	migration.Register("CrearEsquema_20191015_093728", m)
}

// Run the migrations
func (m *CrearEsquema_20191015_093728) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS necesidades;")
	m.SQL("SET search_path TO pg_catalog,public,necesidades;")

}

// Reverse the migrations
func (m *CrearEsquema_20191015_093728) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS necesidades")

}
