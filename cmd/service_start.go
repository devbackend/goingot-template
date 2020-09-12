package cmd

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
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/devbackend/goingot/internal/handler"
)

// WithServiceStart return instance of "service start" command
func WithServiceStart() Command {
	return func(serviceCmd *cobra.Command) {
		serviceCmd.AddCommand(
			&cobra.Command{
				Use:   "start",
				Short: "Run service",
				Run: func(cmd *cobra.Command, args []string) {
					_ = godotenv.Load()

					port := os.Getenv("PORT")
					if port == "" {
						log.Fatal("Empty port")
					}

					uptimeHandler := handler.UptimeHandler{Start: time.Now()}

					router := mux.NewRouter()
					router.HandleFunc("/", uptimeHandler.Handle)

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
			},
		)
	}
}
