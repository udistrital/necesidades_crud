package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaProductoCatalogoNecesidad_20191015_113553 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaProductoCatalogoNecesidad_20191015_113553{}
	m.Created = "20191015_113553"

	migration.Register("CrearTablaProductoCatalogoNecesidad_20191015_113553", m)
}

// Run the migrations
func (m *CrearTablaProductoCatalogoNecesidad_20191015_113553) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.producto_catalogo_necesidad (id serial NOT NULL,catalogo_id integer NOT NULL,necesidad_id integer NOT NULL,unidad_id integer NOT NULL,iva_id integer,cantidad integer NOT NULL,valor numeric(20,7) NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_producto_catalogo_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaProductoCatalogoNecesidad_20191015_113553) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.producto_catalogo_necesidad;")

}
