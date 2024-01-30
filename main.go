package main

import (
	"context"
	"github.com/gorilla/mux"
	"gocourse17/core/grpc/adapters"
	"gocourse17/core/rest/handlers"
	"gocourse17/core/service/patients"
	"log"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	pb "gocourse17/core/grpc/generated"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer cancelFunc()

	waitForTheEnd := &sync.WaitGroup{}

	patientsService := patients.NewService()

	// start the http server
	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		server := adapters.NewPatients(patientsService)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterPatientServiceServer(s, server)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// start the scheduler service
	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		handler := handlers.NewPatients(patientsService)

		router := mux.NewRouter()

		router.HandleFunc("/patients", handler.AddPatient).Methods("POST")
		router.HandleFunc("/patients/{id}", handler.GetPatient).Methods("GET")
		router.HandleFunc("/patients/{id}", handler.UpdatePatient).Methods("PUT")

		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	waitForTheEnd.Wait()

	return
}
