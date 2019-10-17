package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDetallePrestacionServicio_20191015_114859 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDetallePrestacionServicio_20191015_114859{}
	m.Created = "20191015_114859"

	migration.Register("CrearTablaDetallePrestacionServicio_20191015_114859", m)
}

// Run the migrations
func (m *CrearTablaDetallePrestacionServicio_20191015_114859) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.detalle_prestacion_servicio (id serial NOT NULL,perfil_id integer NOT NULL,nucleo_conocimiento_id integer,necesidad_id integer NOT NULL,cantidad integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_detalle_prestacion_servicio PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.detalle_prestacion_servicio OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaDetallePrestacionServicio_20191015_114859) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.detalle_prestacion_servicio;")

}
