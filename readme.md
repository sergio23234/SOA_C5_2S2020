# CORTO 5
## Nombre: Sergio Fernando De los Rios Roque.
## Carnet: 201213282. 
### Lenguaje Utilizado: Golag. 
### version: 0.0.2
### Descripcion:
Se tiene en cada una de los archivos un servidor:
- en localhost : 9093 estar corriendo el orquestador de servicios, sera el encargado de redirigir los peididos de los demas servicios, ahora los demas servicios en vez de comunicarse entre sí, se comunican unicamente al orquestador, y el orquestador es el unico que conoce las direcciones de servicio de todos. 
- en localhost : 9092 esta corriendo el servidor del cliente para recibir los mensajes y enviar los pedidos. 
- en localhost : 9091 esta corriendo el servidor del restaurante para recibir el pedido del cliente y enviarlo al restaurante
- en localhost : 9090 esta corriendo el servidor del repartidor para recibir los pedidos del restaurante/cliente y enviar los mensajes al ciente
### link del video:
https://youtu.be/XcYs6tjvMzk
### generar ejecutable:
En el cmd o terminal dirigirse a la carpeta donde esta el archivo y ejecutar el siguiente comando: 
`go build nombre_del_archivo.go`
### ejecutar:
En el cmd o terminal dirigirse a la carpeta donde teta el archivo y ejecutar el siguiente comando: 
-si desea probar :`go run nombre_del_archivo.go`
-si desea ejecutar el . exe: `nombre_del_archivo.exe` 

**dentro del archivo ya se encuentra generados los .exe para poderse ejecutar en windows**

