package logger

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &ResponseRecorderWriter{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		log.Println(
			NewRequestLoggerBuilder().
				SetMethod(r.Method).
				SetStatus(writer.Status).
				SetPath(r.URL.Path).
				SetSince(time.Since(start)),
		)
	})
}
