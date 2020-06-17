package data

var Avengers = Marvel{}

var tempAvengers = []Characters{}

func GetAvengers() Marvel {
	return Avengers

}

func AddAllAvengers(marvel Marvel) {
	Avengers = marvel
	for i, _ := range Avengers.Character {
		Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower / 2
		Avengers.Character[i].Count = 0
	}
}

func AddAvengerChallenge(character Characters) {
	if len(Avengers.Character) == maxCharacters {
		tempAvengers = append(tempAvengers, character)
	} else {
		character.CurrentPower = character.MaxPower / 2
		character.Count = 0

		Avengers.Character = append(Avengers.Character, character)
	}
}

func AddAvenger(character Characters) string {
	character.CurrentPower = character.MaxPower / 2
	character.Count = 0
	var removed string
	if len(Avengers.Character) == maxCharacters {
		index := findLeastUsed(Avengers)
		removed = Avengers.Character[index].Name
		Avengers.Character[index] = character
	} else {
		Avengers.Character = append(Avengers.Character, character)
	}
	return removed
}

func UpdateAvenger(character Characters) error {
	character.CurrentPower = character.MaxPower
	for i, value := range Avengers.Character {
		if value.Name == character.Name {
			Avengers.Character[i] = character
			return nil
		}
	}
	return ErrorNotFound
}

func DeleteAvengers(name string) error {
	for i, value := range Avengers.Character {
		if value.Name == name {
			Avengers.Character = append(Avengers.Character[:i], Avengers.Character[i+1:]...)
			return nil
		}
	}
	return ErrorNotFound
}

func UpdatePowerAvenger(name string, power int) error {
	for i, value := range Avengers.Character {
		if value.Name == name {
			if Avengers.Character[i].CurrentPower+power < Avengers.Character[i].MaxPower {
				Avengers.Character[i].CurrentPower += power
			} else {
				Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower
			}
			//Avengers.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func UseAvengerChallenge(name string) (error, string) {
	var err = UseAvenger(name)
	var removed string
	if err == ErrorNotFound {
		return addAndUseAvenger(name)
	}
	return err, removed
}

func addAndUseAvenger(name string) (error, string) {
	var removed string
	for i, value := range tempAvengers {
		if value.Name == name {
			tempAvengers[i].CurrentPower = tempAvengers[i].MaxPower / 2
			tempAvengers[i].Count = 1
			index := findLeastPower(Avengers)
			removed = Avengers.Character[index].Name
			Avengers.Character[index] = tempAvengers[i]

			return nil, removed
		}
	}
	return ErrorNotFound, removed
}

func UseAvenger(name string) error {
	for i, value := range Avengers.Character {
		if value.Name == name {
			Avengers.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func RestoreAvenger(name string) error {
	for i, value := range Avengers.Character {
		if value.Name == name {
			Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower / 2
			return nil
		}
	}
	return ErrorNotFound
}
