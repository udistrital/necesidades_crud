package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Necesidad_20220208_232314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Necesidad_20220208_232314{}
	m.Created = "20220208_232314"

	migration.Register("Necesidad_20220208_232314", m)
}

// Run the migrations
func (m *Necesidad_20220208_232314) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE necesidades.necesidad RENAME COLUMN consecutivo_solicitud TO consecutivo;")
	m.SQL("ALTER TABLE necesidades.necesidad DROP COLUMN consecutivo_necesidad;")
}

// Reverse the migrations
func (m *Necesidad_20220208_232314) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE necesidades.necesidad ADD COLUMN consecutivo_necesidad integer NULL;")
	m.SQL("ALTER TABLE necesidades.necesidad RENAME COLUMN consecutivo TO consecutivo_solicitud;")
}
