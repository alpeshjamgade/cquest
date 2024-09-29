package middlewares

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/logger"
	"cquest/internal/utils"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
		Logger := logger.CreateLoggerWithCtx(ctx)
		Logger.Infof("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
