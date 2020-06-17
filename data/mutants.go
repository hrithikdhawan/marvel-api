package data

var Mutants = Marvel{}

func GetMutants() Marvel {
	return Mutants
}

func AddAllMutants(marvel Marvel) {
	Mutants = marvel
	for i, _ := range Mutants.Character {
		Mutants.Character[i].CurrentPower = Mutants.Character[i].MaxPower
		Mutants.Character[i].Count = 0
	}
}

func AddMutant(character Characters) {
	character.CurrentPower = character.MaxPower
	character.Count = 0
	if len(Mutants.Character) == maxCharacters {
		index := findLeastUsed(Mutants)
		Mutants.Character[index] = character
	} else {
		Mutants.Character = append(Mutants.Character, character)
	}
}

func UpdateMutant(character Characters) error {
	character.CurrentPower = character.MaxPower
	for i, value := range Mutants.Character {
		if value.Name == character.Name {
			Mutants.Character[i] = character
			return nil
		}
	}
	return ErrorNotFound
}

func DeleteMutant(name string) error {
	for i, value := range Mutants.Character {
		if value.Name == name {
			Mutants.Character[i] = Characters{}
			return nil
		}
	}
	return ErrorNotFound
}

func UpdatePowerMutant(name string, power int) error {
	for i, value := range Mutants.Character {
		if value.Name == name {
			Mutants.Character[i].CurrentPower += power
			Mutants.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func UseMutant(name string) error {
	for i, value := range Mutants.Character {
		if value.Name == name {
			Mutants.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func RestoreMutant(name string) error {
	for i, value := range Mutants.Character {
		if value.Name == name {
			Mutants.Character[i].CurrentPower = Mutants.Character[i].MaxPower
			return nil
		}
	}
	return ErrorNotFound
}
