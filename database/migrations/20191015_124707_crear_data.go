package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearData_20191015_124707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearData_20191015_124707{}
	m.Created = "20191015_124707"

	migration.Register("CrearData_20191015_124707", m)
}

// Run the migrations
func (m *CrearData_20191015_124707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO necesidades.tipo_financiacion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Inversión', 'Inversión', 'I', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_financiacion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Funcionamiento', 'Funcionamiento', 'F', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Licitación pública', 'Licitación pública', 'LP', 1, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Selección abreviada', 'Selección abreviada', 'SA', 2, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Concurso de méritos', 'Concurso de méritos', 'CM', 3, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Mínima cuantía', 'Mínima cuantía', 'MC', 4, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Contratación directa', 'Contratación directa', 'CD', 5, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Contratación directa por urgencia manifiesta (*)', 'Contratación directa por urgencia manifiesta (*)', 'UM', 6, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Convocatoria pública', 'Convocatoria pública', 'CP', 7, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.modalidad_seleccion (nombre, descripcion, codigo_abreviacion, numero_orden, activo, fecha_creacion, fecha_modificacion) VALUES('Otra', 'Otra', 'O', 8, true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Solicitada', 'Solicitud registrada', 'S', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Aprobada', 'Necesidad aprobada por el ordenador del gasto', 'A', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Rechazada', 'Necesidad rechazada por el ordenador del gasto', 'R', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Anulada', 'Necesidad anulada luego de haber sido aprobada o en estado Nueva si se hizo cambio de vigencia', 'AN', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Modificada', 'Necesidad modificada luego de estar en estado nueva o rechazada', 'M', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Enviada', 'Necesidad enviada al ordenador del gasto para que este la califique, aqui ya no es posible modificarla', 'E', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Cdp Solicitado', 'Solicitud de CDP creada', 'CS', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('Guardada', 'Guardada sin enviar', 'G', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('CDP Expedido', 'CDP expedido y por aprobar', 'CDPE', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.estado_necesidad (nombre, descripcion, codigo_abreviacionn, activo, fecha_creacion, fecha_modificacion) VALUES('CDP Aprobado', 'CDP aprobado', 'CDPA', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_contrato_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Compra', 'Aqui van las compras (Arka)', 'C', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_contrato_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Prestación de servicio', 'Aqui van las contrataciones de Prestación de Servicio', 'CPS', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_contrato_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('No aplica', 'Parámetro para las necesidades que no sean de contratación', 'NA', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_contrato_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Compra y servicio', 'Aqui van las compras y servicio', 'CS', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_contrato_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Servicio', 'Aqui van los servicios generales', 'S', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Contratación', 'Necesidades de compras o adquisición de servicios que serán estipuladas en un contrato', 'C', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Servicios públicos', 'Servicios públicos', 'SP', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_duracion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Duración', 'Duración', 'D', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_duracion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Único pago', 'Único pago', 'UP', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_duracion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Único pago / Hasta agotar presupuesto', 'Único pago / Hasta agotar presupuesto', 'UPAP', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.tipo_duracion_necesidad (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES('Hasta agotar presupuesto', 'Hasta agotar presupuesto', 'AP', true, current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.marco_legal (nombre_documento, enlace, activo, fecha_creacion, fecha_modificacion) VALUES('Acuerdo 03 del 11 de Marzo de 2015 del CSU', 'http://sgral.udistrital.edu.co/xdata/csu/acu_2015-03.pdf', true,  current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.marco_legal (nombre_documento, enlace, activo, fecha_creacion, fecha_modificacion) VALUES('Circular 025 de 2015 Evaluación de Proveedores', 'https://www.udistrital.edu.co/files/CIRCULAR_025_de_2015_Evaluacion_de_Proveedores.pdf', true,  current_timestamp, current_timestamp);")
	m.SQL("INSERT INTO necesidades.marco_legal (nombre_documento, enlace, activo, fecha_creacion, fecha_modificacion) VALUES('Resolución 132 del 25 de Marzo de 2015', 'http://sgral.udistrital.edu.co/xdata/rec/res_2015-263.pdf', true,  current_timestamp, current_timestamp);")
}

// Reverse the migrations
func (m *CrearData_20191015_124707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
