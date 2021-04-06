package main

import (
    "log"
    "github.com/mberoiza/twittor/handlers"
    "github.com/mberoiza/twittor/bd"
)

func main() {
    if bd.ChequeoConnection()==0 {
        log.Fatal("Sin conexión a la BD")
        return
    } 
    handlers.Manejadores()
}
