package main

import (
	"Product/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.Api()
}
