package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearLlavesForaneas_20191015_120209 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearLlavesForaneas_20191015_120209{}
	m.Created = "20191015_120209"

	migration.Register("CrearLlavesForaneas_20191015_120209", m)
}

// Run the migrations
func (m *CrearLlavesForaneas_20191015_120209) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_tipo_duracion_necesidad FOREIGN KEY (tipo_duracion_necesidad_id) REFERENCES necesidades.tipo_duracion_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_tipo_necesidad FOREIGN KEY (tipo_necesidad_id) REFERENCES necesidades.tipo_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_tipo_contrato_necesidad FOREIGN KEY (tipo_contrato_necesidad_id) REFERENCES necesidades.tipo_contrato_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_estado_necesidad FOREIGN KEY (estado_necesidad_id) REFERENCES necesidades.estado_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_estado_modalidad_seleccion FOREIGN KEY (modalidad_seleccion_id) REFERENCES necesidades.modalidad_seleccion (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_tipo_financiacion_necesidad FOREIGN KEY (tipo_financiacion_necesidad_id) REFERENCES necesidades.tipo_financiacion_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad ADD CONSTRAINT fk_necesidad_dependencia_necesidad FOREIGN KEY (dependencia_necesidad_id) REFERENCES necesidades.dependencia_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.marco_legal_necesidad ADD CONSTRAINT fk_marco_legal_necesidad_marco_legal FOREIGN KEY (marco_legal_id) REFERENCES necesidades.marco_legal (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.marco_legal_necesidad ADD CONSTRAINT fk_marco_legal_necesidad_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.producto_rubro_necesidad ADD CONSTRAINT fk_producto_rubro_necesidad_rubro_necesidad FOREIGN KEY (rubro_necesidad_id) REFERENCES necesidades.rubro_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.fuente_rubro_necesidad ADD CONSTRAINT fk_fuente_rubro_necesidad_rubro_necesidad FOREIGN KEY (rubro_necesidad_id) REFERENCES necesidades.rubro_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.actividad_meta ADD CONSTRAINT fk_actividad_meta_rubro_necesidad FOREIGN KEY (meta_rubro_necesidad_id) REFERENCES necesidades.meta_rubro_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.producto_catalogo_necesidad ADD CONSTRAINT fk_producto_catalogo_necesidad_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.rubro_necesidad ADD CONSTRAINT fk_rubro_necesidad_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.necesidad_rechazada ADD CONSTRAINT fk_necesidad_rechazada_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.requisito_minimo ADD CONSTRAINT fk_requisito_minimo_producto_catalogo_necesidad FOREIGN KEY (producto_catalogo_necesidad_id) REFERENCES necesidades.producto_catalogo_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.detalle_prestacion_servicio ADD CONSTRAINT fk_detalle_prestacion_servicio_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.actividad_especifica_necesidad ADD CONSTRAINT fk_actividad_especifica_necesidad_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.actividad_economica_necesidad ADD CONSTRAINT fk_actividad_economica_necesidad_necesidad FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.meta_rubro_necesidad ADD CONSTRAINT fk_meta_rubro_necesidad_rubro_necesidad FOREIGN KEY (rubro_necesidad_id) REFERENCES necesidades.rubro_necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.detalle_servicio_necesidad ADD CONSTRAINT fk_detalle_prestacion_servicio_necesidad_cp FOREIGN KEY (necesidad_id) REFERENCES necesidades.necesidad (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE necesidades.fuente_actividad ADD CONSTRAINT fk_fuente_actividad_meta FOREIGN KEY (actividad_meta_necesidad_id) REFERENCES necesidades.actividad_meta(id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")

}

// Reverse the migrations
func (m *CrearLlavesForaneas_20191015_120209) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
