package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaActividadEspecificaNecesidad_20191015_115027 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaActividadEspecificaNecesidad_20191015_115027{}
	m.Created = "20191015_115027"

	migration.Register("CrearTablaActividadEspecificaNecesidad_20191015_115027", m)
}

// Run the migrations
func (m *CrearTablaActividadEspecificaNecesidad_20191015_115027) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.actividad_especifica_necesidad (id serial NOT NULL,necesidad_id integer NOT NULL,descripcion varchar NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_actividad_especifica_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.actividad_especifica_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaActividadEspecificaNecesidad_20191015_115027) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.actividad_especifica_necesidad;")

}
