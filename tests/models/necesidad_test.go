package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOperationNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.Necesidad{
		ConsecutivoNecesidad: 1,
		Justificacion:        "testNecesidad",
		EstudioMercado:       "testNecesidad",
		AnalisisRiesgo:       "testNecesidad",
		DiasDuracion:         1,
		TipoContratoId:       1,
		JustificacionRechazo: 1,
	}
	testUpdateObj := models.Necesidad{
		ConsecutivoNecesidad: 1,
		Justificacion:        "updateNecesidad",
		EstudioMercado:       "updateNecesidad",
		AnalisisRiesgo:       "updateNecesidad",
		DiasDuracion:         1,
		TipoContratoId:       1,
		JustificacionRechazo: 1,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
