package handlers

import (
	"github.com/berkantcen/pokedex-fullstack/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Pokemons = []models.Pokemon{
	{
		ID:      1,
		Name:    "Bulbasaur",
		Type:    "Grass",
		Image:   "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
		Creator: "Original",
	},
	{
		ID:      2,
		Name:    "Ivysaur",
		Type:    "Grass",
		Image:   "https://assets.pokemon.com/assets/cms2/img/pokedex/full/002.png",
		Creator: "Original",
	},
}

// GetPokemonsHandler returns all pokemons
func GetPokemonsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Pokemons)
}

// GetPokemonHandler returns a pokemon by id
func GetPokemonHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, pokemon := range Pokemons {
		if pokemon.ID == id {
			c.JSON(http.StatusOK, pokemon)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}

// CreatePokemonHandler creates a new pokemon
func CreatePokemonHandler(c *gin.Context) {
	var newPokemon models.Pokemon
	if err := c.ShouldBindJSON(&newPokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPokemon.ID = len(Pokemons) + 1
	Pokemons = append(Pokemons, newPokemon)
	c.JSON(http.StatusCreated, Pokemons)
}

// DeletePokemonHandler deletes a pokemon by id
func DeletePokemonHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, pokemon := range Pokemons {
		if pokemon.ID == id {
			Pokemons = append(Pokemons[:index], Pokemons[index+1:]...)
			c.JSON(http.StatusOK, Pokemons)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}
