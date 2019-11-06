package crudoperationstestcases

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationTipoCOntratoNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.TipoContratoNecesidad{
		Nombre:            "test",
		CodigoAbreviacion: "tst",
		Descripcion:       "test",
	}
	testUpdateObj := models.TipoContratoNecesidad{
		Nombre:            "testUpd",
		CodigoAbreviacion: "tst",
		Descripcion:       "test",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
