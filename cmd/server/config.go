package main

import (
	"errors"
	"fmt"
	"os"
)

// AppConfig centraliza as configurações de ambiente da aplicação
type AppConfig struct {
	Scope string // "local" ou "prod"
	Port  string // porta que o servidor vai usar
}

// NewConfig lê e valida variáveis de ambiente
func NewConfig() (*AppConfig, error) {
	scope := os.Getenv("SCOPE")
	if scope == "" {
		return nil, errors.New("SCOPE não definido. Use 'local' ou 'prod'")
	}
	if scope != "local" && scope != "prod" {
		return nil, fmt.Errorf("SCOPE inválido: %s. Use 'local' ou 'prod'", scope)
	}

	port := os.Getenv("PORT")
	if port == "" {
		if scope == "prod" {
			port = "80"
		} else {
			port = "8080"
		}
	}

	return &AppConfig{
		Scope: scope,
		Port:  port,
	}, nil
}
