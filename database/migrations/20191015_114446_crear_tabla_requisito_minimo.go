package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRequisitoMinimo_20191015_114446 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRequisitoMinimo_20191015_114446{}
	m.Created = "20191015_114446"

	migration.Register("CrearTablaRequisitoMinimo_20191015_114446", m)
}

// Run the migrations
func (m *CrearTablaRequisitoMinimo_20191015_114446) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.requisito_minimo (id serial NOT NULL,producto_catalogo_necesidad_id integer NOT NULL,descripcion varchar NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_requisito_minimo PRIMARY KEY (id));")
	m.SQL("CREATE TABLE necesidades.requisito_minimo_necesidad (id serial NOT NULL,necesidad_id integer NOT NULL,descripcion varchar NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_requisito_minimo_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.requisito_minimo OWNER TO desarrollooas;")
	m.SQL("ALTER TABLE necesidades.requisito_minimo_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaRequisitoMinimo_20191015_114446) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.requisito_minimo;")
	m.SQL("DROP TABLE IF EXISTS necesidades.requisito_minimo_necesidad;")

}
