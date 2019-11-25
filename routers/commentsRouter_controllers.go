package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEconomicaNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadEspecificaNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ActividadMetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DependenciaNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetallePrestacionServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:DetalleServicioNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:EstadoNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteActividadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:FuenteRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MarcoLegalNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:MetaRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:NecesidadRechazadaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoCatalogoNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:ProductoRubroNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RequisitoMinimoNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:RubroNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoContratoNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoDuracionNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoFinanciacionNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/necesidades_crud/controllers:TipoNecesidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
