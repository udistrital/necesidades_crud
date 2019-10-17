package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMarcoLegalNecesidad_20191015_110040 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMarcoLegalNecesidad_20191015_110040{}
	m.Created = "20191015_110040"

	migration.Register("CrearTablaMarcoLegalNecesidad_20191015_110040", m)
}

// Run the migrations
func (m *CrearTablaMarcoLegalNecesidad_20191015_110040) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.marco_legal_necesidad (id serial NOT NULL,marco_legal_id integer NOT NULL,necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_marco_legal_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.marco_legal_necesidad OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaMarcoLegalNecesidad_20191015_110040) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.marco_legal_necesidad;")

}
