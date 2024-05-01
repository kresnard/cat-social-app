package main

import (
	"cat_social_app/config"
	"cat_social_app/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Conservice-xyzfig error: %s", err)
	}
	log.Println("start running app")

	app.Run(cfg)
}
