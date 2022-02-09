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
	m.SQL("EXEC sp_rename 'necesidades.necesidad.consecutivo_solicitud', 'consecutivo', 'COLUMN';")
	m.SQL("ALTER TABLE necesidades.necesidad DROP COLUMN consecutivo_necesidad;")
}

// Reverse the migrations
func (m *Necesidad_20220208_232314) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("EXEC sp_rename 'necesidades.necesidad.consecutivo', 'consecutivo_solicitud', 'COLUMN';")
	m.SQL("ALTER TABLE necesidades.necesidad ADD consecutivo_necesidad int4 NULL;")
}
