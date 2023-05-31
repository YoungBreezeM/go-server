package main

import "server/routes"

const (
	port = ":8080"
)

func main() {
	r := routes.SetupRouter()
	r.Run(port)
}
