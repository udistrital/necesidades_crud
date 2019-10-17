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
	m.SQL("CREATE TABLE necesidades.necesidad (id serial NOT NULL,consecutivo_solicitud integer NOT NULL,consecutivo_necesidad integer,vigencia integer NOT NULL,objeto varchar NOT NULL,fecha_solicitud timestamp NOT NULL,valor numeric(20,7) NOT NULL,area_funcional integer NOT NULL,tipo_duracion_necesidad_id integer NOT NULL,modalidad_seleccion_id integer,tipo_contrato_id integer,plan_anual_adquisiciones_id integer NOT NULL,estudio_mercado varchar,tipo_contrato_necesidad_id integer NOT NULL,tipo_financiacion_necesidad_id smallint,analisis_riesgo varchar,tipo_necesidad_id integer NOT NULL,justificacion_rechazo integer,dependencia_necesidad_id integer NOT NULL,estado_necesidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_necesidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE necesidades.necesidad OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaNecesidad_20191015_105521) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS necesidades.necesidad")
}
