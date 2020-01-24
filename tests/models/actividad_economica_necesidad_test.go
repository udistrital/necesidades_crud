package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationActividadEconomicaNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ActividadEconomicaNecesidad{
		ActividadEconomicaId: 1,
		Activo:               true,
	}
	testUpdateObj := models.ActividadEconomicaNecesidad{
		ActividadEconomicaId: 2,
		Activo:               false,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
