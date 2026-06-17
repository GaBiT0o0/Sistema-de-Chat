package main

import (
	"fmt"
	"log"
	"net/http"

	"SistemadeChat/gochatservices/routes"
)

func main() {

	// registrar rutas
	routes.SetupRoutes()

	fmt.Println("Servidor de chat iniciado en http://localhost:8080")

	// iniciar servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
