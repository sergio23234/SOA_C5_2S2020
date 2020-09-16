package main

import (
	"testing"
)

func TestEnviar_signal(t *testing.T) {
	var ped pedido
	ped.Desc = "hamburguesa"
	ped.Id = "35"
	var resultado string = enviar_signal(ped)
	var deseado string = "enviado al restaurante"
	if deseado == resultado {
		t.Log("Test finalizado correctamente")
	} else {
		t.Error("Error en el test")
		t.Fail()
	}
}
