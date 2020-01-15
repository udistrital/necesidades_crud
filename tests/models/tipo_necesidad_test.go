package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationTipoNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.TipoNecesidad{
		Nombre:            "testTipoNecesidad",
		CodigoAbreviacion: "testTipoNecesidad",
		Descripcion:       "testTipoNecesidad",
	}
	testUpdateObj := models.TipoNecesidad{
		Nombre:            "updateTipoNecesidad",
		CodigoAbreviacion: "testTipoNecesidad",
		Descripcion:       "testTipoNecesidad",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
