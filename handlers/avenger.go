package handlers

import (
	"api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *Universe) GetAllAvengers(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle get all request ")
	avengers := data.GetAvengers()
	err := avengers.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to JSON ", http.StatusInternalServerError)
		u.l.Println("Error: 500 ")
	}
	u.l.Print("\n")
}

func (u *Universe) AddAllAvengers(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add all request ")
	avenger := data.Marvel{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	} else {
		data.AddAllAvengers(avenger)
		u.l.Print("\n")
	}
}

func (u *Universe) AddAvengerChallenge(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add request for challenge ", r.URL)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
	} else {
		data.AddAvengerChallenge(avenger)

	}

}

func (u *Universe) AddAvenger(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle add request ", r.URL)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
	} else {
		name := data.AddAvenger(avenger)
		if name != "" {
			rw.Write([]byte(name))
		}
		u.l.Print("\n")
	}
}

func (u *Universe) UpdateAvenger(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle update request ", r.RequestURI)
	avenger := data.Characters{}
	err := avenger.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		u.l.Print("Error: 400 ")
	}
	err = data.UpdateAvenger(avenger)
	if err == data.ErrorNotFound {
		http.Error(rw, "AntiHero ${AntiHero.Name} not found", http.StatusNoContent)
		u.l.Print("Error: 204 ")
	} else if err != nil {
		http.Error(rw, "AntiHero ${AntiHero.Name} not found", http.StatusInternalServerError)
		u.l.Print("Error: 500 ")
	}
	u.l.Print("\n")
	//	u.l.Println("Updated $avenger")
}

func (u *Universe) DeleteAvenger(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle DELETE request ", r.URL)
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.DeleteAvengers(name)
	if err == nil {
		rw.Write([]byte(name))
	}
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UpdatePowerAvenger(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle update power request ", r.RequestURI)
	vars := mux.Vars(r)
	name := vars["name"]
	power, err := strconv.Atoi(vars["power"])
	if err != nil || power < -5 || power > 5 {
		http.Error(rw, "Unable to convert power/power change between -5 and +5", http.StatusBadRequest)
		u.l.Println("Error: 400 ")
		return
	}
	err = data.UpdatePowerAvenger(name, power)
	//	u.l.Print(name+" "+power)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UseAvengerChallenge(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Hande Use challnege request ", r.URL)
	vars := mux.Vars(r)
	name := vars["name"]
	err, removed := data.UseAvengerChallenge(name)
	if removed != "" {
		rw.Write([]byte(removed))
	}
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UseAvenger(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	u.l.Print("Handle USE request ", name)
	err := data.UseAvenger(name)
	//u.l.Print(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) RestoreAvenger(rw http.ResponseWriter, r *http.Request) {
	u.l.Print("Handle Patch request ", r.RequestURI)
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.RestoreAvenger(name)
	u.checkErrorName(rw, name, err)
}
