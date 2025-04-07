package health

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK")) // verifica se a API está viva, Útil pra monitoramento (ex: Load Balancer, Kubernetes, etc.)
}
