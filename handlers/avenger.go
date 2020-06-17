package handlers

import (
	"api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *Universe) GetAll(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle get all request ")
	avengers := data.Get(u.name)
	err := avengers.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to JSON ", http.StatusInternalServerError)
		u.l.Println("Error: 500 ")
	}
	u.l.Print("\n")
}

func (u *Universe) AddAll(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add all request ")
	avenger := data.Get(u.name)
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	} else {
		data.AddAll(avenger)
		u.l.Print("\n")
	}
}

func (u *Universe) AddChallenge(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add request for challenge ", r.URL)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
	} else {
		data.AddChallenge(avenger, u.name)

	}

}

func (u *Universe) Add(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add request ", r.URL)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
	} else {
		name := data.Add(avenger, u.name)
		if name != "" {
			rw.Write([]byte(name))
		}
		u.l.Print("\n")
	}
}

func (u *Universe) Update(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle update request ", r.RequestURI)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Print("Error: 400 ")
	}
	err = data.Update(avenger, u.name)
	if err == data.ErrorNotFound {
		http.Error(rw, "AntiHero ${u.marvel.Name} not found", http.StatusNoContent)
		u.l.Print("Error: 204 ")
	} else if err != nil {
		http.Error(rw, "AntiHero ${u.marvel.Name} not found", http.StatusInternalServerError)
		u.l.Print("Error: 500 ")
	}
	u.l.Print("\n")
	//	u.l.Println("Updated $avenger")
}

func (u *Universe) Delete(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle DELETE request ", r.URL)
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.Delete(name, u.name)
	if err == nil {
		rw.Write([]byte(name))
	}
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UpdatePower(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle update power request ", r.RequestURI)
	vars := mux.Vars(r)
	name := vars["name"]
	power, err := strconv.Atoi(vars["power"])
	if err != nil || power < -5 || power > 5 {
		http.Error(rw, "Unable to convert power/power change between -5 and +5", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
		return
	}
	err = data.UpdatePower(name, power, u.name)
	//	u.l.Print(name+" "+power)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UseChallenge(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Hande Use challnege request ", r.URL)
	vars := mux.Vars(r)
	name := vars["name"]
	err, removed := data.UseChallenge(name, u.name)
	if removed != "" {
		rw.Write([]byte(removed))
	}
	u.checkErrorName(rw, name, err)
}

func (u *Universe) Use(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	u.l.Print("Handle USE request ", name)
	err := data.Use(name, u.name)
	//u.l.Print(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) Restore(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle Patch request ", r.RequestURI)
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.Restore(name, u.name)
	u.checkErrorName(rw, name, err)
}
