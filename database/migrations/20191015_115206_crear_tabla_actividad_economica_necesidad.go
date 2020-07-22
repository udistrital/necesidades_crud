package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaActividadEconomicaNecesidad_20191015_115206 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaActividadEconomicaNecesidad_20191015_115206{}
	m.Created = "20191015_115206"

	migration.Register("CrearTablaActividadEconomicaNecesidad_20191015_115206", m)
}

// Run the migrations
func (m *CrearTablaActividadEconomicaNecesidad_20191015_115206) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.actividad_economica_necesidad (id serial NOT NULL,actividad_economica_id integer NOT NULL,necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_actividad_economica_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.actividad_economica_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaActividadEconomicaNecesidad_20191015_115206) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.actividad_economica_necesidad;")

}
