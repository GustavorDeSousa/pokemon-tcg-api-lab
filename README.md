# 🌟 Pokemon TCG API Lab (com Go, Fx e Chi)

Bem-vindo ao **Pokemon TCG API Lab**, um projeto de estudo feito em Go (🐍) com foco em boas práticas, organização em camadas, injeção de dependência com [Uber Fx](https://github.com/uber-go/fx), e roteamento limpo com [Chi](https://github.com/go-chi/chi).

Este projeto simula uma API para consulta de cartas do jogo **Pokémon TCG** com dados mockados.

> ✅ **Objetivo**: estudar estrutura de aplicação Go idiomática, modularização por domínio e boas práticas com Fx + Chi.

---

## 💡 O que você vai encontrar aqui:

- Estrutura com `cmd/`, `internal/`, e injeção de dependências
- Rotas idiomáticas com Chi
- Mock de cartas do Pokémon por expansão ("journey-together", "scarlet-violet")
- Health check para monitoramento
- Uso de `SCOPE=local` e `SCOPE=prod` para ambientes

---

## 🚀 Como rodar localmente

```bash
# clone o repositório
$ git clone https://github.com/SEU_USUARIO/pokemon-tcg-api-lab.git
$ cd pokemon-tcg-api-lab

# defina o ambiente
$ export SCOPE=local

# rode a aplicação
$ go run cmd/server/main.go
```

---

## 🔗 Endpoints disponíveis

### ✉️ Health check
```
GET /health
Retorna: 200 OK com "OK"
```

### 🎯 Buscar cartas por expansão
```
GET /api/{locale}/expansions/{set}
Exemplo: /api/pt-br/expansions/journey-together
```

---

## 🛏️ Estrutura do projeto
```
cmd/
  server/          
    main.go          -> ponto de entrada da aplicação
    config.go
    lifecycle.go

internal/
  cards/           -> domínio das cartas (model, service, handler)
  health/          -> rota de health check
  routes/          -> roteador principal com chi
```

---

## ⚖️ Tecnologias utilizadas
- [Go](https://golang.org)
- [Uber Fx](https://github.com/uber-go/fx)
- [Chi Router](https://github.com/go-chi/chi)

---

## 🤔 Por que esse projeto existe?
Este é um projeto pessoal de aprendizado, explorando:
- Como estruturar uma API idiomática em Go
- Injeção de dependências com Fx
- Uso de interfaces, construtores, ciclo de vida e modularização
- Práticas reais usadas em produção

---

## 📚 Aprendizados até agora
- ✅ Interfaces em Go funcionam por implementação implícita
- ✅ Chi permite rotas limpas e agrupamento por domínio
- ✅ Fx injeta dependências automaticamente por tipo
- ✅ É possível trocar implementações e testar com facilidade

---
