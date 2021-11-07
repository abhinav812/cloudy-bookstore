package app

import "net/http"

// HandleLive - A http.HandlerFunc that handles liveness checks by
// immediately responding with an HTTP 200 status.
func HandleLive(w http.ResponseWriter, _ *http.Request) {
	writeHealthy(w)
}

func writeHealthy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("healthy"))
}

// HandleReady - A http.HandlerFunc that handles readiness checks by
// responding with an HTTP 200 status if it is healthy, 500 otherwise.
func (s *Server) HandleReady(w http.ResponseWriter, _ *http.Request) {
	db, err := s.db.DB()
	if err != nil {
		writeUnhealthy(s, err, w)
		return
	}
	if err := db.Ping(); err != nil {
		writeUnhealthy(s, err, w)
		return
	}
}

func writeUnhealthy(s *Server, err error, w http.ResponseWriter) {
	s.Logger().Fatal().Err(err).Msg("")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("unhealthy"))
}
