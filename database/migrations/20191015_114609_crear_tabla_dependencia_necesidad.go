package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDependenciaNecesidad_20191015_114609 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDependenciaNecesidad_20191015_114609{}
	m.Created = "20191015_114609"

	migration.Register("CrearTablaDependenciaNecesidad_20191015_114609", m)
}

// Run the migrations
func (m *CrearTablaDependenciaNecesidad_20191015_114609) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.dependencia_necesidad (id serial NOT NULL,jefe_dep_solicitante_id integer NOT NULL,jefe_dep_destino_id integer NOT NULL,supervisor_id integer,interventor_id integer,ordenador_gasto_id integer,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_responsables_necesidad PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaDependenciaNecesidad_20191015_114609) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.dependencia_necesidad;")

}
