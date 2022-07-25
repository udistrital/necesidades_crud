package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDetalleServicioNecesidad_20191015_115948 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDetalleServicioNecesidad_20191015_115948{}
	m.Created = "20191015_115948"

	migration.Register("CrearTablaDetalleServicioNecesidad_20191015_115948", m)
}

// Run the migrations
func (m *CrearTablaDetalleServicioNecesidad_20191015_115948) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.detalle_servicio_necesidad (id serial NOT NULL,tipo_servicio_id integer NOT NULL,valor numeric(20,7) NOT NULL,descripcion varchar NOT NULL,necesidad_id integer NOT NULL,iva_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_detalle_prestacion_servicio_cp PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaDetalleServicioNecesidad_20191015_115948) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.detalle_servicio_necesidad;")

}
