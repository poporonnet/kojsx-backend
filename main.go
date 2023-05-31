package main

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/server/router"
)

func main() {
	fmt.Println(`
 ∩_____∩   KOJS v6
 | 0 0 |   "Kemomimi" Online Judge System
 |  ω  |   (C) 2023 Poporon Network / Tatsuto Yamamoto
    `)

	router.StartServer(3060)
}
