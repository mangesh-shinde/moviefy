package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mangesh-shinde/moviefy/models"
	"github.com/mangesh-shinde/moviefy/utils"
)

func GetAllMovies(c *gin.Context) {
	movies := utils.ReadMovies("movies.json")
	c.IndentedJSON(200, gin.H{"movies": movies})
}

func GetMovieById(c *gin.Context) {
	movieId := c.Param("id")
	movies := utils.ReadMovies("movies.json")
	for _, movie := range movies {
		if movie.ID == movieId {
			c.JSON(200, gin.H{"movie": movie})
			return
		}
	}
	c.JSON(404, gin.H{"message": "No movie found with given id"})
}

func CreateMovie(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer c.Request.Body.Close()
	var movie models.Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		log.Fatal(err.Error())
	}

	movies := utils.ReadMovies("movies.json")
	// err = utils.AppendMovie(movie)
	movies = append(movies, movie)
	fmt.Println(movies)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(201, gin.H{"message": "Movie created successfully", "data": movie})
}

func DeleteMovieById(c *gin.Context) {
	movieId := c.Param("id")
	movies := utils.ReadMovies("movies.json")

	for idx, movie := range movies {
		if movieId == movie.ID {
			movies = append(movies[:idx], movies[idx+1:]...)
			fmt.Println(movies)
			c.JSON(200, gin.H{"message": "Movie deleted successfully!"})
			return
		}
	}

	fmt.Println(movies)
	c.JSON(404, gin.H{"error": "No movie found to delete!"})

}

func DeleteAllMovies(c *gin.Context) {
	mov := utils.ReadMovies("movies.json")
	fmt.Println(mov)
	mov = nil
	c.JSON(200, gin.H{"message": "Deleted all movies"})
}

func UpdateMovieDetails(c *gin.Context) {
	movieId := c.Param("id")

	// get all movies
	movies := utils.ReadMovies("movies.json")

	var IsMovieAvailable bool = false
	var movieIndex int = -1

	// Check if given movie is present in DB
	for idx, movie := range movies {
		if movie.ID == movieId {
			IsMovieAvailable = true
			movieIndex = idx
		}
	}

	if !IsMovieAvailable {
		c.JSON(404, gin.H{"error": "Movie with given ID doesn't exist!"})
		return
	}

	// get movie details from request body
	bytesData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer c.Request.Body.Close()

	var inputMovie models.Movie
	err = json.Unmarshal(bytesData, &inputMovie)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Remove existing movie details
	movieList := append(movies[:movieIndex], movies[movieIndex+1:]...)

	// Add updated movie details
	movieList = append(movieList, inputMovie)

	c.JSON(200, gin.H{"message": "Movie details updated successfully!", "data": movieList})

}
