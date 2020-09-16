package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

type pedido struct {
	Desc string `json:"Desc"`
	Id   string `json:"Id"`
}
type resultado struct {
	Resultado string `json:"respuesta"`
}

func main() {
	go menu()
	servidor()
}
func menu() {
	var numero int = 0
	for numero != 4 {
		fmt.Println("1.Enviar pedido")
		fmt.Println("2.Verificar pedido en restaurante")
		fmt.Println("3.Verficar pedidio al repartidor")
		fmt.Println("4.Salir")
		fmt.Scanf("%d\n", &numero)
		if numero == 1 {
			leer_pedidio()
		} else if numero == 2 {
			fmt.Println(recibir_signal1())
		} else if numero == 3 {
			fmt.Println(recibir_signal2())
		} else if numero == 4 {
			os.Exit(1)
		} else {
			fmt.Println(numero)
		}
	}
}
func leer_pedidio() {
	var envio pedido
	fmt.Println("Cual es ID?")
	fmt.Scanf("%s\n", &envio.Id)
	fmt.Println("Descripcion?")
	fmt.Scanf("%s\n", &envio.Desc)
	fmt.Println(enviar_signal(envio))
}

func servidor() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hola\": \"mundo\"}"))
	})
	mux.HandleFunc("/recibido", func(w http.ResponseWriter, r *http.Request) {
		var p resultado
		w.Header().Set("Content-Type", "application/json")
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		}
		err = json.Unmarshal(s, &p)
		if err != nil {
			panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		}
		fmt.Println("recibido:" + p.Resultado)
		w.Write([]byte("{\"respuesta\": \"confirmado\"}"))
	})
	mux.HandleFunc("/entregado", func(w http.ResponseWriter, r *http.Request) {
		var p resultado
		w.Header().Set("Content-Type", "application/json")
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		}
		err = json.Unmarshal(s, &p)
		if err != nil {
			panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		}
		fmt.Println("recibido:" + p.Resultado)
		w.Write([]byte("{\"respuesta\": \"confirmado\"}"))
	})
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":9092", handler)
}
func enviar_signal(buf pedido) string { //enviar pedido al restaurante
	var r resultado
	jjson := `{"Id":"` + buf.Id + `","Desc":"` + buf.Desc + `"}`
	b := strings.NewReader(jjson)
	resp, err := http.Post("http://localhost:9093/idpedido-cli", "application/json", b)
	if err != nil {
		fmt.Println(err)
		return "error" // handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		return "error"
	}
	return r.Resultado
}
func recibir_signal1() string { //restaurante
	var r resultado
	resp, err := http.Get("http://localhost:9093/cli-res")
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
	}
	return r.Resultado
}
func recibir_signal2() string { //repartidor
	var r resultado
	resp, err := http.Get("http://localhost:9093/cli-rep")
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
	}
	return r.Resultado
}
