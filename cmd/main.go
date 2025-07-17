package main

import (
	"govno/internal/db"
	"govno/internal/routers"
	"log"
)

func main() {
	router := routers.GetAPIv1()

	_, err := db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
