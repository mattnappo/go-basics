package networking

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// makeMuxRouter - Initialize a mux server
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

// StartServer - Start a web server
func StartServer() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("== SERVER == listening on ", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
