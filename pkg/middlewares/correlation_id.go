package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	context_keys "github.com/sunitha/wheels-away-iam/internal/core/contexts"
)

func (h *Middleware) Handle(next Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-CORRELATION-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
			h.logger.Info().Msgf("No correlationID found in the request. Assigning one now: %s", correlationID)
			r.Header.Set("X-CORRELATION-ID", correlationID)
		}
		ctx := context.WithValue(r.Context(), context_keys.CorrelationID, correlationID)
		r = r.WithContext(ctx)

		h.logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(string(context_keys.CorrelationID), correlationID)
		})

		r = r.WithContext(h.logger.WithContext(r.Context()))
		next.ServeHTTP(w, r)
	})
}
