package main

import route2 "github.com/juliofilizzola/simulator-api/app/route"

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	println(stringjson[1])
}
