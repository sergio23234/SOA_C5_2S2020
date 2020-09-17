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

func TestRecibir_signal1(t *testing.T) {
	var resultado string = recibir_signal1()
	var deseado string = "ya en el repartidor"
	var deseado1 string = "aun en restaurante"
	if deseado == resultado {
		t.Log("Test finalizado correctamente")
	} else if deseado1 == resultado {
		t.Log("Test finalizado correctamente")
	} else {
		t.Error("Error en el test")
		t.Fail()
	}
}

func TestRecibir_signal2(t *testing.T) {
	var resultado string = recibir_signal2()
	var deseado string = "entregado o aun no recibido"
	var deseado1 string = "en camino"
	if deseado == resultado {
		t.Log("Test finalizado correctamente")
	} else if deseado1 == resultado {
		t.Log("Test finalizado correctamente")
	} else {
		t.Error("Error en el test")
		t.Fail()
	}
}
