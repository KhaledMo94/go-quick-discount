package handler

import (
	"context"
	"encoding/json"
	"net/http"
)


type Pinger interface {
	Ping(ctx context.Context) error
}

type Handler struct{
	db    Pinger
	redis Pinger
	ms Pinger
}

func New(dbConn Pinger, redisConn Pinger , msConn Pinger) *Handler{
	return  &Handler{
		db:    dbConn,
		redis: redisConn,
		ms : msConn,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health",h.health)
	return mux
}

func (h *Handler) health(w http.ResponseWriter , r *http.Request) {
	ctx := r.Context()
	dbError := h.db.Ping(ctx)
	redisError := h.redis.Ping(ctx)
	msError := h.ms.Ping(ctx)

	status := "ok"
	code := http.StatusOK

	if dbError != nil || redisError != nil || msError != nil {
		status = "degraded"
		code = http.StatusServiceUnavailable
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"status":status})
}