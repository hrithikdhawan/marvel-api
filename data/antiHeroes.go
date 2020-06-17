package data

var AntiHeroes = Marvel{}

func GetAntiHeroes() Marvel {
	return AntiHeroes
}

func AddAllAntiHeroes(marvel Marvel) {
	AntiHeroes = marvel
	for i, _ := range AntiHeroes.Character {
		AntiHeroes.Character[i].CurrentPower = AntiHeroes.Character[i].MaxPower
		AntiHeroes.Character[i].Count = 0
	}
}

func AddAntiHero(character Characters) {
	character.CurrentPower = character.MaxPower
	character.Count = 0
	if len(AntiHeroes.Character) == maxCharacters {
		index := findLeastUsed(AntiHeroes)
		AntiHeroes.Character[index] = character
	} else {
		AntiHeroes.Character = append(AntiHeroes.Character, character)
	}
}

func UpdateAntiHero(character Characters) error {
	character.CurrentPower = character.MaxPower
	for i, value := range AntiHeroes.Character {
		if value.Name == character.Name {
			AntiHeroes.Character[i] = character
			return nil
		}
	}
	return ErrorNotFound
}

func DeleteAntiHero(name string) error {
	for i, value := range AntiHeroes.Character {
		if value.Name == name {
			AntiHeroes.Character[i] = Characters{}
			return nil
		}
	}
	return ErrorNotFound
}

func UpdatePowerAntiHero(name string, power int) error {
	for i, value := range AntiHeroes.Character {
		if value.Name == name {
			AntiHeroes.Character[i].CurrentPower += power
			AntiHeroes.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func UseAntiHero(name string) error {
	for i, value := range AntiHeroes.Character {
		if value.Name == name {
			AntiHeroes.Character[i].Count += 1
			return nil
		}
	}
	return ErrorNotFound
}

func RestoreAntiHero(name string) error {
	for i, value := range AntiHeroes.Character {
		if value.Name == name {
			AntiHeroes.Character[i].CurrentPower = AntiHeroes.Character[i].MaxPower
			return nil
		}
	}
	return ErrorNotFound
}
