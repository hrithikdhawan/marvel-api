package handlers

import (
	"api/data"
	"log"
	"net/http"
)

type Universe struct {
	l *log.Logger
	//	marvel data.Marvel
}

type keyAntiHeroes struct{}

func NewUniverse(l *log.Logger) *Universe {
	return &Universe{l}
}

func (u *Universe) checkErrorName(rw http.ResponseWriter, name string, err error) {
	if err == data.ErrorNotFound {
		http.Error(rw, "AntiHero $name not found", http.StatusNoContent)
		u.l.Print("Error: 204 ")
	} else if err != nil {
		http.Error(rw, "AntiHero $name not found", http.StatusInternalServerError)
		u.l.Println("Error: 500 ")
	}
	u.l.Print("\n")
}
