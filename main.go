package main

import (
	database "program-rme-be/Database"
	route "program-rme-be/Route"
)

func main() {
	database.ConnectDatabase()

	route.Routes()
}
