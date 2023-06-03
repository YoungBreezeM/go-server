package main

import "server/routes"

const (
	port = ":6666"
)

func main() {
	r := routes.SetupRouter()
	r.Run(port)
}
