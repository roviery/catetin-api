package main

import "github.com/roviery/catetin-api/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start("localhost:8080"))
}
