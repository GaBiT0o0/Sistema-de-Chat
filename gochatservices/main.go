package main

import (
	"SistemadeChat/gochatservices/routes"
	"fmt"
	"net/http"
)

func main() {

	routes.SetupRoutes()

	fmt.Println("Servidor iniciado en puerto 8080")

	http.ListenAndServe(":8080", nil)
}
