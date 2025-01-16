package main

import (
	"log"

	"github.com/BassicallyDrums/filmbase/api/apiServer"
)

func main() {

	log.Fatal(apiServer.RunAPIServer())

	// filmsList, err := tmbd.GetFilms()
	// fmt.Println(err)
	// fmt.Println(filmsList)
	// fmt.Print()
}
