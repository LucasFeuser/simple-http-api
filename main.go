package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // Para carregar o.env
)

func main() {
	// Carrega variáveis do arquivo.env no diretório atual
	// Ignora erro se o arquivo não existir (útil se a variável for passada de outra forma)
	_ = godotenv.Load()

	// Lê a variável de ambiente ORIGEM
	origem := os.Getenv("ORIGEM")
	if origem == "" {
		origem = "Indefinida" // Valor padrão caso não seja definida
	}

	// Define o handler para a rota "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá do servidor Go! ORIGEM: %s\n", origem)
	})

	// Define o handler para /healthz (para health checks)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})

	// Define o handler para /readyz (para readiness probes)
    // Neste exemplo simples, é igual ao healthz, mas poderia verificar dependências.
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})


	port := "8080"
	log.Printf("Servidor Go escutando na porta %s com ORIGEM=%s\n", port, origem)

	// Inicia o servidor HTTP na porta 8080 em todas as interfaces (0.0.0.0)
	if err := http.ListenAndServe(":"+port, nil); err!= nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
