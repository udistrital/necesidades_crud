package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationFuenteRubroNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.FuenteRubroNecesidad{
		FuenteId:     "testFuenteRubroNecesidad",
		MontoParcial: 1,
		Activo:       true,
	}
	testUpdateObj := models.FuenteRubroNecesidad{
		FuenteId:     "updateFuenteRubroNecesidad",
		MontoParcial: 1,
		Activo:       true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
