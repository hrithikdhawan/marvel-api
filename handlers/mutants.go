package handlers

import (
	"api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *Universe) GetAllMutants(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle get all request")
	Mutants := data.GetMutants()
	err := Mutants.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to JSON ", http.StatusInternalServerError)
	}
}

func (u *Universe) AddAllMutants(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle add all request")
	Mutant := data.Marvel{}
	err := Mutant.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	data.AddAllMutants(Mutant)
}

func (u *Universe) AddMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle add request")
	Mutant := data.Characters{}
	err := Mutant.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	data.AddMutant(Mutant)
}

func (u *Universe) UpdateMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle update request")
	Mutant := data.Characters{}
	err := Mutant.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	}
	err = data.UpdateMutant(Mutant)
	if err == data.ErrorNotFound {
		http.Error(rw, "Mutant ${Mutant.Name} not found", http.StatusNoContent)
	}
	if err != nil {
		http.Error(rw, "Mutant ${Mutant.Name} not found", http.StatusInternalServerError)
	}
}

func (u *Universe) DeleteMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle DELETE request")
	vars := mux.Vars(r)
	name := vars["name"]
	// avenger := data.Characters{}
	// err := avenger.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
	// }
	err := data.DeleteMutant(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UpdatePowerMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle update power request")
	vars := mux.Vars(r)
	name := vars["name"]
	power, err := strconv.Atoi(vars["power"])
	if err != nil || power < -5 || power > 5 {
		http.Error(rw, "Unable to convert power/power change between -5 and +5", http.StatusBadRequest)
		return
	}
	err = data.UpdatePowerMutant(name, power)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) UseMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle USE request")
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.UseMutant(name)
	u.checkErrorName(rw, name, err)
}

func (u *Universe) RestoreMutant(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle Patch request")
	vars := mux.Vars(r)
	name := vars["name"]
	err := data.RestoreMutant(name)
	u.checkErrorName(rw, name, err)

}
