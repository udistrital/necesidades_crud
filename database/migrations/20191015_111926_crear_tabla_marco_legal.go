package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMarcoLegal_20191015_111926 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMarcoLegal_20191015_111926{}
	m.Created = "20191015_111926"

	migration.Register("CrearTablaMarcoLegal_20191015_111926", m)
}

// Run the migrations
func (m *CrearTablaMarcoLegal_20191015_111926) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.marco_legal (id serial NOT NULL,nombre_documento varchar NOT NULL,enlace varchar NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_marco_legal PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.marco_legal OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaMarcoLegal_20191015_111926) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.marco_legal;")

}
