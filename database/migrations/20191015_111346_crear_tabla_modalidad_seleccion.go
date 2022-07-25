package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaModalidadSeleccion_20191015_111346 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaModalidadSeleccion_20191015_111346{}
	m.Created = "20191015_111346"

	migration.Register("CrearTablaModalidadSeleccion_20191015_111346", m)
}

// Run the migrations
func (m *CrearTablaModalidadSeleccion_20191015_111346) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.modalidad_seleccion (id serial NOT NULL,nombre varchar,descripcion varchar,codigo_abreviacion varchar NOT NULL,numero_orden integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_modalidad_seleccion PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaModalidadSeleccion_20191015_111346) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.modalidad_seleccion;")

}
