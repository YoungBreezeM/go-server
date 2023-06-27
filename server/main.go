package main

import (
	"fmt"
	"log"
	"server/config"
	"server/routes"
)

func main() {
	r := routes.SetupRouter()
	if err := r.Run(fmt.Sprintf("%s:%d", config.Cfg.Host, config.Cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
