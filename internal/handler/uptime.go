package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UptimeHandler struct {
	Start time.Time
}

type jsonResponse struct {
	Response string `json:"response"`
}

func (h *UptimeHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	uptime := time.Since(h.Start)

	resp, _ := json.Marshal(jsonResponse{"Uptime: " + uptime.String()})
	_, err := w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}
