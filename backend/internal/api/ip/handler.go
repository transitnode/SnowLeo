package ipapi

import (
	"encoding/json"
	"net/http"

	"github.com/transitnode/SnowLeo/internal/services/ip"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}

	result, err := ip.Lookup(query)

	if err != nil {
		http.Error(w, "lookup failed: no such host", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
