package main

import (
    "net/http"
    "log"
)

func main() {
    // Servir los archivos estáticos desde la carpeta "build"
    fs := http.FileServer(http.Dir("./frontend/build"))
    http.Handle("/", fs)

    // Iniciar el servidor en el puerto 8080
    println("Servidor corriendo en http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v\n", err)
    }
}