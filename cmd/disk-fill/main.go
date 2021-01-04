package main

import (
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	length := 1024 * 1024

	lengthEnv := os.Getenv("LEN")
	if lengthEnv != "" {
		value, err := strconv.ParseInt(lengthEnv, 10, 32)
		if err != nil {
			panic(err)
		}

		length = int(value)
	}

	rng := rand.New(rand.NewSource(0))

	createdFiles := make([]string, 0)
	defer removeFiles(createdFiles)

	for {
		filename := generateFilename()
		data := generateData(length, rng)

		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		createdFiles = append(createdFiles, filename)

		err = ioutil.WriteFile(filename, data, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func generateFilename() string {
	return uuid.New().String()
}

func generateData(length int, rng *rand.Rand) []byte {
	b := make([]byte, length)

	_, err := rng.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

func removeFiles(names []string) {
	for _, filename := range names {
		err := os.Remove(filename)
		if err != nil {
			log.Println(err)
		}
	}
}
