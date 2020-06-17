package main

import (
	"api/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var servemux = mux.NewRouter()

func main() {
	l := log.New(os.Stdout, "marvel ", log.LstdFlags)
	handler := handlers.NewUniverse(l)

	setUpAvengers(handler)

	setUpAntiHeroes(handler)

	setUpMutants(handler)

	s := http.Server{
		Addr:     ":9876",
		Handler:  servemux,
		ErrorLog: l,
	}

	go func() {
		l.Println("Starting server on port 9876")
		error := s.ListenAndServe()
		if error != nil {
			l.Println("Error starting server: $error")
			os.Exit(1)
		}
	}()
	//	tests.Test()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	_ = <-c
	l.Println("Got signal")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}

func setUpAvengers(handler *handlers.Universe) {
	getRouter := servemux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/avengers", handler.GetAllAvengers)
	getRouter.HandleFunc("/avengers/character/{name}", handler.UseAvenger)
	getRouter.HandleFunc("/avengers/character/challenge/{name}", handler.UseAvengerChallenge)

	postRouter := servemux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/avengers/character", handler.AddAvenger)
	postRouter.HandleFunc("/avengers/character/challenge", handler.AddAvengerChallenge)

	postRouter.HandleFunc("/avengers", handler.AddAllAvengers) // update All also

	putRouter := servemux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/avengers", handler.UpdateAvenger) // single update
	getRouter.HandleFunc("/avengers/character/{name}/{power}", handler.UpdatePowerAvenger)

	deleteRouter := servemux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/avengers/character/{name}", handler.DeleteAvenger)

	patchtRouter := servemux.Methods(http.MethodPatch).Subrouter()
	patchtRouter.HandleFunc("/avengers/character/{name}", handler.RestoreAvenger)

}

func setUpAntiHeroes(handler *handlers.Universe) {
	getRouter := servemux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/anti-heroes", handler.GetAllAntiHeroes)
	getRouter.HandleFunc("/anti-heroes/character/{name}", handler.UseAntiHero)

	postRouter := servemux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/anti-heroes/character", handler.AddAntiHero)
	postRouter.HandleFunc("/anti-heroes", handler.AddAllAntiHeroes) // update All also

	putRouter := servemux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/anti-heroes", handler.UpdateAntiHero) // single update
	getRouter.HandleFunc("/anti-heroes/character/{name}/{power}", handler.UpdatePowerAntiHero)

	deleteRouter := servemux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/anti-heroes/character/{name}", handler.DeleteAntiHero)

	patchtRouter := servemux.Methods(http.MethodPatch).Subrouter()
	patchtRouter.HandleFunc("/anti-heroes/character/{name}", handler.RestoreAntiHero)

}

func setUpMutants(handler *handlers.Universe) {
	getRouter := servemux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/mutants", handler.GetAllMutants)
	getRouter.HandleFunc("/mutants/character/{name}", handler.UseMutant)

	postRouter := servemux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/mutants/character", handler.AddMutant)
	postRouter.HandleFunc("/mutants", handler.AddAllMutants) // update All also

	putRouter := servemux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/mutants", handler.UpdateMutant) // single update
	getRouter.HandleFunc("/mutants/character/{name}/{power}", handler.UpdatePowerMutant)

	deleteRouter := servemux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/mutants/character/{name}", handler.DeleteMutant)

	patchtRouter := servemux.Methods(http.MethodPatch).Subrouter()
	patchtRouter.HandleFunc("/mutants/character/{name}", handler.RestoreMutant)

}
