package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationRequisitoMinimo(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.RequisitoMinimo{
		Descripcion: "testRequisitoMinimo",
		Activo:      true,
	}
	testUpdateObj := models.RequisitoMinimo{
		Descripcion: "updateRequisitoMinimo",
		Activo:      true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
