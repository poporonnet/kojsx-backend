package main

import (
	"fmt"
	"os"

	"github.com/poporonnet/kojsx-backend/pkg/server/router"
)

var (
	VERSION  = "v9.0.0-alpha"
	REVISION = "dev"
)

func main() {
	fmt.Printf(`
   .__,  _   
  / n | |∩|   KOJS v9 (%s @%s)
  ∪ | |_| |   "Kemomimi" Online Judge System
   /       \  (C) 2025 Poporon Network / Tatsuto Yamamoto
  |  ε  , ε | 
   \    ω  /

`, VERSION, REVISION)
	mode := os.Getenv("KOJS_MODE")
	router.StartServer(3060, mode)
}
