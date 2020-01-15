package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationDetalleServicioNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.DetalleServicioNecesidad{
		Valor:          1,
		TipoServicioId: 1,
		Descripcion:    "testTipoNecesidad",
	}
	testUpdateObj := models.DetalleServicioNecesidad{
		Valor:          1,
		TipoServicioId: 2,
		Descripcion:    "updateTipoNecesidad",
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
