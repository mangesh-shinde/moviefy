package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mangesh-shinde/moviefy/models"
)

func AppendMovie(movie models.Movie) error {
	movies := ReadMovies("movies.json")
	movies = append(movies, movie)

	tmpFileName := fmt.Sprintf("movies.json.%d", time.Now().UnixMicro())
	fmt.Println(tmpFileName)
	err := WriteMovieDataToTempFile(tmpFileName, movies)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Rename 'movies.json' to some new name
	dstFile := fmt.Sprintf("%s%s", tmpFileName, ".bkp")
	err = RenameFile("movies.json", dstFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Rename updated file to 'movies.json'
	err = RenameFile(tmpFileName, "movies.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Remove(dstFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func WriteMovieDataToTempFile(fileName string, movies []models.Movie) error {
	// Create a temp file
	tmpFile, err := os.Create(filepath.Join("data", fileName))
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	// transform data struture to bytes for writing to file
	bytesData, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		return err
	}

	// Write data to file
	_, err = tmpFile.Write(bytesData)
	if err != nil {
		return err
	}

	fmt.Printf("Data writtern to tmp file.\n")

	return nil
}

func RenameFile(srcFile string, destFile string) error {

	oldFileName := filepath.Join("data", srcFile)
	newFileName := filepath.Join("data", destFile)

	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		return err
	}

	return nil
}
