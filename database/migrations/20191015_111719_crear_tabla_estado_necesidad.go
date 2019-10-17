package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEstadoNecesidad_20191015_111719 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEstadoNecesidad_20191015_111719{}
	m.Created = "20191015_111719"

	migration.Register("CrearTablaEstadoNecesidad_20191015_111719", m)
}

// Run the migrations
func (m *CrearTablaEstadoNecesidad_20191015_111719) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.estado_necesidad (id serial NOT NULL,nombre varchar,descripcion varchar,codigo_abreviacionn varchar,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_estado_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.estado_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaEstadoNecesidad_20191015_111719) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.estado_necesidad;")

}
