package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"
)

// StartServer é chamado pelo Fx quando a aplicação sobe
// Ele usa as configurações de AppConfig para decidir porta e escopo
func StartServer(lc fx.Lifecycle, cfg *AppConfig, handler http.Handler) {
	server := &http.Server{
		Addr:         ":" + cfg.Port, // porta definida com base no ambiente
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("🚀 [%s] Servidor rodando na porta %s\n", cfg.Scope, cfg.Port)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Erro no servidor: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("🛑 Encerrando servidor...")
			return server.Shutdown(ctx)
		},
	})
}
