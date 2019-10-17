package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearData_20191015_124707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearData_20191015_124707{}
	m.Created = "20191015_124707"

	migration.Register("CrearData_20191015_124707", m)
}

// Run the migrations
func (m *CrearData_20191015_124707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CrearData_20191015_124707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
