package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Service represents a service that needs to be gracefully shut down.
type Service struct {
	Name string
}

// Start starts the service.
func (s *Service) Start() {
	fmt.Printf("[%s] Service started\n", s.Name)
	// Simulate some work
	for {
		select {
		case <-time.After(time.Second):
			fmt.Printf("[%s] Doing some work...\n", s.Name)
		}
	}
}

// Stop stops the service.
func (s *Service) Stop() {
	fmt.Printf("[%s] Service stopped\n", s.Name)
}

func main() {
	// Create a quit channel to receive termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Create services
	service1 := &Service{Name: "Service 1"}
	service2 := &Service{Name: "Service 2"}

	// Start services in separate goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		service1.Start()
	}()
	go func() {
		defer wg.Done()
		service2.Start()
	}()

	// Wait for termination signal
	<-quit
	fmt.Println("Received termination signal. Shutting down...")

	// Stop services
	service1.Stop()
	service2.Stop()

	// Wait for services to stop gracefully
	wg.Wait()

	fmt.Println("All services stopped. Exiting.")
}
