package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoFinanciacionNecesidad_20191015_110534 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoFinanciacionNecesidad_20191015_110534{}
	m.Created = "20191015_110534"

	migration.Register("CrearTablaTipoFinanciacionNecesidad_20191015_110534", m)
}

// Run the migrations
func (m *CrearTablaTipoFinanciacionNecesidad_20191015_110534) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.tipo_financiacion_necesidad (id serial NOT NULL,nombre varchar,descripcion varchar,codigo_abreviacion varchar,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_tipo_financiacion_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.tipo_financiacion_necesidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaTipoFinanciacionNecesidad_20191015_110534) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.tipo_financiacion_necesidad;")

}
