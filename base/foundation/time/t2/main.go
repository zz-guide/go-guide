package main

import (
	"log"
	"time"
)

func main() {
	var v time.Duration = time.Minute
	log.Print(v.String())
	log.Print(v.Hours())
	log.Print(v.Minutes())
}
