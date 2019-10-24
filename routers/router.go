// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/necesidades_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_necesidad",
			beego.NSInclude(
				&controllers.EstadoNecesidadController{},
			),
		),

		beego.NSNamespace("/producto_rubro_necesidad",
			beego.NSInclude(
				&controllers.ProductoRubroNecesidadController{},
			),
		),

		beego.NSNamespace("/marco_legal",
			beego.NSInclude(
				&controllers.MarcoLegalController{},
			),
		),

		beego.NSNamespace("/fuente_rubro_necesidad",
			beego.NSInclude(
				&controllers.FuenteRubroNecesidadController{},
			),
		),

		beego.NSNamespace("/tipo_financiacion_necesidad",
			beego.NSInclude(
				&controllers.TipoFinanciacionNecesidadController{},
			),
		),

		beego.NSNamespace("/modalidad_seleccion",
			beego.NSInclude(
				&controllers.ModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/necesidad",
			beego.NSInclude(
				&controllers.NecesidadController{},
			),
		),

		beego.NSNamespace("/marco_legal_necesidad",
			beego.NSInclude(
				&controllers.MarcoLegalNecesidadController{},
			),
		),

		beego.NSNamespace("/tipo_contrato_necesidad",
			beego.NSInclude(
				&controllers.TipoContratoNecesidadController{},
			),
		),

		beego.NSNamespace("/producto_catalogo_necesidad",
			beego.NSInclude(
				&controllers.ProductoCatalogoNecesidadController{},
			),
		),

		beego.NSNamespace("/actividad_meta",
			beego.NSInclude(
				&controllers.ActividadMetaController{},
			),
		),

		beego.NSNamespace("/necesidad_rechazada",
			beego.NSInclude(
				&controllers.NecesidadRechazadaController{},
			),
		),

		beego.NSNamespace("/requisito_minimo",
			beego.NSInclude(
				&controllers.RequisitoMinimoController{},
			),
		),

		beego.NSNamespace("/rubro_necesidad",
			beego.NSInclude(
				&controllers.RubroNecesidadController{},
			),
		),

		beego.NSNamespace("/meta_rubro_necesidad",
			beego.NSInclude(
				&controllers.MetaRubroNecesidadController{},
			),
		),

		beego.NSNamespace("/actividad_especifica_necesidad",
			beego.NSInclude(
				&controllers.ActividadEspecificaNecesidadController{},
			),
		),

		beego.NSNamespace("/actividad_economica_necesidad",
			beego.NSInclude(
				&controllers.ActividadEconomicaNecesidadController{},
			),
		),

		beego.NSNamespace("/detalle_servicio_necesidad",
			beego.NSInclude(
				&controllers.DetalleServicioNecesidadController{},
			),
		),

		beego.NSNamespace("/dependencia_necesidad",
			beego.NSInclude(
				&controllers.DependenciaNecesidadController{},
			),
		),

		beego.NSNamespace("/tipo_necesidad",
			beego.NSInclude(
				&controllers.TipoNecesidadController{},
			),
		),

		beego.NSNamespace("/detalle_prestacion_servicio",
			beego.NSInclude(
				&controllers.DetallePrestacionServicioController{},
			),
		),

		beego.NSNamespace("/tipo_duracion_necesidad",
			beego.NSInclude(
				&controllers.TipoDuracionNecesidadController{},
			),
		),
		beego.NSNamespace("/fuente_actividad",
		beego.NSInclude(
			&controllers.FuenteActividadController{},
		),
	),
	)
	beego.AddNamespace(ns)
}
