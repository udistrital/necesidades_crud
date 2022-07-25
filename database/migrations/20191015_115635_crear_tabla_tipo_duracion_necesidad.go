package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoDuracionNecesidad_20191015_115635 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoDuracionNecesidad_20191015_115635{}
	m.Created = "20191015_115635"

	migration.Register("CrearTablaTipoDuracionNecesidad_20191015_115635", m)
}

// Run the migrations
func (m *CrearTablaTipoDuracionNecesidad_20191015_115635) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.tipo_duracion_necesidad (id serial NOT NULL,nombre varchar,descripcion varchar,codigo_abreviacion varchar,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_tipo_duracion_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaTipoDuracionNecesidad_20191015_115635) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.tipo_duracion_necesidad;")

}
