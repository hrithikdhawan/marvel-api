package data

var Avengers = Marvel{Name: "Avengers"}

var AntiHeroes = Marvel{Name: "AntiHeroes"}

var Mutants = Marvel{Name: "Mutants"}

var tempAvengers = []Characters{}

func Get(name string) Marvel {
	switch name {
	case "Avengers":
		return Avengers
	case "AntiHeroes":
		return AntiHeroes
	case "Mutants":
		return Mutants

	}
	return Marvel{}
}

func AddAll(marvel Marvel) {
	switch marvel.Name {
	case "Avengers":
		Avengers = marvel
		for i, _ := range Avengers.Character {
			Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower / 2
			Avengers.Character[i].Count = 0
		}
	case "Anti Heroes":
		AntiHeroes = marvel
		for i, _ := range AntiHeroes.Character {
			AntiHeroes.Character[i].CurrentPower = AntiHeroes.Character[i].MaxPower / 2
			AntiHeroes.Character[i].Count = 0
		}
	case "Mutants":
		Mutants = marvel
		for i, _ := range Mutants.Character {
			Mutants.Character[i].CurrentPower = Mutants.Character[i].MaxPower / 2
			AntiHeroes.Character[i].Count = 0
		}
	}
}

func AddChallenge(character Characters, name string) {
	character.CurrentPower = character.MaxPower / 2
	character.Count = 0
	switch name {
	case "Avengers":
		if len(Avengers.Character) == maxCharacters {
			tempAvengers = append(tempAvengers, character)
		} else {
			Avengers.Character = append(Avengers.Character, character)
		}
	case "AntiHeroes":
		if len(AntiHeroes.Character) == maxCharacters {
			tempAvengers = append(tempAvengers, character)
		} else {
			AntiHeroes.Character = append(AntiHeroes.Character, character)
		}
	case "Mutants":
		if len(Mutants.Character) == maxCharacters {
			tempAvengers = append(tempAvengers, character)
		} else {
			Mutants.Character = append(Mutants.Character, character)
		}
	}
}

func Add(character Characters, name string) string {
	character.CurrentPower = character.MaxPower / 2
	character.Count = 0
	var removed string
	switch name {
	case "Avengers":
		if len(Avengers.Character) == maxCharacters {
			index := findLeastUsed(Avengers)
			removed = Avengers.Character[index].Name
			Avengers.Character[index] = character
		} else {
			Avengers.Character = append(Avengers.Character, character)
		}
	case "AntiHeroes":
		if len(AntiHeroes.Character) == maxCharacters {
			index := findLeastUsed(AntiHeroes)
			removed = AntiHeroes.Character[index].Name
			AntiHeroes.Character[index] = character
		} else {
			AntiHeroes.Character = append(AntiHeroes.Character, character)
		}
	case "Mutants":
		if len(Mutants.Character) == maxCharacters {
			index := findLeastUsed(Mutants)
			removed = Mutants.Character[index].Name
			Mutants.Character[index] = character
		} else {
			Mutants.Character = append(Mutants.Character, character)
		}
	}
	return removed
}

func Update(character Characters, name string) error {
	character.CurrentPower = character.MaxPower / 2
	switch name {
	case "Avengers":
		for i, value := range Avengers.Character {
			if value.Name == character.Name {
				Avengers.Character[i] = character
				return nil
			}
		}
		return ErrorNotFound
	case "AntiHeroes":
		for i, value := range AntiHeroes.Character {
			if value.Name == character.Name {
				AntiHeroes.Character[i] = character
				return nil
			}
		}
		return ErrorNotFound
	case "Mutants":
		for i, value := range Mutants.Character {
			if value.Name == character.Name {
				Mutants.Character[i] = character
				return nil
			}
		}
		return ErrorNotFound
	}
	return ErrorNotFound
}

