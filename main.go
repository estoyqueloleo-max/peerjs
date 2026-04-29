package main

import (
	"log"
	"os"
	"strconv"

	peerjs "github.com/muka/peerjs-go/server"
)

func main() {
	// 1. Render inyecta el puerto en la variable de entorno PORT
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "9000" // Puerto por defecto para local
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Puerto inválido: %v", err)
	}

	// 2. Configuración del servidor PeerJS
	opts := peerjs.NewOptions()
	opts.Port = port
	opts.Host = "0.0.0.0"
	opts.Path = "/myapp" // Puedes cambiar esto por el nombre de tu app
	opts.AllowDiscovery = true

	server := peerjs.New(opts)
	defer server.Stop()

	// 3. Canal para capturar errores de inicio
	log.Printf("Iniciando servidor PeerJS en Go sobre %s:%d%s...", opts.Host, port, opts.Path)

	// El servidor de muka maneja internamente el servidor HTTP
	if err := server.Start(); err != nil {
		log.Fatalf("Error crítico en el servidor: %v", err)
	}
}
