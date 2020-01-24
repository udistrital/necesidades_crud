package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationDependenciaNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.DependenciaNecesidad{
		SupervisorId:     1,
		InterventorId:    1,
		OrdenadorGastoId: 1,
	}
	testUpdateObj := models.DependenciaNecesidad{
		SupervisorId:     2,
		InterventorId:    2,
		OrdenadorGastoId: 2,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
