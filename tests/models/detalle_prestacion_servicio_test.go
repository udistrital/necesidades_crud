package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationDetallePrestacionServicio(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.DetallePrestacionServicio{
		NucleoConocimientoId: 1,
	}
	testUpdateObj := models.DetallePrestacionServicio{
		NucleoConocimientoId: 2,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
