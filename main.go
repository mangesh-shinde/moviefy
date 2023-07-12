package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mangesh-shinde/moviefy/controllers"
)

func main() {
	fmt.Println("Welcome to Moviefy!")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Moviefy!"})
	})

	// Route to fetch all movies list
	r.GET("/api/movies", controllers.GetAllMovies)

	// Route to fetch movie by id
	r.GET("/api/movies/:id", controllers.GetMovieById)

	// Route to create a movie
	r.POST("/api/movies", controllers.CreateMovie)

	// Route to update movie details
	r.PUT("/api/movies/:id", controllers.UpdateMovieDetails)

	// Route to delete a movie (by id)
	r.DELETE("/api/movies/:id", controllers.DeleteMovieById)

	// Route to delete all movies
	r.DELETE("/api/movies", controllers.DeleteAllMovies)

	fmt.Printf("Starting server on port 8080")
	r.Run(":8080")

}
