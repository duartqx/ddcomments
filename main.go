package main

import (
	"log"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"
)

func main() {

	db, err := r.GetDBConnection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("Hello World!")

}
