package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRubroNecesidad_20191015_114039 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRubroNecesidad_20191015_114039{}
	m.Created = "20191015_114039"

	migration.Register("CrearTablaRubroNecesidad_20191015_114039", m)
}

// Run the migrations
func (m *CrearTablaRubroNecesidad_20191015_114039) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.rubro_necesidad (id serial NOT NULL,rubro_id varchar NOT NULL,necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_rubro_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.rubro_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaRubroNecesidad_20191015_114039) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.rubro_necesidad;")

}
