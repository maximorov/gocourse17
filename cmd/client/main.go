package main

import (
	"context"
	"log"
	"time"

	pb "gocourse17/core/grpc/generated"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Не вдалося підключитися: %v", err)
	}
	defer conn.Close()

	c := pb.NewPatientServiceClient(conn)

	// Використання клієнта
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Додавання пацієнта
	r, err := c.AddPatient(ctx, &pb.AddPatientRequest{Patient: &pb.Patient{Id: "1", Name: "Іван Іванов", Age: "30", Diagnosis: "Діагноз"}})
	if err != nil {
		log.Fatalf("Не вдалося створити пацієнта: %v", err)
	}
	log.Printf("Відповідь сервера: %s", r.GetMessage())

	// Отримання пацієнта
	r2, err := c.GetPatient(ctx, &pb.GetPatientRequest{Id: "1"})
	if err != nil {
		log.Fatalf("Не вдалося отримати пацієнта: %v", err)
	}
	log.Printf("Дані пацієнта: %v", r2.GetPatient())

	// Оновлення пацієнта
	r3, err := c.UpdatePatient(ctx, &pb.UpdatePatientRequest{Patient: &pb.Patient{Id: "1", Name: "Іван Петров", Age: "31", Diagnosis: "Оновлений діагноз"}})
	if err != nil {
		log.Fatalf("Не вдалося оновити пацієнта: %v", err)
	}
	log.Printf("Відповідь сервера: %s", r3.GetMessage())
}
