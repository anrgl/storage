package main

import (
	"fmt"
	"log"

	"github.com/anrgl/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()

	f, err := storage.Upload("filename.txt", []byte("test hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it works", storage, f)

	f, err = storage.GetByID(f.FileID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
