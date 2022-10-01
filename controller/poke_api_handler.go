package controller

import (
	"encoding/json"
	"fmt"
	"go-testing-pokeAPI/models"
	"go-testing-pokeAPI/util"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//respondWithJSON write json response format

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}

}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["idPoke"]

	req := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/%s", id)

	res, err := http.Get(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apiPoke models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPoke)
	if err != nil {
		log.Fatal(err)
	}

	parsedPoke, err := util.ParsePokemon(apiPoke)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("err found: %s", err.Error()))
	}

	respondWithJSON(w, http.StatusOK, parsedPoke)
}
