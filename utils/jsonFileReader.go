package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mangesh-shinde/moviefy/models"
)

func ReadMovies(filename string) []models.Movie {
	fileName := filepath.Join("data", filename)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while opening file ", file.Name())
	}
	defer file.Close()

	bytesData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error while reading movies data.", err.Error())
	}

	var Movies []models.Movie
	err = json.Unmarshal(bytesData, &Movies)
	if err != nil {
		log.Fatal("Error while unmarshalling movies data.", err.Error())
	}

	return Movies
}
