package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devbackend/goingot/pkg/http/sender"
)

type UptimeHandler struct {
	Start time.Time
	sender.Sender
}

type jsonResponse struct {
	Response string `json:"response"`
}

func (h *UptimeHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	h.SendOK(w, jsonResponse{fmt.Sprintf("Uptime: %s", time.Since(h.Start).String())})
}
