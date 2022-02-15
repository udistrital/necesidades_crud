package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Necesidad struct {
	Id                          int                        `orm:"column(id);pk;auto"`
	Consecutivo                 int                        `orm:"column(consecutivo)"`
	Vigencia                    string                     `orm:"column(vigencia)"`
	Objeto                      string                     `orm:"column(objeto)"`
	Justificacion               string                     `orm:"column(justificacion);null"`
	EstudioMercado              string                     `orm:"column(estudio_mercado);null"`
	AnalisisRiesgo              string                     `orm:"column(analisis_riesgo);null"`
	FechaSolicitud              time.Time                  `orm:"column(fecha_solicitud);type(timestamp without time zone)"`
	Valor                       float64                    `orm:"column(valor)"`
	AreaFuncional               int                        `orm:"column(area_funcional)"`
	TipoDuracionNecesidadId     *TipoDuracionNecesidad     `orm:"column(tipo_duracion_necesidad_id);rel(fk)"`
	DiasDuracion                int                        `orm:"column(dias_duracion);null"`
	ModalidadSeleccionId        *ModalidadSeleccion        `orm:"column(modalidad_seleccion_id);rel(fk)"`
	TipoContratoId              int                        `orm:"column(tipo_contrato_id);null"`
	PlanAnualAdquisicionesId    int                        `orm:"column(plan_anual_adquisiciones_id)"`
	TipoContratoNecesidadId     *TipoContratoNecesidad     `orm:"column(tipo_contrato_necesidad_id);rel(fk)"`
	TipoFinanciacionNecesidadId *TipoFinanciacionNecesidad `orm:"column(tipo_financiacion_necesidad_id);rel(fk)"`
	TipoNecesidadId             *TipoNecesidad             `orm:"column(tipo_necesidad_id);rel(fk)"`
	JustificacionRechazo        int                        `orm:"column(justificacion_rechazo);null"`
	DependenciaNecesidadId      *DependenciaNecesidad      `orm:"column(dependencia_necesidad_id);rel(fk)"`
	EstadoNecesidadId           *EstadoNecesidad           `orm:"column(estado_necesidad_id);rel(fk)"`
	Activo                      bool                       `orm:"column(activo)"`
	FechaCreacion               time.Time                  `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion           time.Time                  `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *Necesidad) TableName() string {
	return "necesidad"
}

func init() {
	orm.RegisterModel(new(Necesidad))
}

// AddNecesidad insert a new Necesidad into database and returns
// last inserted Id on success.
func AddNecesidad(m *Necesidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNecesidadById retrieves Necesidad by Id. Returns error if
// Id doesn't exist
func GetNecesidadById(id int) (v *Necesidad, err error) {
	o := orm.NewOrm()
	v = &Necesidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNecesidad retrieves all Necesidad matches certain condition. Returns empty list if
// no records exist
func GetAllNecesidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Necesidad)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Necesidad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNecesidad updates Necesidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateNecesidadById(m *Necesidad) (err error) {
	o := orm.NewOrm()
	v := Necesidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNecesidad deletes Necesidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNecesidad(id int) (err error) {
	o := orm.NewOrm()
	v := Necesidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Necesidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
