package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationNecesidadRechazada(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.NecesidadRechazada{
		Justificacion: "testNecesidadRechazada",
		Activo:        true,
	}
	testUpdateObj := models.NecesidadRechazada{
		Justificacion: "updateNecesidadRechazada",
		Activo:        true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
