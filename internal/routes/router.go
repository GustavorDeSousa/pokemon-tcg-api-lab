package routes

import (
	"net/http"

	"github.com/GustavorDeSousa/pokemon-tcg-api-lab/internal/cards"
	"github.com/GustavorDeSousa/pokemon-tcg-api-lab/internal/health"
	"github.com/go-chi/chi/v5"
)

// Usada pelo Fx pra montar e injetar o roteador com todas as rotas
func NewRouter(cardHandler *cards.Handler, healthHandler *health.Handler) http.Handler {
	route := chi.NewRouter()

	route.Get("/health", healthHandler.Check) // Rota apenas para validar a saude
	route.Route("/api", func(r chi.Router) {
		// As chaves `{}` representam variáveis dinâmicas
		r.Get("/{locale}/expansions/{set}", cardHandler.GetBySet)
	})

	return route
}

// -------------------------
// 📌 Sobre o uso do chi.Router
//
// Esta aplicação usa a biblioteca `chi` (https://github.com/go-chi/chi) para roteamento HTTP.
//
// ✅ Por que usamos chi:
// - Leve e idiomático (usa net/http puro)
// - Suporte nativo a parâmetros na URL: /api/{locale}/expansions/{set}
// - Permite agrupar rotas com middlewares
// - Fácil de manter e escalar em projetos modulares
//
// 🚀 Outras bibliotecas populares:
//
// 1. Gin (https://github.com/gin-gonic/gin)
//    - Super rápida e popular
//    - Sintaxe simples tipo Express
//    - Menos idiomática (usa wrappers próprios)
//
// 2. Echo (https://github.com/labstack/echo)
//    - Muito rápida e com recursos embutidos
//    - Ideal para APIs com muita estrutura pronta
//
// 3. Gorilla Mux (https://github.com/gorilla/mux)
//    - Foi padrão por anos
//    - Está em modo manutenção
//    - Ainda usado em muitos sistemas legados
//
//
// -------------------------
