package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationModalidadSeleccion(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ModalidadSeleccion{
		Nombre:            "testModalidadSeleccion",
		CodigoAbreviacion: "testModalidadSeleccion",
	}
	testUpdateObj := models.ModalidadSeleccion{
		Nombre:            "updateModalidadSeleccion",
		CodigoAbreviacion: "testModalidadSeleccion",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
