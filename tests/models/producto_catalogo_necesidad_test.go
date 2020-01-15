package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationProductoCatalogoNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ProductoCatalogoNecesidad{
		UnidadId: 1,
		IvaId:    1,
		Cantidad: 1,
	}
	testUpdateObj := models.ProductoCatalogoNecesidad{
		UnidadId: 1,
		IvaId:    2,
		Cantidad: 1,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
