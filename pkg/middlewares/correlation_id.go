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
		logger := h.logger
		if correlationID == "" {
			correlationID = uuid.NewString()
			logger.Info().Msgf("No correlationID found in the request. Assigning one now: %s", correlationID)
			r.Header.Set("X-CORRELATION-ID", correlationID)
		}
		ctx := context.WithValue(r.Context(), context_keys.CorrelationID, correlationID)
		r = r.WithContext(ctx)

		logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			f := map[string]string{"correlation_id": correlationID}
			return c.Fields(f)
		})
		r = r.WithContext(logger.WithContext(r.Context()))
		next.ServeHTTP(w, r)
	})
}
