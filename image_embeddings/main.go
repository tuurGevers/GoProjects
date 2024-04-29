package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/adaptor/v2"

	admin "admin-service/cmd"
	user "user-service/cmd"
)

type App struct {
	weaver.Implements[weaver.Main]
	UserService  weaver.Listener
	AdminService weaver.Listener
}

func setupLogger() {
	log.SetOutput(os.Stdout) // Direct all log output to stdout
}
func serve(ctx context.Context, a *App) error {
	var wg sync.WaitGroup

	// Initialize User Service
	userFiberApp, err := user.NewService()
	if err != nil {
		log.Printf("Failed to initialize User Service: %v", err)
		return err
	}
	userHandler := adaptor.FiberApp(userFiberApp)
	log.Printf("User Service available on %v\n", a.UserService)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.Serve(a.UserService, userHandler); err != nil {
			log.Fatalf("Failed to start User Service: %v", err)
		}
	}()

	// Initialize Admin Service
	adminFiberApp, err := admin.NewService()
	if err != nil {
		log.Printf("Failed to initialize Admin Service: %v", err)
		return err
	}
	adminHandler := adaptor.FiberApp(adminFiberApp)
	log.Printf("Admin Service available on %v\n", a.AdminService)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.Serve(a.AdminService, adminHandler); err != nil {
			log.Fatalf("Failed to start Admin Service: %v", err)
		}
	}()

	// Wait for both servers to exit
	wg.Wait()
	return nil
}

func main() {

	setupLogger()
	// db.FetchCollection()
	// Handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	log.Println("Starting services...")
	if err := weaver.Run(ctx, serve); err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}

	<-ctx.Done() // Block until context is done
	stop()       // Stop the signal notifier
	log.Println("Shutting down gracefully...")
}
