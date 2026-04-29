package main

import (
	"log"
	"os"

	peerjs "github.com/muka/peerjs-go/server"
)

func main() {
	// 1. Render inyecta el puerto en la variable de entorno PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Puerto por defecto para local
	}

	// 2. Configuración del servidor PeerJS
	opts := peerjs.NewOptions()
	opts.Port = port
	opts.Path = "/myapp" // Puedes cambiar esto por el nombre de tu app
	opts.AllowDiscovery = true

	server := peerjs.NewServer(opts)
	defer server.Stop()

	// 3. Canal para capturar errores de inicio
	log.Printf("Iniciando servidor PeerJS en Go sobre el puerto %s...", port)

	// El servidor de muka maneja internamente el servidor HTTP
	if err := server.Start(); err != nil {
		log.Fatalf("Error crítico en el servidor: %v", err)
	}
}
