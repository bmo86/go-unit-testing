package util

import (
	"encoding/json"
	"go-testing-pokeAPI/models"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)
	body, err := ioutil.ReadFile("samples/pokeApi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expectedPoke models.Pokemon

	err = json.Unmarshal([]byte(body), &expectedPoke)
	c.NoError(err)

	c.Equal(expectedPoke, parsedPokemon)

}
