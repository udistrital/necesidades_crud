package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationEstadoNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.EstadoNecesidad{
		Nombre:             "testEstadoNecesidad",
		CodigoAbreviacionn: "testEstadoNecesidad",
		Descripcion:        "testEstadoNecesidad",
	}
	testUpdateObj := models.EstadoNecesidad{
		Nombre:             "updateEstadoNecesidad",
		CodigoAbreviacionn: "testEstadoNecesidad",
		Descripcion:        "testEstadoNecesidad",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
