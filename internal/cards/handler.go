package cards

import (
	"encoding/json" // para serializar structs em JSON
	"net/http"      // para lidar com requisições/respostas HTTP

	"github.com/go-chi/chi/v5" // para extrair parâmetros da URL (ex: {set})
)

type Handler struct { // Esse Handler é o "ponteiro" entre HTTP e a regra de negócio. (entender melhor aqui)
	Service Service
}

func NewHandler(s Service) *Handler { // Fx vai usar isso pra criar o handler com o Service já injetado.
	return &Handler{Service: s}
}

func (h *Handler) GetBySet(w http.ResponseWriter, r *http.Request) {
	set := chi.URLParam(r, "set")    // pega o {set} da URL
	cards := h.Service.GetBySet(set) // chama a lógica do serviço e busca os cards filtrados

	// Define o tipo de retorno como JSON
	// Converte os cards encontrados em JSON e envia como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}
