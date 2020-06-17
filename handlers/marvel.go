package handlers

import (
	"api/data"
	"log"
	"net/http"
)

type Universe struct {
	l    *log.Logger
	name string
}

type keyAntiHeroes struct{}

func NewMarvel(l *log.Logger, name string) *Universe {
	return &Universe{l, name}
}

func (u *Universe) checkErrorName(rw http.ResponseWriter, name string, err error) {
	if err == data.ErrorNotFound {
		http.Error(rw, "marvel $name not found", http.StatusNoContent)
		u.l.Print("Error: 204 ")
	} else if err != nil {
		http.Error(rw, "marvel $name not found", http.StatusInternalServerError)
		u.l.Println("Error: 500 ")
	}
	u.l.Print("\n")
}
