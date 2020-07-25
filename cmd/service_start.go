package cmd

import (
	"context"
	"encoding/json"
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

type JSONResponse struct {
	Response string `json:"response"`
}

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run service",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("Empty port")
		}

		router := mux.NewRouter()
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)

			resp, _ := json.Marshal(JSONResponse{"Current time: " + time.Now().Format("2006-01-02T15:04:05-0700")})
			_, err := w.Write(resp)
			if err != nil {
				log.Println(err)
			}
		})

		log.Println("Start on port", port)

		serv := http.Server{
			Addr:    net.JoinHostPort("", port),
			Handler: router,
		}

		go func() {
			err := serv.ListenAndServe()
			if err != nil {
				log.Fatal()
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
