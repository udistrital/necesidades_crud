package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaFuenteRubroNecesidad_20191015_112214 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaFuenteRubroNecesidad_20191015_112214{}
	m.Created = "20191015_112214"

	migration.Register("CrearTablaFuenteRubroNecesidad_20191015_112214", m)
}

// Run the migrations
func (m *CrearTablaFuenteRubroNecesidad_20191015_112214) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.fuente_rubro_necesidad (id serial NOT NULL,fuente_id varchar NOT NULL,rubro_necesidad_id integer NOT NULL,monto_parcial numeric(20,7) NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_fuente_rubro_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaFuenteRubroNecesidad_20191015_112214) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.fuente_rubro_necesidad;")

}
