package main

import (
	"fmt"
	"os"

	"github.com/poporonnet/kojsx-backend/pkg/server/router"
)

var (
	VERSION  = "v6.0.0"
	REVISION = "dev"
)

func main() {
	fmt.Printf(`
 ∩_____∩   KOJS v6 (%s @%s)
 | 0 0 |   "Kemomimi" Online Judge System
 |  ω  |   (C) 2023 Poporon Network / Tatsuto Yamamoto
`, VERSION, REVISION)
	mode := os.Getenv("KOJS_MODE")
	router.StartServer(3060, mode)
}
