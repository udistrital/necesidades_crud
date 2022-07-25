package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNecesidadRechazada_20191015_114240 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNecesidadRechazada_20191015_114240{}
	m.Created = "20191015_114240"

	migration.Register("CrearTablaNecesidadRechazada_20191015_114240", m)
}

// Run the migrations
func (m *CrearTablaNecesidadRechazada_20191015_114240) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.necesidad_rechazada (id serial NOT NULL,necesidad_id integer NOT NULL,justificacion varchar NOT NULL,fecha_rechazo timestamp NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_necesidad_rechazada PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaNecesidadRechazada_20191015_114240) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.necesidad_rechazada")

}
