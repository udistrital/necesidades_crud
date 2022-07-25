package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNecesidad_20191015_105521 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNecesidad_20191015_105521{}
	m.Created = "20191015_105521"

	migration.Register("CrearTablaNecesidad_20191015_105521", m)
}

// Run the migrations
func (m *CrearTablaNecesidad_20191015_105521) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE necesidades.necesidad (id serial NOT NULL,consecutivo_solicitud int4 NOT NULL,consecutivo_necesidad int4 NULL,vigencia varchar NOT NULL,objeto varchar NOT NULL,justificacion varchar NULL,estudio_mercado varchar NULL,analisis_riesgo varchar NULL,fecha_solicitud timestamp NOT NULL,valor numeric(20,7) NOT NULL,area_funcional int4 NOT NULL,tipo_duracion_necesidad_id int4 NOT NULL,dias_duracion int4 NULL,modalidad_seleccion_id int4 NULL,tipo_contrato_id int4 NULL,plan_anual_adquisiciones_id int4 NOT NULL,tipo_contrato_necesidad_id int4 NOT NULL,tipo_financiacion_necesidad_id int4 NULL,tipo_necesidad_id int4 NOT NULL,justificacion_rechazo int4 NULL,dependencia_necesidad_id int4 NOT NULL,estado_necesidad_id int4 NOT NULL,activo bool NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_necesidad PRIMARY KEY (id));")
}

// Reverse the migrations
func (m *CrearTablaNecesidad_20191015_105521) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.necesidad")
}
