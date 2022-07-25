package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaActividadMeta_20191015_112436 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaActividadMeta_20191015_112436{}
	m.Created = "20191015_112436"

	migration.Register("CrearTablaActividadMeta_20191015_112436", m)
}

// Run the migrations
func (m *CrearTablaActividadMeta_20191015_112436) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.actividad_meta (id serial NOT NULL,actividad_id varchar NOT NULL,meta_rubro_necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_actividad_meta PRIMARY KEY (id));")
	m.SQL("CREATE TABLE necesidades.fuente_actividad (id serial NOT NULL,fuente_id varchar NOT NULL,actividad_meta_necesidad_id int4 NOT NULL,monto_parcial numeric(20,7) NOT NULL,activo bool NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_fuente_actividad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaActividadMeta_20191015_112436) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.actividad_meta;")
	m.SQL("DROP TABLE IF EXISTS necesidades.fuente_actividad;")

}
