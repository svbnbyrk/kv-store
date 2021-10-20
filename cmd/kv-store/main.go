package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/svbnbyrk/kv-store/handlers"
	"github.com/svbnbyrk/kv-store/internal"
)

func main() {
	l := log.New(os.Stdout, "kv-store ", log.LstdFlags)

	kvs := internal.NewStore()
	hp := handlers.NewStore(l, kvs)
	//read store data in file
	kvs.Read(l)

	sm := http.NewServeMux()
	sm.Handle("/", hp)

	sm.HandleFunc("/flush", hp.FlushStore)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//hardcoded N time interval
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			kvs.Save(l)
		}
	}()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//catch the shutdown signal for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	l.Println("Recieved terminate, graceful shutdown is beginning", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
