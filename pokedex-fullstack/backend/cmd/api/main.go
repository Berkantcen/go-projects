package main

import (
	"github.com/berkantcen/pokedex-fullstack/cmd/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", handlers.HomeHandler)
	router.GET("/pokemons", handlers.GetPokemonsHandler)
	router.GET("/pokemons/:id", handlers.GetPokemonHandler)
	router.POST("/pokemons", handlers.CreatePokemonHandler)
	router.DELETE("/pokemons/:id", handlers.DeletePokemonHandler)
	_ = router.Run()
}
