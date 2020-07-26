package cmd

import (
	"context"
	"github.com/devbackend/goingot/internal/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run service",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("Empty port")
		}

		handler := handler.UptimeHandler{Start: time.Now()}

		router := mux.NewRouter()
		router.HandleFunc("/", handler.Handle)

		log.Println("Start on port", port)

		serv := http.Server{
			Addr:    net.JoinHostPort("", port),
			Handler: router,
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
	},
}

func init() {
	serviceCmd.AddCommand(serviceStartCmd)
}
