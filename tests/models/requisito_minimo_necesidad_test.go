package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationRequisitoMinimoNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.RequisitoMinimoNecesidad{
		Descripcion: "testRequisitoMinimoNecesidad",
		Activo:      true,
	}
	testUpdateObj := models.RequisitoMinimoNecesidad{
		Descripcion: "updateRequisitoMinimoNecesidad",
		Activo:      false,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
