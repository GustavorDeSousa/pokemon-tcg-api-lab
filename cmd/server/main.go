package main

import (
	"github.com/GustavorDeSousa/pokemon-tcg-api-lab/internal/cards"
	"github.com/GustavorDeSousa/pokemon-tcg-api-lab/internal/health"
	"github.com/GustavorDeSousa/pokemon-tcg-api-lab/internal/routes"
	"go.uber.org/fx"
)

func main() {
	fx.New( // Cria o container de dependências
		fx.Provide( // constroi as dependências
			NewConfig, // ✅ registra AppConfig

			// Bloco de Cards
			cards.NewService,
			cards.NewHandler,

			// Bloco da Saude da API
			health.NewHandler,

			routes.NewRouter, // injeta roteador com todas as rotas
		),

		// - fx.Invoke(StartServer) inicia o servidor HTTP
		fx.Invoke(
			StartServer,
		),
	).Run()
}
