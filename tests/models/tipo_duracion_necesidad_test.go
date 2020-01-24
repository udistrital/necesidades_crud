package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationTipoDuracionNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.TipoDuracionNecesidad{
		Nombre:            "testTipoDuracionNecesidad",
		CodigoAbreviacion: "testTipoDuracionNecesidad",
		Descripcion:       "testTipoDuracionNecesidad",
	}
	testUpdateObj := models.TipoDuracionNecesidad{
		Nombre:            "udpateTipoDuracionNecesidad",
		CodigoAbreviacion: "testTipoDuracionNecesidad",
		Descripcion:       "testTipoDuracionNecesidad",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
