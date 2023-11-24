package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"room_read/internal/domain/ports"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"

	"github.com/google/uuid"
)

type RestController interface {
	RequestHandler(w http.ResponseWriter, r *http.Request)
}

type restController struct {
	configuration configuration.Configuration
	restPort      ports.RestPort
}

func NewRestController(configuration configuration.Configuration, restPort ports.RestPort) RestController {
	return &restController{
		configuration: configuration,
		restPort:      restPort,
	}
}

func (r *restController) RequestHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(context.Background(), "traceId", uuid.New().String())

	logging.Debug(ctx, "incoming request", slog.String("url", req.URL.Path))

	messages, err := r.restPort.GetMessages(req.Context())

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(messages)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
