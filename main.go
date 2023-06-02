package main

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/server/router"
)

var (
	VERSION  = "v6.0.0"
	REVISION = ""
)

func main() {
	fmt.Printf(`
 ∩_____∩   KOJS v6 (%s @%s)
 | 0 0 |   "Kemomimi" Online Judge System
 |  ω  |   (C) 2023 Poporon Network / Tatsuto Yamamoto
    `, VERSION, REVISION)
	router.StartServer(3060)
}
