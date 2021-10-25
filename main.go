package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andreastihor/stockbit/config"
	"github.com/andreastihor/stockbit/pkg/movie"
	routes "github.com/andreastihor/stockbit/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var (
	addr string = ":3000"
)

func main() {
	l := log.New(os.Stdout, "Movie API ", log.LstdFlags)
	router := mux.NewRouter()
	//router.Host("http://localhost/")
	corsMw := mux.CORSMethodMiddleware(router)
	router.Use(corsMw)
	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), //
		handlers.AllowedHeaders([]string{"Origin", "Authorization", "X-Requested-With", "Content-Type", "Accept"}),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
		handlers.AllowCredentials(),
	)

	// Initiate Package for Injection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Env.DSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Print(err)
	}

	repo := movie.NewRepository(db)
	service := movie.NewService(repo)
	handler := movie.NewHandler(service)
	routes.NewRouter(router, handler)

	// create a new server
	s := http.Server{
		Addr:         addr,              // configure the bind address
		Handler:      ch(router),        // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 50 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port " + addr)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//trap sigterm or interuppt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = s.Shutdown(ctx)
	if err != nil {
		log.Println("Error on shutting down : ", err)
	}

}
