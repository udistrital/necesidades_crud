package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMetaRubroNecesidad_20191015_115808 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMetaRubroNecesidad_20191015_115808{}
	m.Created = "20191015_115808"

	migration.Register("CrearTablaMetaRubroNecesidad_20191015_115808", m)
}

// Run the migrations
func (m *CrearTablaMetaRubroNecesidad_20191015_115808) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.meta_rubro_necesidad (id serial NOT NULL,meta_id varchar NOT NULL,rubro_necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_meta_rubro_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaMetaRubroNecesidad_20191015_115808) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.meta_rubro_necesidad;")

}
