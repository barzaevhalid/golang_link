package stat

import (
	"net/http"
	"rest_api/configs"
	"rest_api/pkg/res"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	*configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	*configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.HandleFunc("GET /stat", handler.GetStat())
}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))

		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))

		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")

		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := h.StatRepository.GetStats(by, from, to)
		res.Json(w, stats, http.StatusOK)
	}
}
