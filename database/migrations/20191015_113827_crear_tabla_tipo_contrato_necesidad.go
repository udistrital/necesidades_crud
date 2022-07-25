package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoContratoNecesidad_20191015_113827 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoContratoNecesidad_20191015_113827{}
	m.Created = "20191015_113827"

	migration.Register("CrearTablaTipoContratoNecesidad_20191015_113827", m)
}

// Run the migrations
func (m *CrearTablaTipoContratoNecesidad_20191015_113827) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.tipo_contrato_necesidad (id serial NOT NULL,nombre varchar,descripcion varchar,codigo_abreviacion varchar,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_tipo_contrato_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaTipoContratoNecesidad_20191015_113827) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.tipo_contrato_necesidad;")

}
