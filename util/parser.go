package util

import (
	"errors"
	"go-testing-pokeAPI/models"
)

var (
	// error occurs when the type array in pokeAPi response it's not found
	ErrNotFoundPokemonType = errors.New("pokemon type array not found")
	// errr occurs when we found type struct but no name
	ErrNotFoundPokemonTypeName = errors.New("pokemons type name not found")
)

func ParsePokemon(apiPokemon models.PokeApiPokemonResponse) (models.Pokemon, error) {
	if len(apiPokemon.PokemonType) < 1 {
		return models.Pokemon{}, ErrNotFoundPokemonType
	}

	if apiPokemon.PokemonType[0].RefType.NAme == "" {
		return models.Pokemon{}, ErrNotFoundPokemonTypeName
	}

	pokemonType := apiPokemon.PokemonType[0].RefType.NAme

	abilitiesMap := map[string]int{}

	for _, stat := range apiPokemon.Stats {
		parsedAbilityName, ok := models.AllowedAbilities[stat.Stat.NAme]
		if !ok {
			continue
		}
		abilitiesMap[parsedAbilityName] = stat.BaseStat
	}

	parsedPokemon := models.Pokemon{
		Id:        apiPokemon.Id,
		Name:      apiPokemon.Name,
		Power:     pokemonType,
		Abilities: abilitiesMap,
	}

	return parsedPokemon, nil

}
