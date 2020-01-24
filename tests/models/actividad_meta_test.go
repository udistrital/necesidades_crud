package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationActividadMeta(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ActividadMeta{
		ActividadId: "testActividadMeta",
		Activo:      true,
	}
	testUpdateObj := models.ActividadMeta{
		ActividadId: "updateActividadMeta",
		Activo:      true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
