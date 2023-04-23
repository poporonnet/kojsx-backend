package main

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/server/router"
)

func main() {
	fmt.Println(`
KOJS
version 6.0.0.pre-alpha.0
(C) 2023 Poporon Network
    `)

	router.StartServer(3060)
}
