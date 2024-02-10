package main

import (
	"fmt"
	"log"

	"github.com/Snorkin/auth_service/configs"
)

func main() {
	log.Println("Starting auth service")

	cfg, err := configs.GetConfig(".env")
	if err != nil {
		log.Fatalln("Cannot load env variables")
	}
	fmt.Println(cfg)
}
