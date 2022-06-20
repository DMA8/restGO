package entrypoint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testTask/internal/domain"
)

type prop string

func (h *Handler) validateQuery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var InpMsg domain.Props
		var ctx context.Context
		err := json.NewDecoder(r.Body).Decode(&InpMsg)
		defer r.Body.Close()
		if err != nil {
			writeAnswer(w, http.StatusBadRequest, fmt.Sprintf("decoding query error: %s", err.Error()))
			return
		}
		err = validateInput(&InpMsg)
		if err != nil {
			writeAnswer(w, http.StatusBadRequest, fmt.Sprintf("validation query error: %s", err.Error()))
			return
		}
		initHeaders(w)
		ctx = context.WithValue(r.Context(), prop("prop"), &InpMsg)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
