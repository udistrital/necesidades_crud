package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaProductoRubroNecesidad_20191015_111055 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaProductoRubroNecesidad_20191015_111055{}
	m.Created = "20191015_111055"

	migration.Register("CrearTablaProductoRubroNecesidad_20191015_111055", m)
}

// Run the migrations
func (m *CrearTablaProductoRubroNecesidad_20191015_111055) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.producto_rubro_necesidad (id serial NOT NULL,rubro_necesidad_id integer NOT NULL,producto_id varchar NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_producto_rubro_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.producto_rubro_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaProductoRubroNecesidad_20191015_111055) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.producto_rubro_necesidad;")

}
