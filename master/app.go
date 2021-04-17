package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fredrikluo/distworker/master/service"
	"github.com/gorilla/mux"
)

type MasterService struct {
	router      *mux.Router
	server      *http.Server
	coordinator *service.Coordinator
}

func (ms *MasterService) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
}

func (ms *MasterService) handleHeartBeat(w http.ResponseWriter, r *http.Request) {
}

func (ms *MasterService) handleGetTask(w http.ResponseWriter, r *http.Request) {
}

func (ms *MasterService) handleReport(w http.ResponseWriter, r *http.Request) {
}

func (ms *MasterService) Run() {
	ms.coordinator = service.NewCoordinator()

	ms.router.HandleFunc("/health", ms.handleHealthCheck).
		Methods("GET")

	ms.router.HandleFunc("/dw/heartbeat", ms.handleHeartBeat).
		Methods("POST")

	ms.router.HandleFunc("/dw/task", ms.handleGetTask).
		Methods("GET")

	ms.router.HandleFunc("/dw/report/{taskid}", ms.handleReport).
		Methods("POST")

	ms.server = &http.Server{
		Handler:      ms.router,
		Addr:         "127.0.0.1:8010",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(ms.server.ListenAndServe())
}

func main() {
	service := &MasterService{}
	service.Run()
}
