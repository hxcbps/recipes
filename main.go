// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//  Schemes: http
//  Host: localhost:8080
//  BasePath: /
//  Version: 1.0.0
//  Contact: Team Mr.Meseeks
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

// swagger:operation POST /recipes recipes newRecipe
// ---
// summary: Crear una nueva receta
// description: Crear una nueva receta
// tags:
// - recipes
// produces:
// - application/json
// parameters:
//   - name: recipe
//     in: body
//     description: Recipe object
//     required: true
//
// responses:
//
//	'200':
//	  description: OK
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

// swagger:operation GET /recipes recipes listRecipes
// ---
// summary: Obtener lista de recetas
// description: Obtener todas las recetas existentes
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation GET /recipes/search recipes searchRecipes
// ---
// summary: Buscar recetas por tag
// description: Devuelve una lista de recetas filtradas por una etiqueta
// produces:
// - application/json
// parameters:
//   - in: query
//     name: tag
//     description: Tag to filter by
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	  description: List of recipes filtered by tag
//	  schema:
//	    type: array
//	    items:
//	      "$ref": "#/definitions/Recipe"
//	'400':
//	  description: Bad request
func SearchRecipesHandler(c *gin.Context) {

	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)
	for _, recipe := range recipes {
		found := false
		for _, t := range recipe.Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipe)
		}
	}
	c.JSON(http.StatusOK, listOfRecipes)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// ---
// summary: Actualizar una receta existente
// description: Actualiza una receta existente mediante su ID
// tags:
// - recipes
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID de la receta a actualizar
//     required: true
//     type: string
//   - name: recipe
//     in: body
//     description: Objeto de la receta a actualizar
//     required: true
//
// responses:
//
//	'200':
//	  description: Receta actualizada
//	'404':
//	  description: Receta no encontrada
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedRecipe Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	for index, recipe := range recipes {
		if recipe.ID == id {
			updatedRecipe.ID = recipe.ID
			updatedRecipe.PublishedAt = recipe.PublishedAt
			recipes[index] = updatedRecipe
			c.JSON(http.StatusOK, updatedRecipe)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Recipe not found"})
}

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// ---
// summary: Eliminar receta por ID
// description: Eliminar recetas que coincidan con un ID
// produces:
// - application/json
// parameters:
//   - in: path
//     name: id
//     description: ID of the recipe to delete
//     required: true
//     type: string
//
// responses:
//
//	'204':
//	  description: Recipe deleted successfully
//	'404':
//	  description: Recipe not found
//	'500':
//	  description: Internal server error
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	for index, recipe := range recipes {
		if recipe.ID == id {
			recipes = append(recipes[:index], recipes[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"status": "Recipe deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Recipe not found"})
}

func main() {
	router := gin.Default()
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/search", SearchRecipesHandler)
	router.POST("/recipes", NewRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.Run()
}
