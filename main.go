package main

import (
	"fmt"
	"github.com/joho/godotenv"
	route2 "github.com/juliofilizzola/simulator-api/app/route"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Printf(stringjson[1])
}