func Delete(name string, marvel string) error {
	switch marvel {
	case "Avengers":
		for i, value := range Avengers.Character {
			if value.Name == name {
				Avengers.Character = append(Avengers.Character[:i], Avengers.Character[i+1:]...)
				return nil
			}
		}
		return ErrorNotFound
	case "AntiHeroes":
		for i, value := range AntiHeroes.Character {
			if value.Name == name {
				AntiHeroes.Character = append(AntiHeroes.Character[:i], AntiHeroes.Character[i+1:]...)
				return nil
			}
		}
		return ErrorNotFound
	case "Mutants":
		for i, value := range Mutants.Character {
			if value.Name == name {
				Mutants.Character = append(Mutants.Character[:i], Mutants.Character[i+1:]...)
				return nil
			}
		}
		return ErrorNotFound
	}
	return ErrorNotFound
}

func UpdatePower(name string, power int, marvel string) error {
	switch marvel {
	case "Avengers":
		for i, value := range Avengers.Character {
			if value.Name == name {
				if Avengers.Character[i].CurrentPower+power < Avengers.Character[i].MaxPower {
					Avengers.Character[i].CurrentPower += power
				} else {
					Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower
				}
				return nil
			}
		}
		return ErrorNotFound
	case "AntiHeroes":
		for i, value := range AntiHeroes.Character {
			if value.Name == name {
				if AntiHeroes.Character[i].CurrentPower+power < AntiHeroes.Character[i].MaxPower {
					AntiHeroes.Character[i].CurrentPower += power
				} else {
					AntiHeroes.Character[i].CurrentPower = AntiHeroes.Character[i].MaxPower
				}
				return nil
			}
		}
		return ErrorNotFound
	case "Mutants":
		for i, value := range Mutants.Character {
			if value.Name == name {
				if Mutants.Character[i].CurrentPower+power < Mutants.Character[i].MaxPower {
					Mutants.Character[i].CurrentPower += power
				} else {
					Mutants.Character[i].CurrentPower = Mutants.Character[i].MaxPower
				}
				return nil
			}
		}
		return ErrorNotFound
	}
	return ErrorNotFound
}

func UseChallenge(name string, marvel string) (error, string) {
	var err = Use(name, marvel)
	var removed string
	if err == ErrorNotFound {
		return addAndUse(name, marvel)
	}
	return err, removed
}

func addAndUse(name string, marvel string) (error, string) {
	var removed string
	for i, value := range tempAvengers {
		if value.Name == name {
			tempAvengers[i].CurrentPower = tempAvengers[i].MaxPower / 2
			tempAvengers[i].Count = 1
			switch marvel {
			case "Avengers":
				index := findLeastPower(Avengers)
				removed = Avengers.Character[index].Name
				Avengers.Character[index] = tempAvengers[i]
				return nil, removed
			case "AntiHeroes":
				index := findLeastPower(AntiHeroes)
				removed = AntiHeroes.Character[index].Name
				AntiHeroes.Character[index] = tempAvengers[i]
				return nil, removed
			case "Mutants":
				index := findLeastPower(Mutants)
				removed = Mutants.Character[index].Name
				Mutants.Character[index] = tempAvengers[i]
				return nil, removed
			}

		}
	}
	return ErrorNotFound, removed
}

func Use(name string, marvel string) error {
	switch marvel {
	case "Avengers":
		for i, value := range Avengers.Character {
			if value.Name == name {
				Avengers.Character[i].Count += 1
				return nil
			}
		}
	case "AntiHeroes":
		for i, value := range AntiHeroes.Character {
			if value.Name == name {
				AntiHeroes.Character[i].Count += 1
				return nil
			}
		}
	case "Mutants":
		for i, value := range Mutants.Character {
			if value.Name == name {
				Mutants.Character[i].Count += 1
				return nil
			}
		}
	}
	return ErrorNotFound
}

func Restore(name string, marvel string) error {
	switch marvel {
	case "Avengers":
		for i, value := range Avengers.Character {
			if value.Name == name {
				Avengers.Character[i].CurrentPower = Avengers.Character[i].MaxPower / 2
				return nil
			}
		}
	case "AntiHeroes":
		for i, value := range AntiHeroes.Character {
			if value.Name == name {
				AntiHeroes.Character[i].CurrentPower = AntiHeroes.Character[i].MaxPower / 2
				return nil
			}
		}
	case "Mutants":
		for i, value := range Mutants.Character {
			if value.Name == name {
				Mutants.Character[i].CurrentPower = Mutants.Character[i].MaxPower / 2
				return nil
			}
		}
	}
	return ErrorNotFound
}
