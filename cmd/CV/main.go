package main

import (
    "net/http"
)

func main() {
    // Servir los archivos estáticos desde la carpeta "build"
    fs := http.FileServer(http.Dir("./frontend/build"))
    http.Handle("/", fs)

    // Iniciar el servidor en el puerto 8080
    println("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}