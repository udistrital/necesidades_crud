package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationTipoFinanciacionNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.TipoFinanciacionNecesidad{
		Nombre:            "testTipoFinanciancionNecesidad",
		CodigoAbreviacion: "testTipoFinanciancionNecesidad",
		Descripcion:       "testTipoFinanciancionNecesidad",
	}
	testUpdateObj := models.TipoDuracionNecesidad{
		Nombre:            "updateTipoFinanciancionNecesidad",
		CodigoAbreviacion: "testTipoFinanciancionNecesidad",
		Descripcion:       "testTipoFinanciancionNecesidad",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
