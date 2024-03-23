package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/devbackend/goingot/internal/handler"
	"github.com/devbackend/goingot/pkg/env"
	"github.com/devbackend/goingot/pkg/http/sender"
)

func main() {
	env.MustLoad()

	port := env.MustNotEmpty("PORT")

	uptimeHandler := handler.UptimeHandler{
		Sender: sender.Sender{},
		Start:  time.Now(),
	}

	router := mux.NewRouter()
	router.HandleFunc("/", uptimeHandler.Handle)

	log.Println("Start on port", port)

	serv := http.Server{
		Addr:                         net.JoinHostPort("", port),
		Handler:                      router,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            time.Second,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	go func() {
		err := serv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Println("Stopping...")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	err := serv.Shutdown(timeout)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Stopped")
}
