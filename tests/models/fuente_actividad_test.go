package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationFuenteActividad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.FuenteActividad{
		FuenteId:     "testFuenteActividad",
		MontoParcial: 1,
		Activo:       true,
	}
	testUpdateObj := models.FuenteActividad{
		FuenteId:     "updateFuenteActividad",
		MontoParcial: 1,
		Activo:       true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
