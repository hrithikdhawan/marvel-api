package handlers

import (
	"api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *Universe) GetAllAntiHeroes(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle get all request")
	AntiHeroes := data.GetAntiHeroes()
	err := AntiHeroes.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to JSON ", http.StatusInternalServerError)
	}
}

func (u *Universe) AddAllAntiHeroes(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle add all request")
	AntiHero := data.Marvel{}
	err := AntiHero.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	data.AddAllAntiHeroes(AntiHero)
}

func (u *Universe) AddAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle add request")
	AntiHero := data.Characters{}
	err := AntiHero.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	data.AddAntiHero(AntiHero)
}

func (u *Universe) UpdateAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle update request")
	AntiHero := data.Characters{}
	err := AntiHero.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	err = data.UpdateAntiHero(AntiHero)
	if err == data.ErrorNotFound {
		http.Error(rw, "AntiHero ${AntiHero.Name} not found", http.StatusNoContent)
	}
	if err != nil {
		http.Error(rw, "AntiHero ${AntiHero.Name} not found", http.StatusInternalServerError)
	}
}

func (u *Universe) DeleteAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle DELETE request")
	vars := mux.Vars(r)
	name := vars["name"]
	// avenger := data.Characters{}
	// err := avenger.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	// }
	err := data.DeleteAntiHero(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UpdatePowerAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle update power request")
	vars := mux.Vars(r)
	name := vars["name"]
	power, err := strconv.Atoi(vars["power"])
	if err != nil || power < -5 || power > 5 {
		http.Error(rw, "Unable to convert power/power change between -5 and +5", http.StatusBadRequest)
		return
	}
	err = data.UpdatePowerAntiHero(name, power)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UseAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle USE request")
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.UseAntiHero(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) RestoreAntiHero(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle Patch request")
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.RestoreAntiHero(name)
	u.checkErrorName(rw, name, err)

}
