package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationProductoRubroNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ProductoRubroNecesidad{
		ProductoId:   "testProductoRubroNecesidad",
		MontoParcial: 1,
		Activo:       true,
	}
	testUpdateObj := models.ProductoRubroNecesidad{
		ProductoId:   "updateProductoRubroNecesidad",
		MontoParcial: 1,
		Activo:       true,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
