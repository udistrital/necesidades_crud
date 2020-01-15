package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationMarcoLegal(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.MarcoLegal{
		NombreDocumento: "testMarcoLegal",
		Enlace:          "testMarcoLegal",
		Activo:          true,
	}
	testUpdateObj := models.MarcoLegal{
		NombreDocumento: "updateMarcoLegal",
		Enlace:          "testMarcoLegal",
		Activo:          true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
